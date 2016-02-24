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

// control flags
var (
	dbFilePath = flag.String("db", "./db", "database file path")
	port       = flag.String("port", ":8080", "server port")
	migrate    = flag.Bool("migrate", false, "migrate database")
)

func main() {
	flag.Parse()
	log.Println("starting...")
	// connect to datastore
	err := datastore.Init(*dbFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if *migrate {
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
	}
	// serve HTTP web server
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(*port, nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
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
