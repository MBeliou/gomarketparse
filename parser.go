package gomarketparse

import (
	"strings"
	"time"

	"github.com/gocolly/colly"
)

const BASE_URL = "https://coinmarketcap.com/currencies"

type Parser struct {
	c *colly.Collector
}

func NewParser() Parser {
	return Parser{
		c: colly.NewCollector(),
	}
}

func (p *Parser) GetHistoricalData(currency string, from time.Time, to time.Time) Records {
	url := createURL(currency, from, to)
	records := Records{}

	p.c.OnHTML(".cmc-tab-historical-data", func(e *colly.HTMLElement) {

		e.ForEach("tr", func(i int, el *colly.HTMLElement) {
			// NOTE: Skipping table headers.
			if i != 0 {

				record := Record{
					Date:      el.ChildText("td:nth-of-type(1)"),
					Open:      strings.ReplaceAll(el.ChildText("td:nth-of-type(2)"), ",", ""),
					High:      strings.ReplaceAll(el.ChildText("td:nth-of-type(3)"), ",", ""),
					Low:       strings.ReplaceAll(el.ChildText("td:nth-of-type(4)"), ",", ""),
					Close:     strings.ReplaceAll(el.ChildText("td:nth-of-type(5)"), ",", ""),
					Volume:    strings.ReplaceAll(el.ChildText("td:nth-of-type(6)"), ",", ""),
					MarketCap: strings.ReplaceAll(el.ChildText("td:nth-of-type(7)"), ",", ""),
				}
				records = append(records, record)
			}
		})
	})

	p.c.Visit(url)

	return records
}
