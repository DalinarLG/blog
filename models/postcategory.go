package models
//structure for sql rows
type PostCategory struct {
	ID_Category int    `json:"id_category"`
	Name        string `json:"name"`
	Posts       struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	
}
