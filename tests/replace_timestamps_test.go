package tests

import (
	"squawkmarketbackend/utils"
	"testing"
)

func TestReplaceTimestamps(t *testing.T) {
	headline := "03:52AM HSBC Posts Higher Profit After Rise in Global Interest Rates"
	expected := "HSBC Posts Higher Profit After Rise in Global Interest Rates"
	actual := utils.ReplaceTimestamps(headline)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	headline = "  19:49PM Exclusive: Warburg Pincus raising $439 mln in maiden yuan fund for China deals -sources   "
	expected = "Exclusive: Warburg Pincus raising $439 mln in maiden yuan fund for China deals -sources"
	actual = utils.ReplaceTimestamps(headline)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
