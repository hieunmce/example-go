package book

import (
	"context"

	"github.com/luquehuong/example-go/domain"
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

func (mw validationMiddleware) Create(ctx context.Context, book *domain.Book) (err error) {

	// Validate name of a book is not empty and length > 5 characters. if not reject it with error message
	if book.Name == "" {
		return ErrNameIsRequired
	}
	if len(book.Name) <= 5 {
		return ErrNameIsToShort
	}

	// Validate description of a book is not empty and length > 5 characters. if not reject it with error message
	if book.Description == "" {
		return ErrDescriptionIsRequired
	}
	if len(book.Description) <= 5 {
		return ErrDescriptionIsToShort
	}

	// If pass validate, accept create new book.
	return mw.Service.Create(ctx, book)
}
func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.Book, error) {
	return mw.Service.FindAll(ctx)
}
func (mw validationMiddleware) Find(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	return mw.Service.Find(ctx, book)
}

func (mw validationMiddleware) Update(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	// Validate name of a book is not empty and length > 5 characters. if not reject it with error message
	if book.Name == "" {
		return nil, ErrNameIsRequired
	}
	if len(book.Name) <= 5 {
		return nil, ErrNameIsToShort
	}

	// Validate description of a book is not empty and length > 5 characters. if not reject it with error message
	if book.Description == "" {
		return nil, ErrDescriptionIsRequired
	}
	if len(book.Description) <= 5 {
		return nil, ErrDescriptionIsToShort
	}

	// If pass validate, accept update the book.
	return mw.Service.Update(ctx, book)
}
func (mw validationMiddleware) Delete(ctx context.Context, book *domain.Book) error {
	return mw.Service.Delete(ctx, book)
}
