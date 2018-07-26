package lendbook

import (
	"context"

	"github.com/neverdiefc/example-go/domain"
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

func (mw validationMiddleware) Create(ctx context.Context, lendbook *domain.Lendbook) (err error) {

	return mw.Service.Create(ctx, lendbook)
}
func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.Lendbook, error) {
	return mw.Service.FindAll(ctx)
}
func (mw validationMiddleware) Find(ctx context.Context, lendbook *domain.Lendbook) (*domain.Lendbook, error) {
	return mw.Service.Find(ctx, lendbook)
}

func (mw validationMiddleware) Update(ctx context.Context, lendbook *domain.Lendbook) (*domain.Lendbook, error) {

	return mw.Service.Update(ctx, lendbook)
}
func (mw validationMiddleware) Delete(ctx context.Context, lendbook *domain.Lendbook) error {
	return mw.Service.Delete(ctx, lendbook)
}
