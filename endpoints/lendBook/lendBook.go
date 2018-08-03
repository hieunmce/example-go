package lendBook

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/luquehuong/example-go/domain"
	"github.com/luquehuong/example-go/service"
)

// CreateData data for CreateUser
type CreateData struct {
	Book_id domain.UUID `json:"book_id"`
	User_id domain.UUID `json:"user_id"`
	From    time.Time   `json:"from"`
	To      time.Time   `json:"to"`
}

// CreateRequest request struct for CreateUser
type CreateRequest struct {
	LendBook CreateData `json:"lendBook"`
}

// CreateResponse response struct for CreateUser
type CreateResponse struct {
	LendBook domain.LendBook `json:"lendBook"`
}

// StatusCode customstatus code for success create LendBook
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a LendBook
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(CreateRequest)
			lendBook = &domain.LendBook{
				Book_id: req.LendBook.Book_id,
				User_id: req.LendBook.User_id,
				From:    req.LendBook.From,
				To:      req.LendBook.To,
			}
		)

		err := s.LendBookService.Create(ctx, lendBook)
		if err != nil {
			return nil, err
		}

		return CreateResponse{LendBook: *lendBook}, nil
	}
}

// FindRequest request struct for Find a LendBook
type FindRequest struct {
	LendBookID domain.UUID
}

// FindResponse response struct for Find a LendBook
type FindResponse struct {
	LendBook *domain.LendBook `json:"lendBook"`
}

// MakeFindEndPoint make endpoint for find LendBook
func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var userlendbookFind domain.LendBook
		req := request.(FindRequest)
		userlendbookFind.ID = req.LendBookID

		lendBook, err := s.LendBookService.Find(ctx, &userlendbookFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{LendBook: lendBook}, nil
	}
}

// FindAllRequest request struct for FindAll LendBook
type FindAllRequest struct{}

// FindAllResponse request struct for find all LendBook
type FindAllResponse struct {
	LendBooks []domain.LendBook `json:"userlendbooks"`
}

// MakeFindAllEndpoint make endpoint for find all LendBook
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		userlendbooks, err := s.LendBookService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{LendBooks: userlendbooks}, nil
	}
}

// UpdateData data for Create
type UpdateData struct {
	ID      domain.UUID `json:"-"`
	Book_id domain.UUID `json:"book_id"`
	User_id domain.UUID `json:"user_id"`
	From    time.Time   `json:"from"`
	To      time.Time   `json:"to"`
}

// UpdateRequest request struct for update
type UpdateRequest struct {
	LendBook UpdateData `json:"lendBook"`
}

// UpdateResponse response struct for Create
type UpdateResponse struct {
	LendBook domain.LendBook `json:"lendBook"`
}

// MakeUpdateEndpoint make endpoint for update a LendBook
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(UpdateRequest)
			lendBook = domain.LendBook{
				Model:   domain.Model{ID: req.LendBook.ID},
				Book_id: req.LendBook.Book_id,
				User_id: req.LendBook.User_id,
				From:    req.LendBook.From,
				To:      req.LendBook.To,
			}
		)

		res, err := s.LendBookService.Update(ctx, &lendBook)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{LendBook: *res}, nil
	}
}

// DeleteRequest request struct for delete a LendBook
type DeleteRequest struct {
	LendBookID domain.UUID
}

// DeleteResponse response struct for Find a LendBook
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for update a LendBook
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			userlendbookFind = domain.LendBook{}
			req              = request.(DeleteRequest)
		)
		userlendbookFind.ID = req.LendBookID

		err := s.LendBookService.Delete(ctx, &userlendbookFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
