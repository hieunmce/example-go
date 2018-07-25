package domain

// User describe book in systenm
type Book struct {
	Model
	Name        string `json:"name"`
	CategoryID  UUID   `sql:",type:uuid" json:"id"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
