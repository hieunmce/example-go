package domain

// Book is type
type Book struct {
	Model
	Category_id UUID   `json: "category_id"`
	Name        string `json: "name"`
	Author      string `json: "author"`
	Description string `json: "description"`
}
