package models

type Item struct {
	ID          string `xml:"uniqueobjectid"`
	Street      string `xml:"street"`
	HouseNum    int    `xml:"houseNumber"`
	HouseNumAdd int    `xml:"houseNumberAddtion"`
	PostCode    int    `xml:"postalCode"`
	City        string `xml:"City"`
	SubArea     int    `xml:"SubArea"`
	/* house status legends:
	1=Woonhuis, 2=Appartement
	3=Garage, 4=Kamer
	5=Villa, 6=Woonboot
	7=Studio, 8=Seniorenwoning
	*/
	Type   int `xml:"houseType"`
	Tenant int `xml:"tenant"`
	// Estate owner status legends
	//1 = wekelijks, 2= maandelijks
	EstateOwner        string    `xml:"estateOwner"`
	Stats              int       `xml:"stats"`
	Furnished          bool      `xml:"furnished"`
	MinPrice           int       `xml:"minprice"`
	HidePrice          bool      `xml:"hideprice"`
	NrOfRooms          int       `xml:"NrOfRooms"`
	NrOfLivingRooms    int       `xml:"NrOfLivingRooms"`
	ProjectName        string    `xml:"Projectnaam"`
	HouseTypeInProject string    `xml:"WoningtypeInProject"`
	Available          time.Time `xml:"Available"`
	InsertDate         time.Time `xml:"insertDate"`
	DescriptionNL      string    `xml:"description_nl"`
	DescriptionFR      string    `xml:"description_fr"`
	DescriptionEN      string    `xml:"description_en"`
	DescriptionDE      string    `xml:"description_de"`
	DescriptionES      string    `xml:"description_es"`
	DescriptionIT      string    `xml:"description_it"`
	Photos             []string  `xml: "photos>photo"`
	UpdatePhotos       bool      `xml: "updatePhotos"`
	BrochureURL        string    `xml: "brochure"`
	PlanURL            string    `xml:"plattegrond"`
	Area               int       `xml:"size_m2"`
	NrOfBathrooms      int       `xml:"numberOfBathrooms"`
	ContractLength     int       `xml:"contractLentgh_months"`
	MinContractLength  int       `xml:"minContractLentgh"`
	BuildYear          time.Time `xml:"buildYear"`
	Parking            bool      `xml:"Paring"`
	Bath               bool      `xml:"bath"`
	SeparateShower     bool      `xml:"separateShower"`
	SeparateToilet     bool      `xml:"separateToilet"`
	Lift               bool      `xml:"lift"`
	Garden             bool      `xml:"garden"`
	// garden and roof terrass status legends:
	// N=Noord, O=Oost, Z=Zuid, W=West, NO=Noord Oost
	// NW=Noord West, ZO=Zuid Oost, ZW=Zuid West
	GardenLigging      string `xml:"garden"`
	GardenArea         int    `xml:"gardenSizeM2"`
	RoofTerrass        bool   `xml:"roofTerrass"`
	RoofTerrassLigging string `xml:"roofTerrassLigging"`
	RoofTerrassArea    int    `xml:"roofTerrassSizeM2"`
	Balcony            bool   `xml:"balcony"`
	BalconyLigging     string `xml:"BalconyLigging"`
	BalconyArea        int    `xml:"BaclconySizeM2"`
	SwimmingPool       bool   `xml:"swimmingPool"`
	AirConditioning    bool   `xml:"airConditioning"`
	FirePlace          bool   `xml:"firePlace"`
	Garage             bool   `xml:"garage"`
	Cellar             bool   `xml:"cellar"`
	// Public Transport quality status legends:
	// 0=Geen voorkeur 2=Slecht 3=Middelmatig 4=Goed
	PTQualityID int `xml:"publicTransportQualityID"`
	// 0=niet tonen, 1=tonen
	ShowHouseNum bool `xml:"showhouseNumber"`
	// 0=nee tonen, 1=ja, 3=geen voorkeur
	GroundFloor bool `xml:"groundFloor"`

	// 0000000=geen voorkeur, 0100000=hout, 0010000=laminaat,
	// 0001000=kunststof,0000100=tegelvloer,0000010=tapijt,
	// 0000001=kaal
	FloorQuality int `xml:"floorQuality"`
	// 1xxxxx=Gas,x1xxxx=Water,xx1xxx=Electra
	// xxx1xx=Televisie xxxx1x=Internet xxxxx1=Service Kosten
	RentIncluded int `xml:"rentIncluded"`
}
