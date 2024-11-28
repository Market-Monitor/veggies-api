package api

import (
	"context"
	"fmt"
	"log"

	"github.com/Market-Monitor/veggies-api/api/database"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var mongoClient *mongo.Client = database.Connect()

func Start() {
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatalln(err)
		}
	}()

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := mongoClient.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString(
			"Hello, World!",
		)
	})

	api := app.Group("/api")
	api.Get("/history-prices", GetHistoryPrices)
	api.Get("/prices", GetVeggiePrices)

	log.Fatal(app.Listen(":7000"))
}
