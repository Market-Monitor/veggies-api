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

func GetTradingCenters() ([]types.TradingCenter, error) {
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
	fmt.Println(tradingCenter)
	databaseName := strings.ToUpper(fmt.Sprintf("MM_%s", tradingCenter))

	return mongoClient.Database(databaseName)
}
