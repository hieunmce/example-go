package service

import (
	"github.com/ntp13495/example-go/service/book"
	"github.com/ntp13495/example-go/service/category"
	"github.com/ntp13495/example-go/service/user"
)

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
	BookService     book.Service
}
