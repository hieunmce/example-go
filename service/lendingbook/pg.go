package lendingbook

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/ntp13495/example-go/domain"
)

// pgService implmenter for LendingBook serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for LendingBook service
func (s *pgService) Create(_ context.Context, p *domain.LendingBook) error {
	res := []domain.LendingBook{}
	s.db.Find(&res)
	for _, iterator := range res {
		if p.BookID == iterator.BookID {
			return ErrBookInUse
		}
	}

	return s.db.Create(p).Error
}

// Update implement Update for LendingBook service
func (s *pgService) Update(_ context.Context, p *domain.LendingBook) (*domain.LendingBook, error) {
	old := domain.LendingBook{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	resBook := []domain.Book{}
	s.db.Find(&resBook)
	flag := 0
	for _, iterator := range resBook {
		if p.BookID == iterator.ID {
			flag = 1
		}
	}
	if flag == 0 {
		return nil, ErrInvalidBook
	}

	resUser := []domain.User{}
	s.db.Find(&resUser)
	flag = 0
	for _, iterator := range resUser {
		if p.UserID == iterator.ID {
			flag = 1
		}
	}
	if flag == 0 {
		return nil, ErrInvalidUser
	}

	old.BookID = p.BookID
	old.UserID = p.UserID
	old.From = p.From
	old.To = p.To

	return &old, s.db.Save(&old).Error
}

// Find implement Find for LendingBook service
func (s *pgService) Find(_ context.Context, p *domain.LendingBook) (*domain.LendingBook, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for LendingBook service
func (s *pgService) FindAll(_ context.Context) ([]domain.LendingBook, error) {
	res := []domain.LendingBook{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for LendingBook service
func (s *pgService) Delete(_ context.Context, p *domain.LendingBook) error {
	old := domain.LendingBook{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}
