package book

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound                 = errNotFound{}
	ErrUnknown                  = errUnknown{}
	ErrNameIsRequired           = errNameIsRequired{}
	ErrRecordNotFound           = errRecordNotFound{}
	ErrNotExistCategoryID       = errNotExistCategoryID{}
	ErrCategoryIDIsRequired     = errCategoryIDIsRequired{}
	ErrMinimumLengthName        = errMinimumLengthName{}
	ErrDescriptionIsRequired    = errDescriptionIsRequired{}
	ErrMinimumLengthDescription = errMinimumLengthDescription{}
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

type errNotExistCategoryID struct{}

func (errNotExistCategoryID) Error() string {
	return "The ID for category not exist in table Categories"
}

func (errNotExistCategoryID) StatusCode() int {
	return http.StatusNotFound
}

type errCategoryIDIsRequired struct{}

func (errCategoryIDIsRequired) Error() string {
	return "Category id is required"
}
func (errCategoryIDIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errMinimumLengthName struct{}

func (errMinimumLengthName) Error() string {
	return "Minimum Length for Name is 5 characters "
}
func (errMinimumLengthName) StatusCode() int {
	return http.StatusBadRequest
}

type errDescriptionIsRequired struct{}

func (errDescriptionIsRequired) Error() string {
	return "Description is required"
}
func (errDescriptionIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errMinimumLengthDescription struct{}

func (errMinimumLengthDescription) Error() string {
	return "Minimum length for description is 5 characters"
}
