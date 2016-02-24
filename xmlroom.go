package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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
	m := r.URL.Query().Get("m")
	log.Println(m)
	memberID, err := strconv.Atoi(m)
	if err != nil {
		http.Error(w, "please provide member id", http.StatusBadRequest)
	}
	house := models.Housing{}
	items, err := datastore.ReadItemsByMember(memberID)
	if err != nil {
		log.Fatal(err)
	}
	house.Member.EstateAgentID = memberID
	house.Member.Items = items
	itemsJ, err := json.Marshal(house)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(itemsJ))
}
