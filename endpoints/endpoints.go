package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/phungvandat/example-go/service"

	"github.com/phungvandat/example-go/endpoints/category"
	"github.com/phungvandat/example-go/endpoints/user"
)

// Endpoints .
type Endpoints struct {
	FindCategory    endpoint.Endpoint
	FindAllCategory endpoint.Endpoint
	CreateCategory  endpoint.Endpoint
	UpdateCategory  endpoint.Endpoint
	DeleteCategory  endpoint.Endpoint

	FindUser    endpoint.Endpoint
	FindAllUser endpoint.Endpoint
	CreateUser  endpoint.Endpoint
	UpdateUser  endpoint.Endpoint
	DeleteUser  endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		FindCategory:    category.MakeFindEndPoint(s),
		FindAllCategory: category.MakeFindAllEndpoint(s),
		CreateCategory:  category.MakeCreateEndpoint(s),
		UpdateCategory:  category.MakeUpdateEndpoint(s),
		DeleteCategory:  category.MakeDeleteEndpoint(s),

		FindUser:    user.MakeFindEndPoint(s),
		FindAllUser: user.MakeFindAllEndpoint(s),
		CreateUser:  user.MakeCreateEndpoint(s),
		UpdateUser:  user.MakeUpdateEndpoint(s),
		DeleteUser:  user.MakeDeleteEndpoint(s),
	}
}
