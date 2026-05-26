package http1parser

import (
	"errors"
	"net/textproto"
)

var ErrBadProto = errors.New("bad protocol")

// Http1ExtractHeaders is an HTTP/1.0 and HTTP/1.1 header-only parser,
// to extract the original header names for the received request.
// Fully inspired by readMIMEHeader() in
// https://github.com/golang/go/blob/master/src/net/textproto/reader.go
func Http1ExtractHeaders(r *textproto.Reader) ([]string, error) {
	_ = "STUB: not implemented"
	// Discard first line, it doesn't contain useful information, and it has
	// already been validated in http.ReadRequest()
	return nil, nil
}

// The first line cannot start with a leading space.

// We have finished to parse the headers if we receive empty
// data without an error

// Key ends at first colon.
