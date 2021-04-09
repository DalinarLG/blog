package routers

import (
	"encoding/json"
	"net/http"
	"text/template"
	"time"

	"github.com/DalinarLG/blog/db"
	"github.com/DalinarLG/blog/jwt"
	"github.com/DalinarLG/blog/models"
	"github.com/DalinarLG/blog/validations"
)



func LoginUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")	
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "email is required", 400)
		return
	}

	if !validations.ValidateEmail(user.Email) {
		http.Error(w, "not valid email", 400)
		return
	}

	if len(user.Password) == 0 || len(user.Password) <6 {
		http.Error(w, "password must have 6 characters", 400)
		return
	}

	user.Password = template.HTMLEscapeString(user.Password)

	result, status := db.LoginUser(user.Email, user.Password)
	if !status {
		http.Error(w, "email or password incorrect", 400)
		return
	}

	token, err := jwt.GenerateToken(result)
	if err != nil {
		http.Error(w, "error generating token "+err.Error(), 400)
		return
	}

	tklogin := models.Login{
		Token: token,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tklogin)

	http.SetCookie(w, &http.Cookie{
		Value: token,
		Name: "token",
		Expires: time.Now().Add(24*time.Hour),
	})

}