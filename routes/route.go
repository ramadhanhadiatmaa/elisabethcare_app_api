package routes

import (
	"elisabethapi/controllers/doktercontroller"
	"elisabethapi/controllers/policontroller"
	"elisabethapi/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {
	dokter := r.Group("/dokter")
	poli := r.Group("/poli")

	dokter.Get("/", middlewares.AuthMiddleware, doktercontroller.Index)
	dokter.Get("/:kode_dokter", middlewares.AuthMiddleware, doktercontroller.Show)
	dokter.Post("/", middlewares.AuthMiddleware, doktercontroller.Create)
	dokter.Put("/:kode_dokter", middlewares.AuthMiddleware, doktercontroller.Update)
	dokter.Delete("/:kode_dokter", middlewares.AuthMiddleware, doktercontroller.Delete)

	poli.Get("/", middlewares.AuthMiddleware, policontroller.Index)
	poli.Get("/:kode_poli", middlewares.AuthMiddleware, policontroller.Show)
	poli.Post("/", middlewares.AuthMiddleware, policontroller.Create)
	poli.Put("/:kode_poli", middlewares.AuthMiddleware, policontroller.Update)
	poli.Delete("/:kode_poli", middlewares.AuthMiddleware, policontroller.Delete)
}