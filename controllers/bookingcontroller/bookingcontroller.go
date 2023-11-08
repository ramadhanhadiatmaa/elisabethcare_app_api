package bookingcontroller

import (
	"elisabethapi/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error{

	var booking []models.Booking
	models.DB.Find(&booking)

	return c.Status(fiber.StatusOK).JSON(booking)
}
func Show(c *fiber.Ctx) error{
	
	no_booking := c.Params("no_booking")

	var booking models.Booking

	if err := models.DB.Model(&booking).Where("no_booking = ?", no_booking).First(&booking).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Nomor pesanan tidak ditemukan",
			})
		}
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Server sedang mengalami gangguan",
			})
	}
	return c.JSON(booking)
}
func Create(c *fiber.Ctx) error{

	var booking models.Booking

	if err := c.BodyParser(&booking); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	booking.CreatedAt = time.Now()
	booking.UpdatedAt = time.Now()

	if err := models.DB.Create(&booking).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message" : err.Error(),
		})
	}

	return c.JSON(fiber.Map {
		"message": "Pesanan berhasil dibuat",
	})
}
func Update(c *fiber.Ctx) error{

	no_booking := c.Params("no_booking")

	var booking models.Booking

	if err := c.BodyParser(&booking); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map {
			"message": err.Error(),
		})
	}

	booking.UpdatedAt = time.Now()

	if models.DB.Where("no_booking = ?", no_booking).Updates(&booking).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map {
			"message": "Data pesanan tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map {
		"message": "Data berhasil diupdate",
	})
}
func Delete(c *fiber.Ctx) error{
	no_booking := c.Params("no_booking")
	var booking models.Booking

	if models.DB.Where("no_booking = ?", no_booking).Delete(&booking).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map {
			"message": "Data pesanan tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map {
		"message": "Berhasil menghapus data pesanan",
	})
}
