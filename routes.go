package main

import (
	"sschneemelcher/artificialacademy/controllers"
	"sschneemelcher/artificialacademy/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	// Chat routes
	app.Get("/", middleware.RequireAuth, controllers.ChatIndex)
	app.Get("/chat/:chatid", middleware.RequireAuth, controllers.ChatByID)
	app.Post("/chat", middleware.RequireAuth, controllers.ChatPost)
	app.Post("/new", middleware.RequireAuth, controllers.ChatNew)
	app.Delete("/chat", middleware.RequireAuth, controllers.ChatClear)

	// User routes
	app.Get("/login", controllers.UserIndex)
	app.Post("/signup", controllers.Signup)
	app.Post("/login", controllers.Login)
	app.Delete("/logout", controllers.Logout)
}
