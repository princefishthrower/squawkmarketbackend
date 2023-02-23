package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"squawkmarketbackend/http_helper"
)

func SendSlackMessage(text string) {
	body, err := json.Marshal(map[string]string{"text": fmt.Sprintf("(%s) %s", os.Getenv("ENVIRONMENT"), text)})
	if err != nil {
		return
	}

	http_helper.MakeHTTPRequest(os.Getenv("SLACK_PAYMENT_ACTIVITY_WEBHOOK_URL"), "POST", nil, nil, bytes.NewBuffer(body))
}
