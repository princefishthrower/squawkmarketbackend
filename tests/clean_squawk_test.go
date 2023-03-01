package tests

import (
	"squawkmarketbackend/utils"
	"testing"
)

func TestCleanSquawk(t *testing.T) {
	squawk := "03:52AM HSBC Posts Higher Profit After Rise in Global Interest Rates"
	expected := "HSBC Posts Higher Profit After Rise in Global Interest Rates"
	actual := utils.CleanSquawk(squawk)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	squawk = "  19:49PM Exclusive: Warburg Pincus raising $439 mln in maiden yuan fund for China deals -sources   "
	expected = "Warburg Pincus raising $439 mln in maiden yuan fund for China deals"
	actual = utils.CleanSquawk(squawk)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	squawk = "04:35AM                Stock Futures Fall After Dow Posts Losses for Week"
	expected = "Stock Futures Fall After Dow Posts Losses for Week"
	actual = utils.CleanSquawk(squawk)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	squawk = "Exclusive: Warburg Pincus raising $439 mln in maiden yuan fund for China deals - sources"
	expected = "Warburg Pincus raising $439 mln in maiden yuan fund for China deals"
	actual = utils.CleanSquawk(squawk)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	squawk = "UPDATE 1-Europe's carbon price hits record high of 100 euros"
	expected = "Europe's carbon price hits record high of 100 euros"
	actual = utils.CleanSquawk(squawk)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
