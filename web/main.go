package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Maksadbek/xmlroom/src/datastore"
)

var (
	p      = flag.String("port", ":8080", "server host and port")
	dbFile = flag.String("db", "./db", "database file path")
)

func main() {
	flag.Parse()
	// initialize database
	err := datastore.Init(*dbFile)
	if err != nil {
		log.Fatal(err)
	}
	// serve HTTP web server
	log.Println("starting web server on", *p)
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(*p, nil))
}

// index handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	// get all data from datastore
	housing, err := datastore.ReadAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}
	// marshal to JSON
	jsonResp, err := json.Marshal(housing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, string(jsonResp))
}
