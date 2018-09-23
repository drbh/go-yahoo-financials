package main

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
