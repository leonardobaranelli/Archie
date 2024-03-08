package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/leonardobaranelli/rest-api-backend-app/database"
	"github.com/leonardobaranelli/rest-api-backend-app/models"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}

	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

func GetFact(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	fact := models.Fact{}
	database.DB.Db.First(&fact, id)
	if fact.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Fact not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fact)
}

func UpdateFact(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	fact := models.Fact{}
	if err := c.BodyParser(&fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Model(&fact).Where("id = ?", id).Updates(&fact)

	return c.Status(fiber.StatusOK).JSON(fact)
}

func DeleteFact(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	fact := models.Fact{}
	database.DB.Db.First(&fact, id)
	if fact.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Fact not found",
		})
	}

	database.DB.Db.Delete(&fact)

	return c.SendStatus(fiber.StatusNoContent)
}
