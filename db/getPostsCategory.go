package db

import (
	"log"

	"github.com/DalinarLG/blog/models"
)

func GetPostCategories(id int) (models.ResponsePostCategory, error) {
	var posts []models.PostCategory //structure
	var catres models.ResponsePostCategory

	stmt, err := DB.Prepare("SELECT c.id_category, c.name, p.id_post, p.title, p.body FROM category as c INNER JOIN post as p on p.post_category_id = c.id_category WHERE c.id_category = ? ")
	if err != nil {
		log.Println(err.Error())
		return catres, err
	}
	defer stmt.Close()

	result, err := stmt.Query(id)
	if err != nil {
		log.Println(err.Error())
		return catres, err
	}

	for result.Next() {
		var post models.PostCategory
		err := result.Scan(&post.ID_Category, &post.Name, &post.Posts.ID, &post.Posts.Title, &post.Posts.Body)

		if err != nil {
			log.Println(err.Error())
			return catres, err
		}

		posts = append(posts, post)
	}

	var postscat []models.Post

	for _, p := range posts {
		var post models.Post
		catres.ID_Category = id
		catres.Name = p.Name

		post.ID_Post = p.Posts.ID
		post.Title = p.Posts.Title
		post.Body = p.Posts.Body

		postscat = append(postscat, post)
		catres.Posts = postscat

	}

	return catres, nil

}
