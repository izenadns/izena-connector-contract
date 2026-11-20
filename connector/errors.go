package connector

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrInvalidServerURL = errors.New("invalid provider URL")
	ErrInvalidServerTLS = errors.New("invalid provider CA bundle")
	ErrInvalidServer    = errors.New("invalid provider configuration")
	ErrAuthentication   = errors.New("provider authentication failed")
	ErrNotFound         = errors.New("provider resource not found")
	ErrConflict         = errors.New("provider request conflicted")
	ErrValidation       = errors.New("provider request was rejected")
	ErrTimeout          = errors.New("provider request timed out")
	ErrUnavailable      = errors.New("provider is unavailable")
	ErrTLS              = errors.New("provider TLS negotiation failed")
	ErrRedirect         = errors.New("provider redirect rejected")
	ErrMalformed        = errors.New("provider response was malformed")
	ErrResponseTooLarge = errors.New("provider response exceeded its size limit")
	ErrUnsupported      = errors.New("provider operation is unsupported")
)

// Error classifies a provider failure. Code should be a fixed diagnostic code.
// Detail is reserved for sanitized upstream rejection reasons. Neither field
// may contain credentials or arbitrary dependency errors.
type Error struct {
	Kind       error
	StatusCode int
	Code       string
	Detail     string
	// Cause may carry context cancellation or a deadline.
	Cause error
}

func (e *Error) Error() string {
	if e == nil || e.Kind == nil {
		return "provider request failed"
	}
	if e.Code == "" {
		return e.Kind.Error()
	}
	return fmt.Sprintf("%s (%s)", e.Kind, e.Code)
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}
	switch {
	case errors.Is(e.Cause, context.Canceled):
		return errors.Join(e.Kind, context.Canceled)
	case errors.Is(e.Cause, context.DeadlineExceeded):
		return errors.Join(e.Kind, context.DeadlineExceeded)
	default:
		return e.Kind
	}
}

// ErrorDetail returns the sanitized upstream rejection reason, when available.
func ErrorDetail(err error) string {
	var value *Error
	if !errors.As(err, &value) || value == nil {
		return ""
	}
	return value.Detail
}

// ErrorStatusCode returns the upstream HTTP status, when available.
func ErrorStatusCode(err error) int {
	var value *Error
	if !errors.As(err, &value) || value == nil {
		return 0
	}
	return value.StatusCode
}
