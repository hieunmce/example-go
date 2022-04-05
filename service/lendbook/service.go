package lendbook

import (
	"context"

	"example.com/m/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.LendBook) error
	Update(ctx context.Context, p *domain.LendBook) (*domain.LendBook, error)
	Find(ctx context.Context, p *domain.LendBook) (*domain.LendBook, error)
	FindAll(ctx context.Context) ([]domain.LendBook, error)
	Delete(ctx context.Context, p *domain.LendBook) error
}
