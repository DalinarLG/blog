package models

type Category struct {
	ID_Category int    `json:"id_category,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
