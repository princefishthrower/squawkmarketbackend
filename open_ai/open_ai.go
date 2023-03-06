package open_ai

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func AskGPT(prompt string) (*string, error) {
	c := openai.NewClient(os.Getenv("OPEN_AI_SECRET_KEY"))
	ctx := context.Background()
	req := openai.CompletionRequest{
		Model:            openai.GPT3TextDavinci003,
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

func SpeechToText(filePath string) (*string, error) {
	c := openai.NewClient(os.Getenv("OPEN_AI_SECRET_KEY"))
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: filePath,
	}
	resp, err := c.CreateTranscription(ctx, req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return nil, err
	}
	return &resp.Text, nil
}
