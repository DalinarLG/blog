package db

import "log"

func DeleteCategory(id int) error {
	stmt, err := DB.Prepare("delete from category where id_category =?")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err

}
