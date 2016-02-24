package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Maksadbek/xmlroom/datastore"
	"github.com/Maksadbek/xmlroom/models"
)

var (
	dbFilePath = flag.String("db", "./db", "database file path")
	port       = flag.String("port", ":8080", "server port")
)

func main() {
	flag.Parse()
	log.Println("starting...")
	err := datastore.Init(*dbFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// open file
	f, err := os.Open("./testdata/challenge.xml")
	if err != nil {
		log.Fatal(err)
	}

	// decode xml
	dec := xml.NewDecoder(f)
	housing := models.Housing{}
	err = dec.Decode(&housing)
	if err != nil {
		log.Fatal(err)
	}

	// insert data
	err = datastore.CreateMemberWithItems(housing.Member)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(*port, nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	housing, err := datastore.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	jsonResp, err := json.Marshal(housing)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(jsonResp))
}
