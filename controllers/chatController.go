package controllers

import (
	"fmt"
	"net/http"
	"os"

	"sschneemelcher/artificialacademy/helpers"
	"sschneemelcher/artificialacademy/initializers"
	"sschneemelcher/artificialacademy/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ChatNew(c *fiber.Ctx) error {
	// Get user from jwt
	user := c.Locals("user").(models.User)
	if &user == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "User not set"})
	}

	chat := models.Chat{Title: "Untitled"}
	initializers.DB.Create(&chat)
	userChat := models.UserChat{UserID: user.ID, ChatID: chat.ID}
	initializers.DB.Create(&userChat)

	return c.Redirect(fmt.Sprintf("/chat/%d", chat.ID))
}

func ChatByID(c *fiber.Ctx) error {
	// Get user from jwt
	user := c.Locals("user").(models.User)

	if &user == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "User not set"})
	}

	var selectedChatID = c.Params("chatid")
	var selectedChat models.UserChat
	result := initializers.DB.First(&selectedChat, "user_id = ? AND chat_id = ?", user.ID, selectedChatID)
	if result.Error != nil {

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to lookup chat with ID %s", selectedChatID),
		})
	}

	history, err := helpers.GetHistory(selectedChat)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Error looking up history"})
	}

	chats, err := helpers.GetChatList(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "No chats found"})
	}

	return c.Render("chat/index", fiber.Map{
		"companyName": os.Getenv("COMPANY_NAME"),
		"history":     history,
		"chatID":      selectedChat.ChatID,
		"userID":      user.ID,
		"userName":    user.Name,
		"chats":       chats,
	})

}

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
			chat := models.Chat{Title: "Untitled"}
			initializers.DB.Create(&chat)
			userChat := models.UserChat{UserID: user.ID, ChatID: chat.ID}
			initializers.DB.Create(&userChat)
			lastChat = userChat
		} else {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to lookup latest chat"})
		}
	}

	history, err := helpers.GetHistory(lastChat)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Error looking up history"})
	}

	chats, err := helpers.GetChatList(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "No chats found"})
	}

	return c.Render("chat/index", fiber.Map{
		"companyName": os.Getenv("COMPANY_NAME"),
		"history":     history,
		"chatID":      lastChat.ChatID,
		"userID":      user.ID,
		"userName":    user.Name,
		"chats":       chats,
	})
}

func ChatPost(c *fiber.Ctx) error {
	// Parse message text

	type Body struct {
		Content string `json:"content" xml:"content" form:"content"`
		ChatID  uint   `json:"chat_id" xml:"chat_id" form:"chat_id"`
	}

	body := new(Body)
	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse body"})
	}

	user := c.Locals("user").(models.User)

	var userChat models.UserChat
	result := initializers.DB.First(&userChat, "user_id = ? AND chat_id = ?", user.ID, body.ChatID)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to lookup latest chat"})
	}

	var history []models.Message
	result = initializers.DB.Find(&history, "chat_id = ?", userChat.ChatID)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to lookup history"})
	}

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
	completion := helpers.GetCompletion(prompt)

	// Save completion as message in DB
	completionMessage := models.Message{Content: completion, UserID: 0, ChatID: userChat.ChatID}
	result = initializers.DB.Create(&completionMessage)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to save response"})
	}

	return c.JSON(fiber.Map{"message": completion, "user": "StudyBot"})
}

func ChatClear(c *fiber.Ctx) error {
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

	// This only soft deletes the messages
	result = initializers.DB.Where("chat_id = ?", userChat.ChatID).Delete(&models.Message{})

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to delete messages"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{})
}
