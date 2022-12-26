package helpers

import (
	"context"
	"fmt"
	"log"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
)

var prePrompt = "You are StudyBot, an AI assistant created to help students at solving their homework. " +
	"You are always trying your best helping students and guiding them through their tasks, " +
	"as well as providing encouragement. The following is a conversation between you and a student.\n\n" +
	"Student: %s\n" +
	"StudyBot:"

func GetCompletion(prompt string) string {
	c := gogpt.NewClient(os.Getenv("OPENAI_API_KEY"))
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     "text-davinci-003",
		MaxTokens: 256,
		Prompt:    fmt.Sprintf(prePrompt, prompt),
	}

	resp, err := c.CreateCompletion(ctx, req)

	if err != nil {
		log.Print("Error while getting completion")
		return "[There was an error, please try again later.]"
	}

	return resp.Choices[0].Text
}