package main

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/Market-Monitor/veggies-api/api"
)

func main() {
	api.Start()
}
