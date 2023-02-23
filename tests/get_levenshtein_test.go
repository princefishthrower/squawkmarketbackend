package tests

import (
	"testing"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

func TestGetLevenshtein(t *testing.T) {
	squawk1 := "Stock Futures Rise After Four-Day Selloff for S&P 500"
	squawk2 := "Stocks Open Higher After Four-Day Selloff for S&P 500"

	expected := 0.792453
	actual := strutil.Similarity(squawk1, squawk2, metrics.NewLevenshtein())
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	squawk1 = "Same"
	squawk2 = "Same"

	expected = 1.000000
	actual = strutil.Similarity(squawk1, squawk2, metrics.NewLevenshtein())
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	squawk1 = "ABCD"
	squawk2 = "ZYXW"

	expected = 0.000000
	actual = strutil.Similarity(squawk1, squawk2, metrics.NewLevenshtein())
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	squawk1 = "Stocks gain after back-to-back losses for S&P 500, Dow"
	squawk2 = "Stock Futures Rise After Four-Day Selloff for S&P 500"
	expected = 0.388889
	actual = strutil.Similarity(squawk1, squawk2, metrics.NewLevenshtein())
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

	squawk1 = "New top loser: Bonso Electronics International Inc., symbol BNSO, down 31.55%."
	squawk2 = "New top loser: Bonso Electronics International Inc., symbol BNSO, down 29.12%."
	expected = 0.948718
	actual = strutil.Similarity(squawk1, squawk2, metrics.NewLevenshtein())
	if actual != expected {
		t.Errorf("Expected %f, got %f", expected, actual)
	}

}
