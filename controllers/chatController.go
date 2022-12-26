package controllers

import (
	"fmt"
	// "log"
	"sschneemelcher/artificialacademy/helpers"
	"sschneemelcher/artificialacademy/initializers"
	"sschneemelcher/artificialacademy/models"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Content string `json:"content" xml:"content" form:"content"`
}

func ChatIndex(c *fiber.Ctx) error {
	// get history from DB
	var history []models.Message
	result := initializers.DB.Find(&history)

	if result.Error != nil {
		return result.Error
	}

	return c.Render("chat/index", fiber.Map{
		"history": history,
	})
}

func ChatPost(c *fiber.Ctx) error {
	// Parse message text
	m := new(Message)
	if err := c.BodyParser(m); err != nil {
		return err
	}

	// log.Println(m.Content)

	// Save message content in DB
	message := models.Message{Content: m.Content, IsResponse: false}
	result := initializers.DB.Create(&message)

	if result.Error != nil {
		return result.Error
	}

	// Generate completion
	completion := helpers.GetCompletionDummy(m.Content)

	// Save completion as message in DB
	completionMessage := models.Message{Content: completion, IsResponse: true}
	completionResult := initializers.DB.Create(&completionMessage)

	if completionResult.Error != nil {
		return result.Error
	}

	return c.SendString(fmt.Sprintf("%s", completion))
}

func ChatClear(c *fiber.Ctx) error {
	// delete all messages
	// result := initializers.DB.Delete(&models.Message{}, )
	result := initializers.DB.Where("1 = 1").Delete(&models.Message{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
