package domain

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Model base model for domain type
type Model struct {
	ID        UUID       `sql:",type:uuid" json:"id"`
	CreatedAt time.Time  `sql:"default:now()" json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// BeforeCreate prepare data before create data
func (m *Model) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}
