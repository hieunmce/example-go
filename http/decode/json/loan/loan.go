package loan

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/luantranminh/example-go/domain"
	loanEndpoint "github.com/luantranminh/example-go/endpoints/loan"
)

// FindRequest .
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	loanID, err := domain.UUIDFromString(chi.URLParam(r, "loan_id"))
	if err != nil {
		return nil, err
	}
	return loanEndpoint.FindRequest{LoanID: loanID}, nil
}

// FindAllRequest .
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return loanEndpoint.FindAllRequest{}, nil
}

// CreateRequest .
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req loanEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest .
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	loanID, err := domain.UUIDFromString(chi.URLParam(r, "loan_id"))
	if err != nil {
		return nil, err
	}

	var req loanEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.Loan.ID = loanID

	return req, nil
}

// DeleteRequest .
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	loanID, err := domain.UUIDFromString(chi.URLParam(r, "loan_id"))
	if err != nil {
		return nil, err
	}
	return loanEndpoint.DeleteRequest{LoanID: loanID}, nil
}
