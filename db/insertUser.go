package db

import (
	"log"

	"github.com/DalinarLG/blog/models")


func InsertUser(u models.User)(bool, error ){

	stmt, err := DB.Prepare("insert into user (name, email, password, avatar) values(?, ?, ?, ?) ")
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	defer stmt.Close()
	u.Password = Encrypt(u.Password)
	_, err = stmt.Exec(u.Name, u.Email, u.Password, u.Avatar)
	if err != nil {
		//log.Println(err.Error())
		return false, err
	}

	return true, nil
}