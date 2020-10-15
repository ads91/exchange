package orders

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// AddJSONOrderFromDir : add a JSON order residing in a local directory to an order table instance
func AddJSONOrderFromDir(ot *OrderTable, fpath string, delete bool) {
	var oj orderJSON
	// open the file
	jsonFile, err := os.Open(fpath)
	if err != nil {
		log.Fatalf("couldn't open the file %s, error was %s", fpath, err)
	}
	// parse the file
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// read the order from the JSON
	err = json.Unmarshal(byteValue, &oj)
	if err != nil {
		log.Fatal(err)
	}
	// add order to table
	addOrderToTable(newOrderFromJSON(&oj), ot)
	// close and delete order file, if required
	closeFile(jsonFile, fpath, delete)
}

// AddJSONOrderFromHTTP : add a JSON order through a request to an HTTP server
func (ot *OrderTable) AddJSONOrderFromHTTP(w http.ResponseWriter, r *http.Request) {
	var oj orderJSON
	// parse form from request
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}
	// deserialise the JSON into a struct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&oj)
	if err != nil {
		log.Fatal(err)
	}
	// add order to table
	addOrderToTable(newOrderFromJSON(&oj), ot)
}

// newOrderFromJSON : create an order instance from a JSON
func newOrderFromJSON(ojson *orderJSON) interface{} {
	// type conversions from JSON file, arguments must align
	return newOrder(ojson.Type, ojson.Client, ojson.Amount, ojson.Price)
}
