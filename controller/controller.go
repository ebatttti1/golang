package controllers

import (
	"main/database"
	"main/models"
	"main/worker"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateConfig(c *fiber.Ctx) error {
	var config models.CommandLineConfig
	if err := c.BodyParser(&config); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if config.Name == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Name must be provided"})
	}

	if config.Command == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Command must be provided"})
	}

	if config.Limit == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Limit must be greater than 0"})
	}

	if config.Interval == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Interval must be greater than 0"})
	}

	config.Interval = int(time.Duration(config.Interval) * time.Second)

	if err := database.DB.Create(&config).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create config"})
	}

	worker.ConfigChannel <- config
	return c.JSON(config)
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
	if err := database.DB.First(&config, id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Command not found"})
	}

	if err := c.BodyParser(&config); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if config.Limit > 0 {
		config.Limit--
	}

	config.Interval = 30
	if err := database.DB.Save(&config).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update command"})
	}

	worker.ConfigChannel <- config
	return c.JSON(config)
}

func DeleteConfig(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.CommandLineConfig{}, id)
	return c.JSON(fiber.Map{"message": "config deleted deleted successfully"})
}