package service

import (
	"github.com/luantranminh/example-go/service/book"
	"github.com/luantranminh/example-go/service/category"
	"github.com/luantranminh/example-go/service/loan"
	"github.com/luantranminh/example-go/service/user"
)

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
	BookService     book.Service
	LoanService     loan.Service
}
