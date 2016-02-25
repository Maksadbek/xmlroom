package models

import "time"

var testXML string = `
<Member>
   <items>
	 <item>
	   <uniqueobjectid>1</uniqueobjectid>
	   <street>teststreet</street>
	   <HouseNumber>1</HouseNumber>
	   <houseNumberAddtion>1</houseNumberAddtion>
	   <postalCode>1</postalCode>
	   <City>1</City>
	   <SubArea>1</SubArea>
	   <estateOwner>testOwner</estateOwner>
	   <MinPrice>1</MinPrice>
	   <Projectnaam>testProjectName</Projectnaam>
	   <WoningtypeInProject>testHouseType</WoningtypeInProject>
	   <DescriptionNL>nl_description</DescriptionNL>
	   <DescriptionFR>fr_desctiption</DescriptionFR>
	   <description_en>en_description</description_en>
	   <description_de>de_description</description_de>
	   <description_es>es_description</description_es>
	   <description_it>it_description</description_it>
	   <BrochureURL>testBrochureUrl</BrochureURL>
	   <plattegrond>testPlanURL</plattegrond>
	   <buildYear>testBuildYear</buildYear>
	   <furnished>1</furnished>
	   <hideprice>1</hideprice>
	   <UpdatePhotos>1</UpdatePhotos>
	   <Paring>1</Paring>
	   <bath>1</bath>
	   <separateShower>1</separateShower>
	   <separateToilet>1</separateToilet>
	   <lift>1</lift>
	   <swimmingPool>1</swimmingPool>
	   <airConditioning>1</airConditioning>
	   <firePlace>1</firePlace>
	   <garage>1</garage>
	   <cellar>1</cellar>
	   <showhouseNumber>1</showhouseNumber>
	   <photos>
		 <photo>test_1</photo>
		 <photo>test_2</photo>
	   </photos>
	   <tenant>1</tenant>
	   <size_m2>1</size_m2>
	   <numberOfBathrooms>1</numberOfBathrooms>
	   <contractLentgh_months>1</contractLentgh_months>
	   <minContractLentgh>1</minContractLentgh>
	   <NrOfRooms>1</NrOfRooms>
	   <NrOfLivingRooms>1</NrOfLivingRooms>
	   <stats>testStats</stats>
	   <Available>2006-01-02</Available>
	   <insertDate>2006-01-02</insertDate>
	   <houseType>1</houseType>
	   <publicTransportQualityID>0</publicTransportQualityID>
	   <groundFloor>3</groundFloor>
	   <garden>1</garden>
	   <gardenLigging>ZW</gardenLigging>
	   <gardenSizeM2>1</gardenSizeM2>
	   <roofTerrass>1</roofTerrass>
	   <roofTerrassLigging>O</roofTerrassLigging>
	   <roofTerrassSizeM2>1</roofTerrassSizeM2>
	   <balcony>1</balcony>
	   <BalconyLigging>Z</BalconyLigging>
	   <BaclconySizeM2>1</BaclconySizeM2>
	   <floorQuality>00000001</floorQuality>
	   <rentIncluded>111111</rentIncluded>
	 </item>
   </items>
   <estateAgentID>123</estateAgentID>
 </Member>
`

var testItem Item = Item{
	ID:                 "1",
	Street:             "teststreet",
	HouseNumber:        "1",
	HouseNumAdd:        "1",
	PostCode:           "1",
	City:               "1",
	SubArea:            "1",
	EstateOwner:        "testOwner",
	MinPrice:           NullInt(1),
	ProjectName:        "testProjectName",
	HouseTypeInProject: "testHouseType",
	DescriptionNL:      "nl_description",
	DescriptionFR:      "fr_desctiption",
	DescriptionEN:      "en_description",
	DescriptionDE:      "de_description",
	DescriptionES:      "es_description",
	DescriptionIT:      "it_description",
	BrochureURL:        "testBrochureUrl",
	PlanURL:            "testPlanURL",
	BuildYear:          "testBuildYear",
	Furnished:          true,
	HidePrice:          true,
	UpdatePhotos:       true,
	Parking:            true,
	Bath:               true,
	SeparateShower:     true,
	SeparateToilet:     true,
	Lift:               true,
	SwimmingPool:       true,
	AirConditioning:    true,
	FirePlace:          true,
	Garage:             true,
	Cellar:             true,
	ShowHouseNum:       true,
	Photos:             []string{"test_1", "test_2"},
	Tenant:             NullInt(1),
	Area:               NullInt(1),
	NrOfBathrooms:      NullInt(1),
	ContractLength:     NullInt(1),
	MinContractLength:  NullInt(1),
	NrOfRooms:          NullInt(1),
	NrOfLivingRooms:    NullInt(1),
	Stats:              Period(""),
	Available:          Date(time.Now()),
	InsertDate:         Date(time.Now()),
	Type:               HouseType("woonhouis"),
	PTQuality:          PTQuality("Geen voorkeur"),
	GroundFloor:        Preference("geen voorkeur"),
	Garden:             true,
	GardenLoc:          Loc("Zuid West"),
	GardenArea:         NullInt(1),
	RoofTerrass:        true,
	RoofTerrassLoc:     Loc("Oost"),
	RoofTerrassArea:    1,
	Balcony:            true,
	BalconyLoc:         Loc("Zuid"),
	BalconyArea:        NullInt(1),
	FloorQuality:       FQuality("kaal"),
	RentIncluded:       RentStatus("111111"),
}

var testMembers []Member = []Member{
	Member{
		Items:         []Item{testItem},
		EstateAgentID: 123,
	},
}
