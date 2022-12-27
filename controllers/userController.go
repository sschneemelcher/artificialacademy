package controllers

import (
	"log"
	"net/http"
	"os"
	"sschneemelcher/artificialacademy/initializers"
	"sschneemelcher/artificialacademy/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Body struct {
	Name     string `json:"name" xml:"name" form:"name"`
	Password string `json:"pass" xml:"pass" form:"pass"`
}

func UserIndex(c *fiber.Ctx) error {
	return c.Render("user/login", fiber.Map{
		"companyName": os.Getenv("COMPANY_NAME"),
	})
}

func Signup(c *fiber.Ctx) error {
	// Get name/pass off request body
	body := new(Body)
	if c.BodyParser(body) != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to read body"})
	}

	log.Println(body.Password)

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	user := models.User{Name: body.Name, Password: string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{})
}

func Login(c *fiber.Ctx) error {
	// Get name/pass off request body
	body := new(Body)
	if c.BodyParser(body) != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to read body",
		})
	}

	// Lookup requested user
	var user models.User
	initializers.DB.First(&user, "name = ?", body.Name)

	if user.ID == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Incorrect username or password",
		})
	}

	// Compare password with saved hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Incorrect username or password",
		})
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to create token",
		})
	}

	// Set cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.HTTPOnly = true
	c.Cookie(cookie)

	return c.Status(http.StatusOK).JSON(fiber.Map{})
}