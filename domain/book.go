package domain

//struct Book
type Book struct {
	Model
	CategoryID  UUID   `json:"category_id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
