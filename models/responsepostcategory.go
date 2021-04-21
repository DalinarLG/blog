package models

//Response category with array of posts
type ResponsePostCategory struct {
	ID_Category int    `json:"id_category"`
	Name        string `json:"name"`
	Posts       []Post 
}
