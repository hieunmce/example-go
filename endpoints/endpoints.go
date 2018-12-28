package endpoints

import (
	"github.com/go-kit/kit/endpoint"

	"github.com/minhkhiemm/example-go/endpoints/account"
	"github.com/minhkhiemm/example-go/endpoints/order"
	"github.com/minhkhiemm/example-go/service"
)

// Endpoints .
type Endpoints struct {
	// Orders
	GetAllOrderByDate endpoint.Endpoint
	CreateOrder       endpoint.Endpoint
	GetOrderByID      endpoint.Endpoint
	// Accounts
	CreateAccount endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		// Orders
		GetAllOrderByDate: order.GetAllByDate(s),
		CreateOrder:       order.Create(s),
		GetOrderByID:      order.Get(s),

		// Accounts
		CreateAccount: account.Create(s),
	}
}
