package main

import (
	"strconv"
)

// BALANCE

type BalanceSheetHistory struct {
	BalanceSheetStatements []BalanceStatement `json:"balanceSheetStatements"`
	MaxAge                 int                `json:"maxAge"`
}

type BalanceSheetHistoryQuarterly struct {
	BalanceSheetStatements []BalanceStatement `json:"balanceSheetStatements"`
	MaxAge                 int                `json:"maxAge"`
}
type BalanceStatement struct {
	TotalLiab               CellValue `json:"totalLiab"`
	TotalStockholderEquity  CellValue `json:"totalStockholderEquity"`
	OtherCurrentLiab        CellValue `json:"otherCurrentLiab"`
	TotalAssets             CellValue `json:"totalAssets"`
	EndDate                 CellValue `json:"endDate"`
	CommonStock             CellValue `json:"commonStock"`
	OtherCurrentAssets      CellValue `json:"otherCurrentAssets"`
	RetainedEarnings        CellValue `json:"retainedEarnings"`
	OtherLiab               CellValue `json:"otherLiab"`
	TreasuryStock           CellValue `json:"treasuryStock"`
	OtherAssets             CellValue `json:"otherAssets"`
	Cash                    CellValue `json:"cash"`
	TotalCurrentLiabilities CellValue `json:"totalCurrentLiabilities"`
	ShortLongTermDebt       CellValue `json:"shortLongTermDebt"`
	OtherStockholderEquity  CellValue `json:"otherStockholderEquity"`
	PropertyPlantEquipment  CellValue `json:"propertyPlantEquipment"`
	TotalCurrentAssets      CellValue `json:"totalCurrentAssets"`
	LongTermInvestments     CellValue `json:"longTermInvestments"`
	NetTangibleAssets       CellValue `json:"netTangibleAssets"`
	ShortTermInvestments    CellValue `json:"shortTermInvestments"`
	NetReceivables          CellValue `json:"netReceivables"`
	LongTermDebt            CellValue `json:"longTermDebt"`
	Inventory               CellValue `json:"inventory"`
	AccountsPayable         CellValue `json:"accountsPayable"`
}

func (bs *BalanceStatement) cols() string {
	return "TotalLiab" + "," +
		"TotalStockholderEquity" + "," +
		"OtherCurrentLiab" + "," +
		"TotalAssets" + "," +
		"EndDate" + "," +
		"CommonStock" + "," +
		"OtherCurrentAssets" + "," +
		"RetainedEarnings" + "," +
		"OtherLiab" + "," +
		"TreasuryStock" + "," +
		"OtherAssets" + "," +
		"Cash" + "," +
		"TotalCurrentLiabilities" + "," +
		"ShortLongTermDebt" + "," +
		"OtherStockholderEquity" + "," +
		"PropertyPlantEquipment" + "," +
		"TotalCurrentAssets" + "," +
		"LongTermInvestments" + "," +
		"NetTangibleAssets" + "," +
		"ShortTermInvestments" + "," +
		"NetReceivables" + "," +
		"LongTermDebt" + "," +
		"Inventory" + "," +
		"AccountsPayable"
}

func (bs *BalanceStatement) toCSV() string {
	return strconv.Itoa(bs.TotalLiab.Raw) + "," +
		strconv.Itoa(bs.TotalStockholderEquity.Raw) + "," +
		strconv.Itoa(bs.OtherCurrentLiab.Raw) + "," +
		strconv.Itoa(bs.TotalAssets.Raw) + "," +
		strconv.Itoa(bs.EndDate.Raw) + "," +
		strconv.Itoa(bs.CommonStock.Raw) + "," +
		strconv.Itoa(bs.OtherCurrentAssets.Raw) + "," +
		strconv.Itoa(bs.RetainedEarnings.Raw) + "," +
		strconv.Itoa(bs.OtherLiab.Raw) + "," +
		strconv.Itoa(bs.TreasuryStock.Raw) + "," +
		strconv.Itoa(bs.OtherAssets.Raw) + "," +
		strconv.Itoa(bs.Cash.Raw) + "," +
		strconv.Itoa(bs.TotalCurrentLiabilities.Raw) + "," +
		strconv.Itoa(bs.ShortLongTermDebt.Raw) + "," +
		strconv.Itoa(bs.OtherStockholderEquity.Raw) + "," +
		strconv.Itoa(bs.PropertyPlantEquipment.Raw) + "," +
		strconv.Itoa(bs.TotalCurrentAssets.Raw) + "," +
		strconv.Itoa(bs.LongTermInvestments.Raw) + "," +
		strconv.Itoa(bs.NetTangibleAssets.Raw) + "," +
		strconv.Itoa(bs.ShortTermInvestments.Raw) + "," +
		strconv.Itoa(bs.NetReceivables.Raw) + "," +
		strconv.Itoa(bs.LongTermDebt.Raw) + "," +
		strconv.Itoa(bs.Inventory.Raw) + "," +
		strconv.Itoa(bs.AccountsPayable.Raw)
}
