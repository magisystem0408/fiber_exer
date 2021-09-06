package routes

import (
	"fiber_first/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register",controllers.Register)
}