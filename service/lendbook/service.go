package lendbook

import (
	"context"

	"github.com/neverdiefc/example-go/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.Lendbook) error
	Update(ctx context.Context, p *domain.Lendbook) (*domain.Lendbook, error)
	Find(ctx context.Context, p *domain.Lendbook) (*domain.Lendbook, error)
	FindAll(ctx context.Context) ([]domain.Lendbook, error)
	Delete(ctx context.Context, p *domain.Lendbook) error
}
