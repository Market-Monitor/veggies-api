package api

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/Market-Monitor/veggies-api/api/types"
	"github.com/gofiber/fiber/v2"
)

// GetTradingCenters returns all trading centers from the database
// @Summary Get trading centers
// @Description Get all trading centers
// @Tags trading_centers
// @Accept json
// @Produce json
// @Success 200 {object} utils.HTTPSuccessResponse
// @Failure 500 {object} utils.HTTPErrorResponse
// @Router /api/trading-centers [get]
func GetTradingCenters() ([]types.TradingCenter, error) {
	return getTradingCenters()
}

func getTradingCenters() ([]types.TradingCenter, error) {
	coll := mongoClient.Database(DATABASE).Collection(COLL_TRADING_CENTERS)

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var tradingCenters []types.TradingCenter

	if err := cursor.All(context.TODO(), &tradingCenters); err != nil {
		return nil, err
	}

	return tradingCenters, nil
}

func GetTD_DB(c *fiber.Ctx) *mongo.Database {
	tradingCenter := c.Locals("tradingCenter")

	tcStr, _ := tradingCenter.(string)
	return getDb(tcStr)
}

func getDb(dbName string) *mongo.Database {
	databaseName := strings.ToUpper(fmt.Sprintf("MM_%s", dbName))

	return mongoClient.Database(databaseName)
}
