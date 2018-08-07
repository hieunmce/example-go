package service

import (
	"github.com/luquehuong/example-go/service/book"
	"github.com/luquehuong/example-go/service/category"
	"github.com/luquehuong/example-go/service/lendBook"
	"github.com/luquehuong/example-go/service/user"
)

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
	BookService     book.Service
	LendBookService lendBook.Service
}
