package db

import (
	"fmt"
	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"log"
)

var db driver.Database

func ConnectDB() error {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		return fmt.Errorf("http.NewConnection: %v", err)
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "openSesame"),
	})
	if err != nil {
		return fmt.Errorf("driver.NewClient: %v", err)
	}

	db, err = c.Database(nil, "liquidtracker")
	if err != nil {
		return fmt.Errorf("c.Database: %v", err)
	}

	log.Println("Connected to ArangoDB database")

	return nil
}

func CreateCollections() error {
	if db == nil {
		return fmt.Errorf("database connection is nil")
	}

	collections := []string{"brands", "ratings", "brand_suggestions"}

	for _, name := range collections {
		exists, err := db.CollectionExists(nil, name)
		if err != nil {
			return fmt.Errorf("CollectionExists: %v", err)
		}
		if !exists {
			_, err := db.CreateCollection(nil, name, nil)
			if err != nil {
				return fmt.Errorf("CreateCollection: %v", err)
			}
		}
	}

	return nil
}

func GetDB() driver.Database {
	if db == nil {
		log.Fatal("Database not initialized")
	}
	return db
}
