package lendBook

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound           = errNotFound{}
	ErrUnknown            = errUnknown{}
	ErrRecordNotFound     = errRecordNotFound{}
	ErrBookIDNotFound     = errBookIDNotFound{}
	ErrUserIDNotFound     = errUserIDNotFound{}
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
	return "user name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errNameIsToShort struct{}

func (errNameIsToShort) Error() string {
	return "name must longer than 5 characters"
}

func (errNameIsToShort) StatusCode() int {
	return http.StatusBadRequest
}

type errBookIDNotFound struct{}

func (errBookIDNotFound) Error() string {
	return "Book_id not found"
}
func (errBookIDNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errUserIDNotFound struct{}

func (errUserIDNotFound) Error() string {
	return "User_id not found"
}
func (errUserIDNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errBookIsNotAvailable struct{}

func (errBookIsNotAvailable) Error() string {
	return "this book is not available to lend now"
}

func (errBookIsNotAvailable) StatusCode() int {
	return http.StatusBadRequest
}
