package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"

	"github.com/Maksadbek/xmlroom/models"
)

func main() {
	log.Println("starting...")
	b, err := ioutil.ReadFile("./testdata/challenge.xml")
	if err != nil {
		log.Println(err)
	}
	housing := models.Housing{}
	err = xml.Unmarshal(b, &housing)
	if err != nil {
		log.Println(err)
	}

	log.Printf("%+v\n", housing)
}
