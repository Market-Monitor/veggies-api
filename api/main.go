package api

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func Start() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString(
			"Hello, World!",
		)
	})

	log.Fatal(app.Listen(":7000"))
}
