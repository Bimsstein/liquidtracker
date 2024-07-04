package collections

import (
	arango "LiquidTracker/db"
	"LiquidTracker/models"
	"fmt"
	driver "github.com/arangodb/go-driver"
)

type brandRepository struct {
	//TODO ausserhalb Repository Layer
}

func AddBrand(name string) error { //TODO Klasse zwischenpacken fürs Repository Layer
	db := arango.GetDB()

	if db == nil {
		return fmt.Errorf("database connection is nil")
	}

	collection, err := db.Collection(nil, "brands")
	if err != nil {
		return fmt.Errorf("db.Collection: %v", err)
	}

	brand := models.Brand{
		Name: name,
	}

	_, err = collection.CreateDocument(nil, brand)
	if err != nil {
		return fmt.Errorf("CreateDocument: %v", err)
	}

	return nil
}

func GetBrands() ([]models.Brand, error) {
	db := arango.GetDB()

	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := "FOR b IN brands RETURN b"
	cursor, err := db.Query(nil, query, nil)
	if err != nil {
		return nil, fmt.Errorf("db.Query: %v", err)
	}
	defer cursor.Close()

	var brands []models.Brand
	for {
		var doc models.Brand
		_, err := cursor.ReadDocument(nil, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, fmt.Errorf("cursor.ReadDocument: %v", err)
		}
		brands = append(brands, doc)
	}

	return brands, nil
}
