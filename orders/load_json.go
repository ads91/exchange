package orders

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// AddJSONOrderFromDir : add a JSON order residing in a local directory to an order table instance
func AddJSONOrderFromDir(ot *OrderTable, fpath string, delete bool) {
	var ojson orderJSON
	// open the file
	jsonFile, err := os.Open(fpath)
	if err != nil {
		log.Fatalf("couldn't open the file %s, error was %s", fpath, err)
	}
	// parse the file
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// read the order from the JSON
	err = json.Unmarshal(byteValue, &ojson)
	if err != nil {
		log.Fatal(err)
	}
	// add order to table
	addOrderToTable(newOrderFromJSON(&ojson), ot)
	// close and delete order file, if required
	closeFile(jsonFile, fpath, delete)
}

// newOrderFromJSON : create an order instance from a JSON
func newOrderFromJSON(ojson *orderJSON) interface{} {
	// type conversions from JSON file, arguments must align
	return newOrder(ojson.Type, ojson.Client, ojson.Amount, ojson.Price)
}
