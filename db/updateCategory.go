package db

import (
	"log"
	"github.com/DalinarLG/blog/models"
)


func UpdateCategory(cat models.Category)(bool, error){

	stmt, err := DB.Prepare("update category set name = ?, description=? where id =?")
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(cat.Name, cat.Description, cat.ID_Category)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return true, nil
}