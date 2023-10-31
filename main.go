package main

import (
	"elisabethapi/models"
	"elisabethapi/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	models.ConnectionDatabase()

	app := fiber.New()

	app.Use(cors.New())

	routes.Route(app)

	app.Listen(":8081")
}