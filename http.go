package goproxy

import (
	"net/http"
)

func (proxy *ProxyHttpServer) handleHttp(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// http.ResponseWriter will take care of filling the correct response length
// Setting it now, might impose wrong value, contradicting the actual new
// body the user returned.
// We keep the original body to remove the header only if things changed.
// This will prevent problems with HEAD requests where there's no body, yet,
// the Content-Length header should be set.

// Announce trailers known at this point (HTTP/1.1 with pre-announced
// Trailer header). Setting "Trailer" before WriteHeader makes
// http.Server commit to chunked encoding (h1) or a trailing HEADERS
// frame (h2), which is required for any trailers to be forwarded.
// Mirrors net/http/httputil.ReverseProxy.

// We have already written the "101 Switching Protocols" response,
// now we hijack the connection to send WebSocket data

// Content-Type header may also contain charset definition, so here we need to check the prefix.
// Transfer-Encoding can be a list of comma separated values, so we use Contains() for it.

// server-side events, flush the buffered data to the client.

// Forward upstream response trailers. Two cases:
//   1. resp.Trailer count == announcedTrailers: every trailer was
//      pre-announced, so http.Server is already looking for them
//      under the unprefixed names — write values there.
//   2. resp.Trailer count > announcedTrailers (HTTP/2 servers, or
//      late additions): use http.TrailerPrefix so http.Server emits
//      them as trailers without needing the leading announcement.
//      We still need a Flush below to force chunked encoding for
//      bodies short enough that http.Server would otherwise inline
//      them with Content-Length and silently drop trailers.

// Force chunking even when the body is small / fully buffered.
