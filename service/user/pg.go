package user

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/minhkhiemm/example-go/domain"
)

// pgService implmenter for User serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for User service
func (s *pgService) Create(_ context.Context, p *domain.User) error {
	return s.db.Create(p).Error
}

// Update implement Update for User service
func (s *pgService) Update(_ context.Context, p *domain.User) (*domain.User, error) {
	old := domain.User{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	old.Name = p.Name
	old.Email = p.Email

	return &old, s.db.Save(&old).Error
}

// Find implement Find for User service
func (s *pgService) Find(_ context.Context, p *domain.User) (*domain.User, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for User service
func (s *pgService) FindAll(_ context.Context) ([]domain.User, error) {
	res := []domain.User{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for User service
func (s *pgService) Delete(_ context.Context, p *domain.User) error {
	old := domain.User{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}
