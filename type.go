package gomarketcap

import (
	"encoding/csv"
	"fmt"
	"os"
)

type mode int

type Record struct {
	Date      string
	Open      string
	High      string
	Low       string
	Close     string
	Volume    string
	MarketCap string
}

type Records []Record

func (r *Records) WriteTo(filename string, mode int) error {
	// mode is a flag for opening files. see: O_APPEND / O_CREATE
	if mode&os.O_APPEND != 0 {
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("Can't append to file %s - %s\n", filename, err)
			} else {
				// Convert for writing
				writer := csv.NewWriter(file)
				defer file.Close()
				defer writer.Flush()

				err = writer.WriteAll(r.toSlice())
				if err != nil {
					return fmt.Errorf("couldn't append to file %s - %s", filename, err)
				}
				fmt.Printf("Records appended to file %s\n", filename)
				return nil
			}
		}

	}

	if mode&os.O_CREATE != 0 {
		if file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0664); err == nil {
			writer := csv.NewWriter(file)
			defer file.Close()
			defer writer.Flush()

			err := writer.WriteAll(r.toSlice())
			if err != nil {
				return fmt.Errorf("couldn't create file %s - %s", filename, err)
			}
			fmt.Printf("Records saved to file %s\n", filename)
		}
	}
	return nil
}

func (r *Record) toSlice() []string {
	slice := []string{}
	slice = append(slice, r.Date, r.Open, r.High, r.Low, r.Close, r.Volume, r.MarketCap)
	return slice
}

func (r *Records) toSlice() [][]string {
	slice := [][]string{}
	for _, record := range *r {
		slice = append(slice, record.toSlice())
	}
	return slice
}
