package main

import (
	"elisabethapi/models"
	"elisabethapi/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectionDatabase()

	app := fiber.New()

	routes.Route(app)

	app.Listen(":8081")
}