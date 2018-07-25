package domain

// Book is type
type Book struct {
	Model
	Category_id UUID
	Name        string `json: "Name"`
	Author      string `json: "Author"`
	Description string `json: "Description"`
}
