package main

import (
	"sschneemelcher/artificialacademy/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	// Chat routes
	app.Get("/chat", controllers.ChatIndex)
	app.Post("/chat", controllers.ChatPost)
	app.Delete("/chat", controllers.ChatClear)

	// User routes
	app.Post("/signup", controllers.Signup)
	app.Post("/login", controllers.Login)
}
