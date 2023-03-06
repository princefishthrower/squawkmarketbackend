package tests

import (
	"log"
	"squawkmarketbackend/videototext"
	"testing"

	"github.com/joho/godotenv"
)

func TestYoutubeToVideoIdToText(t *testing.T) {
	// load env so we have google credentials
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = videototext.YoutubeVideoIdToText("dp8PhLsUcFE")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("success")
}
