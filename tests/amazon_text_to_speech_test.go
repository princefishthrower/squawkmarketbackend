package tests

import (
	"log"
	"squawkmarketbackend/amazontexttospeech"
	"testing"

	"github.com/joho/godotenv"
)

func TestAmazonTextToSpeech(t *testing.T) {
	// read env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mp3data, err := amazontexttospeech.TextToSpeech("hello world")
	if err != nil {
		t.Error(err)
		return
	}

	if mp3data == nil {
		t.Error("mp3 data is nil")
	}

}
