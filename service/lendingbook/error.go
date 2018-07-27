package lendingbook

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound           = errNotFound{}
	ErrUnknown            = errUnknown{}
	ErrNameIsRequired     = errNameIsRequired{}
	ErrNameLengthRequired = errNameIsLengthRequired{}
	ErrRecordNotFound     = errRecordNotFound{}
	ErrInvalidBook        = errInvalidBook{}
	ErrInvalidUser        = errInvalidUser{}
	ErrBookInUse          = errBookInUse{}
	// ErrRecordExisted      = errRecordExisted{}
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
	return "lendingbook name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errNameIsLengthRequired struct{}

func (errNameIsLengthRequired) Error() string {
	return "lendingbook name must have more than 5 characters."
}

func (errNameIsLengthRequired) StatusCode() int {
	return http.StatusLengthRequired
}

type errInvalidBook struct{}

func (errInvalidBook) Error() string {
	return "Invalid book is required"
}

func (errInvalidBook) StatusCode() int {
	return http.StatusBadRequest
}

type errInvalidUser struct{}

func (errInvalidUser) Error() string {
	return "Invalid user is required"
}

func (errInvalidUser) StatusCode() int {
	return http.StatusBadRequest
}

type errBookInUse struct{}

func (errBookInUse) Error() string {
	return "this book is unavailable"
}

func (errBookInUse) StatusCode() int {
	return http.StatusBadRequest
}

// type errRecordExisted struct{}

// func (errRecordExisted) Error() string {
// 	return "lendingbook existed"
// }

// func (errRecordExisted) StatusCode() int {
// 	return http.StatusConflict
// }
