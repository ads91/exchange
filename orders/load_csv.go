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
	// create orders
	order := newOrderFromCSVRow(row)
	// check for bid
	bid, ok := order.(Bid)
	if ok {
		log.Printf("added bid to table: (%s, %s, %s, %s)", row[0], row[1], row[2], row[3])
		ot.Bids = append(ot.Bids, bid) // [TODO]: index into order table, rather than sorting
	}
	// check for offer
	offer, ok := order.(Offer)
	if ok {
		log.Printf("added offer to table: (%s, %s, %s, %s)", row[0], row[1], row[2], row[3])
		ot.Offers = append(ot.Offers, offer) // [TODO]: index into order table, rather than sorting
	}
	// close and delete order file, if required
	csvFile.Close()
	if delete {
		log.Print("deleting order at ", fpath)
		err := os.Remove(fpath)
		if err != nil {
			log.Fatalf("couldn't delete order at %s, error was %s", fpath, err)
		}
	}
}

// newOrderFromCSVRow : convert CSV row into an order instance
func newOrderFromCSVRow(row []string) interface{} {
	// type conversions from CSV file, arguments must align
	amount, _ := strconv.Atoi(row[2])
	price, _ := strconv.ParseFloat(row[3], 64)
	return newOrder(row[1], row[0], amount, price)
}
