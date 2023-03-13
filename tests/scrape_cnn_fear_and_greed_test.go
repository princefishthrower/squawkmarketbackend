package tests

import (
	"fmt"
	fear_and_greed "squawkmarketbackend/fear_and_greed"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetCnnFearAndGreed(t *testing.T) {

	// load .env
	err := godotenv.Load("../.env")
	if err != nil {
		t.Errorf("Error loading .env file")
	}

	// here the config is the finviz sector config
	squawk, err := fear_and_greed.GetFearAndGreedSquawk()
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	if squawk == "" {
		t.Error("squawk is empty")
		return
	}

	fmt.Println("SQUAWK: ", squawk)
}
