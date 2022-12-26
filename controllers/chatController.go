package controllers

import (
	"fmt"
	"log"
	"sschneemelcher/artificialacademy/helpers"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Text string `json:"text" xml:"text" form:"text"`
}

func ChatIndex(c *fiber.Ctx) error {
	return c.Render("chat/index", fiber.Map{})
}

func ChatPost(c *fiber.Ctx) error {
	// Parse message text
	m := new(Message)
	if err := c.BodyParser(m); err != nil {
		return err
	}

	log.Println(m.Text)

	// Generate completion
	res := helpers.GetCompletion(m.Text)

	return c.SendString(fmt.Sprintf("%s", res))
}
