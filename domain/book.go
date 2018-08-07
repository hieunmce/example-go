package domain

// User describe user in systenm
type Book struct {
	Model
	Name        string `json:"name"`
	Category_id UUID   `json:"category_id"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
