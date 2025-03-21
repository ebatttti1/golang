package controllers

import (
	"main/database"
	"main/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateConfig(c *fiber.Ctx) error {
	config := new(models.CommandLineConfig)
	if err := c.BodyParser(config); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "wrong format"})
	}

	if config.Limit == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "limit can not be 0"})
	}

	if config.Interval == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "interval can not be 0"})
	}

	config.Interval = config.Interval * int(time.Hour)
	database.DB.Create(config)
	c.Status(http.StatusOK).JSON(fiber.Map{"message": "config created successfully"})
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "new user created successfully"})
}

func GetConfigs(c *fiber.Ctx) error {
	var configs []models.CommandLineConfig
	database.DB.Find(&configs)
	return c.JSON(configs)
}

func GetConfig(c *fiber.Ctx) error {
	id := c.Params("id")
	var config models.CommandLineConfig
	if result := database.DB.First(&config, id); result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	return c.JSON(config)
}

func UpdateConfig(c *fiber.Ctx) error {
	id := c.Params("id")
	var config models.CommandLineConfig
	if result := database.DB.First(&config, id); result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "congig not found"})
	}

	if err := c.BodyParser(&config); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "wrong format"})
	}

	if config.Limit == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "limited reached"})
	}

	if config.Limit > 0 {
		config.Limit--
	}

	database.DB.Save(&config)
	c.Status(http.StatusOK).JSON(fiber.Map{"message": "config updated successfully"})
	return c.JSON(config)
}

func DeleteConfig(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.CommandLineConfig{}, id)
	return c.JSON(fiber.Map{"message": "config deleted deleted successfully"})
}