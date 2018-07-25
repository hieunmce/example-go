package domain

import "time"

// Loan describe action: user lend a book in systenm
type Loan struct {
	Model
	BookID UUID       `json:"book_id"`
	UserID UUID       `json:"user_id"`
	To     *time.Time `json:"to,omitempty"`
}
