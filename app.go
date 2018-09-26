package main

import (
	"bufio"
	// "encoding/csv"
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/andlabs/ui/winmanifest"
	// "io"
	// "github.com/davecgh/go-spew/spew"
	"github.com/andlabs/ui"
	// "io/ioutil"
	"log"
	// "net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	// "time"
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

	fmt.Println("Reading cleaned_stocks")

	// fmt.Println(string(data))

	// csvFile, _ := os.Open(myLocation + "/cleaned_stocks.csv")
	// reader := csv.NewReader(bufio.NewReader(csvFile))

	// for {
	// 	line, error := reader.Read()
	// 	if error == io.EOF {
	// 		break
	// 	} else if error != nil {
	// 		log.Fatal(error)
	// 	}
	// 	companies = append(companies, Company{
	// 		Symbol:    line[0],
	// 		Sector:    line[1],
	// 		SubSector: line[2],
	// 	})
	// }
	data, err := Asset("cleaned_stocks.csv")
	if err != nil {
		// Asset was not found.
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		// fmt.Println(line)
		companies = append(companies, Company{
			Symbol:    string(line[0]),
			Sector:    string(line[1]),
			SubSector: string(line[2]),
		})
	}

	databasePath := myLocation + "/bolt.db"
	fmt.Println(databasePath)
	client.OpenBoltDb(databasePath)

	// sym := "AAPL"
	// fmt.Println(companies)
	// GetTechnicalQuote("SRPT")
	ui.Main(func() {
		window = ui.NewWindow("go yahoo financial", 500, 700, true)
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

var currentTicker string

// func updadeTechnicals(entry1 *ui.ProgressBar) {
func updadeTechnicals() {
	for i := 1; i < len(companies); i++ {
		comp := companies[i]
		// fmt.Println("\n")
		fmt.Println(i, comp.Symbol)
		// defer
		GetTechnicalQuote(comp.Symbol)
		currentTicker = comp.Symbol
		val := int((float64(i) / float64(len(companies))) * 100)
		fmt.Println(val) //, i, len(companies))
		// fmt.Printf("%T\n", val)
		// fmt.Printf("%T\n", i)
		// fmt.Printf("%T\n", len(companies))
		// entry1.SetValue(val)
	}
}

// func updadeTechnicals(entry1 *ui.ProgressBar) {
func updadeStatements() {
	for i := 1; i < len(companies); i++ {
		comp := companies[i]
		fmt.Println("\n")
		fmt.Println(i, comp.Symbol)
		GetFinancials(comp.Symbol)
		// entry1.SetText(comp.Symbol)
	}
}

func makeBasicControlsPage() ui.Control {

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	// button1 := ui.NewButton("Button")
	button2 := ui.NewButton("Statements")
	button3 := ui.NewButton("Technical")

	// hbox.Append(ui.NewCheckbox("Checkbox"), false)

	// hbox.Append(button1, false)
	hbox.Append(button2, false)
	hbox.Append(button3, false)

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

	vbox.Append(ui.NewLabel("Last Statement Update:\nLast Technical Update:"), false)

	// }

	// button1.OnClicked(func(b *ui.Button) {
	// 	sym := entry1.Text()

	// 	// for i := 0; i < 1000; i++ {
	// 	GetFinancials(sym)
	// 	// fmt.Println(i)
	// 	// }
	// 	entry1.SetText("")
	// })

	// ticker := time.NewTicker(500 * time.Millisecond)
	// quit := make(chan struct{})
	// go func() {
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			// do stuff

	// 			entry1.SetText(currentTicker)
	// 		case <-quit:
	// 			ticker.Stop()
	// 			return
	// 		}
	// 	}
	// }()

	entry1 := ui.NewEntry()
	label1 := ui.NewLabel("RSI Limit")
	entry1.SetReadOnly(false)
	entry1.SetText("5.0")
	vbox.Append(label1, false)
	vbox.Append(entry1, false)

	entry2 := ui.NewEntry()
	label2 := ui.NewLabel("MFI limit")
	entry2.SetReadOnly(false)
	entry2.SetText("0.5")
	vbox.Append(label2, false)
	vbox.Append(entry2, false)

	entry3 := ui.NewEntry()
	label3 := ui.NewLabel("Statement Score Limit (out of 6)")
	entry3.SetReadOnly(false)
	entry3.SetText("5")
	vbox.Append(label3, false)
	vbox.Append(entry3, false)

	// entry4 := ui.NewEntry()
	label4 := ui.NewLabel("Results")
	mline := ui.NewMultilineEntry()
	mline.SetReadOnly(true)
	// entry4.SetReadOnly(true)
	// entry4.SetText("waiting")
	vbox.Append(label4, false)
	vbox.Append(mline, false)
	// vbox.Append(entry4, false)

	button2.OnClicked(func(b *ui.Button) {
		go updadeStatements()
	})
	button3.OnClicked(func(b *ui.Button) {

		// entry1 := ui.NewProgressBar()
		// entry1.SetReadOnly(false)
		// entry1.SetText("")
		// hbox.Append(entry1, false)
		// go func() {
		go updadeTechnicals()

		// }()

	})
	button4 := ui.NewButton("Evaluate")
	vbox.Append(button4, false)
	button4.OnClicked(func(b *ui.Button) {

		Medians()

		rsiLimit, _ := strconv.ParseFloat(entry1.Text(), 64)
		mfiLimit, _ := strconv.ParseFloat(entry2.Text(), 64)
		scoreLimit, _ := strconv.Atoi(entry3.Text())

		val := Check(rsiLimit, mfiLimit, scoreLimit)

		var prettyJSON bytes.Buffer
		error := json.Indent(&prettyJSON, val, "", "\t")
		if error != nil {
			return
		}

		mline.SetText(string(prettyJSON.Bytes()))
	})

	return vbox
}
