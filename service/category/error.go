package category

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound        = errNotFound{}
	ErrUnknown         = errUnknown{}
	ErrNameIsRequired  = errNameIsRequired{}
	ErrEmailIsRequired = errEmailIsRequired{}
	ErrEmailIsInvalid  = errEmailIsInvalid{}
	ErrRecordNotFound  = errRecordNotFound{}
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

type errEmailIsRequired struct{}

func (errEmailIsRequired) Error() string {
	return "email is required"
}
func (errEmailIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errEmailIsInvalid struct{}

func (errEmailIsInvalid) Error() string {
	return "email address is invalid"
}
func (errEmailIsInvalid) StatusCode() int {
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
	return "user name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}
