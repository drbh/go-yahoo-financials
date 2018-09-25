package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	// "fmt"
	// "github.com/davecgh/go-spew/spew"
	"io"
	"log"
	"os"
	"strconv"
)

type Evaluated struct {
	Symbol                 string
	RsiTrigger             bool
	MfiTrigger             bool
	ReturnOnAssetstTrigger bool
	ReturnOnEquityTrigger  bool
	ProfitMarginTrigger    bool
	QuickRatioTrigger      bool
	CurrentRatioTrigger    bool
	DebtToEquityTrigger    bool
}

func (ev *Evaluated) ToCSV() {
	filecontent := ""
	filecontent = filecontent + ev.Symbol + "," +
		strconv.FormatBool(ev.RsiTrigger) + "," +
		strconv.FormatBool(ev.MfiTrigger) + "," +
		strconv.FormatBool(ev.ReturnOnAssetstTrigger) + "," +
		strconv.FormatBool(ev.ReturnOnEquityTrigger) + "," +
		strconv.FormatBool(ev.ProfitMarginTrigger) + "," +
		strconv.FormatBool(ev.QuickRatioTrigger) + "," +
		strconv.FormatBool(ev.CurrentRatioTrigger) + "," +
		strconv.FormatBool(ev.DebtToEquityTrigger)

	dumpToFile(myLocation+"/evals/"+ev.Symbol+"_"+"eval.csv", filecontent)
}

func Check() {

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

	var tape []Evaluated

	for i := 1; i < len(companies); i++ {
		score := 0
		comp := companies[i]
		data := ReadRatios(comp.Symbol)
		targets := ReadTargets(comp.Sector)
		technicals := ReadTechnicals(comp.Symbol)
		//

		var ratios ComputedRatios
		json.Unmarshal([]byte(data), &ratios)

		var targs SectorTargets
		json.Unmarshal([]byte(targets), &targs)

		var techs TechnicalQuote
		json.Unmarshal([]byte(technicals), &techs)

		if techs.Symbol == "" {

		} else {

			// spew.Dump(techs)

			evalut := Evaluated{
				techs.Symbol,
				false, false, false, false,
				false, false, false, false,
			}

			if techs.Rsi <= .005 {
				score++
				evalut.RsiTrigger = true
			}

			if techs.Mfi <= .005 {
				score++
				evalut.MfiTrigger = true
			}

			if ratios.ReturnOnAssets >= targs.ReturnOnAssetsTarget {
				score++
				evalut.ReturnOnAssetstTrigger = true
			}

			if ratios.ReturnOnEquity >= targs.ReturnOnEquityTarget {
				score++
				evalut.ReturnOnEquityTrigger = true
			}

			if ratios.ProfitMargin >= targs.ProfitMarginTarget {
				score++
				evalut.ProfitMarginTrigger = true
			}

			if 1 <= ratios.QuickRatio && ratios.QuickRatio <= 2.5 {
				score++
				evalut.QuickRatioTrigger = true
			}

			if 1 <= ratios.CurrentRatio && ratios.CurrentRatio <= 2 {
				score++
				evalut.CurrentRatioTrigger = true
			}

			if ratios.DebtToEquity <= targs.DebtToEquityTarget {
				score++
				evalut.DebtToEquityTrigger = true
			}

			if score > 5 {
				evalut.ToCSV()
				// fmt.Println(comp.Symbol, score)
				// fmt.Println(evalut)
				// if evalut.RsiTrigger == true && evalut.MfiTrigger == true {
				// if evalut.MfiTrigger == true {
				// if evalut.RsiTrigger == true {
				if true {
					// fmt.Println(techs)
					tape = append(tape, evalut)
				}

			}

		}
		// fmt.Println(targets)
	}
	// fmt.Println(tape)
	jsond, _ := json.Marshal(tape)
	dumpToFile(myLocation+"/tape.json", string(jsond))
}
