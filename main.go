package main

import (
	"os"
	"sschneemelcher/artificialacademy/initializers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	// Load templates
	engine := html.New("./views", ".html")

	// Setup app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Configure app
	app.Static("/", "./public")

	// Routes
	Routes(app)

	// Start app
	app.Listen(":" + os.Getenv("PORT"))
}
