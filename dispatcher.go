package goproxy

import (
	"net"
	"net/http"
	"regexp"
)

// ReqCondition.HandleReq will decide whether or not to use the ReqHandler on an HTTP request
// before sending it to the remote server.
type ReqCondition interface {
	RespCondition
	HandleReq(req *http.Request, ctx *ProxyCtx) bool
}

// RespCondition.HandleReq will decide whether or not to use the RespHandler on an HTTP response
// before sending it to the proxy client. Note that resp might be nil, in case there was an
// error sending the request.
type RespCondition interface {
	HandleResp(resp *http.Response, ctx *ProxyCtx) bool
}

// ReqConditionFunc.HandleReq(req,ctx) <=> ReqConditionFunc(req,ctx).
type ReqConditionFunc func(req *http.Request, ctx *ProxyCtx) bool

// RespConditionFunc.HandleResp(resp,ctx) <=> RespConditionFunc(resp,ctx).
type RespConditionFunc func(resp *http.Response, ctx *ProxyCtx) bool

func (c ReqConditionFunc) HandleReq(req *http.Request, ctx *ProxyCtx) bool {
	_ = "STUB: not implemented"
	return false

	// ReqConditionFunc cannot test responses. It only satisfies RespCondition interface so that
	// to be usable as RespCondition.
}

func (c ReqConditionFunc) HandleResp(resp *http.Response, ctx *ProxyCtx) bool {
	_ = "STUB: not implemented"
	return false
}

func (c RespConditionFunc) HandleResp(resp *http.Response, ctx *ProxyCtx) bool {
	_ = "STUB: not implemented"
	return false

	// UrlHasPrefix returns a ReqCondition checking wether the destination URL the proxy client has requested
	// has the given prefix, with or without the host.
	// For example UrlHasPrefix("host/x") will match requests of the form 'GET host/x', and will match
	// requests to url 'http://host/x'
}

func UrlHasPrefix(prefix string) ReqConditionFunc {
	_ = "STUB: not implemented"
	return *new(ReqConditionFunc)
}

// Make sure to include the / as the first path character when we do a match
// using the host

// We use the original value to distinguish between "" and "/" in the user specified string

// Scheme value is something like "https", we must include the :// characters

// UrlIs returns a ReqCondition, testing whether or not the request URL is one of the given strings
// with or without the host prefix.
// UrlIs("google.com/","foo") will match requests 'GET /' to 'google.com', requests `'GET google.com/' to
// any host, and requests of the form 'GET foo'.
func UrlIs(urls ...string) ReqConditionFunc {
	_ = "STUB: not implemented"
	return *new(ReqConditionFunc)
}

// ReqHostMatches returns a ReqCondition, testing whether the host to which the request was directed to matches
// any of the given regular expressions.
func ReqHostMatches(regexps ...*regexp.Regexp) ReqConditionFunc {
	_ = "STUB: not implemented"
	return *new(ReqConditionFunc)
}

// ReqHostIs returns a ReqCondition, testing whether the host to which the request is directed to equal
// to one of the given strings.
func ReqHostIs(hosts ...string) ReqConditionFunc {
	_ = "STUB: not implemented"
	return *new(ReqConditionFunc)
}

// IsLocalHost checks whether the destination host is localhost.
var IsLocalHost ReqConditionFunc = func(req *http.Request, ctx *ProxyCtx) bool {
	h := req.URL.Hostname()
	if h == "localhost" {
		return true
	}
	if ip := net.ParseIP(h); ip != nil {
		return ip.IsLoopback()
	}

	// In case of IPv6 without a port number Hostname() sometimes returns the invalid value.
	if ip := net.ParseIP(req.URL.Host); ip != nil {
		return ip.IsLoopback()
	}

	return false
}

// UrlMatches returns a ReqCondition testing whether the destination URL
// of the request matches the given regexp, with or without prefix.
func UrlMatches(re *regexp.Regexp) ReqConditionFunc {
	_ = "STUB: not implemented"
	return *new(ReqConditionFunc)
}

// DstHostIs returns a ReqCondition testing wether the host in the request url is the given string.
func DstHostIs(host string) ReqConditionFunc {
	_ = "STUB: not implemented"
	// Make sure to perform a case-insensitive host check
	return *new(ReqConditionFunc)
}

// Check if the user specified a custom port that we need to match

// Check port matching only if it was specified

// SrcIpIs returns a ReqCondition testing whether the source IP of the request is one of the given strings.
func SrcIpIs(ips ...string) ReqCondition { _ = "STUB: not implemented"; return *new(ReqCondition) }

// Not returns a ReqCondition negating the given ReqCondition.
func Not(r ReqCondition) ReqConditionFunc { _ = "STUB: not implemented"; return *new(ReqConditionFunc) }

// ContentTypeIs returns a RespCondition testing whether the HTTP response has Content-Type header equal
// to one of the given strings.
func ContentTypeIs(typ string, types ...string) RespCondition {
	_ = "STUB: not implemented"
	return *new(RespCondition)
}

// StatusCodeIs returns a RespCondition, testing whether or not the HTTP status
// code is one of the given ints.
func StatusCodeIs(codes ...int) RespCondition {
	_ = "STUB: not implemented"
	return *new(RespCondition)
}

// ProxyHttpServer.OnRequest Will return a temporary ReqProxyConds struct, aggregating the given condtions.
// You will use the ReqProxyConds struct to register a ReqHandler, that would filter
// the request, only if all the given ReqCondition matched.
// Typical usage:
//
//	proxy.OnRequest(UrlIs("example.com/foo"),UrlMatches(regexp.MustParse(`.*\.exampl.\com\./.*`)).Do(...)
func (proxy *ProxyHttpServer) OnRequest(conds ...ReqCondition) *ReqProxyConds {
	_ = "STUB: not implemented"
	return nil
}

// ReqProxyConds aggregate ReqConditions for a ProxyHttpServer.
// Upon calling Do, it will register a ReqHandler that would
// handle the request if all conditions on the HTTP request are met.
type ReqProxyConds struct {
	proxy    *ProxyHttpServer
	reqConds []ReqCondition
}

// DoFunc is equivalent to proxy.OnRequest().Do(FuncReqHandler(f)).
func (pcond *ReqProxyConds) DoFunc(f func(req *http.Request, ctx *ProxyCtx) (*http.Request, *http.Response)) {
	_ = "STUB: not implemented"
	return
}

// ReqProxyConds.Do will register the ReqHandler on the proxy,
// the ReqHandler will handle the HTTP request if all the conditions
// aggregated in the ReqProxyConds are met. Typical usage:
//
//	proxy.OnRequest().Do(handler) // will call handler.Handle(req,ctx) on every request to the proxy
//	proxy.OnRequest(cond1,cond2).Do(handler)
//	// given request to the proxy, will test if cond1.HandleReq(req,ctx) && cond2.HandleReq(req,ctx) are true
//	// if they are, will call handler.Handle(req,ctx)
func (pcond *ReqProxyConds) Do(h ReqHandler) { _ = "STUB: not implemented"; return }

// HandleConnect is used when proxy receives an HTTP CONNECT request,
// it'll then use the HttpsHandler to determine what should it
// do with this request. The handler returns a ConnectAction struct, the Action field in the ConnectAction
// struct returned will determine what to do with this request. ConnectAccept will simply accept the request
// forwarding all bytes from the client to the remote host, ConnectReject will close the connection with the
// client, and ConnectMitm, will assume the underlying connection is an HTTPS connection, and will use Man
// in the Middle attack to eavesdrop the connection. All regular handler will be active on this eavesdropped
// connection.
// The ConnectAction struct contains possible tlsConfig that will be used for eavesdropping. If nil, the proxy
// will use the default tls configuration.
//
//	proxy.OnRequest().HandleConnect(goproxy.AlwaysReject) // rejects all CONNECT requests
func (pcond *ReqProxyConds) HandleConnect(h HttpsHandler) { _ = "STUB: not implemented"; return }

// HandleConnectFunc is equivalent to HandleConnect,
// for example, accepting CONNECT request if they contain a password in header
//
//	io.WriteString(h,password)
//	passHash := h.Sum(nil)
//	proxy.OnRequest().HandleConnectFunc(func(host string, ctx *ProxyCtx) (*ConnectAction, string) {
//		c := sha1.New()
//		io.WriteString(c,ctx.Req.Header.Get("X-GoProxy-Auth"))
//		if c.Sum(nil) == passHash {
//			return OkConnect, host
//		}
//		return RejectConnect, host
//	})
func (pcond *ReqProxyConds) HandleConnectFunc(f func(host string, ctx *ProxyCtx) (*ConnectAction, string)) {
	_ = "STUB: not implemented"
	return
}

// HijackConnect registers a handler that takes full control of the raw net.Conn
// for CONNECT requests that match the aggregated conditions.
// The handler receives the original HTTP request, the raw client connection, and the proxy context.
// It is the handler's responsibility to write an HTTP response (e.g. "HTTP/1.1 200 OK\r\n\r\n")
// and close the connection when done.
func (pcond *ReqProxyConds) HijackConnect(f func(req *http.Request, client net.Conn, ctx *ProxyCtx)) {
	_ = "STUB: not implemented"
	return
}

// ProxyConds is used to aggregate RespConditions for a ProxyHttpServer.
// Upon calling ProxyConds.Do, it will register a RespHandler that would
// handle the HTTP response from remote server if all conditions on the HTTP response are met.
type ProxyConds struct {
	proxy    *ProxyHttpServer
	reqConds []ReqCondition
	respCond []RespCondition
}

// ProxyConds.DoFunc is equivalent to proxy.OnResponse().Do(FuncRespHandler(f)).
func (pcond *ProxyConds) DoFunc(f func(resp *http.Response, ctx *ProxyCtx) *http.Response) {
	_ = "STUB: not implemented"
	return
}

// ProxyConds.Do will register the RespHandler on the proxy, h.Handle(resp,ctx) will be called on every
// request that matches the conditions aggregated in pcond.
func (pcond *ProxyConds) Do(h RespHandler) { _ = "STUB: not implemented"; return }

// OnResponse is used when adding a response-filter to the HTTP proxy, usual pattern is
//
//	proxy.OnResponse(cond1,cond2).Do(handler) // handler.Handle(resp,ctx) will be used
//				// if cond1.HandleResp(resp) && cond2.HandleResp(resp)
func (proxy *ProxyHttpServer) OnResponse(conds ...RespCondition) *ProxyConds {
	_ = "STUB: not implemented"
	return nil
}

// AlwaysMitm is a HttpsHandler that always eavesdrop https connections, for example to
// eavesdrop all https connections to www.google.com, we can use
//
//	proxy.OnRequest(goproxy.ReqHostIs("www.google.com")).HandleConnect(goproxy.AlwaysMitm)
var AlwaysMitm FuncHttpsHandler = func(host string, ctx *ProxyCtx) (*ConnectAction, string) {
	return MitmConnect, host
}

// AlwaysReject is a HttpsHandler that drops any CONNECT request, for example, this code will disallow
// connections to hosts on any other port than 443
//
//	proxy.OnRequest(goproxy.Not(goproxy.ReqHostMatches(regexp.MustCompile(":443$"))).
//		HandleConnect(goproxy.AlwaysReject)
var AlwaysReject FuncHttpsHandler = func(host string, ctx *ProxyCtx) (*ConnectAction, string) {
	return RejectConnect, host
}

// HandleBytes will return a RespHandler that read the entire body of the request
// to a byte array in memory, would run the user supplied f function on the byte arra,
// and will replace the body of the original response with the resulting byte array.
func HandleBytes(f func(b []byte, ctx *ProxyCtx) []byte) RespHandler {
	_ = "STUB: not implemented"
	return *new(RespHandler)
}
