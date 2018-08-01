package book

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound              = errNotFound{}
	ErrUnknown               = errUnknown{}
	ErrNameIsRequired        = errNameIsRequired{}
	ErrRecordNotFound        = errRecordNotFound{}
	ErrCategoryNotFound      = errCategoryNotFound{}
	ErrNameIsInvalid         = errNameIsInvalid{}
	ErrDescriptionIsInvalid  = errDescriptionIsInvalid{}
	ErrDescriptionIsRequired = errDescriptionIsRequired{}
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

type errRecordNotFound struct{}

func (errRecordNotFound) Error() string {
	return "client record not found"
}
func (errRecordNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errNameIsRequired struct{}

func (errNameIsRequired) Error() string {
	return "book name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errCategoryNotFound struct{}

func (errCategoryNotFound) Error() string {
	return "category is not exist"
}

func (errCategoryNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errNameIsInvalid struct{}

func (errNameIsInvalid) Error() string {
	return "book name is invalid, need to have more than 5 character"
}
func (errNameIsInvalid) StatusCode() int {
	return http.StatusBadRequest
}

type errDescriptionIsInvalid struct{}

func (errDescriptionIsInvalid) Error() string {
	return "description is invalid, need to have more than 5 character"
}
func (errDescriptionIsInvalid) StatusCode() int {
	return http.StatusBadRequest
}

type errDescriptionIsRequired struct{}

func (errDescriptionIsRequired) Error() string {
	return "description is required"
}
func (errDescriptionIsRequired) StatusCode() int {
	return http.StatusBadRequest
}
