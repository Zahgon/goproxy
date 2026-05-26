package goproxy

import (
	"crypto/tls"
	"errors"
	"io"
	"net"
	"net/http"

	"golang.org/x/net/http2"
)

var ErrInvalidH2Frame = errors.New("invalid H2 frame")

// H2Transport is an implementation of RoundTripper that abstracts an entire
// HTTP/2 session, sending all client frames to the server and responses back
// to the client.
type H2Transport struct {
	ClientReader io.Reader
	ClientWriter io.Writer
	TLSConfig    *tls.Config
	Host         string
}

// RoundTrip executes an HTTP/2 session (including all contained streams).
// The request and response are ignored but any error encountered during the
// proxying from the session is returned as a result of the invocation.
func (r *H2Transport) RoundTrip(_ *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ensure that we only advertise HTTP/2 as the accepted protocol.

// Initiate TLS and check remote host name against certificate.

// Send new client preface to match the one parsed in req.

func dial(network, addr string) (c net.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// proxyFrame reads a single frame from the Framer and, when successful, writes
// a ~identical one back to the Framer.
func proxyFrame(fr *http2.Framer) error { _ = "STUB: not implemented"; return nil }

// NOTE: If we want to parse headers, need to handle
// settings where s.ID == http2.SettingHeaderTableSize and
// accordingly update the Framer options.
