package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"

	"github.com/DalinarLG/blog/db"
	"github.com/DalinarLG/blog/models"
)


func InsertPost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Error: "+ err.Error(), 400)
		return
	}
	cat_id := r.URL.Query().Get("id_category")
	id_cat,_ := strconv.Atoi(cat_id) 

	if len(post.Title) == 0 {
		http.Error(w, "Title is required ", 400)
		return
	}
	if len(post.Body) == 0 {
		http.Error(w, "Body is required ", 400)
		return
	}

	post.Title = template.HTMLEscapeString(post.Title)
	post.Body = template.HTMLEscapeString(post.Body)
	post.Post_user_id = UserID

	status, err := db.InsertPost(post, id_cat)
	if !status {
		http.Error(w, "Could not register ", 400)
		return
	}

	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}