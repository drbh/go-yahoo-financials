package main

import (
	"strconv"
)

// CASHFLOW

type CashflowStatementHistory struct {
	CashflowStatements []CashFlowStatement `json:"cashflowStatements"`
	MaxAge             int                 `json:"maxAge"`
}

type CashflowStatementHistoryQuarterly struct {
	CashflowStatements []CashFlowStatement `json:"cashflowStatements"`
	MaxAge             int                 `json:"maxAge"`
}

type CashFlowStatement struct {
	Investments                           CellValue `json:"investments"`
	ChangeToLiabilities                   CellValue `json:"changeToLiabilities"`
	TotalCashflowsFromInvestingActivities CellValue `json:"totalCashflowsFromInvestingActivities"`
	NetBorrowings                         CellValue `json:"netBorrowings"`
	TotalCashFromFinancingActivities      CellValue `json:"totalCashFromFinancingActivities"`
	ChangeToOperatingActivities           CellValue `json:"changeToOperatingActivities"`
	IssuanceOfStock                       CellValue `json:"issuanceOfStock"`
	NetIncome                             CellValue `json:"netIncome"`
	ChangeInCash                          CellValue `json:"changeInCash"`
	EndDate                               CellValue `json:"endDate"`
	RepurchaseOfStock                     CellValue `json:"repurchaseOfStock"`
	TotalCashFromOperatingActivities      CellValue `json:"totalCashFromOperatingActivities"`
	Depreciation                          CellValue `json:"depreciation"`
	OtherCashflowsFromInvestingActivities CellValue `json:"otherCashflowsFromInvestingActivities"`
	DividendsPaid                         CellValue `json:"dividendsPaid"`
	ChangeToInventory                     CellValue `json:"changeToInventory"`
	ChangeToAccountReceivables            CellValue `json:"changeToAccountReceivables"`
	OtherCashflowsFromFinancingActivities CellValue `json:"otherCashflowsFromFinancingActivities"`
	ChangeToNetincome                     CellValue `json:"changeToNetincome"`
	CapitalExpenditures                   CellValue `json:"capitalExpenditures"`
}

func (cs *CashFlowStatement) cols() string {
	return "Investments" + "," +
		"ChangeToLiabilities" + "," +
		"TotalCashflowsFromInvestingActivities" + "," +
		"NetBorrowings" + "," +
		"TotalCashFromFinancingActivities" + "," +
		"ChangeToOperatingActivities" + "," +
		"IssuanceOfStock" + "," +
		"NetIncome" + "," +
		"ChangeInCash" + "," +
		"EndDate" + "," +
		"RepurchaseOfStock" + "," +
		"TotalCashFromOperatingActivities" + "," +
		"Depreciation" + "," +
		"OtherCashflowsFromInvestingActivities" + "," +
		"DividendsPaid" + "," +
		"ChangeToInventory" + "," +
		"ChangeToAccountReceivables" + "," +
		"OtherCashflowsFromFinancingActivities" + "," +
		"ChangeToNetincome" + "," +
		"CapitalExpenditures"
}

func (cs *CashFlowStatement) toCSV() string {
	return strconv.Itoa(cs.Investments.Raw) + "," +
		strconv.Itoa(cs.ChangeToLiabilities.Raw) + "," +
		strconv.Itoa(cs.TotalCashflowsFromInvestingActivities.Raw) + "," +
		strconv.Itoa(cs.NetBorrowings.Raw) + "," +
		strconv.Itoa(cs.TotalCashFromFinancingActivities.Raw) + "," +
		strconv.Itoa(cs.ChangeToOperatingActivities.Raw) + "," +
		strconv.Itoa(cs.IssuanceOfStock.Raw) + "," +
		strconv.Itoa(cs.NetIncome.Raw) + "," +
		strconv.Itoa(cs.ChangeInCash.Raw) + "," +
		strconv.Itoa(cs.EndDate.Raw) + "," +
		strconv.Itoa(cs.RepurchaseOfStock.Raw) + "," +
		strconv.Itoa(cs.TotalCashFromOperatingActivities.Raw) + "," +
		strconv.Itoa(cs.Depreciation.Raw) + "," +
		strconv.Itoa(cs.OtherCashflowsFromInvestingActivities.Raw) + "," +
		strconv.Itoa(cs.DividendsPaid.Raw) + "," +
		strconv.Itoa(cs.ChangeToInventory.Raw) + "," +
		strconv.Itoa(cs.ChangeToAccountReceivables.Raw) + "," +
		strconv.Itoa(cs.OtherCashflowsFromFinancingActivities.Raw) + "," +
		strconv.Itoa(cs.ChangeToNetincome.Raw) + "," +
		strconv.Itoa(cs.CapitalExpenditures.Raw)

}
