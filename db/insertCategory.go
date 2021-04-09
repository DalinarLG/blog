package db

import (
	"log"

	"github.com/DalinarLG/blog/models"
)

func InsertCategory(c models.Category) error {
	stmt, err := DB.Prepare("insert into category (name, description) values (?,?)")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.Name, c.Description)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
