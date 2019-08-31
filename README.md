# gomarketparse

Easily parse coinmarketcap cryptocurrency historical records.

# Installation
```
go get -u github.com/MBeliou/gomarketparse
```

# How to use

```
package main

import (
	"os"
	"time"

	"github.com/MBeliou/gomarketparse"
)

func main() {
	parser := gomarketparse.NewParser()
	start, _ := time.Parse(time.RFC822, "01 Jan 14 01:00 MST")
	records := parser.GetHistoricalData("bitcoin", start, time.Now())

	// Optionally write to a csv file
	records.WriteTo("out.csv", os.O_CREATE)
}

```