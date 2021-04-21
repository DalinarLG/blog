package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DalinarLG/blog/db"
)


func GetPostCategory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	id := r.URL.Query().Get("id_category")
	if len(id) == 0 {
		http.Error(w, "Category id is required", 400)
		return
	}

	cat_id, _ := strconv.Atoi(id)

	result, err := db.GetPostCategories(cat_id)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}