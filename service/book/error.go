package book

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound              = errNotFound{}
	ErrUnknown               = errUnknown{}
	ErrNameIsRequired        = errNameIsRequired{}
	ErrAuthorIsRequired      = errAuthorIsRequired{}
	ErrDescriptionIsRequired = errDescriptionIsRequired{}
	ErrRecordNotFound        = errRecordNotFound{}
)

type errNotFound struct{}

func (errNotFound) Error() string {
	return "record not found"
}
func (errNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errUnknown struct{}

func (errUnknown) Error() string {
	return "unknown error"
}
func (errUnknown) StatusCode() int {
	return http.StatusBadRequest
}

type errAuthorIsRequired struct{}

func (errAuthorNotFound) Error() string {
	return "Author not found"
}
func (errAuthorNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errDescriptionIsRequired struct{}

func (errDescriptionNotFound) Error() string {
	return "Description not found"
}
func (errAuthorNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errRecordNotFound struct{}

func (errRecordNotFound) Error() string {
	return "client record not found"
}
func (errRecordNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errNameIsRequired struct{}

func (errNameIsRequired) Error() string {
	return "user name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}
