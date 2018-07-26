package service

import (
	"github.com/neverdiefc/example-go/service/book"
	"github.com/neverdiefc/example-go/service/category"
	"github.com/neverdiefc/example-go/service/lendbook"
	"github.com/neverdiefc/example-go/service/user"
)

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
	BookService     book.Service
	LendbookService lendbook.Service
}
