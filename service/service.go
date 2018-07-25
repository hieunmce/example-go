package service

import "github.com/luantranminh/example-go/service/user"
import "github.com/luantranminh/example-go/service/category"
import "github.com/luantranminh/example-go/service/book"

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
	BookService     book.Service
}
