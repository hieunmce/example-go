package loan

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/luantranminh/example-go/domain"
)

// pgService implmenter for Loan serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for Loan service
func (s *pgService) Create(_ context.Context, p *domain.Loan) error {
	books := []domain.Book{}
	users := []domain.User{}
	loans := []domain.Loan{}
	var userExisted = false
	var bookExisted = false

	s.db.Find(&users)
	s.db.Find(&books)
	s.db.Find(&loans)

	for _, element := range loans {
		if p.BookID == element.BookID {
			return ErrBookIsNotAvailable
		}
	}

	for _, element := range users {
		if p.UserID == element.ID {
			userExisted = true
			break
		}
	}

	for _, element := range books {
		if p.BookID == element.ID {
			bookExisted = true
			break
		}
	}

	if bookExisted == true && userExisted == true {
		return s.db.Create(p).Error
	}

	return ErrRecordNotFound
}

// Update implement Update for Loan service
func (s *pgService) Update(_ context.Context, p *domain.Loan) (*domain.Loan, error) {
	old := domain.Loan{Model: domain.Model{ID: p.ID}}

	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	old.UserID = p.UserID
	old.BookID = p.BookID
	old.To = p.To

	return &old, s.db.Save(&old).Error
}

// Find implement Find for Loan service
func (s *pgService) Find(_ context.Context, p *domain.Loan) (*domain.Loan, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for Loan service
func (s *pgService) FindAll(_ context.Context) ([]domain.Loan, error) {
	res := []domain.Loan{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Loan service
func (s *pgService) Delete(_ context.Context, p *domain.Loan) error {
	old := domain.Loan{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return err
	}

	return s.db.Delete(old).Error
}
