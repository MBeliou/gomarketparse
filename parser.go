package gomarketparse

import (
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
			// Skip table headers
			if i != 0 {
				record := Record{
					Date:      el.ChildText("td:nth-of-type(1)"),
					Open:      el.ChildText("td:nth-of-type(2)"),
					High:      el.ChildText("td:nth-of-type(3)"),
					Low:       el.ChildText("td:nth-of-type(4)"),
					Close:     el.ChildText("td:nth-of-type(5)"),
					Volume:    el.ChildText("td:nth-of-type(6)"),
					MarketCap: el.ChildText("td:nth-of-type(7)"),
				}
				records = append(records, record)
			}
		})
	})

	p.c.Visit(url)

	return records
}
