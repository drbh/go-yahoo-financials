package main

// INCOME
type IncomeStatementHistory struct {
	IncomeStatements []IncomeStatement `json:"incomeStatementHistory"`
	MaxAge           int               `json:"maxAge"`
}

type IncomeStatementHistoryQuarterly struct {
	IncomeStatements []IncomeStatement `json:"incomeStatementHistory"`
	MaxAge           int               `json:"maxAge"`
}

type IncomeStatement struct {
	ResearchDevelopment               CellValue `json:"researchDevelopment"`
	EffectOfAccountingCharges         CellValue `json:"effectOfAccountingCharges"`
	IncomeBeforeTax                   CellValue `json:"incomeBeforeTax"`
	MinorityInterest                  CellValue `json:"minorityInterest"`
	NetIncome                         CellValue `json:"netIncome"`
	SellingGeneralAdministrative      CellValue `json:"sellingGeneralAdministrative"`
	GrossProfit                       CellValue `json:"grossProfit"`
	Ebit                              CellValue `json:"ebit"`
	EndDate                           CellValue `json:"endDate"`
	OperatingIncome                   CellValue `json:"operatingIncome"`
	OtherOperatingExpenses            CellValue `json:"otherOperatingExpenses"`
	InterestExpense                   CellValue `json:"interestExpense"`
	ExtraordinaryItems                CellValue `json:"extraordinaryItems"`
	NonRecurring                      CellValue `json:"nonRecurring"`
	OtherItems                        CellValue `json:"otherItems"`
	IncomeTaxExpense                  CellValue `json:"incomeTaxExpense"`
	TotalRevenue                      CellValue `json:"totalRevenue"`
	TotalOperatingExpenses            CellValue `json:"totalOperatingExpenses"`
	CostOfRevenue                     CellValue `json:"costOfRevenue"`
	TotalOtherIncomeExpenseNet        CellValue `json:"totalOtherIncomeExpenseNet"`
	DiscontinuedOperations            CellValue `json:"discontinuedOperations"`
	NetIncomeFromContinuingOps        CellValue `json:"netIncomeFromContinuingOps"`
	NetIncomeApplicableToCommonShares CellValue `json:"netIncomeApplicableToCommonShares"`
}
