package routers

import (
	"encoding/json"
	"net/http"	
	"text/template"

	"github.com/DalinarLG/blog/db"
	"github.com/DalinarLG/blog/models"
)



func UpdateCategory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	var cat models.Category
	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		http.Error(w, "Error "+err.Error(), 400)
		return
	}

	cat.Name = template.HTMLEscapeString(cat.Name)
	cat.Description = template.HTMLEscapeString(cat.Description)

	status, err := db.UpdateCategory(cat)
	if err != nil {
		http.Error(w, "Error updating category "+err.Error(), 400 )
		return
	}
	if !status {
		http.Error(w, "Error updating category", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}