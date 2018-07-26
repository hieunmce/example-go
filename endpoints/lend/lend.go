package lend

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/minhkhiemm/example-go/domain"
	"github.com/minhkhiemm/example-go/service"
)

// CreateData data for CreateLend
type CreateData struct {
	BookID domain.UUID `json:"book_id"`
	UserID domain.UUID `json:"user_id"`
	Name   string      `json:"name"`
	From   time.Time   `json:"from"`
	To     time.Time   `json:"to"`
}

// CreateRequest request struct for CreateLend
type CreateRequest struct {
	Lend CreateData `json:"lend"`
}

// CreateResponse response struct for CreateLend
type CreateResponse struct {
	Lend domain.Lend `json:"lend"`
}

// StatusCode customstatus code for success create Lend
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a Lend
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(CreateRequest)
			lend = &domain.Lend{
				Name:   req.Lend.Name,
				BookID: req.Lend.BookID,
				From:   req.Lend.From,
				To:     req.Lend.To,
				UserID: req.Lend.UserID,
			}
		)

		err := s.LendService.Create(ctx, lend)
		if err != nil {
			return nil, err
		}

		return CreateResponse{Lend: *lend}, nil
	}
}

// FindRequest request struct for Find a Lend
type FindRequest struct {
	LendID domain.UUID
}

// FindResponse response struct for Find a Lend
type FindResponse struct {
	Lend *domain.Lend `json:"lend"`
}

// MakeFindEndPoint make endpoint for find Lend
func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var lendFind domain.Lend
		req := request.(FindRequest)
		lendFind.ID = req.LendID

		lend, err := s.LendService.Find(ctx, &lendFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{Lend: lend}, nil
	}
}

// FindAllRequest request struct for FindAll Lend
type FindAllRequest struct{}

// FindAllResponse request struct for find all Lend
type FindAllResponse struct {
	Lends []domain.Lend `json:"lends"`
}

// MakeFindAllEndpoint make endpoint for find all Lend
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		lends, err := s.LendService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Lends: lends}, nil
	}
}

// UpdateData data for Create
type UpdateData struct {
	ID     domain.UUID `json:"-"`
	Name   string      `json:"name"`
	From   time.Time   `json:"from"`
	To     time.Time   `json:"to"`
	BookID domain.UUID `json:"book_id"`
	UserID domain.UUID `json:"user_id"`
}

// UpdateRequest request struct for update
type UpdateRequest struct {
	Lend UpdateData `json:"lend"`
}

// UpdateResponse response struct for Create
type UpdateResponse struct {
	Lend domain.Lend `json:"lend"`
}

// MakeUpdateEndpoint make endpoint for update a Lend
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(UpdateRequest)
			lend = domain.Lend{
				Model:  domain.Model{ID: req.Lend.ID},
				Name:   req.Lend.Name,
				BookID: req.Lend.BookID,
				From:   req.Lend.From,
				To:     req.Lend.To,
				UserID: req.Lend.UserID,
			}
		)

		res, err := s.LendService.Update(ctx, &lend)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{Lend: *res}, nil
	}
}

// DeleteRequest request struct for delete a Lend
type DeleteRequest struct {
	LendID domain.UUID
}

// DeleteResponse response struct for Find a Lend
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for update a Lend
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			lendFind = domain.Lend{}
			req      = request.(DeleteRequest)
		)
		lendFind.ID = req.LendID

		err := s.LendService.Delete(ctx, &lendFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
