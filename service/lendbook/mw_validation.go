package lendbook

import (
	"context"

	"example.com/m/domain"
)

// Declare Regex
const (
	emailRegex = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
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

func (mw validationMiddleware) Create(ctx context.Context, lendbook *domain.LendBook) (err error) {

	return mw.Service.Create(ctx, lendbook)
}
func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.LendBook, error) {
	return mw.Service.FindAll(ctx)
}
func (mw validationMiddleware) Find(ctx context.Context, lendbook *domain.LendBook) (*domain.LendBook, error) {
	return mw.Service.Find(ctx, lendbook)
}

func (mw validationMiddleware) Update(ctx context.Context, lendbook *domain.LendBook) (*domain.LendBook, error) {

	return mw.Service.Update(ctx, lendbook)
}
func (mw validationMiddleware) Delete(ctx context.Context, lendbook *domain.LendBook) error {
	return mw.Service.Delete(ctx, lendbook)
}
