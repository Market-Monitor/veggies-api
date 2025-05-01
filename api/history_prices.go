package api

import (
	"context"

	"github.com/Market-Monitor/veggies-api/api/types"
	"github.com/Market-Monitor/veggies-api/api/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
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
// @Router /api/{tradingCenter}/history_prices [get]
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

// GetLatestHistoryPrices returns the latest history prices of all veggies
// @Summary Get latest history prices
// @Description Get the latest history prices of all veggies
// @Tags history_prices
// @Accept json
// @Produce json
// @Param tradingCenter path string true "Trading Center"
// @Success 200 {object} utils.HTTPSuccessResponse
// @Failure 500 {object} utils.HTTPErrorResponse
// @Router /api/{tradingCenter}/latest_history_prices [get]
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

	filterImageSource := func(veggieId string) string {
		for _, veggie := range veggies {
			if veggie.ID == veggieId {
				return veggie.ImageURL
			}
		}

		return ""
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

		data = append(data, types.LatestHistoryPrice{
			ImageSource:     filterImageSource(historyPrice.ParentId),
			ParentId:        historyPrice.ParentId,
			ParentName:      historyPrice.ParentName,
			Category:        historyPrice.Category,
			TradingCenterId: historyPrice.TradingCenterId,
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
