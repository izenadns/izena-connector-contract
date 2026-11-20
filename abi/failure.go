package abi

import (
	"context"
	"errors"

	"github.com/izenadns/izena-connector-contract/connector"
)

type Failure struct {
	Kind   string
	Status int
	Code   string
	Detail string
	Cause  string
}

var failureKinds = []struct {
	name string
	err  error
}{
	{"invalid_url", connector.ErrInvalidServerURL},
	{"invalid_tls", connector.ErrInvalidServerTLS},
	{"invalid_server", connector.ErrInvalidServer},
	{"authentication", connector.ErrAuthentication},
	{"not_found", connector.ErrNotFound},
	{"conflict", connector.ErrConflict},
	{"validation", connector.ErrValidation},
	{"timeout", connector.ErrTimeout},
	{"unavailable", connector.ErrUnavailable},
	{"tls", connector.ErrTLS},
	{"redirect", connector.ErrRedirect},
	{"malformed", connector.ErrMalformed},
	{"too_large", connector.ErrResponseTooLarge},
	{"unsupported", connector.ErrUnsupported},
}

func EncodeFailure(err error) *Failure {
	if err == nil {
		return nil
	}
	result := &Failure{Kind: "unavailable"}
	for _, kind := range failureKinds {
		if errors.Is(err, kind.err) {
			result.Kind = kind.name
			break
		}
	}
	var classified *connector.Error
	if errors.As(err, &classified) {
		result.Status, result.Code, result.Detail = classified.StatusCode, classified.Code, classified.Detail
	}
	switch {
	case errors.Is(err, context.Canceled):
		result.Cause = "canceled"
	case errors.Is(err, context.DeadlineExceeded):
		result.Cause = "deadline"
	}
	return result
}

func (f *Failure) Decode() error {
	if f == nil {
		return nil
	}
	kind := connector.ErrUnavailable
	for _, known := range failureKinds {
		if known.name == f.Kind {
			kind = known.err
			break
		}
	}
	var cause error
	switch f.Cause {
	case "canceled":
		cause = context.Canceled
	case "deadline":
		cause = context.DeadlineExceeded
	}
	return &connector.Error{Kind: kind, StatusCode: f.Status, Code: f.Code, Detail: f.Detail, Cause: cause}
}
