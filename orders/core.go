package orders

import (
	"io/ioutil"
	"log"
	"sync"
	"time"
)

func ScanDir(dir string) []string {
	var orders []string
	files, err := ioutil.ReadDir(dir)
	// check for errors
	if err != nil {
		log.Fatal(err)
	}
	// loop the files
	for _, file := range files {
		orders = append(orders, dir+file.Name())
	}
	return orders
}

// ListToDir : listen to a local directory for orders
func ListenToDir(ot *OrderTable, wg *sync.WaitGroup, f func(ot *OrderTable, fpath string, delete bool), dir string, waitTimeSecs int, delete bool) {
	var fpaths []string
	defer wg.Done()
	// run indefinitely
	for {
		// wait
		time.Sleep(time.Duration(waitTimeSecs) * time.Second)
		// scan the dir for files
		fpaths = ScanDir(dir)

	}
}
