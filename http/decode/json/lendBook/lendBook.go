package lendBook

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/luquehuong/example-go/domain"
	lendBookEndpoint "github.com/luquehuong/example-go/endpoints/lendBook"
)

// FindRequest .
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	lendBookID, err := domain.UUIDFromString(chi.URLParam(r, "lendBook_id"))
	if err != nil {
		return nil, err
	}
	return lendBookEndpoint.FindRequest{LendBookID: lendBookID}, nil
}

// FindAllRequest .
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return lendBookEndpoint.FindAllRequest{}, nil
}

// CreateRequest .
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req lendBookEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest .
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	lendBookID, err := domain.UUIDFromString(chi.URLParam(r, "lendBook_id"))
	if err != nil {
		return nil, err
	}

	var req lendBookEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.LendBook.ID = lendBookID

	return req, nil
}

// DeleteRequest .
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	lendBookID, err := domain.UUIDFromString(chi.URLParam(r, "lendBook_id"))
	if err != nil {
		return nil, err
	}
	return lendBookEndpoint.DeleteRequest{LendBookID: lendBookID}, nil
}
