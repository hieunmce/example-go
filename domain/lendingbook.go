package domain

import (
	"time"
)

// LendingBook decribes book in system
type LendingBook struct {
	Model
	BookID UUID      `json:"book_id"`
	UserID UUID      `json:"user_id"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}
