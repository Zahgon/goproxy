package goproxy

import (
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"sync"
)

// ConnectActionLiteral defines the action the proxy should take
// when it receives an HTTP CONNECT request from a client.
type ConnectActionLiteral int

const (
	// ConnectAccept instructs the proxy to accept the CONNECT request
	// and establish a transparent TCP tunnel to the destination host.
	// The proxy will forward raw bytes in both directions without inspecting them.
	ConnectAccept ConnectActionLiteral = iota

	// ConnectReject instructs the proxy to reject the CONNECT request
	// and immediately close the connection with the client.
	ConnectReject

	// ConnectMitm instructs the proxy to perform a Man-in-the-Middle (MITM)
	// attack on the CONNECT tunnel. The proxy generates a dynamic TLS certificate
	// for the target host, signed by its CA (see GoproxyCa), and establishes
	// separate TLS connections with both the client and the destination server.
	// All request and response handlers remain active on this intercepted connection.
	ConnectMitm

	// ConnectHijack instructs the proxy to hand the raw net.Conn to the function
	// defined in ConnectAction.Hijack, giving full low-level control of the
	// connection to the caller. The hijack function is responsible for sending
	// an HTTP response (e.g. "HTTP/1.1 200 OK") back to the client.
	ConnectHijack

	// ConnectHTTPMitm is deprecated: use ConnectMitm instead.
	ConnectHTTPMitm

	// ConnectProxyAuthHijack instructs the proxy to hijack the CONNECT connection
	// after a proxy authentication failure, allowing the handler to send
	// a custom authentication challenge or error response to the client.
	ConnectProxyAuthHijack
)

var (
	// OkConnect is a ready-to-use ConnectAction that accepts the CONNECT request
	// and creates a transparent TCP tunnel to the destination host, using the built-in CA.
	OkConnect = &ConnectAction{Action: ConnectAccept, TLSConfig: TLSConfigFromCA(&GoproxyCa)}

	// MitmConnect is a ready-to-use ConnectAction that performs MITM interception,
	// signing dynamic TLS certificates with the built-in CA (GoproxyCa).
	// Use proxy.CertStore to cache generated certificates and save CPU in production.
	MitmConnect = &ConnectAction{Action: ConnectMitm, TLSConfig: TLSConfigFromCA(&GoproxyCa)}

	// HTTPMitmConnect is deprecated: use MitmConnect instead.
	HTTPMitmConnect = &ConnectAction{Action: ConnectHTTPMitm, TLSConfig: TLSConfigFromCA(&GoproxyCa)}

	// RejectConnect is a ready-to-use ConnectAction that rejects the CONNECT request
	// and closes the connection with the client.
	RejectConnect = &ConnectAction{Action: ConnectReject, TLSConfig: TLSConfigFromCA(&GoproxyCa)}
)

var _errorRespMaxLength int64 = 500

const _tlsRecordTypeHandshake = byte(22)

type readBufferedConn struct {
	net.Conn
	r io.Reader
}

func (c *readBufferedConn) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"

	// ConnectAction enables the caller to override the standard connect flow.
	// When Action is ConnectHijack, it is up to the implementer to send the
	// HTTP 200, or any other valid http response back to the client from within the
	// Hijack func.
	return 0, nil
}

type ConnectAction struct {
	Action    ConnectActionLiteral
	Hijack    func(req *http.Request, client net.Conn, ctx *ProxyCtx)
	TLSConfig func(host string, ctx *ProxyCtx) (*tls.Config, error)
}

func stripPort(s string) string { _ = "STUB: not implemented"; return "" }

// ipv6 address example: [2606:4700:4700::1111]:443
// strip '[' and ']'

// ipv4

func (proxy *ProxyHttpServer) dial(ctx *ProxyCtx, network, addr string) (c net.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// if the user didn't specify any dialer, we just use the default one,
// provided by net package

func (proxy *ProxyHttpServer) connectDial(ctx *ProxyCtx, network, addr string) (c net.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

type halfClosable interface {
	net.Conn
	CloseWrite() error
	CloseRead() error
}

var _ halfClosable = (*net.TCPConn)(nil)

func (proxy *ProxyHttpServer) handleHttps(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// If found a result, break the loop immediately

// Make sure to close the underlying TCP socket.
// CloseRead() and CloseWrite() keep it open until its timeout,
// causing error when there are thousands of requests.

// There is a race with the runtime here. In the case where the
// connection to the target site times out, we cannot control which
// io.Copy loop will receive the timeout signal first. This means
// that in some cases the error passed to the ConnErrorHandler will
// be the timeout error, and in other cases it will be an error raised
// by the use of a closed network connection.
//
// 2020/05/28 23:42:17 [001] WARN: Error copying to client: read tcp 127.0.0.1:33742->127.0.0.1:34763: i/o timeout
// 2020/05/28 23:42:17 [001] WARN: Error copying to client: read tcp 127.0.0.1:45145->127.0.0.1:60494: use of closed
//                                                          network connection
//
// It's also not possible to synchronize these connection closures due to
// TCP connections which are half-closed. When this happens, only the one
// side of the connection breaks out of its io.Copy loop. The other side
// of the connection remains open until it either times out or is reset by
// the client.

// this goes in a separate goroutine, so that the net/http server won't think we're
// still handling the request even after hijacking the connection. Those HTTP CONNECT
// request can take forever, and the server will be stuck when "closed".
// TODO: Allow Server.Close() mechanism to shut down this connection as nicely as possible

// Check if this is an HTTP or an HTTPS MITM request

// Create a TLS connection over the TCP connection

// since we're converting the request, need to carry over the
// original connecting IP as well

// Since we handled the request parsing by our own, we manually
// need to set a cancellable context when we finished the request
// processing (same behaviour of the stdlib)

// explicitly discard request body to avoid data races in certain RoundTripper implementations
// see https://github.com/golang/go/issues/61596#issuecomment-1652345131

// Bug fix which goproxy fails to provide request
// information URL in the context when does HTTPS MITM

// Handle HTTP/2 connections.

// NOTE: As of 1.22, golang's http module will not recognize or
// parse the HTTP Body for PRI requests. This leaves the body of
// the http2.ClientPreface ("SM\r\n\r\n") on the wire which we need
// to clear before setting up the connection.

// Return chunked encoded response when we don't know the length of the resp, if the body
// has been modified by the response handler or if there is no content length in the response.
// We include 0 in resp.ContentLength <= 0 because 0 is the field zero value and some user
// might incorrectly leave it instead of setting it to -1 when the length is unknown (but we
// also check that the Content-Length header is empty, so there is no issue with empty bodies).

// The MITM'd client speaks HTTP/1.1, but the upstream
// response may have been received over HTTP/2. Normalize
// the protocol version so resp.Write() produces a valid
// HTTP/1.1 status line.

// According to resp.Body documentation:
// As of Go 1.12, the Body will also implement io.Writer
// on a successful "101 Switching Protocols" response,
// as used by WebSockets and HTTP/2's "h2c" mode.

// Set Body to nil so resp.Write only writes the headers
// and returns immediately without blocking on the body
// (or else we wouldn't be able to proxy WebSocket data).

func httpError(w io.WriteCloser, ctx *ProxyCtx, err error) { _ = "STUB: not implemented"; return }

func copyOrWarn(ctx *ProxyCtx, dst io.Writer, src io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

// Discard closed connection errors

func copyAndClose(ctx *ProxyCtx, dst, src halfClosable, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

// Fully close dst to unblock any goroutine blocked on
// io.Copy reading from it. Half-close (CloseWrite/CloseRead)
// would not interrupt a pending read, leaving the other
// goroutine stuck and the client connection never closed.

func dialerFromEnv(proxy *ProxyHttpServer) func(network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return nil
}

// NewConnectDialToProxy returns a dial function that establishes TCP connections
// through an upstream HTTP/HTTPS proxy using the CONNECT method.
// Use it to set proxy.ConnectDial when chaining two proxy servers.
// For authentication or other CONNECT request modifications, use NewConnectDialToProxyWithHandler instead.
func (proxy *ProxyHttpServer) NewConnectDialToProxy(httpsProxy string) func(network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return nil
}

// NewConnectDialToProxyWithHandler returns a dial function that establishes TCP connections
// through an upstream HTTP/HTTPS proxy using the CONNECT method, calling connectReqHandler
// before sending the CONNECT request. Use connectReqHandler to add headers such as
// Proxy-Authorization to authenticate with the upstream proxy.
// If connectReqHandler is nil, the behavior is identical to NewConnectDialToProxy.
func (proxy *ProxyHttpServer) NewConnectDialToProxyWithHandler(
	httpsProxy string,
	connectReqHandler func(req *http.Request),
) func(network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return nil
}

// Read response.
// Okay to use and discard buffered reader here, because
// TLS server will not speak until spoken to.

// Read response.
// Okay to use and discard buffered reader here, because
// TLS server will not speak until spoken to.

// TLSConfigFromCA returns a TLSConfig function that generates dynamic TLS certificates
// for each target host, signed by the given CA certificate.
// The generated certificates are used during MITM interception (ConnectMitm).
// If a CertStorage is set on the ProxyCtx, certificates are cached and reused to save CPU.
func TLSConfigFromCA(ca *tls.Certificate) func(host string, ctx *ProxyCtx) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil
}

func (proxy *ProxyHttpServer) initializeTLSconnection(
	ctx *ProxyCtx,
	targetConn net.Conn,
	tlsConfig *tls.Config,
	addr string,
) (net.Conn, error) {
	_ = "STUB: not implemented"
	// Infer target ServerName, it's a copy of implementation inside tls.Dial()
	return *new(net.Conn), nil
}

// Make a copy to avoid polluting argument or default.
