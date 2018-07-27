package lend

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound           = errNotFound{}
	ErrUnknown            = errUnknown{}
	ErrNameIsRequired     = errNameIsRequired{}
	ErrRecordNotFound     = errRecordNotFound{}
	ErrBookNotFound       = errBookNotFound{}
	ErrUserNotFound       = errUserNotFound{}
	ErrBookIsNotAvailable = errBookIsNotAvailable{}
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

type errBookNotFound struct{}

func (errBookNotFound) Error() string {
	return "book is not in category"
}
func (errBookNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errUserNotFound struct{}

func (errUserNotFound) Error() string {
	return "user is not exist"
}
func (errUserNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errBookIsNotAvailable struct{}

func (errBookIsNotAvailable) Error() string {
	return "book is not available to lend"
}
func (errBookIsNotAvailable) StatusCode() int {
	return http.StatusBadRequest
}
