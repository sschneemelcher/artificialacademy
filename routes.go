package main

import (
	"sschneemelcher/artificialacademy/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.ChatIndex)
	app.Post("/chat", controllers.ChatPost)
}