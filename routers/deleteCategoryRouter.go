package routers

import (
	"net/http"
	"strconv"

	"github.com/DalinarLG/blog/db"
)

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		http.Error(w, "id is required", 400)
		return
	}

	idc, _ := strconv.Atoi(id)

	err := db.DeleteCategory(idc)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
