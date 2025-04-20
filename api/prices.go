package api

import (
	"context"

	"github.com/Market-Monitor/veggies-api/api/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GetVeggiePrices returns all history prices of a veggie by its ID and Class
// @Summary Get veggie prices
// @Description Get all history prices of a veggie by its ID and Class
// @Tags history_prices
// @Accept json
// @Produce json
// @Param tradingCenter path string true "Trading Center"
// @Param id query string true "Veggie ID"
// @Param class query string true "Veggie Class"
// @Success 200 {object} utils.HTTPSuccessResponse
// @Failure 500 {object} utils.HTTPErrorResponse
// @Router /api/{tradingCenter}/veggie_prices [get]
func GetVeggiePrices(c *fiber.Ctx) error {
	veggieId := c.Query("id")
	veggieClass := c.Query("class")

	db := GetTD_DB(c)
	coll := db.Collection(COLL_HISTORY_PRICES)

	filter := bson.M{
		"parentId": veggieId,
		"id":       veggieClass,
	}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return utils.ResError(c, 500, err)
	}

	var veggiePrices []map[string]any
	if err := cursor.All(context.TODO(), &veggiePrices); err != nil {
		return utils.ResError(c, 500, err)
	}

	return utils.ResSuccess(c, 200, veggiePrices)
}
