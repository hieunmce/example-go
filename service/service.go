package service

import (
	"example.com/m/service/book"
	"example.com/m/service/category"
	"example.com/m/service/lendbook"
	"example.com/m/service/user"
)

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
	BookService     book.Service
	LendBookService lendbook.Service
}
