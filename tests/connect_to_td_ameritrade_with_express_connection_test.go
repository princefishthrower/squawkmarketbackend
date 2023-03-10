package tests

import (
	"log"
	"os"
	"squawkmarketbackend/jobs"
	"testing"

	"github.com/joho/godotenv"
)

func TestConnectToTdAmeritradeWithExpressConnection(t *testing.T) {
	// load in .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	userPrincipalsString := os.Getenv("TD_AMERITRADE_USER_PRINCIPALS")
	_, _, _, err = jobs.ConnectToTDAmeritradeWithExpressConnection(userPrincipalsString)
	if err != nil {
		t.Errorf("Error connecting to TD Ameritrade with express connection: %v", err)
	}
}
