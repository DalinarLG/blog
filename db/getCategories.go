package db

import (
	"log"

	"github.com/DalinarLG/blog/models")


func GetCategories()([]models.Category, error){
	var cats []models.Category
	stmt, err := DB.Prepare("select * from category")
	if err != nil {
		log.Println(err.Error())
		return cats, err
	}
	defer stmt.Close()

	result, err := stmt.Query()
	if err != nil {
		log.Println(err.Error())
		return cats, err
	}

	for result.Next(){
		var cat models.Category
		err := result.Scan(&cat.ID_Category, &cat.Name, &cat.Description)
		if err != nil {
			log.Println(err.Error())
			return cats, err
		}

		cats = append(cats, cat)
	}

	return cats, nil
}