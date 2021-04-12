package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DalinarLG/blog/db"
)


func GetCategories(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	
	cats, err := db.GetCategories()
	if err != nil {
		http.Error(w, "Error fetching categories: "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cats)
}