package domain

type Book struct {
	Model
	Category UUID
	name     string `json: "name"`
	author   string `json: "author"`
}
