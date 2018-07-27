package lendingbook

import (
	"context"

	"github.com/ntp13495/example-go/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.LendingBook) error
	Update(ctx context.Context, p *domain.LendingBook) (*domain.LendingBook, error)
	Find(ctx context.Context, p *domain.LendingBook) (*domain.LendingBook, error)
	FindAll(ctx context.Context) ([]domain.LendingBook, error)
	Delete(ctx context.Context, p *domain.LendingBook) error
}
