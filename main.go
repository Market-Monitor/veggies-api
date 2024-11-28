package main

import (
	"github.com/Market-Monitor/veggies-api/api"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	api.Start()
}
