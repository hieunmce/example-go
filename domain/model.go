package domain

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/k0kubun/pp"
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

// AfterDelete delete books after delete category
func (category *Category) AfterDelete(tx *gorm.DB) error {

	pp.Print("After Delete")
	fmt.Print("After Delete")

	return nil
}
