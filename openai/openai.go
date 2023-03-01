package openai

import (
	"context"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func AskGPT(prompt string) (*string, error) {
	c := gogpt.NewClient(os.Getenv("OPEN_AI_SECRET_KEY"))
	ctx := context.Background()
	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci003,
		Temperature:      0.7,
		TopP:             1.0,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.0,
		BestOf:           1,
		MaxTokens:        1000,
		Prompt:           prompt,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return nil, err
	}
	return &resp.Choices[0].Text, nil
}
