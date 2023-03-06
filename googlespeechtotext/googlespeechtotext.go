package googlespeechtotext

import (
	"context"
	"fmt"
	"os"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "cloud.google.com/go/speech/apiv1/speechpb"
)

const (
	sampleRateHertz = 16000
	languageCode    = "en-US"
)

func SpeechToText(fileName string) (*string, error) {
	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	// read in the flac file
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// Receive the transcript from the API
	// Send audio data to the Speech API
	req := &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:          speechpb.RecognitionConfig_FLAC,
			SampleRateHertz:   sampleRateHertz,
			LanguageCode:      languageCode,
			AudioChannelCount: 2,
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	}
	resp, err := client.Recognize(ctx, req)
	if err != nil {
		fmt.Printf("Failed to create streaming recognizer: %v", err)
		return nil, err
	}

	// Get the transcript from the response
	transcript := ""
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("Transcript: %v\n", alt.Transcript)
			transcript = alt.Transcript
		}
	}

	return &transcript, nil
}
