package loan

import (
	"context"

	"github.com/luantranminh/example-go/domain"
)

type validationMiddleware struct {
	Service
}

// ValidationMiddleware ...
func ValidationMiddleware() func(Service) Service {
	return func(next Service) Service {
		return &validationMiddleware{
			Service: next,
		}
	}
}

func (mw validationMiddleware) Create(ctx context.Context, loan *domain.Loan) (err error) {
	return mw.Service.Create(ctx, loan)
}
func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.Loan, error) {
	return mw.Service.FindAll(ctx)
}
func (mw validationMiddleware) Find(ctx context.Context, loan *domain.Loan) (*domain.Loan, error) {
	return mw.Service.Find(ctx, loan)
}

func (mw validationMiddleware) Update(ctx context.Context, loan *domain.Loan) (*domain.Loan, error) {
	return mw.Service.Update(ctx, loan)
}
func (mw validationMiddleware) Delete(ctx context.Context, loan *domain.Loan) error {
	return mw.Service.Delete(ctx, loan)
}
