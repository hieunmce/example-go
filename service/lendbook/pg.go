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

	user := domain.User{Model: domain.Model{ID: p.UserID}}
	book := domain.Book{Model: domain.Model{ID: p.BookID}}

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

	if err := s.db.Where("book_id = ?", p.BookID).Find(&domain.Lendbook{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return s.db.Create(p).Error
		}
		return err
	}
	return ErrBookIsBusy

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

	user := domain.User{Model: domain.Model{ID: p.UserID}}
	book := domain.Book{Model: domain.Model{ID: p.BookID}}

	if err := s.db.Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRecordUserNotFound
		}
		return nil, err
	}

	if err := s.db.Find(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRecordBookNotFound
		}
		return nil, err
	}

	if err := s.db.Where("book_id = ?", p.BookID).Find(&domain.Lendbook{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			old.BookID = p.BookID
			old.UserID = p.UserID
			old.From = p.From
			old.To = p.To
			return &old, s.db.Save(&old).Error
		}
		return nil, err
	}
	return nil, ErrBookIsBusy

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
	return s.db.Delete(old).Error
}
