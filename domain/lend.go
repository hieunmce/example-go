package domain

import (
	"time"
)

// User describe user in systenm
type Lend struct {
	Model
	BookID UUID      `json:"book_id"`
	UserID UUID      `json:"user_id"`
	Name   string    `json:"name"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}
