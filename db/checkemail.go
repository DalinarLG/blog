package db

import (
	"log"

	"github.com/DalinarLG/blog/models")


func CheckEmail(email string)(models.User, bool, error){
	
	var user models.User

	stmt, err := DB.Prepare("select id, name, email, avatar from user where email=? ")
	if err != nil {
		log.Println("Error fetching data stmt "+err.Error())
		return user, false, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email,  &user.Avatar)
	if err != nil {
		log.Println("Error fetching data query "+err.Error())
		return user, false, err
	}

	

	return user, true, nil

}