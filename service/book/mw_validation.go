package book

import (
	"context"

	"github.com/phungvandat/example-go/domain"
)

// Declare Regex

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

func (mw validationMiddleware) Create(ctx context.Context, book *domain.Book) (err error) {
	if book.Name == "" {
		return ErrNameIsRequired
	}
	if book.Description == "" {
		return ErrDescriptionIsRequired
	}
	if len(book.Name) < 5 {
		return ErrMinimumLengthName
	}
	if len(book.Description) < 5 {
		return ErrMinimumLengthDescription
	}
	if book.CategoryID.IsZero() {
		return ErrCategoryIDIsRequired
	}
	return mw.Service.Create(ctx, book)
}
func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.Book, error) {
	return mw.Service.FindAll(ctx)
}
func (mw validationMiddleware) Find(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	return mw.Service.Find(ctx, book)
}

func (mw validationMiddleware) Update(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	if book.Name == "" {
		return nil, ErrNameIsRequired
	}
	if len(book.Name) < 5 {
		return nil, ErrMinimumLengthName
	}
	if book.Description == "" {
		return nil, ErrDescriptionIsRequired
	}
	if len(book.Description) < 5 {
		return nil, ErrMinimumLengthDescription
	}

	return mw.Service.Update(ctx, book)
}
func (mw validationMiddleware) Delete(ctx context.Context, book *domain.Book) error {
	return mw.Service.Delete(ctx, book)
}
