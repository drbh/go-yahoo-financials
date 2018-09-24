package main

import (
// "fmt"
)

type ComputedRatios struct {
	ReturnOnAssets float64 `json:"ReturnOnAssets"`
	ReturnOnEquity float64 `json:"ReturnOnEquity"`
	ProfitMargin   float64 `json:"ProfitMargin"`
	QuickRatio     float64 `json:"QuickRatio"`
	CurrentRatio   float64 `json:"CurrentRatio"`
	DebtToEquity   float64 `json:"DebtToEquity"`
}

func (qss *QuoteSummaryStore) ReturnOnAssets() float64 {

	istatements := qss.IncomeStatementHistoryQuarterly.IncomeStatements
	bstatements := qss.BalanceSheetHistoryQuarterly.BalanceSheetStatements

	sumNetIncome := 0
	sumTotalAssets := 0

	for i := 0; i < len(istatements); i++ {
		sumNetIncome = sumNetIncome + istatements[i].NetIncome.Raw
	}

	for i := 0; i < len(bstatements); i++ {
		sumTotalAssets = sumTotalAssets + bstatements[i].TotalAssets.Raw
	}

	roa := float64(sumNetIncome) / float64(float64(sumTotalAssets)/float64(4))

	// fmt.Println("Return on assets", roa)
	return roa

	// fmt.Println(sumNetIncome)
	// fmt.Println(sumTotalAssets)
	// fmt.Println(sumTotalAssets + sumNetIncome)
}

func (qss *QuoteSummaryStore) ReturnOnEquity() float64 {

	istatements := qss.IncomeStatementHistoryQuarterly.IncomeStatements
	bstatements := qss.BalanceSheetHistoryQuarterly.BalanceSheetStatements

	sumNetIncome := 0
	sumTotalEquity := 0

	for i := 0; i < len(istatements); i++ {
		sumNetIncome = sumNetIncome + istatements[i].NetIncome.Raw
	}

	for i := 0; i < len(bstatements); i++ {
		sumTotalEquity = sumTotalEquity + bstatements[i].TotalStockholderEquity.Raw
	}

	roe := float64(sumNetIncome) / float64(float64(sumTotalEquity)/float64(4))

	// fmt.Println("Return on equity", roe)
	return roe
}

func (qss *QuoteSummaryStore) ProfitMargin() float64 {

	istatements := qss.IncomeStatementHistoryQuarterly.IncomeStatements

	sumNetIncome := 0
	sumRevenue := 0

	for i := 0; i < len(istatements); i++ {
		sumNetIncome = sumNetIncome + istatements[i].NetIncome.Raw
		sumRevenue = sumRevenue + istatements[i].TotalRevenue.Raw
	}

	pm := float64(sumNetIncome) / float64(sumRevenue)

	// fmt.Println("Profit Margin", pm)
	return pm
}

func (qss *QuoteSummaryStore) QuickRatio() float64 {
	bstatements := qss.BalanceSheetHistoryQuarterly.BalanceSheetStatements
	i := 0
	lastTotalCurrentLiabilities := float64(bstatements[i].TotalCurrentLiabilities.Raw)
	lastTotalInventory := float64(bstatements[i].Inventory.Raw)

	lastTotalCurrentAssets := float64(bstatements[i].TotalCurrentAssets.Raw)

	qr := (lastTotalCurrentAssets - lastTotalInventory) / lastTotalCurrentLiabilities
	// fmt.Println("Quick Ratio", qr)
	return qr
}

func (qss *QuoteSummaryStore) CurrentRatio() float64 {
	bstatements := qss.BalanceSheetHistoryQuarterly.BalanceSheetStatements
	i := 0
	lastTotalCurrentLiabilities := float64(bstatements[i].TotalCurrentLiabilities.Raw)
	lastTotalCurrentAssets := float64(bstatements[i].TotalCurrentAssets.Raw)
	cr := lastTotalCurrentAssets / lastTotalCurrentLiabilities
	// fmt.Println("Current Ratio", cr)
	return cr
}

func (qss *QuoteSummaryStore) DebtToEquity() float64 {
	bstatements := qss.BalanceSheetHistoryQuarterly.BalanceSheetStatements
	i := 0
	lastShortLongTermDebt := float64(bstatements[i].ShortLongTermDebt.Raw)
	lastLongTermDebt := float64(bstatements[i].LongTermDebt.Raw)

	lastTotalStockholderEquity := float64(bstatements[i].TotalStockholderEquity.Raw)

	totalDebt := lastShortLongTermDebt + lastLongTermDebt
	dte := totalDebt / lastTotalStockholderEquity

	// fmt.Println("Debt To Equity", dte)
	return dte
}
