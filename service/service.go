package service

import "github.com/luquehuong/example-go/service/user"

// Service define list of all services in projects
type Service struct {
	UserService user.Service
}
