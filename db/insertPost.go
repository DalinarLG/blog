package db

import (
	"log"
	"github.com/DalinarLG/blog/models"
)


func InsertPost(post models.Post, cat_id int)(bool, error){
	stmt, err := DB.Prepare("insert into post (title, body, post_user_id, post_category_id) values(?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(post.Title, post.Body, post.Post_user_id, cat_id)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return true, err

}