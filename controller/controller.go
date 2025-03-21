package controllers

import (
	"log"
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
		log.Fatalf("error in parse body, error = %v \n", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if config.Name == "" {
		log.Fatalln("error: Name must be provided")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Name must be provided"})
	}

	if config.Command == "" {
		log.Fatalln("error: Command must be provided")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Command must be provided"})
	}

	if config.Limit == 0 {
		log.Fatalln("error: Limit must be greater than 0")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Limit must be greater than 0"})
	}

	if config.Interval == 0 {
		log.Fatalln("error: Interval must be greater than 0")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Interval must be greater than 0"})
	}

	config.Interval = int(time.Duration(config.Interval) * time.Second)
	log.Fatalf("time duration = %v \n", config.Interval)

	if err := database.DB.Create(&config).Error; err != nil {
		log.Fatalf("error in create Config, error = %v \n", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create Config"})
	}

	log.Fatalln("Config sent to the worker Config channel")
	worker.ConfigChannel <- config
	
	log.Fatalln("Config created successfully")
	log.Fatalln(config)
	return c.JSON(config)
}

func GetConfigs(c *fiber.Ctx) error {
	var configs []models.CommandLineConfig
	database.DB.Find(&configs)
	log.Fatalln("Configs fetched successfully")
	log.Fatalln(configs)
	return c.JSON(configs)
}

func GetConfig(c *fiber.Ctx) error {
	id := c.Params("id")
	var config models.CommandLineConfig
	if result := database.DB.First(&config, id); result.Error != nil {
		log.Fatalf("error in get Config, error = %v \n", result.Error)
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	log.Fatalf("Config %v is available \n", id)
	log.Fatalln(config)
	return c.JSON(config)
}

func UpdateConfig(c *fiber.Ctx) error {
	id := c.Params("id")
	var config models.CommandLineConfig
	if err := database.DB.First(&config, id).Error; err != nil {
		log.Fatalf("error in get Config, error = %v \n", err)
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Config not found"})
	}

	if err := c.BodyParser(&config); err != nil {
		log.Fatalf("error in parse body, error = %v \n", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if config.Limit > 0 {
		config.Limit--
	}

	config.Interval = 30
	if err := database.DB.Save(&config).Error; err != nil {
		log.Fatalf("error in update Config, error = %v \n", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update Config"})
	}

	log.Fatalf("updated Config %v sent to the channel\n", id)
	worker.ConfigChannel <- config
	
	log.Fatalf("Config %v updated successfully \n", id)
	return c.JSON(config)
}

func DeleteConfig(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.CommandLineConfig{}, id)
	log.Fatalf("Config %v deleted successfully \n", id)
	return c.JSON(fiber.Map{"message": "Config deleted deleted successfully"})
}