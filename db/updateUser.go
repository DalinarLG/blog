package db

import (
	"log"

	"github.com/DalinarLG/blog/models")


func UpdateUser(u models.User, id int)(bool, error){
	stmt, err := DB.Prepare("update user set name = Coalesce(?,name), avatar=Coalesce(?,avatar) where id=?" )
	if err != nil {
		log.Println("Error stmt "+err.Error())
		return false, err

	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Name, u.Avatar, id)
	if err != nil {
		log.Println("Error query "+err.Error())
		return false, err
	}

	return true, nil
}