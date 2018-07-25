package service

import "github.com/luantranminh/example-go/service/user"
import "github.com/luantranminh/example-go/service/category"

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
}
