package lendbook

import (
	"context"

	"github.com/jinzhu/gorm"

	"example.com/m/domain"
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
func (s *pgService) Create(_ context.Context, p *domain.LendBook) error {
	user := domain.User{Model: domain.Model{ID: p.User_id}}
	book := domain.Book{Model: domain.Model{ID: p.Book_id}}

	if err := s.db.Find(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

			return ErrRecordBookNotFound
		}
		return err
	}

	if err := s.db.Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

			return ErrRecordUserNotFound
		}
		return err
	}

	if err := s.db.Where("book_id = ?", p.Book_id).Find(&domain.LendBook{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return s.db.Create(p).Error
		}
		return err
	}
	return ErrBookIsBusy

}

// Update implement Update for Lendbook servicep
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

// Find implement Find for Lendbook service
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

// FindAll implement FindAll for Lendbook service
func (s *pgService) FindAll(_ context.Context) ([]domain.LendBook, error) {
	res := []domain.LendBook{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Lendbook service
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
