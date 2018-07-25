package book

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/phungvandat/example-go/domain"
)

// pgService implmenter for Book serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for Book service
func (s *pgService) Create(_ context.Context, p *domain.Book) error {
	return s.db.Create(p).Error
}

// Update implement Update for Book service
func (s *pgService) Update(_ context.Context, p *domain.Book) (*domain.Book, error) {
	old := domain.Book{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	if p.Name != "" {
		old.Name = p.Name
	}
	if !p.CategoryID.IsZero() {
		old.CategoryID = p.CategoryID
	}
	if p.Author != "" {
		old.Author = p.Author
	}
	if p.Description != "" {
		old.Description = p.Description
	}

	return &old, s.db.Save(&old).Error
}

// Find implement Find for Book service
func (s *pgService) Find(_ context.Context, p *domain.Book) (*domain.Book, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for Book service
func (s *pgService) FindAll(_ context.Context) ([]domain.Book, error) {
	res := []domain.Book{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Book service
func (s *pgService) Delete(_ context.Context, p *domain.Book) error {
	old := domain.Book{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}
