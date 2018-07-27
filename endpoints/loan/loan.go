package loan

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/luantranminh/example-go/domain"
	"github.com/luantranminh/example-go/service"
)

// CreateData data for CreateLoan
type CreateData struct {
	BookID domain.UUID `json:"book_id"`
	UserID domain.UUID `json:"user_id"`
	To     *time.Time  `json:"to,omitempty"`
}

// CreateRequest request struct for CreateLoan
type CreateRequest struct {
	Loan CreateData `json:"loan"`
}

// CreateResponse response struct for CreateLoan
type CreateResponse struct {
	Loan domain.Loan `json:"loan"`
}

// StatusCode customstatus code for success create Loan
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a Loan
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(CreateRequest)
			loan = &domain.Loan{
				BookID: req.Loan.BookID,
				UserID: req.Loan.UserID,
				To:     req.Loan.To,
			}
		)

		err := s.LoanService.Create(ctx, loan)
		if err != nil {
			return nil, err
		}

		return CreateResponse{Loan: *loan}, nil
	}
}

// FindRequest request struct for Find a Loan
type FindRequest struct {
	LoanID domain.UUID
}

// FindResponse response struct for Find a Loan
type FindResponse struct {
	Loan *domain.Loan `json:"loan"`
}

// MakeFindEndPoint make endpoint for find Loan
func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var loanFind domain.Loan
		req := request.(FindRequest)
		loanFind.ID = req.LoanID

		loan, err := s.LoanService.Find(ctx, &loanFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{Loan: loan}, nil
	}
}

// FindAllRequest request struct for FindAll Loan
type FindAllRequest struct{}

// FindAllResponse request struct for find all Loan
type FindAllResponse struct {
	Loans []domain.Loan `json:"loans"`
}

// MakeFindAllEndpoint make endpoint for find all Loan
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		loans, err := s.LoanService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Loans: loans}, nil
	}
}

// UpdateData data for Create
type UpdateData struct {
	ID     domain.UUID `json:"-"`
	BookID domain.UUID `json:"book_id"`
	UserID domain.UUID `json:"user_id"`
	To     *time.Time  `json:"to,omitempty"`
}

// UpdateRequest request struct for update
type UpdateRequest struct {
	Loan UpdateData `json:"loan"`
}

// UpdateResponse response struct for Create
type UpdateResponse struct {
	Loan domain.Loan `json:"loan"`
}

// MakeUpdateEndpoint make endpoint for update a Loan
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(UpdateRequest)
			loan = domain.Loan{
				Model:  domain.Model{ID: req.Loan.ID},
				BookID: req.Loan.BookID,
				UserID: req.Loan.UserID,
				To:     req.Loan.To,
			}
		)

		res, err := s.LoanService.Update(ctx, &loan)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{Loan: *res}, nil
	}
}

// DeleteRequest request struct for delete a Loan
type DeleteRequest struct {
	LoanID domain.UUID
}

// DeleteResponse response struct for Find a Loan
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for update a Loan
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			loanFind = domain.Loan{}
			req      = request.(DeleteRequest)
		)
		loanFind.ID = req.LoanID

		err := s.LoanService.Delete(ctx, &loanFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
