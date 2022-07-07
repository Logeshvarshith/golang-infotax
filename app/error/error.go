package error

import (
	"errors"
	"fmt"
	"net/http"
)

type Type string

const (
	Authorization        Type = "AUTHORIZATION"
	BadRequest           Type = "BAD_REQUEST"
	Conflict             Type = "CONFLICT"
	Internal             Type = "INTERNAL"
	NotFound             Type = "NOT_FOUND"
	PayloadTooLarge      Type = "PAYLOAD_TOO_LARGE"
	UnsupportedMediaType Type = "UNSUPPORTED_MEDIA_TYPE"
	ServiceUnavailable   Type = "SERVICE_UNAVAILABLE"
)

type Error struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
} // @name Error

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Status() int {

	switch e.Type {
	case Authorization:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	case UnsupportedMediaType:
		return http.StatusUnsupportedMediaType
	case ServiceUnavailable:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}

func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status()
	}

	return http.StatusInternalServerError
}

func NewAuthorization(reason string) *Error {
	return &Error{
		Type:    Authorization,
		Message: reason,
	}
}

func NewBadRequest(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad Request. Reason: %v", reason),
	}
}

func NewConflict(name, value string) *Error {
	return &Error{
		Type:    Conflict,
		Message: fmt.Sprintf("Resource : %v with value : %v already exists", name, value),
	}
}

func NewBulkInsertConflict(reason string) *Error {
	return &Error{
		Type:    Conflict,
		Message: reason,
	}
}

func NewInternal() *Error {
	return &Error{
		Type:    Internal,
		Message: fmt.Sprintf("Internal server error"),
	}
}

func NewNotFound(name, value string) *Error {
	return &Error{
		Type:    NotFound,
		Message: fmt.Sprintf("Resource : %v with value : %v not found", name, value),
	}
}

func NewPayloadTooLarge(maxContentSize, actualContentSize int64) *Error {
	return &Error{
		Type:    PayloadTooLarge,
		Message: fmt.Sprintf("Payload size more than %v. Actual content size: %v", maxContentSize, actualContentSize),
	}
}

func NewUnsupportMediaType(message string) *Error {
	return &Error{
		Type:    UnsupportedMediaType,
		Message: fmt.Sprint(message),
	}
}

func NewServiceUnavailable() *Error {
	return &Error{
		Type:    ServiceUnavailable,
		Message: fmt.Sprint("Service unavailable or timed out"),
	}
}
