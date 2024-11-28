package api

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetVeggiePrices(c fiber.Ctx) error {
	veggieId := c.Query("id")
	veggieClass := c.Query("class")

	// Get VeggiePrices collection
	coll := mongoClient.Database(DATABASE).Collection(COLL_HISTORY_PRICES)

	// Search
	filter := bson.M{
		"parentId": veggieId,
		"id":       veggieClass,
	}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Iterate through the cursor
	var veggiePrices []map[string]any
	if err := cursor.All(context.TODO(), &veggiePrices); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"veggiePrices": veggiePrices,
	})
}
