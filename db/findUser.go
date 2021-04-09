package db

import (
	"log"

	"github.com/DalinarLG/blog/models")


func FindUser(id int)(models.User, error){
	stmt, err := DB.Prepare("select * from user where id = ?")
	if err != nil {
		log.Println(err.Error())
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Avatar)
	if err != nil {
		log.Println(err.Error())
		return user, err
	}

	return user, nil
}