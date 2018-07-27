package book

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound            = errNotFound{}
	ErrUnknown             = errUnknown{}
	ErrCategoryNotFound    = errCategoryNotFound{}
	ErrDescriptionRequired = errDescriptionRequired{}
	ErrNameIsRequired      = errNameIsRequired{}
	ErrRecordNotFound      = errRecordNotFound{}
)

type errNotFound struct{}

func (errNotFound) Error() string {
	return "record not found"
}
func (errNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errCategoryNotFound struct{}

func (errCategoryNotFound) Error() string {
	return "Category not Found"
}
func (errCategoryNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errUnknown struct{}

func (errUnknown) Error() string {
	return "unknown error"
}
func (errUnknown) StatusCode() int {
	return http.StatusBadRequest
}

type errDescriptionRequired struct{}

func (errDescriptionRequired) Error() string {
	return "length book is empty and length <= 5 characters"
}
func (errDescriptionRequired) StatusCode() int {
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
