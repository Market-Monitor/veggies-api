package api

import (
	"context"

	"github.com/Market-Monitor/veggies-api/api/types"
	"github.com/Market-Monitor/veggies-api/api/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GetAllVeggiesWithClasses returns all veggies with their classes
// @Summary Get all veggies with classes
// @Description Get all veggies with their classes
// @Tags veggies
// @Accept json
// @Produce json
// @Success 200 {object} utils.HTTPSuccessResponse
// @Failure 500 {object} utils.HTTPErrorResponse
// @Router /api/veggies [get]
func GetAllVeggiesWithClasses(c *fiber.Ctx) error {
	tradingCenters, err := getTradingCenters()
	if err != nil {
		return utils.ResError(c, 500, err)
	}

	allVeggies := types.AllVeggieWithClasses{
		TradingCenters: tradingCenters,
		Veggies:        []types.AllVeggies{},
	}

	for _, td := range tradingCenters {
		db := getDb(td.Slug)
		veggieColl := db.Collection(COLL_VEGGIES)
		veggieClassesColl := db.Collection(COLL_VEGGIES_CLASSES)

		veggies := []types.Veggie{}
		vegCursor, err := veggieColl.Find(context.TODO(), bson.M{})
		if err != nil {
			return utils.ResError(c, 500, err)
		}

		if err := vegCursor.All(context.TODO(), &veggies); err != nil {
			return utils.ResError(c, 500, err)
		}

		for _, v := range veggies {
			cursor, err := veggieClassesColl.Find(context.TODO(), bson.M{"parentId": v.ID})
			if err != nil {
				return utils.ResError(c, 500, err)
			}

			var veggieClasses []types.VeggieClass
			if err := cursor.All(context.TODO(), &veggieClasses); err != nil {
				return utils.ResError(c, 500, err)
			}

			allVeggies.Veggies = append(allVeggies.Veggies, types.AllVeggies{
				Data:    v,
				Classes: veggieClasses,
			})
		}
	}

	return utils.ResSuccess(c, 200, allVeggies)
}
