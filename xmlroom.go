package main

import (
	"encoding/xml"
	"flag"
	"log"
	"os"

	"github.com/Maksadbek/xmlroom/datastore"
	"github.com/Maksadbek/xmlroom/models"
)

var dbFilePath = flag.String("db", "./db", "database file path")

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
}
