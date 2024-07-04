package collections

import (
	arango "LiquidTracker/db"
	"LiquidTracker/models"
	"fmt"
	driver "github.com/arangodb/go-driver"
)

func GetBrandSuggestions() ([]models.BrandSuggestion, error) {
	db := arango.GetDB()

	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := "FOR b IN brand_suggestions RETURN b"
	cursor, err := db.Query(nil, query, nil)
	if err != nil {
		return nil, fmt.Errorf("db.Query: %v", err)
	}
	defer func(cursor driver.Cursor) {
		err := cursor.Close()
		if err != nil {
			fmt.Printf("cursor.Close: %v", err)
		}
	}(cursor)

	var brandSuggestions []models.BrandSuggestion
	for {
		var doc models.BrandSuggestion
		_, err := cursor.ReadDocument(nil, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, fmt.Errorf("cursor.ReadDocument: %v", err)
		}
		brandSuggestions = append(brandSuggestions, doc)
	}

	return brandSuggestions, nil
}

func AddBrandSuggestions(brandName string) error {
	db := arango.GetDB()

	if db == nil {
		return fmt.Errorf("database connection is nil")
	}

	collection, err := db.Collection(nil, "brand_suggestions")
	if err != nil {
		return fmt.Errorf("db.Collection: %v", err)
	}

	brandSuggestion := models.BrandSuggestion{
		BrandName: brandName,
	}

	_, err = collection.CreateDocument(nil, brandSuggestion)
	if err != nil {
		return fmt.Errorf("CreateDocument: %v", err)
	}

	return nil
}

func DeleteBrandSuggestions(brandName string) error {
	db := arango.GetDB()

	if db == nil {
		return fmt.Errorf("database connection is nil")
	}

	collection, err := db.Collection(nil, "brand_suggestions")
	if err != nil {
		return fmt.Errorf("db.Collection: %v", err)
	}

	query := "FOR b IN brand_suggestions FILTER b.BrandName == @brandName RETURN b"
	bindVars := map[string]interface{}{
		"brandName": brandName,
	}
	cursor, err := db.Query(nil, query, bindVars)
	if err != nil {
		return fmt.Errorf("db.Query: %v", err)
	}

	var doc models.BrandSuggestion
	meta, err := cursor.ReadDocument(nil, &doc)
	if driver.IsNoMoreDocuments(err) {
		return fmt.Errorf("brand suggestion not found")
	} else if err != nil {
		return fmt.Errorf("cursor.ReadDocument: %v", err)
	}

	_, err = collection.RemoveDocument(nil, meta.Key)
	if err != nil {
		return fmt.Errorf("RemoveDocument: %v", err)
	}

	return nil
}
