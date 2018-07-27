package book

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/ntp13495/example-go/domain"
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
	resBook := []domain.Book{}
	s.db.Find(&resBook)
	for _, iterator := range resBook {
		if p.Name == iterator.Name {
			return ErrRecordExisted
		}
	}
	resCategory := []domain.Category{}
	s.db.Find(&resCategory)
	flag := 0
	for _, iterator := range resCategory {
		if p.CategoryID == iterator.ID {
			flag = 1
		}
	}
	if flag == 0 {
		return ErrInvalidCategory
	}
	if p.Name == "" {
		return ErrBookNameIsRequired
	}
	if len(p.Name) <= 5 {
		return ErrBookNameLengthIsRequired
	}

	if p.Description == "" {
		return ErrDescriptionIsRequired
	}
	if len(p.Description) <= 5 {
		return ErrDescriptionLengthIsRequired
	}
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

	resBook := []domain.Book{}
	s.db.Find(&resBook)
	for _, iterator := range resBook {
		if p.Name == iterator.Name {
			return nil, ErrRecordExisted
		}
	}
	resCategory := []domain.Category{}
	s.db.Find(&resCategory)
	flag := 0
	for _, iterator := range resCategory {
		if p.CategoryID == iterator.ID {
			flag = 1
		}
	}
	if flag == 0 {
		return nil, ErrInvalidCategory
	}
	if p.Name == "" {
		return nil, ErrBookNameIsRequired
	}
	if len(p.Name) <= 5 {
		return nil, ErrBookNameLengthIsRequired
	}

	if p.Description == "" {
		return nil, ErrDescriptionIsRequired
	}
	if len(p.Description) <= 5 {
		return nil, ErrDescriptionLengthIsRequired
	}

	old.Name = p.Name
	old.CategoryID = p.CategoryID
	old.Author = p.Author
	old.Description = p.Description

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
