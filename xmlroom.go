package main

import (
	"encoding/xml"
	"flag"
	"log"
	"os"

	"github.com/Maksadbek/xmlroom/api"
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
		log.Println("starting migration to database")
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
		log.Println("successfully migrated data to database")
	}

	api.Serve(*port)
}
