package category

import (
	"context"
	"fmt"

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
func (mw validationMiddleware) Create(ctx context.Context, category *domain.Category) (err error) {

	if category.Name == "" || len(category.Name) <= 5 {
		return ErrNameIsRequired
	}

	return mw.Service.Create(ctx, category)
}

func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.Category, error) {
	return mw.Service.FindAll(ctx)
}
func (mw validationMiddleware) Find(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	return mw.Service.Find(ctx, category)
}

func (mw validationMiddleware) Update(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	if category.Name == "" {
		return nil, ErrNameIsRequired
	}
	return mw.Service.Update(ctx, category)
}
func (mw validationMiddleware) Delete(ctx context.Context, category *domain.Category) error {
	fmt.Println("assa")
	return mw.Service.Delete(ctx, category)
}
