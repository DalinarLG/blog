package db

import (
	"log"
	"github.com/DalinarLG/blog/models"
)


func UpdateCategory(cat models.Category, id int)(bool, error){

	stmt, err := DB.Prepare("update category set name = ?, description=? where id_category =?")
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(cat.Name, cat.Description, id)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return true, nil
}