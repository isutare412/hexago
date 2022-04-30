package derror

import (
	"errors"
	"net/http"
)

var (
	ErrServiceUnavailable = errors.New("service unavailable")
	ErrUserNotFound       = errors.New("user not found")
)

var DomainErrorToStatusCode = map[error]int{
	ErrServiceUnavailable: http.StatusServiceUnavailable,
	ErrUserNotFound:       http.StatusBadRequest,
}
