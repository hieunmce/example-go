package lendbook

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/neverdiefc/example-go/domain"
)

// pgService implmenter for Lendbook serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for Lendbook service
func (s *pgService) Create(_ context.Context, p *domain.Lendbook) error {

	return s.db.Create(p).Error
}

// Update implement Update for Lendbook service
func (s *pgService) Update(_ context.Context, p *domain.Lendbook) (*domain.Lendbook, error) {

	old := domain.Lendbook{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	old.BookID = p.BookID
	old.UserID = p.UserID
	old.From = p.From
	old.To = p.To

	return &old, s.db.Save(&old).Error

}

// Find implement Find for Lendbook service
func (s *pgService) Find(_ context.Context, p *domain.Lendbook) (*domain.Lendbook, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for Lendbook service
func (s *pgService) FindAll(_ context.Context) ([]domain.Lendbook, error) {
	res := []domain.Lendbook{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Lendbook service
func (s *pgService) Delete(_ context.Context, p *domain.Lendbook) error {
	old := domain.Lendbook{Model: domain.Model{ID: p.ID}}

	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}

	if err := s.db.Where("lendbook_id = ?", p.ID).Delete(&domain.Book{}).Error; err != nil {
		return err
	}

	return s.db.Delete(old).Error
}
