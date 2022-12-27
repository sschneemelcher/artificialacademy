package main

import (
	"sschneemelcher/artificialacademy/controllers"
	"sschneemelcher/artificialacademy/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	// Chat routes
	app.Get("/", middleware.RequireAuth, controllers.ChatIndex)
	app.Post("/chat", middleware.RequireAuth, controllers.ChatPost)
	app.Delete("/chat", middleware.RequireAuth, controllers.ChatClear)

	// User routes
	app.Get("/login", controllers.UserIndex)
	app.Post("/signup", controllers.Signup)
	app.Post("/login", controllers.Login)
}
