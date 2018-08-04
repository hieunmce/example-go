package domain

import (
	"time"
)

// User describe user in systenm
type LendBook struct {
	Model
	Book_id UUID      `json:"book_id"`
	User_id UUID      `json:"user_id"`
	From    time.Time `json:"from"`
	To      time.Time `json:"to"`
}
