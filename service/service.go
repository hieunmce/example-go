package service

import (
	"github.com/phungvandat/example-go/service/book"
	"github.com/phungvandat/example-go/service/category"
	"github.com/phungvandat/example-go/service/user"
)

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
	BookService     book.Service
}
