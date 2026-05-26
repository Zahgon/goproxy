package http1parser

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"net/textproto"
)

type RequestReader struct {
	preventCanonicalization bool
	reader                  *bufio.Reader
	// Used only when preventCanonicalization value is true
	cloned *bytes.Buffer
}

func NewRequestReader(preventCanonicalization bool, conn io.Reader) *RequestReader {
	_ = "STUB: not implemented"
	return nil
}

// IsEOF returns true if there is no more data that can be read from the
// buffer and the underlying connection is closed.
func (r *RequestReader) IsEOF() bool { _ = "STUB: not implemented"; return false }

// Reader is used to take over the buffered connection data
// (e.g. with HTTP/2 data).
// After calling this function, make sure to consume all the data related
// to the current request.
func (r *RequestReader) Reader() *bufio.Reader { _ = "STUB: not implemented"; return nil }

func (r *RequestReader) ReadRequest() (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Just call the HTTP library function if the preventCanonicalization
	// configuration is disabled
}

// Rewrite header keys to the non-canonical parsed value

func getRequestReader(r *bufio.Reader, cloned *bytes.Buffer) *textproto.Reader {
	_ = "STUB: not implemented"
	// "Cloned" buffer uses the raw connection as the data source.
	// However, the *bufio.Reader can read also bytes of another unrelated
	// request on the same connection, since it's buffered, so we have to
	// ignore them before passing the data to our headers parser.
	// Data related to the next request will remain inside the buffer for
	// later usage.
	return nil
}
