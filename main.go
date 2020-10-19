package main

import (
	"exchange/matching"
	"exchange/orders"
	"log"
	"os"
	"sync"
)

// [TODO]
// - allow orders with amount > 1 (i.e. once an order is filled, create a new order and
//	assign back to the orders table).
// - add blocking in to order table to avoid multiple read/writes at the same time (race condition).
// - check race condition in sorting of orders vs. reading/writing to the table (lock on sort).

// === packages ===
// go get golang.org/x/tools/cmd/cover

// === testing ===
// go test ./... -v -coverpkg=./...

func main() {
	var wg sync.WaitGroup
	var ot orders.OrderTable
	// process ID
	log.Printf("process ID is %d", os.Getpid())
	// add a wait for the directory order listener
	wg.Add(1)
	go orders.ListenToDir(&ot, &wg, orders.AddJSONOrderFromDir, "/Users/andrewsanderson/Documents/dev/go/src/exchange/data/orders/", 2, true)
	// add a wait for the HTTP order listener service
	wg.Add(1)
	go orders.ListenToHTTP(&wg, ot.AddJSONOrderFromHTTP, ":8080", "/order")
	// add a wait for the order matcher
	wg.Add(1)
	go matching.SettleOrders(&ot, &wg, matching.WriteToJSON, "/Users/andrewsanderson/Documents/dev/go/src/exchange/data/settlements/", 5)
	// wait
	wg.Wait()
}
