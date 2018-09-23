package main

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
