package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Maksadbek/xmlroom/datastore"
)

// Serve receives host and port, serves web server
func Serve(p string) error {
	// serve HTTP web server
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(p, nil)
	if err != nil {
		return err
	}
	return nil
}

// index handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request %+v\n", r)
	// get all data from datastore
	housing, err := datastore.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// marshal to JSON
	jsonResp, err := json.Marshal(housing)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(jsonResp))
}
