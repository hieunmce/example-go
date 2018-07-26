package domain

import "time"

// User describe user lend a book in system

type Lendbook struct {
	Model
	BookID UUID      `json:"book_id"`
	UserID UUID      `json:"user_id"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}
