package book

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/luquehuong/example-go/domain"
	bookEndpoint "github.com/luquehuong/example-go/endpoints/book"
)

// FindRequest .
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	bookID, err := domain.UUIDFromString(chi.URLParam(r, "book_id"))
	if err != nil {
		return nil, err
	}
	return bookEndpoint.FindRequest{BookID: bookID}, nil
}

// FindAllRequest .
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return bookEndpoint.FindAllRequest{}, nil
}

// CreateRequest .
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req bookEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest .
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	bookID, err := domain.UUIDFromString(chi.URLParam(r, "book_id"))
	if err != nil {
		return nil, err
	}

	var req bookEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	//req.book.ID = bookID
	req.Book.ID = bookID

	return req, nil
}

// DeleteRequest .
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	bookID, err := domain.UUIDFromString(chi.URLParam(r, "book_id"))
	if err != nil {
		return nil, err
	}
	return bookEndpoint.DeleteRequest{BookID: bookID}, nil
}
