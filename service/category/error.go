package category

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound                     = errNotFound{}
	ErrUnknown                      = errUnknown{}
	ErrCategoryNameIsRequired       = errCategoryNameIsRequired{}
	ErrCategoryNameLengthIsRequired = errCategoryNameLengthIsRequired{}
	ErrRecordNotFound               = errRecordNotFound{}
	ErrRecordExisted                = errRecordExisted{}
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

type errCategoryNameIsRequired struct{}

func (errCategoryNameIsRequired) Error() string {
	return "category name is required"
}

func (errCategoryNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errCategoryNameLengthIsRequired struct{}

func (errCategoryNameLengthIsRequired) Error() string {
	return "category name must have more than 5 characters."
}

func (errCategoryNameLengthIsRequired) StatusCode() int {
	return http.StatusLengthRequired
}

type errRecordExisted struct{}

func (errRecordExisted) Error() string {
	return "record already existed"
}

func (errRecordExisted) StatusCode() int {
	return http.StatusConflict
}
