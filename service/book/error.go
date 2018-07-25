package book

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound       = errNotFound{}
	ErrRecordExisted  = errRecordExisted{}
	ErrUnknown        = errUnknown{}
	ErrNameIsRequired = errNameIsRequired{}
	ErrRecordNotFound = errRecordNotFound{}
)

type errRecordExisted struct{}

func (errRecordExisted) Error() string {
	return "record existed"
}
func (errRecordExisted) StatusCode() int {
	return http.StatusFound
}

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
	return "user name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}
