package main

import (
	"encoding/json"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type QuoteSummaryStore struct {

	// Earnings                          string `json:"earnings"`
	// Price                             string `json:"price"`

	IncomeStatementHistory            IncomeStatementHistory            `json:"incomeStatementHistory"`
	IncomeStatementHistoryQuarterly   IncomeStatementHistoryQuarterly   `json:"incomeStatementHistoryQuarterly"`
	BalanceSheetHistory               BalanceSheetHistory               `json:"balanceSheetHistory"`
	BalanceSheetHistoryQuarterly      BalanceSheetHistoryQuarterly      `json:"balanceSheetHistoryQuarterly"`
	CashflowStatementHistory          CashflowStatementHistory          `json:"cashflowStatementHistory"`
	CashflowStatementHistoryQuarterly CashflowStatementHistoryQuarterly `json:"cashflowStatementHistoryQuarterly"`
	Symbol                            string                            `json:"symbol"`
}

type CellValue struct {
	Raw     int    `json:"raw,omitempty"`
	Fmt     string `json:"fmt,omitempty"`
	LongFmt string `json:"longFmt,omitempty"`
}

func dumpToFile(fname string, fileContents string) {
	// print to file for now
	file, errr := os.Create(fname)
	if errr != nil {
		log.Fatal("Cannot create file", errr)
	}
	defer file.Close()

	fmt.Fprintf(file, fileContents)
}

func main() {
	client.OpenBoltDb()
	var rss string

	argsWithoutProg := os.Args[1:]
	// fmt.Println()

	sym := string(argsWithoutProg[0])

	rss = Read(sym)

	if len(rss) == 0 {
		fmt.Println("Fetched From Yahoo")
		// Fetch and save
		// rss = TestParse()

		resp, err := http.Get("https://finance.yahoo.com/quote/" + sym + "/financials?p=" + sym)
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		fmt.Println("Downloaded", sym)
		s := string(body)
		rs := rgx.FindStringSubmatch(s)
		strings := rs[1]
		rss = missingKey.ReplaceAllString(strings, ``)

		Write(sym, []byte(rss))
	} else {
		fmt.Println("Fetched From Local")
		fmt.Println("")
	}

	var yahooFin QuoteSummaryStore
	json.Unmarshal([]byte(rss), &yahooFin)
	dumpToFile("output.json", rss)

}
