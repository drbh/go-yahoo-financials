package main

import (
	"encoding/json"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

type TechnicalQuote struct {
	Rsi       float64 `json:"rsi"`
	Mfi       float64 `json:"mfi"`
	LastPrice float64 `json:"lastPrice"`
	Symbol    string  `json:"symbol"`
}

func GetTechnicalQuote(symbol string) {

	var phistory string
	phistory = ReadOHLC(symbol)
	// fmt.Println(phistory)

	if len(phistory) == 0 {
		fmt.Println("Fetched From Yahoo")

		// symbol := "AAPL"
		spy, _ := quote.NewQuoteFromYahoo(symbol, "2013-09-24", "2018-09-24", quote.Daily, true)

		// Save prices to CSV file
		// fmt.Print(spy.CSV())
		jsond, _ := json.Marshal(spy)
		WriteOHLC(symbol, jsond)

		sl := len(spy.Close)

		if sl < 20 {

		} else {
			rsi := talib.Rsi(spy.Close, 14)
			mfi := talib.Mfi(spy.High, spy.Low, spy.Close, spy.Volume, 10)
			lastPrice := spy.Close[sl-1]

			tq := TechnicalQuote{
				rsi[sl-1],
				mfi[sl-1],
				lastPrice,
				symbol,
			}

			jsond, _ := json.Marshal(tq)

			res := ReadTechnicals(symbol)

			if len(res) < 1 {
				WriteTechnicals(symbol, jsond)
			}

			// spew.Dump(tq)
		}

	} else {
		fmt.Println("Already Stored Locally")
	}

}
