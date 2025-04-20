package api

import (
	"context"

	"github.com/Market-Monitor/veggies-api/api/types"
	"github.com/Market-Monitor/veggies-api/api/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetConfiguration(c *fiber.Ctx) error {
	db := GetTD_DB(c)
	coll := db.Collection(COLL_CONFIGURATIONS)

	var config types.Configuration

	if err := coll.FindOne(context.TODO(), bson.M{
		"configId": 0,
	}).Decode(&config); err != nil {
		return utils.ResError(c, 500, err)
	}

	return utils.ResSuccess(c, 200, config)
}
