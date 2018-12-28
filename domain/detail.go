package domain

type Detail struct {
	Model
	Quantity int  `json:"quantity"`
	DrinkID  UUID `json:"drink_id"`
}
