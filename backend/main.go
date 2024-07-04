// main.go

package main

import (
	"log"
	"net/http"
	"os"

	"LiquidTracker/db"
	"LiquidTracker/handlers"
)

func main() {
	// Initialize database
	err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	err = db.CreateCollections()
	if err != nil {
		log.Fatalf("Failed to create collections: %v", err)
	}

	// HTTP handlers
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/submit", handlers.SubmitHandler)
	http.HandleFunc("/suggest-brand", handlers.SuggestBrandHandler)
	http.HandleFunc("/admin", handlers.AdminHandler)
	http.HandleFunc("/admin/delete-suggested-brand", handlers.DeleteSuggestedBrandHandler)
	http.HandleFunc("/admin/add-suggested-brand", handlers.AddSuggestedBrandHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
