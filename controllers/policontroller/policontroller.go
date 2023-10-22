package policontroller

import (
	"elisabethapi/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index (c *fiber.Ctx) error{

	var poli []models.Poli
	models.DB.Find(&poli)

	return c.Status(fiber.StatusOK).JSON(poli)
}

func Show (c *fiber.Ctx) error{

	kode_poli := c.Params("kode_poli");

	var poli models.Poli

	if err := models.DB.Model(&poli).Where("kode_poli = ?", kode_poli).First(&poli).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Poli tidak ditemukan.",
			})
		}

			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Server sedang mengalami gangguan.",
			})
	}

	return c.JSON(poli)
}

func Create (c *fiber.Ctx) error{

	var poli models.Poli
	if err := c.BodyParser(&poli); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&poli).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data poli berhasil dimasukkan",
	})
}

func Update (c *fiber.Ctx) error{

	kode_poli := c.Params("kode_poli")

	var poli models.Poli

	if err := c.BodyParser(&poli); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("kode_poli = ?", kode_poli).Updates(&poli).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Kode poli " + kode_poli + " tidak ditemukan.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate.",
	})
}

func Delete (c *fiber.Ctx) error{

	kode_poli := c.Params("kode_poli")
	var poli models.Poli

	if models.DB.Where("kode_poli = ?", kode_poli).Delete(&poli).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Poli dengan kode " + kode_poli + " tidak dapat dihapus.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil menghapus poli dengan kode " + kode_poli + ".",
	})
}