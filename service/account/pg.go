package account

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/minhkhiemm/example-go/domain"
)

type pgService struct {
	db *gorm.DB
}

func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

func (s *pgService) Create(_ context.Context, account *domain.Account) (*domain.Account, error) {
	return account, s.db.Create(&account).Error
}
