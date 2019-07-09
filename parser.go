package gomarketcap

import (
	"time"

	"github.com/gocolly/colly"
)

const BASE_URL = "https://coinmarketcap.com/currencies"

type Parser struct {
	*colly.Collector
}

func NewParser() Parser {
	return Parser{
		colly.NewCollector(),
	}
}

func (p *Parser) GetHistoricalData(currency string, from time.Time, to time.Time) Records {
	url := createURL(currency, from, to)
	records := Records{}

	p.OnHTML("#historical-data", func(e *colly.HTMLElement) {

		e.ForEach("tr", func(i int, el *colly.HTMLElement) {
			// Skip table headers
			if i != 0 {
				record := Record{
					Date:      el.ChildText("td:nth-of-type(1)"),
					Open:      el.ChildAttr("td:nth-of-type(2)", "data-format-value"),
					High:      el.ChildAttr("td:nth-of-type(3)", "data-format-value"),
					Low:       el.ChildAttr("td:nth-of-type(4)", "data-format-value"),
					Close:     el.ChildAttr("td:nth-of-type(5)", "data-format-value"),
					Volume:    el.ChildAttr("td:nth-of-type(6)", "data-format-value"),
					MarketCap: el.ChildAttr("td:nth-of-type(7)", "data-format-value"),
				}
				records = append(records, record)
			}
		})
	})

	p.Visit(url)

	return records
}
