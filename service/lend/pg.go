package lend

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/minhkhiemm/example-go/domain"
)

// pgService implmenter for Lend serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for Lend service
func (s *pgService) Create(_ context.Context, p *domain.Lend) error {
	return s.db.Create(p).Error
}

// Update implement Update for Lend service
func (s *pgService) Update(_ context.Context, p *domain.Lend) (*domain.Lend, error) {
	old := domain.Lend{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	old.Name = p.Name
	old.From = p.From
	old.To = p.To

	return &old, s.db.Save(&old).Error
}

// Find implement Find for Lend service
func (s *pgService) Find(_ context.Context, p *domain.Lend) (*domain.Lend, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for Lend service
func (s *pgService) FindAll(_ context.Context) ([]domain.Lend, error) {
	res := []domain.Lend{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Lend service
func (s *pgService) Delete(_ context.Context, p *domain.Lend) error {
	old := domain.Lend{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return err
	}
	return s.db.Delete(old).Error
}
