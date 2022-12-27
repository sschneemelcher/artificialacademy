package middleware

import (
	"fmt"
	"os"
	"sschneemelcher/artificialacademy/initializers"
	"sschneemelcher/artificialacademy/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *fiber.Ctx) error {

	// Get cookie off req
	tokenString := c.Cookies("token")
	if tokenString == "" {
		return c.Redirect("/login")
	}

	// Validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return c.Redirect("/login")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check expiration
		if time.Now().Unix() > int64(claims["exp"].(float64)) {
			return c.Redirect("/login")
		}

		// Get user with token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			return c.Redirect("/login")
		}

		// Make user Object available to next routes TODO
		c.Locals("user", user)

		// Go to next middleware:
		return c.Next()

	} else {
		return c.Redirect("/login")
	}
}
