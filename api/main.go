package api

import (
	"context"
	"log"

	"github.com/Market-Monitor/veggies-api/api/database"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var mongoClient *mongo.Client = database.Connect()

func Start() {
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatalln(err)
		}
	}()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString(
			"Benguet Vegetables & Goods Price Monitoring API (Not Official)",
		)
	})

	api := app.Group("/api")
	api.Get("/history-prices", GetHistoryPrices)
	api.Get("/history-prices/latest", GetLatestHistoryPrices)
	api.Get("/prices", GetVeggiePrices)
	api.Get("/veggies", GetVeggies)
	api.Get("/veggies/classes", GetAllVeggieClasses)
	api.Get("/veggies/:id", GetVeggie)
	api.Get("/config", GetConfiguration)

	log.Fatal(app.Listen(":7000"))
}
