package category

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/luantranminh/example-go/domain"
	"github.com/luantranminh/example-go/service"
)

// CreateData data for CreateCategory
type CreateData struct {
	Name string `json:"name"`
}

// CreateRequest request struct for CreateCategory
type CreateRequest struct {
	Category CreateData `json:"category"`
}

// CreateResponse response struct for CreateCategory
type CreateResponse struct {
	Category domain.Category `json:"category"`
}

// StatusCode customstatus code for success create Category
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a Category
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(CreateRequest)
			category = &domain.Category{
				Name: req.Category.Name,
			}
		)

		err := s.CategoryService.Create(ctx, category)
		if err != nil {
			return nil, err
		}

		return CreateResponse{Category: *category}, nil
	}
}

// FindRequest request struct for Find a Category
type FindRequest struct {
	CategoryID domain.UUID
}

// FindResponse response struct for Find a Category
type FindResponse struct {
	Category *domain.Category `json:"category"`
}

// MakeFindEndPoint make endpoint for find Category
func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var categoryFind domain.Category
		req := request.(FindRequest)
		categoryFind.ID = req.CategoryID

		category, err := s.CategoryService.Find(ctx, &categoryFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{Category: category}, nil
	}
}

// FindAllRequest request struct for FindAll Category
type FindAllRequest struct{}

// FindAllResponse request struct for find all Category
type FindAllResponse struct {
	Categorys []domain.Category `json:"categorys"`
}

// MakeFindAllEndpoint make endpoint for find all Category
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		categorys, err := s.CategoryService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Categorys: categorys}, nil
	}
}

// UpdateData data for Create
type UpdateData struct {
	ID   domain.UUID `json:"-"`
	Name string      `json:"name"`
}

// UpdateRequest request struct for update
type UpdateRequest struct {
	Category UpdateData `json:"category"`
}

// UpdateResponse response struct for Create
type UpdateResponse struct {
	Category domain.Category `json:"category"`
}

// MakeUpdateEndpoint make endpoint for update a Category
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(UpdateRequest)
			category = domain.Category{
				Model: domain.Model{ID: req.Category.ID},
				Name:  req.Category.Name,
			}
		)

		res, err := s.CategoryService.Update(ctx, &category)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{Category: *res}, nil
	}
}

// DeleteRequest request struct for delete a Category
type DeleteRequest struct {
	CategoryID domain.UUID
}

// DeleteResponse response struct for Find a Category
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for update a Category
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			categoryFind = domain.Category{}
			req          = request.(DeleteRequest)
		)
		categoryFind.ID = req.CategoryID

		err := s.CategoryService.Delete(ctx, &categoryFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
