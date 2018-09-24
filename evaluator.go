package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func Check() {

	csvFile, _ := os.Open("cleaned_stocks.csv")
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

	// var sectorList = map[string]*Sector{}

	for i := 1; i < len(companies); i++ {
		score := 0
		comp := companies[i]
		data := ReadRatios(comp.Symbol)
		targets := ReadTargets(comp.Sector)

		var ratios ComputedRatios
		json.Unmarshal([]byte(data), &ratios)

		var targs SectorTargets
		json.Unmarshal([]byte(targets), &targs)

		if ratios.ReturnOnAssets >= targs.ReturnOnAssetsTarget {
			score++
		}

		if ratios.ReturnOnEquity >= targs.ReturnOnEquityTarget {
			score++
		}

		if ratios.ProfitMargin >= targs.ProfitMarginTarget {
			score++
		}

		if 1 <= ratios.QuickRatio && ratios.QuickRatio <= 2.5 {
			score++
		}

		if 1 <= ratios.CurrentRatio && ratios.CurrentRatio <= 2 {
			score++
		}

		if ratios.DebtToEquity <= targs.DebtToEquityTarget {
			score++
		}

		if score > 5 {
			fmt.Println(comp.Symbol, score)
		}

		// fmt.Println(targets)
	}

}
