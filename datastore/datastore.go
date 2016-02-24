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

func CreateItem(item models.Item) error {
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

func CreateMemberWithItems(member models.Member) error {
	// create member
	if _, err := db.Exec(createMemberQuery, member.EstateAgentID); err != nil {
		return err
	}

	for _, item := range member.Items {
		if err := CreateItem(item); err != nil {
			return err
		}
		_, err := db.Exec(createMemberItemsQuery, member.EstateAgentID, item.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadItemsByMember(id int) ([]models.Item, error) {
	items := []models.Item{}

	rows, err := db.Query(readItemsByMemberQuery, id)
	if err != nil {
		return items, err
	}
	for rows.Next() {
		item := models.Item{}
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

func ReadAll() ([]models.Member, error) {
	members := []models.Member{}
	rows, err := db.Query(readAllMembersQuery)
	if err != nil {
		return members, err
	}

	ids := []int{}
	for rows.Next() {
		var m int
		err = rows.Scan(&m)
		if err != nil {
			return members, err
		}
		ids = append(ids, m)
	}
}
