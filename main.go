package main

import (
	"encoding/csv"
	"math"
	"os"

	"github.com/rivo/tview"
)

func main() {
	// get filename from command line args
	var filename string = os.Args[1]
	file, err := os.Open(filename)
	checkNillError(err)
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	checkNillError(err)

	cols := 0
	data := make([][]string, 0)
	for _, line := range lines {
		row := make([]string, 0)
		row = append(row, line...)
		data = append(data, row)
		cols = int(math.Max(float64(cols), float64(len(row))))
	}

	rows := len(data)

	// tview
	app := tview.NewApplication()
	table := tview.NewTable().
		SetBorders(true)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			table.SetCell(r, c, tview.NewTableCell(data[r][c]).SetAlign(tview.AlignCenter))
		}
	}

	if err := app.SetRoot(table, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func checkNillError(err error) {
	if err != nil {
		panic(err)
	}
}
