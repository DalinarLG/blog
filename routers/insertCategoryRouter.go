package routers

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/DalinarLG/blog/db"
	"github.com/DalinarLG/blog/models"
)


func InsertCategory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	var cat models.Category
	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		http.Error(w, "Error "+err.Error(), 400)
		return
	}

	if len(cat.Name)==0{
		http.Error(w, "Name is required",400)
		return
	}

	if len(cat.Description)==0{
		http.Error(w, "Description is required",400)
		return
	}

	cat.Name = template.HTMLEscapeString(cat.Name)
	cat.Description = template.HTMLEscapeString(cat.Description)

	err = db.InsertCategory(cat)
	if err != nil {
		http.Error(w, "Error inserting category "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}