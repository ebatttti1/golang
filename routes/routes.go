package routes

import (
	controllers "main/controller"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	api := app.Group("/schedule")

	api.Post("/add", controllers.CreateConfig)
	api.Get("/select", controllers.GetConfigs)
	api.Post("/update", controllers.UpdateConfig)
	api.Delete("/delete", controllers.DeleteConfig)
	api.Get("/select/:id", controllers.GetConfig)
}