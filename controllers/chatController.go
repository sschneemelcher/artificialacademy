package controllers

import (
	"fmt"
	"net/http"
	// "log"
	"sschneemelcher/artificialacademy/helpers"
	"sschneemelcher/artificialacademy/initializers"
	"sschneemelcher/artificialacademy/models"

	"github.com/gofiber/fiber/v2"
)

func ChatIndex(c *fiber.Ctx) error {
	// get history from DB
	var history []models.Message
	result := initializers.DB.Find(&history)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to lookup history"})
	}

	return c.Render("chat/index", fiber.Map{
		"history": history,
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

	var history []models.Message
	result := initializers.DB.Find(&history)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to lookup history"})
	}

	// log.Println(m.Content)

	// Save message content in DB
	message := models.Message{Content: body.Content, IsResponse: false}
	result = initializers.DB.Create(&message)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to save message"})
	}

	prompt := ""
	for idx, val := range history {
		if idx > 0 {
			if val.IsResponse {
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
	completionMessage := models.Message{Content: completion, IsResponse: true}
	result = initializers.DB.Create(&completionMessage)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to save response"})
	}

	return c.SendString(fmt.Sprintf("%s", completion))
}

func ChatClear(c *fiber.Ctx) error {
	// delete all messages
	// result := initializers.DB.Delete(&models.Message{}, )
	result := initializers.DB.Where("1 = 1").Delete(&models.Message{})

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to delete messages"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{})
}
