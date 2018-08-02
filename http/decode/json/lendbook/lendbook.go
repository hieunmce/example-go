package lendbook

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/neverdiefc/example-go/domain"
	lendbookEndpoint "github.com/neverdiefc/example-go/endpoints/lendbook"
)

// FindRequest .
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	lendbookID, err := domain.UUIDFromString(chi.URLParam(r, "lendbook_id"))
	if err != nil {
		return nil, err
	}
	return lendbookEndpoint.FindRequest{LendbookID: lendbookID}, nil
}

// FindAllRequest .
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return lendbookEndpoint.FindAllRequest{}, nil
}

// CreateRequest .
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req lendbookEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest .
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	lendbookID, err := domain.UUIDFromString(chi.URLParam(r, "lendbook_id"))
	if err != nil {
		return nil, err
	}

	var req lendbookEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.Lendbook.ID = lendbookID

	return req, nil
}

// DeleteRequest .
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	lendbookID, err := domain.UUIDFromString(chi.URLParam(r, "lendbook_id"))
	if err != nil {
		return nil, err
	}
	return lendbookEndpoint.DeleteRequest{LendbookID: lendbookID}, nil
}
