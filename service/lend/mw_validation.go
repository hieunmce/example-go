package lend

import (
	"context"

	"github.com/minhkhiemm/example-go/domain"
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

func (mw validationMiddleware) Create(ctx context.Context, lend *domain.Lend) (err error) {
	if lend.Name == "" {
		return ErrNameIsRequired
	}

	return mw.Service.Create(ctx, lend)
}
func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.Lend, error) {
	return mw.Service.FindAll(ctx)
}
func (mw validationMiddleware) Find(ctx context.Context, lend *domain.Lend) (*domain.Lend, error) {
	return mw.Service.Find(ctx, lend)
}

func (mw validationMiddleware) Update(ctx context.Context, lend *domain.Lend) (*domain.Lend, error) {
	if lend.Name == "" {
		return nil, ErrNameIsRequired
	}

	return mw.Service.Update(ctx, lend)
}
func (mw validationMiddleware) Delete(ctx context.Context, lend *domain.Lend) error {
	return mw.Service.Delete(ctx, lend)
}
