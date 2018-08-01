package lend

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/minhkhiemm/example-go/domain"
	lendEndpoint "github.com/minhkhiemm/example-go/endpoints/lend"
)

// FindRequest .
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	lendID, err := domain.UUIDFromString(chi.URLParam(r, "lend_id"))
	if err != nil {
		return nil, err
	}
	return lendEndpoint.FindRequest{LendID: lendID}, nil
}

// FindAllRequest .
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return lendEndpoint.FindAllRequest{}, nil
}

// CreateRequest .
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req lendEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest .
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	lendID, err := domain.UUIDFromString(chi.URLParam(r, "lend_id"))
	if err != nil {
		return nil, err
	}

	var req lendEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.Lend.ID = lendID

	return req, nil
}

// DeleteRequest .
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	lendID, err := domain.UUIDFromString(chi.URLParam(r, "lend_id"))
	if err != nil {
		return nil, err
	}
	return lendEndpoint.DeleteRequest{LendID: lendID}, nil
}
