package lendbook

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound           = errNotFound{}
	ErrRecordExisted      = errRecordExisted{}
	ErrUserNotExist       = errUserNotExist{}
	ErrUnknown            = errUnknown{}
	ErrBookNotExist       = errBookNotExist{}
	ErrBookNotLend        = errBookNotLend{}
	ErrRecordBookNotFound = errRecordBookNotFound{}
	ErrRecordUserNotFound = errRecordUserNotFound{}
	ErrBookIsBusy         = errBookIsBusy{}
	ErrBookNotFound       = errBookNotFound{}
	ErrRecordNotFound     = errRecordNotFound{}
)

type errBookNotLend struct{}

func (errBookNotLend) Error() string {
	return "Book not lend"
}
func (errBookNotLend) StatusCode() int {
	return http.StatusFound
}

type errUserNotExist struct{}

func (errUserNotExist) Error() string {
	return "User not Exist"
}
func (errUserNotExist) StatusCode() int {
	return http.StatusFound
}

type errBookNotExist struct{}

func (errBookNotExist) Error() string {
	return "Book not Exist"
}
func (errBookNotExist) StatusCode() int {
	return http.StatusFound
}

type errBookNotFound struct{}

func (errBookNotFound) Error() string {
	return "Book Not Found"
}
func (errBookNotFound) StatusCode() int {
	return http.StatusFound
}

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

type errRecordBookNotFound struct{}

func (errRecordBookNotFound) Error() string {
	return "record book not found"
}

func (errRecordBookNotFound) StatusCode() int {
	return http.StatusBadRequest
}

type errRecordUserNotFound struct{}

func (errRecordUserNotFound) Error() string {
	return "record user not found"
}

func (errRecordUserNotFound) StatusCode() int {
	return http.StatusBadRequest
}

type errBookIsBusy struct{}

func (errBookIsBusy) Error() string {
	return "Book has been borrowed"
}

func (errBookIsBusy) StatusCode() int {
	return http.StatusBadRequest
}
