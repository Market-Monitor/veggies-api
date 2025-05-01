package api

import (
	"context"
	"log"

	"github.com/Market-Monitor/veggies-api/api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"go.mongodb.org/mongo-driver/v2/mongo"

	_ "github.com/Market-Monitor/veggies-api/docs"
)

var mongoClient *mongo.Client = database.Connect()

func Start() {
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatalln(err)
		}
	}()

	app := fiber.New()

	// Initialize get trading centers
	tradingCenters, err := GetTradingCenters()
	if err != nil {
		log.Fatalln(err)
	}

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(
			"Benguet Vegetables & Goods Price Monitoring API (Not Official)",
		)
	})

	api := app.Group("/api/:tradingCenter", func(c *fiber.Ctx) error {
		tradingCenter := c.Params("tradingCenter")
		for _, center := range tradingCenters {
			if center.Slug == tradingCenter {
				c.Locals("tradingCenter", center.Slug)
				c.Locals("tradingCenterId", center.M_ID)
				c.Locals("tradingCenterName", center.Name)
				c.Locals("tradingCenterLongName", center.LongName)

				return c.Next()
			}
		}

		return c.Status(404).SendString("Trading Center not found")
	})

	api.Get("/", func(c *fiber.Ctx) error {
		tradingCenter := c.Locals("tradingCenter")
		tradingCenterName := c.Locals("tradingCenterName")

		return c.JSON(fiber.Map{
			"tradingCenter": tradingCenter,
			"name":          tradingCenterName,
			"longName":      c.Locals("tradingCenterLongName"),
		})
	})

	api.Get("/history-prices", GetHistoryPrices)
	api.Get("/history-prices/latest", GetLatestHistoryPrices)
	api.Get("/prices", GetVeggiePrices)
	api.Get("/veggies", GetVeggies)
	api.Get("/veggies/classes", GetAllVeggieClasses)
	api.Get("/veggies/:id", GetVeggie)
	api.Get("/config", GetConfiguration)

	log.Fatal(app.Listen(":7000"))
}
