package user

import (
	"context"
	"example.com/m/domain"
	"regexp"
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

func (mw validationMiddleware) Create(ctx context.Context, user *domain.User) (err error) {

	if user.Name == "" {
		return ErrNameIsRequired
	}

	if user.Email == "" {
		return ErrEmailIsRequired
	}

	if user.Password == "" {
		return ErrPasswordIsRequired
	}

	emailRegexp, _ := regexp.Compile(emailRegex)
	if !emailRegexp.MatchString(user.Email) {
		return ErrEmailIsInvalid
	}
	return mw.Service.Create(ctx, user)
}
func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.User, error) {
	return mw.Service.FindAll(ctx)
}
func (mw validationMiddleware) Find(ctx context.Context, user *domain.User) (*domain.User, error) {
	return mw.Service.Find(ctx, user)
}

func (mw validationMiddleware) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	if user.Name == "" {
		return nil, ErrNameIsRequired
	}

	if user.Email == "" {
		return nil, ErrEmailIsRequired
	}

	emailRegexp, _ := regexp.Compile(emailRegex)
	if !emailRegexp.MatchString(user.Email) {
		return nil, ErrEmailIsInvalid
	}

	return mw.Service.Update(ctx, user)
}
func (mw validationMiddleware) Delete(ctx context.Context, user *domain.User) error {
	return mw.Service.Delete(ctx, user)
}
