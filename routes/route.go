package routes

import (
	"elisabethapi/controllers/bookingcontroller"
	"elisabethapi/controllers/doktercontroller"
	"elisabethapi/controllers/kamarcontroller"
	"elisabethapi/controllers/policontroller"
	"elisabethapi/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {
	dokter := r.Group("/dokter")
	poli := r.Group("/poli")
	booking := r.Group("/booking")
	kamar := r.Group("/kamar")

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

	booking.Get("/", middlewares.AuthMiddleware, bookingcontroller.Index)
	booking.Get("/:no_booking", middlewares.AuthMiddleware, bookingcontroller.Show)
	booking.Post("/", middlewares.AuthMiddleware, bookingcontroller.Create)
	booking.Put("/:no_booking", middlewares.AuthMiddleware, bookingcontroller.Update)
	booking.Delete("/:no_booking", middlewares.AuthMiddleware, bookingcontroller.Delete)

	kamar.Get("/", middlewares.AuthMiddleware, kamarcontroller.Index)
	kamar.Get("/:id", middlewares.AuthMiddleware, kamarcontroller.Show)
	kamar.Post("/", middlewares.AuthMiddleware, kamarcontroller.Create)
	kamar.Put("/:id", middlewares.AuthMiddleware, kamarcontroller.Update)
	kamar.Delete("/:id", middlewares.AuthMiddleware, kamarcontroller.Delete)
}