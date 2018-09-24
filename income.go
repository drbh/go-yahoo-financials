package main

import (
	// "fmt"
	"strconv"
)

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

func (is *IncomeStatement) cols() string {
	return "ResearchDevelopment" + "," +
		"EffectOfAccountingCharges" + "," +
		"IncomeBeforeTax" + "," +
		"MinorityInterest" + "," +
		"NetIncome" + "," +
		"SellingGeneralAdministrative" + "," +
		"GrossProfit" + "," +
		"Ebit" + "," +
		"EndDate" + "," +
		"OperatingIncome" + "," +
		"OtherOperatingExpenses" + "," +
		"InterestExpense" + "," +
		"ExtraordinaryItems" + "," +
		"NonRecurring" + "," +
		"OtherItems" + "," +
		"IncomeTaxExpense" + "," +
		"TotalRevenue" + "," +
		"TotalOperatingExpenses" + "," +
		"CostOfRevenue" + "," +
		"TotalOtherIncomeExpenseNet" + "," +
		"DiscontinuedOperations" + "," +
		"NetIncomeFromContinuingOps" + "," +
		"NetIncomeApplicableToCommonShares"
}

func (is *IncomeStatement) toCSV() string {
	return strconv.Itoa(is.ResearchDevelopment.Raw) + "," +
		strconv.Itoa(is.EffectOfAccountingCharges.Raw) + "," +
		strconv.Itoa(is.IncomeBeforeTax.Raw) + "," +
		strconv.Itoa(is.MinorityInterest.Raw) + "," +
		strconv.Itoa(is.NetIncome.Raw) + "," +
		strconv.Itoa(is.SellingGeneralAdministrative.Raw) + "," +
		strconv.Itoa(is.GrossProfit.Raw) + "," +
		strconv.Itoa(is.Ebit.Raw) + "," +
		strconv.Itoa(is.EndDate.Raw) + "," +
		strconv.Itoa(is.OperatingIncome.Raw) + "," +
		strconv.Itoa(is.OtherOperatingExpenses.Raw) + "," +
		strconv.Itoa(is.InterestExpense.Raw) + "," +
		strconv.Itoa(is.ExtraordinaryItems.Raw) + "," +
		strconv.Itoa(is.NonRecurring.Raw) + "," +
		strconv.Itoa(is.OtherItems.Raw) + "," +
		strconv.Itoa(is.IncomeTaxExpense.Raw) + "," +
		strconv.Itoa(is.TotalRevenue.Raw) + "," +
		strconv.Itoa(is.TotalOperatingExpenses.Raw) + "," +
		strconv.Itoa(is.CostOfRevenue.Raw) + "," +
		strconv.Itoa(is.TotalOtherIncomeExpenseNet.Raw) + "," +
		strconv.Itoa(is.DiscontinuedOperations.Raw) + "," +
		strconv.Itoa(is.NetIncomeFromContinuingOps.Raw) + "," +
		strconv.Itoa(is.NetIncomeApplicableToCommonShares.Raw)
}
