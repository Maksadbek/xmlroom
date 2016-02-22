package main

import (
	"encoding/xml"
	"log"
	"os"

	"github.com/Maksadbek/xmlroom/models"
)

func main() {
	log.Println("starting...")
	// open file
	f, err := os.Open("./testdata/challenge.xml")
	if err != nil {
		log.Fatal(err)
	}
	dec := xml.NewDecoder(f)
	housing := models.Housing{}
	err = dec.Decode(&housing)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", housing.Member.Items[0].RentIncluded)
}
