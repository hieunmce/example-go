package domain

// Book decribes book in system

type LendingABook struct {
	Model
	BookID UUID   `json:"book_id"`
	UserID UUID   `json:"user_id"`
	From   string `json:"from"`
	To     string `json:"to"`
}
