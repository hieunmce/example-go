package category

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"example.com/m/domain"
	categoryEndpoint "example.com/m/endpoints/category"
)

// FindRequest .
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	categoryID, err := domain.UUIDFromString(chi.URLParam(r, "category_id"))
	if err != nil {
		return nil, err
	}
	return categoryEndpoint.FindRequest{CategoryID: categoryID}, nil
}

// FindAllRequest .
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return categoryEndpoint.FindAllRequest{}, nil
}

// CreateRequest .
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req categoryEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest .
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	categoryID, err := domain.UUIDFromString(chi.URLParam(r, "category_id"))
	if err != nil {
		return nil, err
	}

	var req categoryEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.Category.ID = categoryID

	return req, nil
}

// DeleteRequest .
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	categoryID, err := domain.UUIDFromString(chi.URLParam(r, "category_id"))
	if err != nil {
		return nil, err
	}
	return categoryEndpoint.DeleteRequest{CategoryID: categoryID}, nil
}
