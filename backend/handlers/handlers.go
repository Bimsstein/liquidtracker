package handlers

import (
	"LiquidTracker/db/collections"
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("static/index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	brands, err := collections.GetBrands()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, brands)
	if err != nil {
		// At this point, a partial response may have already been written,
		// so it's too late to send an HTTP 500 error.
		log.Printf("Error executing template: %v", err)
	}
}

type Brand struct {
	ID   int
	Name string
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	brandID := r.FormValue("brand")
	variety := r.FormValue("variety")
	rating := r.FormValue("rating")

	if brandID == "" || variety == "" || rating == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	err := collections.AddRating(brandID, variety, rating)
	if err != nil {
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func SuggestBrandHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Brand name is required", http.StatusBadRequest)
		return
	}

	err := collections.AddBrandSuggestions(name)
	if err != nil {
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
