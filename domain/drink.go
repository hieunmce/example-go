package domain

type Drink struct {
	Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
