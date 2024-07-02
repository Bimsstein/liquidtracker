package handlers

import (
	"LiquidTracker/db"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("static/index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.GetDB().Query("SELECT id, name FROM brands")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var brands []Brand
	for rows.Next() {
		var b Brand
		if err := rows.Scan(&b.ID, &b.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		brands = append(brands, b)
	}

	if err := tmpl.Execute(w, brands); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Brand struct {
	ID   int
	Name string
}

func AddBrandHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Brand name is required", http.StatusBadRequest)
		return
	}

	_, err := db.GetDB().Exec("INSERT INTO brands (name) VALUES ($1)", name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
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

	_, err := db.GetDB().Exec("INSERT INTO ratings (brand_id, variety, rating) VALUES ($1, $2, $3)", brandID, variety, rating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	_, err := db.GetDB().Exec("INSERT INTO brand_suggestions (brand_name) VALUES ($1)", name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
