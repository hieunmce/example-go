package category

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound        = errNotFound{}
	ErrUnknown         = errUnknown{}
	ErrNameIsRequired  = errNameIsRequired{}
	ErrNameIsInvalid   = errNameIsInvalid{}
	ErrRecordNotFound  = errRecordNotFound{}
	ErrNameIsDuplicate = errNameIsDuplicate{}
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
	return "category name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errNameIsInvalid struct{}

func (errNameIsInvalid) Error() string {
	return "category name is invalid"
}

func (errNameIsInvalid) StatusCode() int {
	return http.StatusBadRequest
}

type errNameIsDuplicate struct{}

func (errNameIsDuplicate) Error() string {
	return "category name is already exist"
}

func (errNameIsDuplicate) StatusCode() int {
	return http.StatusBadRequest
}
