package orders

import (
	"log"
	"net/http"
	"sync"
	"time"
)

// ListenToDir : listen to a local directory for orders
func ListenToDir(ot *OrderTable, wg *sync.WaitGroup, f func(ot *OrderTable, fpath string, delete bool), dir string, waitTimeSecs int, delete bool) {
	var fpaths []string
	defer wg.Done()
	// run indefinitely
	for {
		// wait
		time.Sleep(time.Duration(waitTimeSecs) * time.Second)
		// scan the dir for files
		fpaths = ScanDir(dir)
		// add orders to order table
		for _, fpath := range fpaths {
			f(ot, fpath, delete)
		}
	}
}

// ListenToHTTP : listen to order through an HTTP server
func ListenToHTTP(wg *sync.WaitGroup, f func(w http.ResponseWriter, r *http.Request), port string, uri string) {
	defer wg.Done()
	// set-up handler
	http.HandleFunc(uri, f)
	// listen indefinitely
	log.Fatal(http.ListenAndServe(port, nil))
}

// newOrder : create a new order instance
func newOrder(typ string, client string, amount int, price float64) interface{} {
	// get timestamp for the order
	t := time.Now().UnixNano()
	// attempt to make a new bid or offer instance
	log.Printf("creating order: (%s, %s, %d, %g)", typ, client, amount, price)
	if typ == BID {
		return Bid{Client: client, Amount: amount, Price: price, Time: t}
	} else if typ == OFFER {
		return Offer{Client: client, Amount: amount, Price: price, Time: t}
	}
	log.Fatalf("can't create order of type %s, expected %s or %s", typ, BID, OFFER)
	// [TODO]: return a boolan and the order instance rather than nil
	return nil
}

// addOrderToTable : add an order to an order table
func addOrderToTable(order interface{}, ot *OrderTable) {
	// check for bid
	bid, ok := order.(Bid)
	if ok {
		ot.Bids = append(ot.Bids, bid) // [TODO]: index into order table, rather than sorting
	}
	// check for offer
	offer, ok := order.(Offer)
	if ok {
		ot.Offers = append(ot.Offers, offer) // [TODO]: index into order table, rather than sorting
	}
	log.Printf("added order to table: ", order)
}
