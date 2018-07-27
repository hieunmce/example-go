package book

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound                    = errNotFound{}
	ErrUnknown                     = errUnknown{}
	ErrBookNameIsRequired          = errBookNameIsRequired{}
	ErrBookNameLengthIsRequired    = errBookNameLengthIsRequired{}
	ErrDescriptionIsRequired       = errDescriptionIsRequired{}
	ErrDescriptionLengthIsRequired = errDescriptionLengthIsRequired{}
	ErrRecordNotFound              = errRecordNotFound{}
	ErrRecordExisted               = errRecordExisted{}
	ErrInvalidCategory             = errInvalidCategory{}
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

type errBookNameIsRequired struct{}

func (errBookNameIsRequired) Error() string {
	return "book name is required"
}

func (errBookNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errBookNameLengthIsRequired struct{}

func (errBookNameLengthIsRequired) Error() string {
	return "book name must have more than 5 characters."
}

func (errBookNameLengthIsRequired) StatusCode() int {
	return http.StatusLengthRequired
}

type errRecordExisted struct{}

func (errRecordExisted) Error() string {
	return "record already existed"
}

func (errRecordExisted) StatusCode() int {
	return http.StatusConflict
}

type errDescriptionIsRequired struct{}

func (errDescriptionIsRequired) Error() string {
	return "book name is required"
}

func (errDescriptionIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errDescriptionLengthIsRequired struct{}

func (errDescriptionLengthIsRequired) Error() string {
	return "book name must have more than 5 characters."
}

func (errDescriptionLengthIsRequired) StatusCode() int {
	return http.StatusLengthRequired
}

type errInvalidCategory struct{}

func (errInvalidCategory) Error() string {
	return "invalid category"
}
func (errInvalidCategory) StatusCode() int {
	return http.StatusNotFound
}
