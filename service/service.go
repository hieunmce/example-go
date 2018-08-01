package service

import (
	"github.com/minhkhiemm/example-go/service/book"
	"github.com/minhkhiemm/example-go/service/category"
	"github.com/minhkhiemm/example-go/service/user"
)

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
	BookService     book.Service
}
