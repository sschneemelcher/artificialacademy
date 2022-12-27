package controllers

import (
	"fmt"
	"net/http"
	"os"

	// "log"
	"sschneemelcher/artificialacademy/helpers"
	"sschneemelcher/artificialacademy/initializers"
	"sschneemelcher/artificialacademy/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ChatIndex(c *fiber.Ctx) error {
	// Get user from jwt
	user := c.Locals("user").(models.User)

	if &user == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "User not set"})
	}

	var lastChat models.UserChat
	result := initializers.DB.Last(&lastChat, "user_id = ?", user.ID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			chat := models.Chat{}
			initializers.DB.Create(&chat)
			userChat := models.UserChat{UserID: user.ID, ChatID: chat.ID}
			initializers.DB.Create(&userChat)
			lastChat = userChat
		} else {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to lookup latest chat"})
		}
	}

	// get history from DB
	var history []models.Message
	result = initializers.DB.Find(&history, "chat_id", lastChat.ChatID)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to lookup history"})
	}

	return c.Render("chat/index", fiber.Map{
		"companyName": os.Getenv("COMPANY_NAME"),
		"history":     history,
		"userId":      user.ID,
	})
}

func ChatPost(c *fiber.Ctx) error {
	// Parse message text

	type Body struct {
		Content string `json:"content" xml:"content" form:"content"`
	}

	body := new(Body)
	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse body"})
	}

	user := c.Locals("user").(models.User)

	var userChat models.UserChat
	result := initializers.DB.Last(&userChat, "user_id = ?", user.ID)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to lookup latest chat"})
	}

	var history []models.Message
	result = initializers.DB.Find(&history, "chat_id = ?", userChat.ChatID)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to lookup history"})
	}

	// log.Println(m.Content)

	// Save message content in DB

	if &user == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "User not set"})
	}

	message := models.Message{Content: body.Content, UserID: user.ID, ChatID: userChat.ChatID}
	result = initializers.DB.Create(&message)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to save message"})
	}

	prompt := ""
	for idx, val := range history {
		if idx > 0 {
			if val.UserID != user.ID {
				prompt += "StudyBot: "
			} else {
				prompt += "Student: "
			}
		}
		prompt += val.Content + "\n"
	}
	prompt += "Student: " + body.Content

	// Generate completion
	completion := helpers.GetCompletionDummy(prompt)

	// Save completion as message in DB
	completionMessage := models.Message{Content: completion, UserID: 0, ChatID: userChat.ChatID}
	result = initializers.DB.Create(&completionMessage)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to save response"})
	}

	return c.SendString(fmt.Sprintf("%s", completion))
}

func ChatClear(c *fiber.Ctx) error {
	// delete all messages

	// Get user from jwt
	user := c.Locals("user").(models.User)

	if &user == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "User not set"})
	}

	var userChat models.UserChat
	result := initializers.DB.Last(&userChat, "user_id = ?", user.ID)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to lookup latest chat"})
	}

	result = initializers.DB.Delete(&models.Message{}, "chat_id = ?", userChat.ChatID)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to delete messages"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{})
}
