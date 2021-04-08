package db

import (
	"github.com/DalinarLG/blog/models"
	"golang.org/x/crypto/bcrypt"
)


func LoginUser(email, pass string)(models.User, bool){
	
	user, found, _ := CheckEmail(email)
	if !found {
		return user, false
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil{
		return user, false
	}

	return user, true
}