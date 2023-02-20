package elevenlabs

import (
	"log"
	"os"
	"squawkmarketbackend/http_helper"
	"strings"
)

func GenerateMP3Data(text string) []byte {
	log.Println("getting voice data...")

	// headers for elevenlabs
	headers := map[string]string{
		"Content-Type": "application/json",
		"xi-api-key":   os.Getenv("ELEVEN_LABS_API_KEY"),
		"Accept":       "audio/mpeg",
	}

	body := strings.NewReader(`{
		"text": "` + text + `", 
		"voice_settings": {
			"stability": 0.15, 
			"similarity_boost": 1.0
		}
	}`)

	// elli voice id MF3mGyEYCl7XYWbV9V6O
	data, err := http_helper.MakeHTTPRequest(
		"https://api.elevenlabs.io/v1/text-to-speech/MF3mGyEYCl7XYWbV9V6O",
		"POST",
		headers,
		nil,
		body,
	)
	if err != nil {
		log.Println(err)
	}

	// return the byte data generated by elevenlabs' boss bots
	return data
}
