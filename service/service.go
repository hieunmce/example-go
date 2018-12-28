package service

import (
	"github.com/minhkhiemm/example-go/service/account"
	"github.com/minhkhiemm/example-go/service/order"
)

// Service define list of all services in projects
type Service struct {
	OrderService   order.Service
	AccountService account.Service
}
