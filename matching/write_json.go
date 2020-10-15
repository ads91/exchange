package matching

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// WriteToJSON : write a settlement to a JSON file
func WriteToJSON(stl Settlement, fname string) {
	fname = fname + ".json"
	// serialise the JSON
	file, err := json.MarshalIndent(stl, "", " ")
	if err != nil {
		log.Fatalf("error trying to create settlement %v, error is %s", stl, err)
	}
	log.Printf("writing settlement to %s", fname)
	// write the file
	err = ioutil.WriteFile(fname, file, 0644)
	if err != nil {
		log.Fatalf("error trying to write settlement %v, error is %s", stl, err)
	}
}
