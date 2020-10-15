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

func newOrder(typ string, client string, amount int, price float64) interface{} {
	return nil
}
