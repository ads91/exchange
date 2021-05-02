package main

import (
	"exchange/config"
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
	// add a wait for the directory order listener if the config demands so
	if config.LOCAL_ORDERS_ENABLED {
		wg.Add(1)
		go orders.ListenToDir(&ot, &wg, orders.AddJSONOrderFromDir, config.LOCAL_ORDERS_DIR, config.LOCAL_ORDERS_SCAN_TIME, config.LOCAL_ORDERS_DELETE_ON_READ)
	}
	// add a wait for the HTTP order listener service
	if config.HTTP_ORDERS_ENABLED {
		wg.Add(1)
		go orders.ListenToHTTP(&wg, ot.AddJSONOrderFromHTTP, getPort(), config.HTTP_ORDERS_END_POINT)
	}
	// add a wait for the order matcher
	wg.Add(1)
	go matching.SettleOrders(&ot, &wg, matching.WriteToJSON, config.SETTLEMENTS_OUTPUT_DIR, config.MATCHING_RATE)
	// wait
	wg.Wait()
}

func getPort() string {
	// test for port
	port := os.Getenv("PORT")
	// default port
	if port == "" {
		port = config.HTTP_ORDERS_PORT
	}
	log.Printf("listening on port %s", port)
	return ":" + port
}
