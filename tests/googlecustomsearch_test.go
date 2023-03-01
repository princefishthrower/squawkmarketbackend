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

	squawkModel, err := googlecustomsearch.CustomSearch("TSLA")
	if err != nil {
		log.Fatalf("Unable to execute search: %v", err)
	}

	if squawkModel == nil {
		t.Errorf("Squawks is nil")
		return
	}

	t.Logf("IN TEST: BEGIN SQUAWKS:\n\n%s", squawkModel.Squawk)

	squawkModel, err = googlecustomsearch.CustomSearch("financial breaking news")
	if err != nil {
		log.Fatalf("Unable to execute search: %v", err)
	}

	if squawkModel == nil {
		t.Errorf("Squawks is nil")
		return
	}

	t.Logf("IN TEST: BEGIN SQUAWK:\n\n%s", squawkModel.Squawk)
}
