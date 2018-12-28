package order

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/minhkhiemm/example-go/domain"
)

type pgService struct {
	db *gorm.DB
}

func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

func (s *pgService) GetAllByDate(_ context.Context, date domain.OrderDate) ([]*domain.Order, error) {
	orders := []*domain.Order{}
	query := s.db
	if date.Date != 0 {
		query = query.Where("date_part('day', created_at) = ?", date.Date)
	}

	if date.Month != 0 {
		query = query.Where("date_part('month', created_at) = ?", date.Month)
	}

	if date.Year != 0 {
		query = query.Where("date_part('year', created_at) = ?", date.Year)
	}

	return orders, query.Find(&orders).Error
}

func (s *pgService) Create(_ context.Context, order *domain.Order) (*domain.Order, error) {
	return order, s.db.Create(&order).Error
}

func (s *pgService) Get(_ context.Context, id domain.UUID) (*domain.Order, error) {
	order := domain.Order{}
	return &order, s.db.Where("id = ?", id).First(&order).Error
}

func (s *pgService) Update(_ context.Context, order *domain.Order) (*domain.Order, error) {
	return order, s.db.Save(&order).Error
}
