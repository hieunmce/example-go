package user

import (
	"context"

	"github.com/trantrongkim98/example-go/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.User) error
	Update(ctx context.Context, p *domain.User) (*domain.User, error)
	Find(ctx context.Context, p *domain.User) (*domain.User, error)
	FindAll(ctx context.Context) ([]domain.User, error)
	Delete(ctx context.Context, p *domain.User) error
}
