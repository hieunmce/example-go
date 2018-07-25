package category

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/luantranminh/example-go/domain"
)

// pgService implmenter for Category serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for Category service
func (s *pgService) Create(_ context.Context, p *domain.Category) error {
	res := []domain.Category{}
	s.db.Find(&res)
	for _, element := range res {
		if p.Name == element.Name {
			return ErrRecordExisted
		}
	}

	return s.db.Create(p).Error
}

// Update implement Update for Category service
func (s *pgService) Update(_ context.Context, p *domain.Category) (*domain.Category, error) {
	old := domain.Category{Model: domain.Model{ID: p.ID}}
	res := []domain.Category{}
	s.db.Find(&res)

	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	for _, element := range res {
		if p.Name == element.Name {
			return nil, ErrRecordExisted
		}
	}

	old.Name = p.Name

	return &old, s.db.Save(&old).Error
}

// Find implement Find for Category service
func (s *pgService) Find(_ context.Context, p *domain.Category) (*domain.Category, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for Category service
func (s *pgService) FindAll(_ context.Context) ([]domain.Category, error) {
	res := []domain.Category{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Category service
func (s *pgService) Delete(_ context.Context, p *domain.Category) error {
	old := domain.Category{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return err
	}

	s.db.Where("category_id = ?", old.ID).Delete(domain.Book{})

	return s.db.Delete(old).Error
}
