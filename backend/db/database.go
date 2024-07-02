package db

import (
	"LiquidTracker/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func InitDB() error {
	// PostgreSQL connection string
	connStr := "host=localhost port=5433 user=postgres password=buschi_8 dbname=liquidtracker sslmode=disable"

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("sql.Open: %v", err)
	}

	// Ping database to ensure it's reachable
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("db.Ping: %v", err)
	}

	log.Println("Connected to PostgreSQL database")

	// Initialize tables
	err = createTables()
	if err != nil {
		return fmt.Errorf("createTables: %v", err)
	}

	return nil
}

func createTables() error {
	if db == nil {
		return fmt.Errorf("database connection is nil")
	}

	queries := []string{
		`CREATE TABLE IF NOT EXISTS brands (
            id SERIAL PRIMARY KEY,
            name TEXT UNIQUE
        )`,
		`CREATE TABLE IF NOT EXISTS ratings (
            id SERIAL PRIMARY KEY,
            brand_id INTEGER REFERENCES brands(id),
            variety TEXT,
            rating INTEGER
        )`,
		`CREATE TABLE IF NOT EXISTS brand_suggestions (
    		id SERIAL PRIMARY KEY,
    		brand_name TEXT
        )`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("db.Exec: %v", err)
		}
	}

	return nil
}

func AddBrand(name string) error {
	query := `INSERT INTO brands (name) VALUES ($1)`

	_, err := db.Exec(query, name)
	if err != nil {
		return fmt.Errorf("db.Exec: %v", err)
	}

	return nil
}

func GetBrandSuggestions() ([]models.BrandSuggestion, error) {
	rows, err := db.Query("SELECT id, brand_name FROM brand_suggestions")
	if err != nil {
		return nil, fmt.Errorf("db.Query: %v", err)
	}
	defer rows.Close()

	var brandSuggestions []models.BrandSuggestion
	for rows.Next() {
		var b models.BrandSuggestion
		if err := rows.Scan(&b.ID, &b.BrandName); err != nil {
			return nil, fmt.Errorf("rows.Scan: %v", err)
		}
		brandSuggestions = append(brandSuggestions, b)
	}

	return brandSuggestions, nil
}
func DeleteBrandSuggestion(brandName string) error {
	_, err := db.Exec("DELETE FROM brand_suggestions WHERE brand_name = $1", brandName)
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *sql.DB {
	if db == nil {
		log.Fatal("Database not initialized")
	}
	return db
}
