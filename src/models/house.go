package models

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"time"
)

// the time form to convert
const FORM = "2006-01-02"

// member structure
type Member struct {
	Items         []Item `xml:"items>item"`
	EstateAgentID int    `xml:"estateAgentID"`
}

// the parent structure
type Housing struct {
	XMLName xml.Name `xml:"housing" json:"-"`
	Member  []Member `xml:"member"`
}

// Custom type for custom XML, JSON marshalling

// Date time type can be used to customize time to JSON and XML marshalling
type Date time.Time
type NullInt int

// House Type legends:
// 1=Woonhuis
// 2=Appartement
// 3=Garage
// 4=Kamer
// 5=Villa
// 6=Woonboot
// 7=Studio
// 8=Seniorenwoning
type HouseType string

// Location status legends
// garden,roof and terrass has such types:
// N=Noord,
// O=Oost,
// Z=Zuid,
// W=West,
// NO=Noord Oost
// ZO=Zuid Oost,
// NW=Noord West,
// ZW=Zuid West
type Loc string

// Public Transport quality status legends:
// 0=Geen voorkeur
// 2=Slecht
// 3=Middelmatig
// 4=Goed
type PTQuality string

// 0=nee tonen
// 1=ja
// 3=geen voorkeur
type Preference string

// Estate owner status legends
// 1 = wekelijks
// 2= maandelijks
type Period string

// 0100000=32 =hout
// 0010000=16 =laminaat
// 0001000=8  =kunststof
// 0000100=4  =tegelvloer
// 0000010=2  =tapijt
// 0000001=1  =kaal
// 0000000=0  =geen voorkeur
type FQuality string

// 1xxxxx=Gas,
// x1xxxx=Water
// xx1xxx=Electra
// xxx1xx=Televisie
// xxxx1x=Internet
// xxxxx1=Service Kosten
type RentStatus string

// House room
type Item struct {
	ID                 string   `xml:"uniqueobjectid"`
	Street             string   `xml:"street"`
	HouseNumber        string   `xml:",cdata"`
	HouseNumAdd        string   `xml:"houseNumberAddtion"`
	PostCode           string   `xml:"postalCode"`
	City               string   `xml:"City"`
	SubArea            string   `xml:"SubArea"`
	EstateOwner        string   `xml:"estateOwner"`
	MinPrice           NullInt  `xml:",cdata"`
	ProjectName        string   `xml:"Projectnaam"`
	HouseTypeInProject string   `xml:"WoningtypeInProject"`
	DescriptionNL      string   `xml:",cdata"`
	DescriptionFR      string   `xml:",cdata"`
	DescriptionEN      string   `xml:"description_en"`
	DescriptionDE      string   `xml:"description_de"`
	DescriptionES      string   `xml:"description_es"`
	DescriptionIT      string   `xml:"description_it"`
	BrochureURL        string   `xml: "brochure"`
	PlanURL            string   `xml:"plattegrond"`
	BuildYear          string   `xml:"buildYear"`
	Furnished          bool     `xml:"furnished"`
	HidePrice          bool     `xml:"hideprice"`
	UpdatePhotos       bool     `xml: "updatePhotos"`
	Parking            bool     `xml:"Paring"`
	Bath               bool     `xml:"bath"`
	SeparateShower     bool     `xml:"separateShower"`
	SeparateToilet     bool     `xml:"separateToilet"`
	Lift               bool     `xml:"lift"`
	SwimmingPool       bool     `xml:"swimmingPool"`
	AirConditioning    bool     `xml:"airConditioning"`
	FirePlace          bool     `xml:"firePlace"`
	Garage             bool     `xml:"garage"`
	Cellar             bool     `xml:"cellar"`
	ShowHouseNum       bool     `xml:"showhouseNumber"`
	Photos             []string `xml:"photos>photo"`

	Tenant            NullInt    `xml:"tenant"`
	Area              NullInt    `xml:"size_m2"`
	NrOfBathrooms     NullInt    `xml:"numberOfBathrooms"`
	ContractLength    NullInt    `xml:"contractLentgh_months"`
	MinContractLength NullInt    `xml:"minContractLentgh"`
	NrOfRooms         NullInt    `xml:"NrOfRooms"`
	NrOfLivingRooms   NullInt    `xml:"NrOfLivingRooms"`
	Stats             Period     `xml:"stats"`
	Available         Date       `xml:"Available"`
	InsertDate        Date       `xml:"insertDate"`
	Type              HouseType  `xml:"houseType"`
	PTQuality         PTQuality  `xml:"publicTransportQualityID"`
	GroundFloor       Preference `xml:"groundFloor"`

	Garden     bool    `xml:"garden"`
	GardenLoc  Loc     `xml:"gardenLigging"`
	GardenArea NullInt `xml:"gardenSizeM2"`

	RoofTerrass     bool `xml:"roofTerrass"`
	RoofTerrassLoc  Loc  `xml:"roofTerrassLigging"`
	RoofTerrassArea int  `xml:"roofTerrassSizeM2"`

	Balcony     bool    `xml:"balcony"`
	BalconyLoc  Loc     `xml:"BalconyLigging"`
	BalconyArea NullInt `xml:"BaclconySizeM2"`

	FloorQuality FQuality `xml:"floorQuality"`

	RentIncluded RentStatus `xml:"rentIncluded"`
}

func (d *Date) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	// input time format
	var val string
	dec.DecodeElement(&val, &start)
	parse, err := time.Parse(FORM, val)
	if err != nil {
		return err
	}
	*d = Date(parse)
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	// get string format of time by required form
	t := time.Time(*d).Format(FORM)
	return json.Marshal(t)

}
func (n *NullInt) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var val string
	dec.DecodeElement(&val, &start)
	if val == "" {
		*n = NullInt(0)
		return nil
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	*n = NullInt(i)
	return nil
}

func (h *HouseType) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var val string
	dec.DecodeElement(&val, &start)
	switch val {
	case "1":
		*h = HouseType("woonhuis")
	case "2":
		*h = HouseType("appartement")
	case "3":
		*h = HouseType("garage")
	case "4":
		*h = HouseType("kamer")
	case "5":
		*h = HouseType("villa")
	case "6":
		*h = HouseType("woonboot")
	case "7":
		*h = HouseType("studio")
	case "8":
		*h = HouseType("seniorenwoning")
	default:
		*h = HouseType("")
	}
	return nil
}

func (h *Loc) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var val string
	dec.DecodeElement(&val, &start)
	switch val {
	case "N":
		*h = Loc("Noord")
	case "O":
		*h = Loc("Oost")
	case "Z":
		*h = Loc("Zuid")
	case "W":
		*h = Loc("West")
	case "NO":
		*h = Loc("Noord Oost")
	case "NW":
		*h = Loc("Noord West")
	case "ZO":
		*h = Loc("Zuid Oost")
	case "ZW":
		*h = Loc("Zuid West")
	default:
		*h = Loc("")
	}
	return nil
}

func (h *PTQuality) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var val string
	dec.DecodeElement(&val, &start)
	switch val {
	case "0":
		*h = PTQuality("Geen voorkeur")
	case "2":
		*h = PTQuality("Slecht")
	case "3":
		*h = PTQuality("Middelmatig")
	case "4":
		*h = PTQuality("Goed")
	default:
		*h = PTQuality("")
	}
	return nil
}

func (p *Preference) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var val string
	dec.DecodeElement(&val, &start)
	switch val {
	case "0":
		*p = Preference("nee tonen")
	case "1":
		*p = Preference("ja")
	case "3":
		*p = Preference("geen voorkeur")
	default:
		*p = Preference("")
	}
	return nil
}

func (p *Period) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var val string
	dec.DecodeElement(&val, &start)
	switch val {
	case "1":
		*p = Period("wekelijks")
	case "2":
		*p = Period("maandelijks")
	default:
		*p = Period("")
	}
	return nil
}

func (q *FQuality) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var val string
	dec.DecodeElement(&val, &start)
	if val == "" {
		*q = ""
		return nil
	}
	statusCode, err := strconv.ParseInt(val, 2, 64)
	if err != nil {
		return err
	}
	switch statusCode {
	case 0:
		*q = FQuality("geen voorkeur")
	case 1:
		*q = FQuality("kaal")
	case 2:
		*q = FQuality("tapijt")
	case 4:
		*q = FQuality("tegelvloer")
	case 8:
		*q = FQuality("kunststof")
	case 16:
		*q = FQuality("laminaat")
	case 32:
		*q = FQuality("hout")
	default:
		*q = FQuality("")
	}
	return nil
}
