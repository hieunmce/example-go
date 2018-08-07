package lendBook

import (
	"context"
	"fmt"
	"time"

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

func (mw validationMiddleware) Create(ctx context.Context, lendBook *domain.LendBook) (err error) {

	//validate books is available to lend, if not available reject with error message
	now := time.Now()
	timeReturn := lendBook.To

	if now.Before(timeReturn) {
		fmt.Println("from,", now)
		fmt.Println("to", timeReturn)
		return ErrBookIsNotAvailable
	}

	// If pass validate, accept create new lendBook.
	return mw.Service.Create(ctx, lendBook)
}
func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.LendBook, error) {
	return mw.Service.FindAll(ctx)
}
func (mw validationMiddleware) Find(ctx context.Context, lendBook *domain.LendBook) (*domain.LendBook, error) {
	return mw.Service.Find(ctx, lendBook)
}

func (mw validationMiddleware) Update(ctx context.Context, lendBook *domain.LendBook) (*domain.LendBook, error) {

	// If pass validate, accept update the lendBook.
	return mw.Service.Update(ctx, lendBook)
}
func (mw validationMiddleware) Delete(ctx context.Context, lendBook *domain.LendBook) error {
	return mw.Service.Delete(ctx, lendBook)
}
