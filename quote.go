package main

import (
	"encoding/json"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"github.com/andlabs/ui"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	// "reflect"
)

var window *ui.Window
var myLocation string

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
	Ratios                            ComputedRatios
}

type CellValue struct {
	Raw     int    `json:"raw,omitempty"`
	Fmt     string `json:"fmt,omitempty"`
	LongFmt string `json:"longFmt,omitempty"`
}

type Company struct {
	Symbol    string `json:"stock"`
	Sector    string `json:"sector"`
	SubSector string `json:"sub.sector"`
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

func (qss *QuoteSummaryStore) saveIS() {
	filecontent := ""
	head := qss.IncomeStatementHistoryQuarterly.IncomeStatements[0].cols() + "\n"
	filecontent = filecontent + head
	for i := 0; i < len(qss.IncomeStatementHistoryQuarterly.IncomeStatements); i++ {
		res := qss.IncomeStatementHistoryQuarterly.IncomeStatements[i]
		line := res.toCSV()
		line = line + "\n"
		filecontent = filecontent + line
	}
	dumpToFile(myLocation+"/data/"+qss.Symbol+"_"+"IncomeStatementHistoryQuarterly.csv", filecontent)
}

func (qss *QuoteSummaryStore) saveBS() {
	filecontent := ""
	head := qss.BalanceSheetHistoryQuarterly.BalanceSheetStatements[0].cols() + "\n"
	filecontent = filecontent + head
	for i := 0; i < len(qss.BalanceSheetHistoryQuarterly.BalanceSheetStatements); i++ {
		res := qss.BalanceSheetHistoryQuarterly.BalanceSheetStatements[i]
		line := res.toCSV()
		line = line + "\n"
		filecontent = filecontent + line
	}
	dumpToFile(myLocation+"/data/"+qss.Symbol+"_"+"BalanceSheetHistoryQuarterly.csv", filecontent)
}

func (qss *QuoteSummaryStore) saveCS() {
	filecontent := ""
	head := qss.CashflowStatementHistoryQuarterly.CashflowStatements[0].cols() + "\n"
	filecontent = filecontent + head
	for i := 0; i < len(qss.CashflowStatementHistoryQuarterly.CashflowStatements); i++ {
		res := qss.CashflowStatementHistoryQuarterly.CashflowStatements[i]
		line := res.toCSV()
		line = line + "\n"
		filecontent = filecontent + line
	}
	dumpToFile(myLocation+"/data/"+qss.Symbol+"_"+"CashflowStatementHistoryQuarterly.csv", filecontent)
}

func fetchLocalorYahoo(sym string) string {
	var rss string
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

		if len(rs) > 0 {
			strings := rs[1]
			rss = missingKey.ReplaceAllString(strings, ``)

			Write(sym, []byte(rss))
		} else {
			fmt.Println("no match in HTML")
		}

	} else {
		fmt.Println("Fetched From Local")
	}
	return rss
}
func GetFinancials(sym string) {
	rss := fetchLocalorYahoo(sym)
	// fmt.Println(len(rss))

	if len(rss) == 0 {
	} else {

		var yahooFin QuoteSummaryStore
		json.Unmarshal([]byte(rss), &yahooFin)

		if len(yahooFin.IncomeStatementHistoryQuarterly.IncomeStatements) > 0 {

			roa := yahooFin.ReturnOnAssets()
			roe := yahooFin.ReturnOnEquity()
			pm := yahooFin.ProfitMargin()

			qr := yahooFin.QuickRatio()
			cr := yahooFin.CurrentRatio()
			dte := yahooFin.DebtToEquity()

			fmt.Println("Compute Ratios!")
			yahooFin.Ratios = ComputedRatios{roa, roe, pm, qr, cr, dte}

			json, _ := json.Marshal(yahooFin.Ratios)

			WriteRatios(yahooFin.Symbol, json)

			yahooFin.saveIS()
			yahooFin.saveBS()
			yahooFin.saveCS()

			// yahooFin.saveRatios()

		} else {
			fmt.Println("SKIP", sym)
		}

	}

}
