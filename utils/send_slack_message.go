package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"squawkmarketbackend/http_helper"
)

func SendSlackMessage(text string) {

	// if staging, prepend "STAGING: " to the message, otherwise just text
	environment := os.Getenv("ENVIRONMENT")
	data := map[string]string{"text": text}
	if environment == "staging" {
		data = map[string]string{"text": fmt.Sprintf("STAGING: %s", text)}
	}

	body, err := json.Marshal(data)
	if err != nil {
		return
	}

	http_helper.MakeHTTPRequest(os.Getenv("SLACK_PAYMENT_ACTIVITY_WEBHOOK_URL"), "POST", nil, nil, bytes.NewBuffer(body))
}
