package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/minhkhiemm/example-go/service"

	"github.com/minhkhiemm/example-go/endpoints/user"
)

// Endpoints .
type Endpoints struct {
	FindUser    endpoint.Endpoint
	FindAllUser endpoint.Endpoint
	CreateUser  endpoint.Endpoint
	UpdateUser  endpoint.Endpoint
	DeleteUser  endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		FindUser:    user.MakeFindEndPoint(s),
		FindAllUser: user.MakeFindAllEndpoint(s),
		CreateUser:  user.MakeCreateEndpoint(s),
		UpdateUser:  user.MakeUpdateEndpoint(s),
		DeleteUser:  user.MakeDeleteEndpoint(s),
	}
}
