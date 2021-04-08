package routers

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/DalinarLG/blog/db"
	"github.com/DalinarLG/blog/models"
)



func UpdateUser( w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error "+err.Error(), 400)
		return
	}

	if len(user.Name) == 0{
		http.Error(w, "name is required", 400)
		return
	}

	user.Name = template.HTMLEscapeString(user.Name)

	status, err := db.UpdateUser(user, UserID)
	if err != nil {
		http.Error(w, "could not update user "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "could not update user ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}