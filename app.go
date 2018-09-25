package main

import (
	"bufio"
	"encoding/csv"
	_ "github.com/andlabs/ui/winmanifest"
	// "encoding/json"
	"fmt"
	"io"
	// "github.com/davecgh/go-spew/spew"
	"github.com/andlabs/ui"
	// "io/ioutil"
	"log"
	// "net/http"
	"os"
	"path/filepath"
	"strconv"
	// "reflect"
)

var companies []Company

func main() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n")
	myLocation = dir

	fmt.Println(myLocation)

	myDataDir := dir + "/data"
	mode := os.FileMode(int(0777))

	if _, err := os.Stat(myDataDir); os.IsNotExist(err) {
		os.Mkdir(myDataDir, mode)
	}

	myDataDir = dir + "/evals"

	if _, err = os.Stat(myDataDir); os.IsNotExist(err) {
		os.Mkdir(myDataDir, mode)
	}

	csvFile, _ := os.Open(myLocation + "/cleaned_stocks.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

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

	databasePath := myLocation + "/bolt.db"
	fmt.Println(databasePath)
	client.OpenBoltDb(databasePath)

	// sym := "AAPL"
	// fmt.Println(sym)

	// GetTechnicalQuote("SRPT")

	ui.Main(func() {
		window = ui.NewWindow("go yahoo financial", 260, 100, true)
		window.SetMargined(true)

		tab := ui.NewTab()
		tab.Append("Basic Controls", makeBasicControlsPage())
		tab.SetMargined(0, true)
		// tab.Append("Numbers and Lists", makeNumbersPage())
		// tab.SetMargined(1, true)
		// tab.Append("Data Choosers", makeDataChoosersPage())
		// tab.SetMargined(2, true)
		window.SetChild(tab)

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()
	})

}

func makeBasicControlsPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	// button1 := ui.NewButton("Button")
	button2 := ui.NewButton("Statements")
	button3 := ui.NewButton("Technical")
	button4 := ui.NewButton("Evaluate")

	// hbox.Append(ui.NewCheckbox("Checkbox"), false)

	// entry1 := ui.NewEntry()
	// entry1.SetReadOnly(false)
	// entry1.SetText("GOOG")
	// hbox.Append(entry1, false)

	// hbox.Append(button1, false)
	hbox.Append(button2, false)
	hbox.Append(button3, false)
	hbox.Append(button4, false)

	vbox.Append(ui.NewLabel("Welcome to go-yahoo-financials\nDownload Balance, Income and Cashflow statements"), false)
	vbox.Append(ui.NewHorizontalSeparator(), false)
	vbox.Append(hbox, false)

	fmt.Println(GetKeys("values"))
	fmt.Println(strconv.Itoa(GetKeys("values")))
	// if GetKeys("values") != nil {
	numValues := strconv.Itoa(GetKeys("values"))
	numRatios := strconv.Itoa(GetKeys("ratios"))
	numTargets := strconv.Itoa(GetKeys("targets"))
	numTechnical := strconv.Itoa(GetKeys("technical"))

	// fmt.Printf("key=%s, value=%s\n", k, v)
	responsestr := fmt.Sprintf("Raw=%s\t\tStatements=%s\nSectors=%s\t\tTechnical=%s", numValues, numRatios, numTargets, numTechnical)

	vbox.Append(ui.NewLabel(responsestr), false)
	// }

	// button1.OnClicked(func(b *ui.Button) {
	// 	sym := entry1.Text()

	// 	// for i := 0; i < 1000; i++ {
	// 	GetFinancials(sym)
	// 	// fmt.Println(i)
	// 	// }

	// 	entry1.SetText("")
	// })
	button2.OnClicked(func(b *ui.Button) {

		for i := 1; i < len(companies); i++ {
			comp := companies[i]
			fmt.Println("\n")
			fmt.Println(i, comp.Symbol)
			GetFinancials(comp.Symbol)
			// entry1.SetText(comp.Symbol)
		}
	})
	button3.OnClicked(func(b *ui.Button) {

		for i := 1; i < len(companies); i++ {
			comp := companies[i]
			fmt.Println("\n")
			fmt.Println(i, comp.Symbol)
			GetTechnicalQuote(comp.Symbol)
			// entry1.SetText(comp.Symbol)
		}
	})
	button4.OnClicked(func(b *ui.Button) {

		Medians()
		Check()

	})

	return vbox
}
