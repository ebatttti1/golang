package main

import (
	"log"
	"main/database"
	"main/models"
	"main/routes"
	"main/worker"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDatabase()
	db := database.GetDB()

	//database.DB.AutoMigrate(&models.CommandLineConfig{})
	var tasks []models.CommandLineConfig
	db.Find(&tasks)
	for _, task := range tasks {
		go worker.ProcessConfig(db, task)
	}
	//select{}

	app := fiber.New()
	routes.SetUpRoutes(app)
	err := app.Listen(":3030")
	if err != nil {
		log.Fatalf("error in run api server, error = %v \n", err)
	}

	select{}
}