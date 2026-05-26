package auth

import (
	"net/http"

	"github.com/elazarl/goproxy"
)

var unauthorizedMsg = []byte("407 Proxy Authentication Required")

func BasicUnauthorized(req *http.Request, realm string) *http.Response {
	_ = "STUB: not implemented"
	// TODO(elazar): verify realm is well formed
	return nil
}

var proxyAuthorizationHeader = "Proxy-Authorization"

func auth(req *http.Request, f func(user, passwd string) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Basic returns a basic HTTP authentication handler for requests
//
// You probably want to use auth.ProxyBasic(proxy) to enable authentication for all proxy activities
func Basic(realm string, f func(user, passwd string) bool) goproxy.ReqHandler {
	_ = "STUB: not implemented"
	return *new(goproxy.ReqHandler)
}

// BasicConnect returns a basic HTTP authentication handler for CONNECT requests
//
// You probably want to use auth.ProxyBasic(proxy) to enable authentication for all proxy activities
func BasicConnect(realm string, f func(user, passwd string) bool) goproxy.HttpsHandler {
	_ = "STUB: not implemented"
	return *new(goproxy.HttpsHandler)
}

// ProxyBasic will force HTTP authentication before any request to the proxy is processed
func ProxyBasic(proxy *goproxy.ProxyHttpServer, realm string, f func(user, passwd string) bool) {
	_ = "STUB: not implemented"
	return
}
