package lendbook

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/neverdiefc/example-go/domain"
	"github.com/neverdiefc/example-go/service"
)

// CreateData data for CreateLendbook
type CreateData struct {
	BookID domain.UUID `json:"book_id"`
	UserID domain.UUID `json:"user_id"`
	From   time.Time   `json:"from"`
	To     time.Time   `json:"to"`
}

// CreateRequest request struct for CreateLendbook
type CreateRequest struct {
	Lendbook CreateData `json:"lendbook"`
}

// CreateResponse response struct for CreateLendbook
type CreateResponse struct {
	Lendbook domain.Lendbook `json:"lendbook"`
}

// StatusCode customstatus code for success create Lendbook
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a Lendbook
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(CreateRequest)
			lendbook = &domain.Lendbook{
				UserID: req.Lendbook.UserID,
				BookID: req.Lendbook.BookID,
				From:   req.Lendbook.From,
				To:     req.Lendbook.To,
			}
		)

		err := s.LendbookService.Create(ctx, lendbook)
		if err != nil {
			return nil, err
		}

		return CreateResponse{Lendbook: *lendbook}, nil
	}
}

// FindRequest request struct for Find a Lendbook
type FindRequest struct {
	LendbookID domain.UUID
}

// FindResponse response struct for Find a Lendbook
type FindResponse struct {
	Lendbook *domain.Lendbook `json:"lendbook"`
}

// MakeFindEndPoint make endpoint for find Lendbook
func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var lendbookFind domain.Lendbook
		req := request.(FindRequest)
		lendbookFind.ID = req.LendbookID

		lendbook, err := s.LendbookService.Find(ctx, &lendbookFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{Lendbook: lendbook}, nil
	}
}

// FindAllRequest request struct for FindAll Lendbook
type FindAllRequest struct{}

// FindAllResponse request struct for find all Lendbook
type FindAllResponse struct {
	Lendbooks []domain.Lendbook `json:"lendbooks"`
}

// MakeFindAllEndpoint make endpoint for find all Lendbook
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		lendbooks, err := s.LendbookService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Lendbooks: lendbooks}, nil
	}
}

// UpdateData data for Create
type UpdateData struct {
	ID     domain.UUID `json:"-"`
	BookID domain.UUID `json:"book_id"`
	UserID domain.UUID `json:"user_id"`
	From   time.Time   `json:"from"`
	To     time.Time   `json:"to"`
}

// UpdateRequest request struct for update
type UpdateRequest struct {
	Lendbook UpdateData `json:"lendbook"`
}

// UpdateResponse response struct for Create
type UpdateResponse struct {
	Lendbook domain.Lendbook `json:"lendbook"`
}

// MakeUpdateEndpoint make endpoint for update a Lendbook
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(UpdateRequest)
			lendbook = domain.Lendbook{
				Model:  domain.Model{ID: req.Lendbook.ID},
				UserID: req.Lendbook.UserID,
				BookID: req.Lendbook.BookID,
				From:   req.Lendbook.From,
				To:     req.Lendbook.To,
			}
		)

		res, err := s.LendbookService.Update(ctx, &lendbook)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{Lendbook: *res}, nil
	}
}

// DeleteRequest request struct for delete a Lendbook
type DeleteRequest struct {
	LendbookID domain.UUID
}

// DeleteResponse response struct for Find a Lendbook
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for update a Lendbook
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			lendbookFind = domain.Lendbook{}
			req          = request.(DeleteRequest)
		)
		lendbookFind.ID = req.LendbookID

		err := s.LendbookService.Delete(ctx, &lendbookFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
