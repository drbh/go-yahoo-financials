package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"github.com/montanaflynn/stats"
	"io"
	"log"
	"os"
)

type Sector struct {
	ReturnOnAssetsList []float64
	ReturnOnEquityList []float64
	ProfitMarginList   []float64
	QuickRatioList     []float64
	CurrentRatioList   []float64
	DebtToEquityList   []float64
}

type SectorTargets struct {
	ReturnOnAssetsTarget float64
	ReturnOnEquityTarget float64
	ProfitMarginTarget   float64
	QuickRatioTarget     float64
	CurrentRatioTarget   float64
	DebtToEquityTarget   float64
}

func Medians() {
	csvFile, _ := os.Open(myLocation + "/cleaned_stocks.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var companies []Company
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		companies = append(companies, Company{
			Symbol:    line[0],
			Sector:    line[1],
			SubSector: line[2],
		})
	}

	var sectorList = map[string]*Sector{}

	for i := 1; i < len(companies); i++ {
		comp := companies[i]
		data := ReadRatios(comp.Symbol)
		var ratios ComputedRatios
		json.Unmarshal([]byte(data), &ratios)
		var x = sectorList[comp.Sector]
		if x == nil {
			x = &Sector{}
		}
		x.ReturnOnAssetsList = append(x.ReturnOnAssetsList, ratios.ReturnOnAssets)
		x.ReturnOnEquityList = append(x.ReturnOnEquityList, ratios.ReturnOnEquity)
		x.ProfitMarginList = append(x.ProfitMarginList, ratios.ProfitMargin)
		x.QuickRatioList = append(x.QuickRatioList, ratios.QuickRatio)
		x.CurrentRatioList = append(x.CurrentRatioList, ratios.CurrentRatio)
		x.DebtToEquityList = append(x.DebtToEquityList, ratios.DebtToEquity)
		sectorList[comp.Sector] = x
	}

	for k, v := range sectorList {
		fmt.Println(k)
		var targets SectorTargets

		roamedian, _ := stats.Percentile(v.ReturnOnAssetsList, 50)
		targets.ReturnOnAssetsTarget = roamedian

		roemedian, _ := stats.Percentile(v.ReturnOnEquityList, 50)
		targets.ReturnOnEquityTarget = roemedian

		pmmedian, _ := stats.Percentile(v.ProfitMarginList, 50)
		targets.ProfitMarginTarget = pmmedian

		qrmedian, _ := stats.Percentile(v.QuickRatioList, 50)
		targets.QuickRatioTarget = qrmedian

		crmedian, _ := stats.Percentile(v.CurrentRatioList, 50)
		targets.CurrentRatioTarget = crmedian

		dtemedian, _ := stats.Percentile(v.DebtToEquityList, 50)
		targets.DebtToEquityTarget = dtemedian

		// spew.Dump(targets)
		jsond, _ := json.Marshal(targets)
		WriteTargets(k, jsond)
		// fmt.Println(median) // 3.5

	}
}
