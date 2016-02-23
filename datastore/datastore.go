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
	_, err := db.Exec(createItemsTableQuery)
	if err != nil {
		return err
	}
	return nil
}
func Create(item models.Item) error {
	photos, err := json.Marshal(item.Photos)
	if err != nil {
		return err
	}
	log.Printf("%+v\n", item)
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
		item.Tenant,
		item.Area,
		item.NrOfBathrooms,
		item.ContractLength,
		item.MinContractLength,
		item.NrOfRooms,
		item.NrOfLivingRooms,
		string(item.Stats),
		time.Time(item.Available),
		time.Time(item.InsertDate),
		item.Type,
		item.PTQuality,
		item.GroundFloor,
		item.Garden,
		item.GardenLoc,
		item.GardenArea,
		item.RoofTerrass,
		item.RoofTerrassLoc,
		item.RoofTerrassArea,
		item.Balcony,
		item.BalconyLoc,
		item.BalconyArea,
		item.FloorQuality,
		item.RentIncluded,
	)
	if err != nil {
		return err
	}
	return nil
}
