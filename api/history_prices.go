package api

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetHistoryPrices(c fiber.Ctx) error {
	// Get HistoryPrices collection
	coll := mongoClient.Database(DATABASE).Collection(COLL_HISTORY_PRICES)

	// Get all documents
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Iterate through the cursor
	var historyPrices []map[string]any
	if err := cursor.All(context.TODO(), &historyPrices); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"historyPrices": historyPrices,
	})
}
