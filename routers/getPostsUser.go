package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DalinarLG/blog/db"
)

func GetPostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(w, "Id is required", 400)
		return
	}
	uid, _ := strconv.Atoi(id)
	posts, err := db.GetPostUser(uid)
	if err != nil {
		http.Error(w, "Error fetching posts "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(posts)

}
