package lendBook

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/luquehuong/example-go/domain"
)

// pgService implmenter for LendBook serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for UserLendBook service
func (s *pgService) Create(_ context.Context, p *domain.LendBook) error {

	// Validate book_id is exist, if not reject it with error message
	bookID := domain.Book{Model: domain.Model{ID: p.Book_id}}
	if err := s.db.Find(&bookID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrBookIDNotFound
		}
		return err
	}

	// Validate user_id is exist, if not reject it with error message
	userID := domain.User{Model: domain.Model{ID: p.User_id}}
	if err := s.db.Find(&userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrUserIDNotFound
		}
		return err
	}

	// If pass validate, create a new user-lend-book
	return s.db.Create(p).Error
}

// Update implement Update for LendBook service
func (s *pgService) Update(_ context.Context, p *domain.LendBook) (*domain.LendBook, error) {
	old := domain.LendBook{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	old.Book_id = p.Book_id
	old.User_id = p.User_id
	old.From = p.From
	old.To = p.To

	return &old, s.db.Save(&old).Error
}

// Find implement Find for LendBook service
func (s *pgService) Find(_ context.Context, p *domain.LendBook) (*domain.LendBook, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for LendBook service
func (s *pgService) FindAll(_ context.Context) ([]domain.LendBook, error) {
	res := []domain.LendBook{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for LendBook service
func (s *pgService) Delete(_ context.Context, p *domain.LendBook) error {
	old := domain.LendBook{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}