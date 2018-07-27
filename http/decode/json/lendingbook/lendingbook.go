package lendingbook

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/ntp13495/example-go/domain"
	lendingbookEndpoint "github.com/ntp13495/example-go/endpoints/lendingbook"
)

// FindRequest .
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	lendingbookID, err := domain.UUIDFromString(chi.URLParam(r, "lendingbook_id"))
	if err != nil {
		return nil, err
	}
	return lendingbookEndpoint.FindRequest{LendingBookID: lendingbookID}, nil
}

// FindAllRequest .
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return lendingbookEndpoint.FindAllRequest{}, nil
}

// CreateRequest .
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req lendingbookEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest .
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	lendingbookID, err := domain.UUIDFromString(chi.URLParam(r, "lendingbook_id"))
	if err != nil {
		return nil, err
	}

	var req lendingbookEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.LendingBook.ID = lendingbookID

	return req, nil
}

// DeleteRequest .
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	lendingbookID, err := domain.UUIDFromString(chi.URLParam(r, "lendingbook_id"))
	if err != nil {
		return nil, err
	}
	return lendingbookEndpoint.DeleteRequest{LendingBookID: lendingbookID}, nil
}
