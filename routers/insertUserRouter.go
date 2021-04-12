package routers

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/DalinarLG/blog/db"
	"github.com/DalinarLG/blog/models"
	"github.com/DalinarLG/blog/validations"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Something went wrong "+err.Error(), 400)
		return
	}

	if len(user.Name) == 0 {
		http.Error(w, "Name is required", 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	if !validations.ValidateEmail(user.Email) {
		http.Error(w, "Not valid email", 400)
		return
	}

	if len(user.Password) == 0 || len(user.Password) < 6 {
		http.Error(w, "At least 6 characters", 400)
		return
	}

	//Escape validation

	user.Name = template.HTMLEscapeString(user.Name)
	user.Password = template.HTMLEscapeString(user.Password)

	_, found, _ := db.CheckEmail(user.Email)

	if found {
		http.Error(w, "Email already in use", 400)
		return
	}

	status, err := db.InsertUser(user)
	if !status {
		http.Error(w, "Could not insert user", 400)
		return
	}

	if err != nil {
		http.Error(w, "Error "+err.Error(), 400)
		return
	}

	err = validations.SendEmail(user.Email)
	if err != nil {
		http.Error(w, "Could not send email", 400)		
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
