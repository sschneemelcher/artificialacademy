package middleware

import "github.com/gofiber/fiber/v2"

func RequireAuth(c *fiber.Ctx) error {
	// TODO

	// Set a custom header on all responses:
	// Go to next middleware:
	return c.Next()
}
