package controllers

import (
	"elisabethapi/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index (c *fiber.Ctx) error {

	var dokter []models.Dokter
	models.DB.Find(&dokter)

	return c.Status(fiber.StatusOK).JSON(dokter)

}

func Show (c *fiber.Ctx) error {

	kode_dokter := c.Params("kode_dokter")

	var dokter models.Dokter

	if err := models.DB.Model(&dokter).Where("kode_dokter = ?", kode_dokter).First(&dokter).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Dokter tidak ditemukan",

			})
		}

			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Server sedang mengalami gangguan",
			})
		}
	
	return c.JSON(dokter)

}

func Create (c *fiber.Ctx) error {

	var dokter models.Dokter
	if err := c.BodyParser(&dokter); err!= nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	dokter.Tentang = "Lorem ipsum dolor sit amet consectetur. Amet urna malesuada curabitur ac eleifend et vitae. Eget faucibus eget lacus egestas. Tellus maecenas at sagittis nec. At nibh dictum in."

	if err := models.DB.Create(&dokter).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	
	return c.JSON(fiber.Map{
		"message": "Data dokter berhasil dimasukkan",
	})
}

func Update (c *fiber.Ctx) error {

	kode_dokter := c.Params("kode_dokter")

	var dokter models.Dokter

	if err := c.BodyParser(&dokter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("kode_dokter = ?", kode_dokter).Updates(&dokter).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Kode dokter dengan nomor " + kode_dokter + " tidak ditemukan.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate.",
	})

}

func Delete (c *fiber.Ctx) error {

	kode_dokter := c.Params("kode_dokter")
	var dokter models.Dokter

	if models.DB.Where("kode_dokter = ?", kode_dokter).Delete(&dokter).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Dokter dengan kode " + kode_dokter + " tidak ditemukan.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil menghapus dokter dengan kode " + kode_dokter + ".",
	})

}