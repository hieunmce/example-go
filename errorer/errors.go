package errorer

import "net/http"

var (
	ErrInvalidReceiveTime = errInvalidReceiveTime{}
)

type errInvalidReceiveTime struct{}

func (errInvalidReceiveTime) Error() string {
	return "receive time must be after order time"
}

func (errInvalidReceiveTime) StatusCode() int {
	return http.StatusBadRequest
}
