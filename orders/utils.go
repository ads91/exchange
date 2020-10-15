package orders

import (
	"io/ioutil"
	"log"
	"os"
)

// ScanDir : scan a directory for files
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

// closeFile : close an open file and delete it, if required
func closeFile(file *os.File, fpath string, delete bool) {
	file.Close()
	if delete {
		log.Print("deleting order at ", fpath)
		err := os.Remove(fpath)
		if err != nil {
			log.Fatalf("couldn't delete order at %s, error was %s", fpath, err)
		}
	}
}
