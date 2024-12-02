package api

import (
	"context"

	"github.com/Market-Monitor/veggies-api/api/types"
	"github.com/Market-Monitor/veggies-api/api/utils"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GetVeggies returns all veggies
func GetVeggies(c fiber.Ctx) error {
	coll := mongoClient.Database(DATABASE).Collection(COLL_VEGGIES)

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return utils.ResError(c, 500, err)
	}

	var veggies []map[string]any
	if err := cursor.All(context.TODO(), &veggies); err != nil {
		return utils.ResError(c, 500, err)
	}

	return utils.ResSuccess(c, 200, veggies)
}

// GetVeggie returns a veggie with its classes by its ID
func GetVeggie(c fiber.Ctx) error {
	veggieId := c.Params("id")

	// Begin getting the veggie
	veggieColl := mongoClient.Database(DATABASE).Collection(COLL_VEGGIES)

	filter := bson.M{
		"id": veggieId,
	}

	var veggie types.Veggie
	if err := veggieColl.FindOne(context.TODO(), filter).Decode(&veggie); err != nil {
		return utils.ResError(c, 500, err)
	}
	// End getting the veggie

	// Begin getting the veggie classes
	veggieClassesColl := mongoClient.Database(DATABASE).Collection(COLL_VEGGIES_CLASSES)

	cursor, err := veggieClassesColl.Find(context.TODO(), bson.M{"parentId": veggieId})
	if err != nil {
		return utils.ResError(c, 500, err)
	}

	var veggieClasses []types.VeggieClass
	if err := cursor.All(context.TODO(), &veggieClasses); err != nil {
		return utils.ResError(c, 500, err)
	}

	return utils.ResSuccess(c, 200, fiber.Map{
		"name":            veggie.Name,
		"id":              veggie.ID,
		"_id":             veggie.M_ID.Hex(),
		"classes":         veggieClasses,
		"tradingCenterId": veggie.TradingCenterId,
	})
	// End getting the veggie classes
}
