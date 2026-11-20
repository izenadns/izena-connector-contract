package abi

import (
	"net/http"
	"net/url"
)

// HTTPRequest addresses a path beneath the configured upstream URL. Parts are
// decoded path segments. The host escapes them and supplies credentials from
// the connection. Limit bounds the response body, in bytes.
type HTTPRequest struct {
	Method  string
	Parts   []string
	Query   url.Values
	Headers http.Header
	Body    []byte
	Limit   int
	Action  bool
}

type HTTPResponse struct {
	Status  int
	Headers http.Header `json:",omitempty"`
	Body    []byte
	Error   *Failure `json:",omitempty"`
}
