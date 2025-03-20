package routes

import (
	controllers "main/controller"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	api := app.Group("/schedule")

	api.Post("/add", controllers.CreateUser)
	api.Get("/select", controllers.GetUsers)
	api.Post("/update", controllers.UpdateUser)
	api.Delete("/delete", controllers.DeleteUser)
	api.Get("/select/:id", controllers.GetUser)
}