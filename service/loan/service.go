package loan

import (
	"context"

	"github.com/luantranminh/example-go/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.Loan) error
	Update(ctx context.Context, p *domain.Loan) (*domain.Loan, error)
	Find(ctx context.Context, p *domain.Loan) (*domain.Loan, error)
	FindAll(ctx context.Context) ([]domain.Loan, error)
	Delete(ctx context.Context, p *domain.Loan) error
}
