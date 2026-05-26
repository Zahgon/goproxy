package goproxy

import (
	"io"
	"net"
	"net/http"
)

func headerContains(header http.Header, name string, value string) bool {
	_ = "STUB: not implemented"
	return false
}

func isWebSocketHandshake(header http.Header) bool { _ = "STUB: not implemented"; return false }

func (proxy *ProxyHttpServer) hijackConnection(ctx *ProxyCtx, w http.ResponseWriter) (net.Conn, error) {
	_ = "STUB: not implemented"
	// Connect to Client
	return *new(net.Conn), nil
}

func (proxy *ProxyHttpServer) proxyWebsocket(ctx *ProxyCtx, remoteConn io.ReadWriter, proxyClient io.ReadWriter) {
	_ = "STUB: not implemented"
	// 2 is the number of goroutines, this code is implemented according to
	// https://stackoverflow.com/questions/52031332/wait-for-one-goroutine-to-finish
	return
}

// Wait until one end closes the connection
