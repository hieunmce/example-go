package lendingbook

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/ntp13495/example-go/domain"
	"github.com/ntp13495/example-go/service"
)

// CreateData data for CreateLendingBook
type CreateData struct {
	BookID domain.UUID `json:"book_id"`
	UserID domain.UUID `json:"user_id"`
	From   time.Time   `json:"from"`
	To     time.Time   `json:"to"`
}

// CreateRequest request struct for CreateLendingBook
type CreateRequest struct {
	LendingBook CreateData `json:"lendingbook"`
}

// CreateResponse response struct for CreateLendingBook
type CreateResponse struct {
	LendingBook domain.LendingBook `json:"lendingbook"`
}

// StatusCode customstatus code for success create LendingBook
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a LendingBook
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req         = request.(CreateRequest)
			lendingbook = &domain.LendingBook{
				BookID: req.LendingBook.BookID,
				UserID: req.LendingBook.UserID,
				From:   req.LendingBook.From,
				To:     req.LendingBook.To,
			}
		)

		err := s.LendingBookService.Create(ctx, lendingbook)
		if err != nil {
			return nil, err
		}

		return CreateResponse{LendingBook: *lendingbook}, nil
	}
}

// FindRequest request struct for Find a LendingBook
type FindRequest struct {
	LendingBookID domain.UUID
}

// FindResponse response struct for Find a LendingBook
type FindResponse struct {
	LendingBook *domain.LendingBook `json:"lendingbook"`
}

// MakeFindEndPoint make endpoint for find LendingBook
func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var lendingbookFind domain.LendingBook
		req := request.(FindRequest)
		lendingbookFind.ID = req.LendingBookID

		lendingbook, err := s.LendingBookService.Find(ctx, &lendingbookFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{LendingBook: lendingbook}, nil
	}
}

// FindAllRequest request struct for FindAll LendingBook
type FindAllRequest struct{}

// FindAllResponse request struct for find all LendingBook
type FindAllResponse struct {
	LendingBooks []domain.LendingBook `json:"lendingbooks"`
}

// MakeFindAllEndpoint make endpoint for find all LendingBook
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		lendingbooks, err := s.LendingBookService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{LendingBooks: lendingbooks}, nil
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
	LendingBook UpdateData `json:"lendingbook"`
}

// UpdateResponse response struct for Create
type UpdateResponse struct {
	LendingBook domain.LendingBook `json:"lendingbook"`
}

// MakeUpdateEndpoint make endpoint for update a LendingBook
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req         = request.(UpdateRequest)
			lendingbook = domain.LendingBook{
				Model:  domain.Model{ID: req.LendingBook.ID},
				BookID: req.LendingBook.BookID,
				UserID: req.LendingBook.UserID,
				From:   req.LendingBook.From,
				To:     req.LendingBook.To,
			}
		)

		res, err := s.LendingBookService.Update(ctx, &lendingbook)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{LendingBook: *res}, nil
	}
}

// DeleteRequest request struct for delete a LendingBook
type DeleteRequest struct {
	LendingBookID domain.UUID
}

// DeleteResponse response struct for Find a LendingBook
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for update a LendingBook
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			lendingbookFind = domain.LendingBook{}
			req             = request.(DeleteRequest)
		)
		lendingbookFind.ID = req.LendingBookID

		err := s.LendingBookService.Delete(ctx, &lendingbookFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
