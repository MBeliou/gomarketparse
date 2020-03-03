package gomarketparse

import (
	"fmt"
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
	fmt.Printf("Got date: %v\n", records[0].Date)
	if records[0].Date == "" {
		t.Errorf("Got an empty date")
	}
}
