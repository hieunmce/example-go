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
	res := []domain.Loan{}
	s.db.Find(&res)
	for _, element := range res {
		if p.UserID == element.UserID {
			return ErrRecordExisted
		}
	}

	return s.db.Create(p).Error
}

// Update implement Update for Loan service
func (s *pgService) Update(_ context.Context, p *domain.Loan) (*domain.Loan, error) {
	old := domain.Loan{Model: domain.Model{ID: p.ID}}
	res := []domain.Loan{}
	s.db.Find(&res)

	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	for _, element := range res {
		if p.UserID == element.UserID {
			return nil, ErrRecordExisted
		}
	}

	old.UserID = p.UserID

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

	s.db.Where("loan_id = ?", old.ID).Delete(domain.Book{})

	return s.db.Delete(old).Error
}
