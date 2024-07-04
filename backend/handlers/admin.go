// handlers/admin.go

package handlers

import (
	"LiquidTracker/db/collections"
	"LiquidTracker/models"
	"html/template"
	"net/http"
)

const (
	adminUsername = "admin"
	adminPassword = "adminpassword"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is authenticated
	username, password, ok := r.BasicAuth()
	if !ok || username != adminUsername || password != adminPassword {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	// Get existing brand suggestions from database
	suggestions, err := collections.GetBrandSuggestions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare data for rendering in template
	data := struct {
		BrandSuggestions []models.BrandSuggestion
	}{
		BrandSuggestions: suggestions,
	}

	// Render admin page template with data
	tmpl := template.Must(template.ParseFiles("static/admin.html"))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteSuggestedBrandHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is authenticated
	username, password, ok := r.BasicAuth()
	if !ok || username != adminUsername || password != adminPassword {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	// Process form submission to delete brand suggestion
	if r.Method == http.MethodPost {
		brandName := r.FormValue("brandName")
		if brandName != "" {
			err := collections.DeleteBrandSuggestions(brandName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	// Redirect back to admin page after deletion
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func AddSuggestedBrandHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is authenticated
	username, password, ok := r.BasicAuth()
	if !ok || username != adminUsername || password != adminPassword {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	// Process form submission to add brand suggestion
	if r.Method == http.MethodPost {
		brandName := r.FormValue("brandName")
		if brandName != "" {
			err := collections.AddBrand(brandName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = collections.DeleteBrandSuggestions(brandName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	// Redirect back to admin page after addition
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
