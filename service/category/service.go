package category

import (
	"context"

	"github.com/luantranminh/example-go/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.Category) error
	Update(ctx context.Context, p *domain.Category) (*domain.Category, error)
	Find(ctx context.Context, p *domain.Category) (*domain.Category, error)
	FindAll(ctx context.Context) ([]domain.Category, error)
	Delete(ctx context.Context, p *domain.Category) error
}
