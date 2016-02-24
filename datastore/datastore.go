package datastore

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Maksadbek/xmlroom/models"
)

var db *sql.DB

// Init function can be used to connect to datastore and create tables
func Init(dbFile string) (err error) {
	log.Println("initializing datastore...")
	db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		return err
	}
	err = CreateTables()
	if err != nil {
		return err
	}
	return
}

// CreteTables is used to migrate database and create the tables
func CreateTables() error {
	log.Println("creating tables...")
	_, err := db.Exec(createItemsTableQuery)
	if err != nil {
		return err
	}
	_, err = db.Exec(createMembersTableQuery)
	if err != nil {
		return err
	}
	_, err = db.Exec(createMemberItemsTableQuery)
	if err != nil {
		return err
	}
	return nil
}

// CreateItem is used to insert Item data to datastore
func CreateItem(item models.Item) error {
	// photos is string array, and kept in json format
	// unmarshal photos
	photos, err := json.Marshal(item.Photos)
	if err != nil {
		return err
	}
	_, err = db.Exec(createItemQuery,
		item.ID,
		item.Street,
		item.HouseNumber,
		item.HouseNumAdd,
		item.PostCode,
		item.City,
		item.SubArea,
		item.EstateOwner,
		item.MinPrice,
		item.ProjectName,
		item.HouseTypeInProject,
		item.DescriptionNL,
		item.DescriptionFR,
		item.DescriptionEN,
		item.DescriptionDE,
		item.DescriptionES,
		item.DescriptionIT,
		item.BrochureURL,
		item.PlanURL,
		item.BuildYear,
		item.Furnished,
		item.HidePrice,
		item.UpdatePhotos,
		item.Parking,
		item.Bath,
		item.SeparateShower,
		item.SeparateToilet,
		item.Lift,
		item.SwimmingPool,
		item.AirConditioning,
		item.FirePlace,
		item.Garage,
		item.Cellar,
		item.ShowHouseNum,
		photos,
		int(item.Tenant),
		int(item.Area),
		int(item.NrOfBathrooms),
		int(item.ContractLength),
		int(item.MinContractLength),
		int(item.NrOfRooms),
		int(item.NrOfLivingRooms),
		string(item.Stats),
		time.Time(item.Available),
		time.Time(item.InsertDate),
		string(item.Type),
		string(item.PTQuality),
		string(item.GroundFloor),
		item.Garden,
		string(item.GardenLoc),
		item.GardenArea,
		item.RoofTerrass,
		string(item.RoofTerrassLoc),
		item.RoofTerrassArea,
		item.Balcony,
		string(item.BalconyLoc),
		int(item.BalconyArea),
		string(item.FloorQuality),
		string(item.RentIncluded),
	)
	if err != nil {
		return err
	}
	return nil
}

// CreateMemberWithItems can be used to insert members data to database
func CreateMemberWithItems(members []models.Member) error {
	// range over members and its items
	for _, m := range members {
		if _, err := db.Exec(createMemberQuery, m.EstateAgentID); err != nil {
			return err
		}

		for _, item := range m.Items {
			if err := CreateItem(item); err != nil {
				return err
			}
			_, err := db.Exec(createMemberItemsQuery, m.EstateAgentID, item.ID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// ReadItemsByMember is used to get member items(rooms) by its ID number
func ReadItemsByMember(id int) ([]models.Item, error) {
	items := []models.Item{}

	rows, err := db.Query(readItemsByMemberQuery, id)
	if err != nil {
		return items, err
	}
	for rows.Next() {
		item := models.Item{}
		// variables with standard Go types to be able to scan them
		// sql.Scan works only this standard types
		// after scan all of them are converted to custom types
		var (
			tenant, area, nrOfBathrooms, balconyArea,
			contractLen, nrOfRooms, nrOfLivingRooms,
			minContractLen int

			dateAvailable, dateInsert time.Time

			stats, hType, PTQuality, groundFloor,
			gardenLoc, balconyLoc, roofTerrassLoc,
			photos, floorQuality, rentIncluded string
		)

		err = rows.Scan(
			&item.ID,
			&item.Street,
			&item.HouseNumber,
			&item.HouseNumAdd,
			&item.PostCode,
			&item.City,
			&item.SubArea,
			&item.EstateOwner,
			&item.MinPrice,
			&item.ProjectName,
			&item.HouseTypeInProject,
			&item.DescriptionNL,
			&item.DescriptionFR,
			&item.DescriptionEN,
			&item.DescriptionDE,
			&item.DescriptionES,
			&item.DescriptionIT,
			&item.BrochureURL,
			&item.PlanURL,
			&item.BuildYear,
			&item.Furnished,
			&item.HidePrice,
			&item.UpdatePhotos,
			&item.Parking,
			&item.Bath,
			&item.SeparateShower,
			&item.SeparateToilet,
			&item.Lift,
			&item.SwimmingPool,
			&item.AirConditioning,
			&item.FirePlace,
			&item.Garage,
			&item.Cellar,
			&item.ShowHouseNum,
			&photos,
			&tenant,
			&area,
			&nrOfBathrooms,
			&contractLen,
			&minContractLen,
			&nrOfRooms,
			&nrOfLivingRooms,
			&stats,
			&dateAvailable,
			&dateInsert,
			&hType,
			&PTQuality,
			&groundFloor,
			&item.Garden,
			&gardenLoc,
			&item.GardenArea,
			&item.RoofTerrass,
			&roofTerrassLoc,
			&item.RoofTerrassArea,
			&item.Balcony,
			&balconyLoc,
			&balconyArea,
			&floorQuality,
			&rentIncluded,
		)

		if err != nil {
			return items, err
		}

		// converting to custom types
		item.Tenant = models.NullInt(tenant)
		item.Area = models.NullInt(area)
		item.NrOfBathrooms = models.NullInt(nrOfBathrooms)
		item.BalconyArea = models.NullInt(balconyArea)
		item.ContractLength = models.NullInt(contractLen)
		item.NrOfRooms = models.NullInt(nrOfRooms)
		item.NrOfLivingRooms = models.NullInt(nrOfLivingRooms)
		item.MinContractLength = models.NullInt(minContractLen)

		item.Available = models.Date(dateAvailable)
		item.InsertDate = models.Date(dateInsert)

		item.Stats = models.Period(stats)
		item.Type = models.HouseType(hType)
		item.PTQuality = models.PTQuality(PTQuality)
		item.GroundFloor = models.Preference(groundFloor)
		item.GardenLoc = models.Loc(gardenLoc)
		item.RoofTerrassLoc = models.Loc(roofTerrassLoc)
		item.Stats = models.Period(stats)
		item.FloorQuality = models.FQuality(floorQuality)
		item.RentIncluded = models.RentStatus(rentIncluded)

		err = json.Unmarshal([]byte(photos), &item.Photos)
		if err != nil {
			return items, err
		}
		items = append(items, item)
	}
	return items, nil
}

// ReadAall reads all member items from database
// it firstly gets all member its ids and then
// reads its items by its id
func ReadAll() (models.Housing, error) {
	housing := models.Housing{}
	rows, err := db.Query(readAllMembersQuery)
	if err != nil {
		return housing, err
	}

	for rows.Next() {
		var (
			id     int
			member models.Member
		)
		err = rows.Scan(&id)
		if err != nil {
			return housing, err
		}

		member.EstateAgentID = id
		i, err := ReadItemsByMember(id)
		if err != nil {
			return housing, err
		}
		member.Items = i
		housing.Member = append(housing.Member, member)
	}

	return housing, nil
}
