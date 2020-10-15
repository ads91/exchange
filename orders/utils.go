package orders

import (
	"io/ioutil"
	"log"
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
