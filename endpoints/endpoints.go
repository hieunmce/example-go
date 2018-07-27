package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/ntp13495/example-go/service"

	book "github.com/ntp13495/example-go/endpoints/book"
	category "github.com/ntp13495/example-go/endpoints/category"
	lendingbook "github.com/ntp13495/example-go/endpoints/lendingbook"
	user "github.com/ntp13495/example-go/endpoints/user"
)

// Endpoints .
type Endpoints struct {
	FindUser    endpoint.Endpoint
	FindAllUser endpoint.Endpoint
	CreateUser  endpoint.Endpoint
	UpdateUser  endpoint.Endpoint
	DeleteUser  endpoint.Endpoint

	FindCategory    endpoint.Endpoint
	FindAllCategory endpoint.Endpoint
	CreateCategory  endpoint.Endpoint
	UpdateCategory  endpoint.Endpoint
	DeleteCategory  endpoint.Endpoint

	FindBook    endpoint.Endpoint
	FindAllBook endpoint.Endpoint
	CreateBook  endpoint.Endpoint
	UpdateBook  endpoint.Endpoint
	DeleteBook  endpoint.Endpoint

	FindLendingBook    endpoint.Endpoint
	FindAllLendingBook endpoint.Endpoint
	CreateLendingBook  endpoint.Endpoint
	UpdateLendingBook  endpoint.Endpoint
	DeleteLendingBook  endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		FindUser:    user.MakeFindEndPoint(s),
		FindAllUser: user.MakeFindAllEndpoint(s),
		CreateUser:  user.MakeCreateEndpoint(s),
		UpdateUser:  user.MakeUpdateEndpoint(s),
		DeleteUser:  user.MakeDeleteEndpoint(s),

		FindCategory:    category.MakeFindEndPoint(s),
		FindAllCategory: category.MakeFindAllEndpoint(s),
		CreateCategory:  category.MakeCreateEndpoint(s),
		UpdateCategory:  category.MakeUpdateEndpoint(s),
		DeleteCategory:  category.MakeDeleteEndpoint(s),

		FindBook:    book.MakeFindEndPoint(s),
		FindAllBook: book.MakeFindAllEndpoint(s),
		CreateBook:  book.MakeCreateEndpoint(s),
		UpdateBook:  book.MakeUpdateEndpoint(s),
		DeleteBook:  book.MakeDeleteEndpoint(s),

		FindLendingBook:    lendingbook.MakeFindEndPoint(s),
		FindAllLendingBook: lendingbook.MakeFindAllEndpoint(s),
		CreateLendingBook:  lendingbook.MakeCreateEndpoint(s),
		UpdateLendingBook:  lendingbook.MakeUpdateEndpoint(s),
		DeleteLendingBook:  lendingbook.MakeDeleteEndpoint(s),
	}
}
