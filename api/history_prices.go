package api

import (
	"context"

	"github.com/Market-Monitor/veggies-api/api/utils"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GetHistoryPrices returns all history prices of all veggies
// TODO: Sort to get only the prices on the day of the request
func GetHistoryPrices(c fiber.Ctx) error {
	coll := mongoClient.Database(DATABASE).Collection(COLL_HISTORY_PRICES)

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return utils.ResError(c, 500, err)
	}

	var historyPrices []map[string]any
	if err := cursor.All(context.TODO(), &historyPrices); err != nil {
		return utils.ResError(c, 500, err)
	}

	return utils.ResSuccess(c, 200, historyPrices)
}
