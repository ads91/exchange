package orders

import (
	"sync"
	"time"
)

func ScanDir(dir string) []string {}

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
