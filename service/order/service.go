package order

import (
	"context"

	"github.com/minhkhiemm/example-go/domain"
)

type Service interface {
	GetAllByDate(ctx context.Context, date domain.OrderDate) ([]*domain.Order, error)
	Create(ctx context.Context, order *domain.Order) (*domain.Order, error)
	Get(ctx context.Context, id domain.UUID) (*domain.Order, error)
	Update(ctx context.Context, order *domain.Order) (*domain.Order, error)
}
