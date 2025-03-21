package main

import (
	"log"
	"main/database"
	"main/models"
	"main/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.CommandLineConfig{})
	app := fiber.New()

	routes.SetUpRoutes(app)

	err := app.Listen(":3030")
	if err != nil {
		log.Fatalf("error in run api server, error = %v \n", err)
	}
}
