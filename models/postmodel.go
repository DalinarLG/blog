package models

type Post struct {
	ID_Post          int    `json:"id_post,omitempty"`
	Title            string `json:"title,omitempty"`
	Body             string `json:"body,omitempty"`
	Post_user_id     int    `json:"post_user_id,omitempty"`
	Post_category_id int    `json:"post_category_id,omitempty"`
}
