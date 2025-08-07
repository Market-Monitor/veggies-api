package api

import (
	"context"
	"fmt"

	"github.com/Market-Monitor/veggies-api/api/types"
	"github.com/Market-Monitor/veggies-api/api/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// GetHistoryPrices returns all history prices of all veggies
// @Summary Get history prices
// @Description Get all history prices of all veggies
// @Tags history_prices
// @Accept json
// @Produce json
// @Param tradingCenter path string true "Trading Center"
// @Success 200 {object} utils.HTTPSuccessResponse
// @Failure 500 {object} utils.HTTPErrorResponse
// @Router /api/{tradingCenter}/history-prices [get]
func GetHistoryPrices(c *fiber.Ctx) error {
	db := GetTD_DB(c)
	coll := db.Collection(COLL_HISTORY_PRICES)

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return utils.ResError(c, 500, err)
	}

	var historyPrices []types.VeggiePrice
	if err := cursor.All(context.TODO(), &historyPrices); err != nil {
		return utils.ResError(c, 500, err)
	}

	return utils.ResSuccess(c, 200, historyPrices)
}

type HistoryIdsBody struct {
	Filters []struct {
		ID            string `json:"id"`
		TradingCenter string `json:"tradingCenter"`
	} `json:"filters" xml:"filters" form:"filters"`
}

// GetLatestHistoryPricesIds returns the latest history prices of veggies by their IDs
// @Summary Get latest history prices by IDs, if the veggie is not found, it will be skipped
// @Description Get the latest history prices of veggies by their IDs
// @Tags history_prices
// @Accept json
// @Produce json
// @Param body body HistoryIdsBody true "Filters for veggies"
// @Success 200 {object} utils.HTTPSuccessResponse
// @Failure 500 {object} utils.HTTPErrorResponse
// @Router /api/veggies/ids [post]
func GetLatestHistoryPricesIds(c *fiber.Ctx) error {
	body := new(HistoryIdsBody)
	if err := c.BodyParser(body); err != nil {
		return utils.ResError(c, 400, fmt.Errorf("invalid request body: %v", err))
	}

	// Get all veggie latest prices & data
	var veggies []types.FilterVeggies

	for _, v := range body.Filters {
		db := getDb(v.TradingCenter)
		veggieColl := db.Collection(COLL_VEGGIES)
		historyColl := db.Collection(COLL_HISTORY_PRICES)

		var veggie types.Veggie
		if err := veggieColl.FindOne(context.TODO(), bson.M{"id": v.ID}).Decode(&veggie); err != nil {
			// skip
			fmt.Println("Veggie not found:", v.ID, "in trading center:", v.TradingCenter)

			continue
		}

		// Get first possible latest history price
		var historyPrice types.VeggiePrice

		hpOpts := options.FindOne().SetSort(bson.D{{Key: "dateUnix", Value: -1}})
		if err := historyColl.FindOne(context.TODO(), bson.M{
			"parentId":      veggie.ID,
			"tradingCenter": v.TradingCenter,
		}, hpOpts).Decode(&historyPrice); err != nil {
			return utils.ResError(c, 500, err)
		}

		// Get latest history price for this veggie
		var latestClassPrices []types.LatestHistoryPriceClass

		cursor, err := historyColl.Find(context.TODO(), bson.M{
			"parentId":      veggie.ID,
			"tradingCenter": v.TradingCenter,
			"dateUnix":      historyPrice.DateUnix,
		})
		if err != nil {
			return utils.ResError(c, 500, err)
		}

		if err = cursor.All(context.TODO(), &latestClassPrices); err != nil {
			return utils.ResError(c, 500, err)
		}

		veggies = append(veggies, types.FilterVeggies{
			ImageSource:      veggie.ImageSource,
			ParentId:         veggie.ID,
			ParentName:       veggie.Name,
			Category:         veggie.TradingCenter,
			TradingCenter:    veggie.TradingCenter,
			PriceUnit:        veggie.PriceUnit,
			LatestUpdateDate: historyPrice.DateUnix,
			Classes:          latestClassPrices,
		})
	}

	return utils.ResSuccess(c, 200, veggies)
}

// GetLatestHistoryPrices returns the latest history prices of all veggies
// @Summary Get latest history prices
// @Description Get the latest history prices of all veggies
// @Tags history_prices
// @Accept json
// @Produce json
// @Param tradingCenter path string true "Trading Center"
// @Success 200 {object} utils.HTTPSuccessResponse
// @Failure 500 {object} utils.HTTPErrorResponse
// @Router /api/{tradingCenter}/history-prices/latest [get]
func GetLatestHistoryPrices(c *fiber.Ctx) error {
	db := GetTD_DB(c)

	// START get config
	configColl := db.Collection(COLL_CONFIGURATIONS)

	var latestConfig types.Configuration

	if err := configColl.FindOne(context.TODO(), bson.M{
		"configId": 0,
	}).Decode(&latestConfig); err != nil {
		return utils.ResError(c, 500, err)
	}
	// END get config

	// START get latest history prices
	coll := db.Collection(COLL_HISTORY_PRICES)
	cursor, err := coll.Find(context.TODO(), bson.M{
		"dateISO": latestConfig.LatestDataDate,
	})
	if err != nil {
		return utils.ResError(c, 500, err)
	}

	var historyPrices []types.VeggiePrice
	if err := cursor.All(context.TODO(), &historyPrices); err != nil {
		return utils.ResError(c, 500, err)
	}
	// END get latest history pricess

	// START get all veggies (for images getting)
	veggieColl := db.Collection(COLL_VEGGIES)
	veggieCollCursor, veggieCollErr := veggieColl.Find(context.TODO(), bson.D{})
	if veggieCollErr != nil {
		return utils.ResError(c, 500, err)
	}

	var veggies []types.Veggie
	if err := veggieCollCursor.All(context.TODO(), &veggies); err != nil {
		return utils.ResError(c, 500, err)
	}
	// END get all veggies

	filterVeggie := func(veggieId string) *types.Veggie {
		for _, veggie := range veggies {
			if veggie.ID == veggieId {
				return &veggie
			}
		}

		return nil
	}

	// START parse latest history
	var data []types.LatestHistoryPrice

	for _, historyPrice := range historyPrices {
		// check if parent exists
		var parentExists bool
		var parentIndex int
		for idx, d := range data {
			if d.ParentId == historyPrice.ParentId {
				parentExists = true
				parentIndex = idx
			}
		}

		if parentExists {
			data[parentIndex].Classes = append(data[parentIndex].Classes, types.LatestHistoryPriceClass{
				ID:    historyPrice.ID,
				Name:  historyPrice.Name,
				Price: historyPrice.Price,
			})

			continue
		}

		veggie := filterVeggie(historyPrice.ParentId)
		veggieImageUrl := ""
		veggiePriceUnit := ""
		if veggie == nil {
			// default values if veggie not found
			veggiePriceUnit = "kilo"
		} else {
			veggieImageUrl = veggie.ImageURL
			veggiePriceUnit = veggie.PriceUnit
		}

		data = append(data, types.LatestHistoryPrice{
			PriceUnit:     veggiePriceUnit,
			ImageSource:   veggieImageUrl,
			ParentId:      historyPrice.ParentId,
			ParentName:    historyPrice.ParentName,
			Category:      historyPrice.Category,
			TradingCenter: historyPrice.TradingCenter,
			Classes: []types.LatestHistoryPriceClass{
				{
					ID:    historyPrice.ID,
					Name:  historyPrice.Name,
					Price: historyPrice.Price,
				},
			},
		})
	}
	// END parse latest history

	return utils.ResSuccess(c, 200, fiber.Map{
		"latestDataDate": latestConfig.LatestDataDate,
		"data":           data,
	})
}
