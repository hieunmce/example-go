package domain

// Book describe Book in systenm
type Book struct {
	Model
	Name        string `json:"name"`
	CategoryID  UUID   `json:"category_id"`
	Author      string `json:"author"`
	Descripsion string `json:"descripsion"`
}
