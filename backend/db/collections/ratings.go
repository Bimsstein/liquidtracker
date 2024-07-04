package collections

import (
	arango "LiquidTracker/db"
	"LiquidTracker/models"
	"fmt"
)

func AddRating(brandID, variety string, rating string) error {
	db := arango.GetDB()

	if db == nil {
		return fmt.Errorf("database connection is nil")
	}

	collection, err := db.Collection(nil, "ratings")
	if err != nil {
		return fmt.Errorf("db.Collection: %v", err)
	}

	ratingDoc := models.Rating{
		BrandID: brandID,
		Variety: variety,
		Rating:  rating,
	}

	_, err = collection.CreateDocument(nil, ratingDoc)
	if err != nil {
		return fmt.Errorf("CreateDocument: %v", err)
	}

	return nil
}
