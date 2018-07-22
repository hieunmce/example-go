package service

import "github.com/hieunmce/example-go/service/user"

// Service define list of all services in projects
type Service struct {
	UserService user.Service
}
