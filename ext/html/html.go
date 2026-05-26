// extension to goproxy that will allow you to easily filter web browser related content.
package goproxy_html

import (
	"io"

	"github.com/elazarl/goproxy"
)

var IsHtml goproxy.RespCondition = goproxy.ContentTypeIs("text/html")

var IsCss goproxy.RespCondition = goproxy.ContentTypeIs("text/css")

var IsJavaScript goproxy.RespCondition = goproxy.ContentTypeIs("text/javascript",
	"application/javascript")

var IsJson goproxy.RespCondition = goproxy.ContentTypeIs("text/json")

var IsXml goproxy.RespCondition = goproxy.ContentTypeIs("text/xml")

var IsWebRelatedText goproxy.RespCondition = goproxy.ContentTypeIs(
	"text/html",
	"text/css",
	"text/javascript", "application/javascript",
	"text/xml",
	"text/json",
)

// HandleString will receive a function that filters a string, and will convert the
// request body to a utf8 string, according to the charset specified in the Content-Type
// header.
// guessing Html charset encoding from the <META> tags is not yet implemented.
func HandleString(f func(s string, ctx *goproxy.ProxyCtx) string) goproxy.RespHandler {
	_ = "STUB: not implemented"
	return *new(goproxy.RespHandler)
}

// Will receive an input stream which would convert the response to utf-8
// The given function must close the reader r, in order to close the response body.
func HandleStringReader(f func(r io.Reader, ctx *goproxy.ProxyCtx) io.Reader) goproxy.RespHandler {
	_ = "STUB: not implemented"
	return *new(goproxy.RespHandler)
}

// Pass UTF-8 data to the callback f() function and convert its
// result back to the original encoding

//no translation is needed, already at utf-8

type readFirstCloseBoth struct {
	r io.ReadCloser
	c io.Closer
}

func (rfcb *readFirstCloseBoth) Read(b []byte) (nr int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rfcb *readFirstCloseBoth) Close() error { _ = "STUB: not implemented"; return nil }
