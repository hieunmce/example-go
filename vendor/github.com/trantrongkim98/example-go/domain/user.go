package domain

// User describe user in systenm
type User struct {
	Model
	Name  string `json:"name"`
	Email string `json:"email"`
}
