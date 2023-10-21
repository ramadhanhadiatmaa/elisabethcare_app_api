package routes

import (
	"elisabethapi/controllers"
	"elisabethapi/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {
	dokter := r.Group("/dokter")

	dokter.Get("/", middlewares.AuthMiddleware, controllers.Index)
	dokter.Get("/:kode_dokter", middlewares.AuthMiddleware, controllers.Show)
	dokter.Post("/", middlewares.AuthMiddleware, controllers.Create)
	dokter.Put("/:kode_dokter", middlewares.AuthMiddleware, controllers.Update)
	dokter.Delete("/:kode_dokter", middlewares.AuthMiddleware, controllers.Delete)
}