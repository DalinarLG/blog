package db

import (
	"log"

	"github.com/DalinarLG/blog/models"
)

func UpdateUser(u models.User, id int) (bool, error) {
	user, _ := FindUser(id)
	stmt, err := DB.Prepare("update user set name = ?,  avatar = ? where id = ?")
	if err != nil {
		log.Println("Error stmt " + err.Error())
		return false, err

	}
	defer stmt.Close()
	if len(u.Name) == 0 {
		u.Name = user.Name
	}
	if len(u.Avatar) == 0{
		u.Avatar = user.Avatar
	}
	_, err = stmt.Exec(u.Name, u.Avatar, id)
	if err != nil {
		log.Println("Error query " + err.Error())
		return false, err
	}

	return true, nil
}
