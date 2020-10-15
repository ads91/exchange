package orders

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

// AddCSVOrderFromDir : add orders to an order table instance
func AddCSVOrderFromDir(ot *OrderTable, fpath string, delete bool) {
	// open the file
	csvFile, err := os.Open(fpath)
	if err != nil {
		log.Fatalf("couldn't open %s, error was %s", fpath, err)
	}
	// parse the file
	r := csv.NewReader(csvFile)
	// read the order from CSV
	row, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
	// add order to table
	addOrderToTable(newOrderFromCSVRow(row), ot)
	// close and delete order file, if required
	closeFile(csvFile, fpath, delete)
}

// newOrderFromCSVRow : convert CSV row into an order instance
func newOrderFromCSVRow(row []string) interface{} {
	// type conversions from CSV file, arguments must align
	amount, _ := strconv.Atoi(row[2])
	price, _ := strconv.ParseFloat(row[3], 64)
	return newOrder(row[1], row[0], amount, price)
}
