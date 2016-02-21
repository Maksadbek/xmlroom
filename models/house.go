package models

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
)

type Item struct {
	ID          string `xml:"uniqueobjectid"`
	Street      string `xml:"street"`
	HouseNumber string `xml:",cdata"`
	HouseNumAdd string `xml:"houseNumberAddtion"`
	PostCode    string `xml:"postalCode"`
	City        string `xml:"City"`
	SubArea     string `xml:"SubArea"`
	// house status legends:
	// 1=Woonhuis, 2=Appartement
	// 3=Garage, 4=Kamer
	// 5=Villa, 6=Woonboot
	// 7=Studio, 8=Seniorenwoning
	Type   NullInt `xml:"houseType"`
	Tenant NullInt `xml:"tenant"`
	// Estate owner status legends
	//1 = wekelijks, 2= maandelijks
	EstateOwner        string  `xml:"estateOwner"`
	Stats              NullInt `xml:"stats"`
	Furnished          bool    `xml:"furnished"`
	MinPrice           string  `xml:",cdata"`
	HidePrice          bool    `xml:"hideprice"`
	NrOfRooms          NullInt `xml:"NrOfRooms"`
	NrOfLivingRooms    NullInt `xml:"NrOfLivingRooms"`
	ProjectName        string  `xml:"Projectnaam"`
	HouseTypeInProject string  `xml:"WoningtypeInProject"`

	Available         Date     `xml:"Available"`
	InsertDate        Date     `xml:"insertDate"`
	DescriptionNL     string   `xml:",cdata"`
	DescriptionFR     string   `xml:",cdata"`
	DescriptionEN     string   `xml:"description_en"`
	DescriptionDE     string   `xml:"description_de"`
	DescriptionES     string   `xml:"description_es"`
	DescriptionIT     string   `xml:"description_it"`
	Photos            []string `xml: "photos>photo"`
	UpdatePhotos      bool     `xml: "updatePhotos"`
	BrochureURL       string   `xml: "brochure"`
	PlanURL           string   `xml:"plattegrond"`
	Area              NullInt  `xml:"size_m2"`
	NrOfBathrooms     NullInt  `xml:"numberOfBathrooms"`
	ContractLength    NullInt  `xml:"contractLentgh_months"`
	MinContractLength NullInt  `xml:"minContractLentgh"`
	BuildYear         string   `xml:"buildYear"`
	Parking           bool     `xml:"Paring"`
	Bath              bool     `xml:"bath"`
	SeparateShower    bool     `xml:"separateShower"`
	SeparateToilet    bool     `xml:"separateToilet"`
	Lift              bool     `xml:"lift"`
	Garden            bool     `xml:"garden"`
	// garden and roof terrass status legends:
	// N=Noord, O=Oost, Z=Zuid, W=West, NO=Noord Oost
	// NW=Noord West, ZO=Zuid Oost, ZW=Zuid West
	GardenLigging      string  `xml:"gardenLigging"`
	GardenArea         NullInt `xml:"gardenSizeM2"`
	RoofTerrass        bool    `xml:"roofTerrass"`
	RoofTerrassLigging string  `xml:"roofTerrassLigging"`
	RoofTerrassArea    NullInt `xml:"roofTerrassSizeM2"`
	Balcony            bool    `xml:"balcony"`
	BalconyLigging     string  `xml:"BalconyLigging"`
	BalconyArea        NullInt `xml:"BaclconySizeM2"`
	SwimmingPool       bool    `xml:"swimmingPool"`
	AirConditioning    bool    `xml:"airConditioning"`
	FirePlace          bool    `xml:"firePlace"`
	Garage             bool    `xml:"garage"`
	Cellar             bool    `xml:"cellar"`
	// Public Transport quality status legends:
	// 0=Geen voorkeur 2=Slecht 3=Middelmatig 4=Goed
	PTQualityID NullInt `xml:"publicTransportQualityID"`
	// 0=niet tonen, 1=tonen
	ShowHouseNum bool `xml:"showhouseNumber"`
	// 0=nee tonen, 1=ja, 3=geen voorkeur
	GroundFloor bool `xml:"groundFloor"`

	// 0000000=geen voorkeur, 0100000=hout, 0010000=laminaat,
	// 0001000=kunststof,0000100=tegelvloer,0000010=tapijt,
	// 0000001=kaal
	FloorQuality string `xml:"floorQuality"`
	// 1xxxxx=Gas,x1xxxx=Water,xx1xxx=Electra
	// xxx1xx=Televisie xxxx1x=Internet xxxxx1=Service Kosten
	RentIncluded string `xml:"rentIncluded"`
}

type Member struct {
	Items         []Item `xml:"items>item"`
	EstateAgentID int    `xml:"estateAgentID"`
}
type Housing struct {
	XMLName xml.Name `xml:"housing"`
	Member  Member   `xml:"member"`
}

type Date time.Time
type NullInt int

func (d *Date) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	const form = "2006-01-02"
	var val string
	dec.DecodeElement(&val, &start)
	parse, err := time.Parse(form, val)
	if err != nil {
		return err
	}
	*d = Date(parse)
	return nil
}

func (n *NullInt) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var val string
	dec.DecodeElement(&val, &start)
	if val == "" {
		*n = NullInt(0)
		return nil
	}
	fmt.Println(val)
	i, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	*n = NullInt(i)
	return nil
}
