package tests

import (
	"log"
	"os"
	"squawkmarketbackend/tdameritrade"
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
	conn, requestId, userPrincipals, err := tdameritrade.ConnectToTDAmeritradeWithExpressConnection(userPrincipalsString)
	if err != nil {
		t.Errorf("Error connecting to TD Ameritrade with express connection: %v", err)
	}
	defer conn.Close()

	// send request to listen to /ES - always open baby
	*requestId += 1
	err = tdameritrade.StreamSymbolQuotes("/ES", *requestId, *conn, *userPrincipals)
	if err != nil {
		t.Errorf("Error streaming symbol quotes: %v", err)
		return
	}

	// wait for a message
	_, message, err := conn.ReadMessage()
	if err != nil {
		t.Errorf("Error reading message: %v", err)
		return
	}

	if len(message) == 0 {
		t.Errorf("Message is empty")
		return
	}

	log.Println("message received:")
	log.Println(string(message))

}
