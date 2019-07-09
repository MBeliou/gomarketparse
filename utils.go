package gomarketcap

import (
	"fmt"
	"time"
)

func createFileName(currency string) string {
	return fmt.Sprintf("%s.csv", currency)
}

func createURL(currency string, start time.Time, end time.Time) string {
	return fmt.Sprintf("%s/%s/historical-data/?start=%s&end=%s", BASE_URL, currency, formatDate(start), formatDate(end))
}

func pad(val int) string {
	if val < 10 {
		return fmt.Sprintf("0%v", val)
	}
	return fmt.Sprintf("%v", val)
}

func formatDate(date time.Time) string {
	y := date.Year()
	m := date.Month()
	d := date.Day()
	return fmt.Sprintf("%v%s%s", y, pad(int(m)), pad(d))
}
