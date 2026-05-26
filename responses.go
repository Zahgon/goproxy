package goproxy

import (
	"net/http"
)

// Will generate a valid http response to the given request the response will have
// the given contentType, and http status.
// Typical usage, refuse to process requests to local addresses:
//
//	proxy.OnRequest(IsLocalHost()).DoFunc(func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request,*http.Response) {
//		return nil,NewResponse(r,goproxy.ContentTypeHtml,http.StatusUnauthorized,
//			`<!doctype html><html><head><title>Can't use proxy for local addresses</title></head><body/></html>`)
//	})
func NewResponse(r *http.Request, contentType string, status int, body string) *http.Response {
	_ = "STUB: not implemented"
	return nil
}

const (
	// ContentTypeText is the MIME type for plain text responses.
	ContentTypeText = "text/plain"
	// ContentTypeHtml is the MIME type for HTML responses.
	ContentTypeHtml = "text/html"
)

// Alias for NewResponse(r,ContentTypeText,http.StatusAccepted,text).
func TextResponse(r *http.Request, text string) *http.Response {
	_ = "STUB: not implemented"
	return nil
}
