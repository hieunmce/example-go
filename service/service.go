package service

import "github.com/minhkhiemm/example-go/service/user"

// Service define list of all services in projects
type Service struct {
	UserService user.Service
}
