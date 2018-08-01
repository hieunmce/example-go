package category

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

func (mw validationMiddleware) Create(ctx context.Context, category *domain.Category) (err error) {
	if category.Name == "" {
		return ErrNameIsRequired
	}
	//check for name is lower than 5 character
	if len(category.Name) < 6 {
		return ErrNameIsInvalid
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
		return nil, nil
	}
	//check for name is lower than 5 character
	if len(category.Name) < 6 {
		return nil, ErrNameIsInvalid
	}

	return mw.Service.Update(ctx, category)
}
func (mw validationMiddleware) Delete(ctx context.Context, category *domain.Category) error {
	return mw.Service.Delete(ctx, category)
}
