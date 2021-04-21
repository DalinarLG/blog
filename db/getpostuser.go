package db

import (
	"log"

	"github.com/DalinarLG/blog/models")



func GetPostUser(id int)([]models.Post, error){
	var posts []models.Post
	stmt, err := DB.Prepare("SELECT id_post, title, body, post_category_id FROM post WHERE post_user_id = ?")
	if err != nil {
		log.Println(err.Error())
		return posts, err
	}
	defer stmt.Close()

	result, err := stmt.Query(id)
	if err != nil {
		log.Println(err.Error())
		return posts, err
	}

	for result.Next(){
		var post models.Post
		err := result.Scan(&post.ID_Post, &post.Title, &post.Body, &post.Post_category_id)
		if err != nil {
			log.Println(err.Error())

		}

		posts = append(posts, post)
	}

	return posts, nil
}