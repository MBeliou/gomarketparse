package gomarketparse

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	parser := NewParser()
	start, _ := time.Parse(time.RFC822, "01 Jan 14 01:00 MST")
	records := parser.GetHistoricalData("bitcoin", start, time.Now())

	if len(records) < 1 {
		t.Errorf("Got an empty result")
	}
}
