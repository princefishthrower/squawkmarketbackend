package tests

import (
	"squawkmarketbackend/utils"
	"testing"
)

func TestCleanHeadline(t *testing.T) {
	headline := "03:52AM HSBC Posts Higher Profit After Rise in Global Interest Rates"
	expected := "HSBC Posts Higher Profit After Rise in Global Interest Rates"
	actual := utils.CleanHeadline(headline)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	headline = "  19:49PM Exclusive: Warburg Pincus raising $439 mln in maiden yuan fund for China deals -sources   "
	expected = "Warburg Pincus raising $439 mln in maiden yuan fund for China deals"
	actual = utils.CleanHeadline(headline)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	headline = "04:35AM                Stock Futures Fall After Dow Posts Losses for Week"
	expected = "Stock Futures Fall After Dow Posts Losses for Week"
	actual = utils.CleanHeadline(headline)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	headline = "Exclusive: Warburg Pincus raising $439 mln in maiden yuan fund for China deals - sources"
	expected = "Warburg Pincus raising $439 mln in maiden yuan fund for China deals"
	actual = utils.CleanHeadline(headline)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	headline = "UPDATE 1-Europe's carbon price hits record high of 100 euros"
	expected = "Europe's carbon price hits record high of 100 euros"
	actual = utils.CleanHeadline(headline)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
