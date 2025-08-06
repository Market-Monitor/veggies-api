package api

import (
	"context"
	"fmt"

	"github.com/Market-Monitor/veggies-api/api/types"
	"github.com/Market-Monitor/veggies-api/api/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GetVeggies returns all veggies
// @Summary Get all veggies
// @Description Get all veggies
// @Tags veggies
// @Accept json
// @Produce json
// @Param tradingCenter path string true "Trading Center"
// @Success 200 {object} utils.HTTPSuccessResponse
// @Failure 500 {object} utils.HTTPErrorResponse
// @Router /api/{tradingCenter}/veggies [get]
func GetVeggies(c *fiber.Ctx) error {
	db := GetTD_DB(c)
	coll := db.Collection(COLL_VEGGIES)

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

// GetVeggieByClassWithPrices returns a veggie with its classes and prices by its ID and Class ID
// @Summary Get veggie by class with prices
// @Description Get veggie by class with prices
// @Tags veggies
// @Accept json
// @Produce json
// @Param tradingCenter path string true "Trading Center"
// @Param id path string true "Veggie ID"
// @Param classId path string true "Veggie Class ID"
// @Success 200 {object} utils.HTTPSuccessResponse
// @Failure 500 {object} utils.HTTPErrorResponse
// @Router /api/{tradingCenter}/veggies/{id}/{classId} [get]
func GetVeggieByClassWithPrices(c *fiber.Ctx) error {
	veggieId := c.Params("id")
	classId := c.Params("classId")

	// Begin getting the veggie
	db := GetTD_DB(c)
	veggieColl := db.Collection(COLL_VEGGIES)

	filter := bson.M{
		"id": veggieId,
	}

	var veggie types.Veggie
	if err := veggieColl.FindOne(context.TODO(), filter).Decode(&veggie); err != nil {
		return utils.ResError(c, 500, err)
	}
	// End getting the veggie

	// Begin getting the veggie classes
	veggieClassesColl := db.Collection(COLL_VEGGIES_CLASSES)

	cursor, err := veggieClassesColl.Find(context.TODO(), bson.M{"parentId": veggieId})
	if err != nil {
		return utils.ResError(c, 500, err)
	}

	var veggieClasses []types.VeggieClass
	if err := cursor.All(context.TODO(), &veggieClasses); err != nil {
		return utils.ResError(c, 500, err)
	}

	// Check if the classId exists in the veggie classes
	var veggieClass *types.VeggieClass
	for _, class := range veggieClasses {
		if class.ID == classId {
			veggieClass = &class
			break
		}
	}
	if veggieClass == nil {
		return utils.ResError(c, 404, fmt.Errorf("Veggie class with ID %s not found", classId))
	}

	// End getting the veggie classes

	// Get veggie class prices
	var veggiePrices []types.VeggiePrice
	veggiePricesColl := db.Collection(COLL_HISTORY_PRICES)
	cursor, err = veggiePricesColl.Find(context.TODO(), bson.M{
		"parentId": veggieId,
		"id":       veggieClass.ID,
	})
	if err != nil {
		return utils.ResError(c, 500, err)
	}
	if err := cursor.All(context.TODO(), &veggiePrices); err != nil {
		return utils.ResError(c, 500, err)
	}

	return utils.ResSuccess(c, 200, fiber.Map{
		"veggie":      veggie,
		"veggieClass": veggieClass,
		"classes":     veggieClasses,
		"prices":      veggiePrices,
	})
}

// GetVeggie returns a veggie with its classes by its ID
// @Summary Get veggie by ID
// @Description Get veggie by ID
// @Tags veggies
// @Accept json
// @Produce json
// @Param tradingCenter path string true "Trading Center"
// @Param id path string true "Veggie ID"
// @Success 200 {object} utils.HTTPSuccessResponse
// @Failure 500 {object} utils.HTTPErrorResponse
// @Router /api/{tradingCenter}/veggies/{id} [get]
func GetVeggie(c *fiber.Ctx) error {
	veggieId := c.Params("id")

	// Begin getting the veggie
	db := GetTD_DB(c)
	veggieColl := db.Collection(COLL_VEGGIES)

	filter := bson.M{
		"id": veggieId,
	}

	var veggie types.Veggie
	if err := veggieColl.FindOne(context.TODO(), filter).Decode(&veggie); err != nil {
		return utils.ResError(c, 500, err)
	}
	// End getting the veggie

	// Begin getting the veggie classes
	veggieClassesColl := db.Collection(COLL_VEGGIES_CLASSES)

	cursor, err := veggieClassesColl.Find(context.TODO(), bson.M{"parentId": veggieId})
	if err != nil {
		return utils.ResError(c, 500, err)
	}

	var veggieClasses []types.VeggieClass
	if err := cursor.All(context.TODO(), &veggieClasses); err != nil {
		return utils.ResError(c, 500, err)
	}

	return utils.ResSuccess(c, 200, fiber.Map{
		"name":          veggie.Name,
		"id":            veggie.ID,
		"_id":           veggie.M_ID.Hex(),
		"classes":       veggieClasses,
		"tradingCenter": veggie.TradingCenter,
		"imageUrl":      veggie.ImageURL,
		"imageSource":   veggie.ImageSource,
		"priceUnit":     veggie.PriceUnit,
	})
	// End getting the veggie classes
}

// GetAllVeggieClasses returns all veggie classes
// @Summary Get all veggie classes
// @Description Get all veggie classes
// @Tags veggies
// @Accept json
// @Produce json
// @Param tradingCenter path string true "Trading Center"
// @Success 200 {object} utils.HTTPSuccessResponse
// @Failure 500 {object} utils.HTTPErrorResponse
// @Router /api/{tradingCenter}/veggies/classes [get]
func GetAllVeggieClasses(c *fiber.Ctx) error {
	db := GetTD_DB(c)
	veggieClassesColl := db.Collection(COLL_VEGGIES_CLASSES)

	cursor, err := veggieClassesColl.Find(context.TODO(), bson.D{})
	if err != nil {
		return utils.ResError(c, 500, err)
	}

	var veggieClasses []types.VeggieClass
	if err := cursor.All(context.TODO(), &veggieClasses); err != nil {
		return utils.ResError(c, 500, err)
	}

	return utils.ResSuccess(c, 200, veggieClasses)
}
