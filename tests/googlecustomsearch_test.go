package tests

import (
	"log"
	"squawkmarketbackend/googlecustomsearch"
	"testing"

	"github.com/joho/godotenv"
)

func TestGoogleCustomSearch(t *testing.T) {

	// load environment for testing
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	squawk, err := googlecustomsearch.CustomSearch("TSLA")
	if err != nil {
		log.Fatalf("Unable to execute search: %v", err)
	}

	if squawk == nil {
		t.Errorf("Squawks is nil")
	}

	t.Logf("IN TEST: BEGIN SQUAWKS:\n\n%v", squawk)

	squawk, err = googlecustomsearch.CustomSearch("financial breaking news")
	if err != nil {
		log.Fatalf("Unable to execute search: %v", err)
	}

	if squawk == nil {
		t.Errorf("Squawks is nil")
	}

	t.Logf("IN TEST: BEGIN SQUAWK:\n\n%v", squawk)
}
