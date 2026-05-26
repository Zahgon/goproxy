// Original implementation from abourget/goproxy, adapted for use as an extension.
// HAR specification: http://www.softwareishard.com/blog/har-12-spec/
package har

import (
	"io"
	"net/http"
	"time"

	"github.com/elazarl/goproxy"
)

type Har struct {
	Log Log `json:"log"`
}

type Log struct {
	Version string   `json:"version"`
	Creator Creator  `json:"creator"`
	Browser *Browser `json:"browser,omitempty"`
	Pages   []Page   `json:"pages,omitempty"`
	Entries []Entry  `json:"entries"`
	Comment string   `json:"comment,omitempty"`
}

func New() *Har { _ = "STUB: not implemented"; return nil }

func makeNewEntries() []Entry { _ = "STUB: not implemented"; return nil }

type Creator struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Comment string `json:"comment,omitempty"`
}

type Browser struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Comment string `json:"comment,omitempty"`
}

type Page struct {
	ID              string      `json:"id,omitempty"`
	StartedDateTime time.Time   `json:"startedDateTime"`
	Title           string      `json:"title"`
	PageTimings     PageTimings `json:"pageTimings"`
	Comment         string      `json:"comment,omitempty"`
}

type Entry struct {
	PageRef         string    `json:"pageref,omitempty"`
	StartedDateTime time.Time `json:"startedDateTime"`
	Time            int64     `json:"time"`
	Request         *Request  `json:"request"`
	Response        *Response `json:"response"`
	Cache           Cache     `json:"cache"`
	Timings         Timings   `json:"timings"`
	ServerIpAddress string    `json:"serverIpAddress,omitempty"`
	Connection      string    `json:"connection,omitempty"`
	Comment         string    `json:"comment,omitempty"`
}

type Cache struct {
	BeforeRequest *CacheEntry `json:"beforeRequest,omitempty"`
	AfterRequest  *CacheEntry `json:"afterRequest,omitempty"`
}

type CacheEntry struct {
	Expires    string `json:"expires,omitempty"`
	LastAccess string `json:"lastAccess"`
	ETag       string `json:"eTag"`
	HitCount   int    `json:"hitCount"`
	Comment    string `json:"comment,omitempty"`
}

type Request struct {
	Method      string          `json:"method"`
	Url         string          `json:"url"`
	HttpVersion string          `json:"httpVersion"`
	Cookies     []Cookie        `json:"cookies"`
	Headers     []NameValuePair `json:"headers"`
	QueryString []NameValuePair `json:"queryString"`
	PostData    *PostData       `json:"postData,omitempty"`
	BodySize    int64           `json:"bodySize"`
	HeadersSize int64           `json:"headersSize"`
}

func (entry *Entry) fillIPAddress(req *http.Request) { _ = "STUB: not implemented"; return }

// try to parse the host as an IP address

// Shared utility function for reading body content
func readBody(ctx *goproxy.ProxyCtx, body io.ReadCloser) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Shared function for handling mime types
func parseMediaType(ctx *goproxy.ProxyCtx, header http.Header) string {
	_ = "STUB: not implemented"
	return ""
}

func parsePostData(ctx *goproxy.ProxyCtx, req *http.Request) *PostData {
	_ = "STUB: not implemented"
	return nil
}

type Response struct {
	Status      int             `json:"status"`
	StatusText  string          `json:"statusText"`
	HttpVersion string          `json:"httpVersion"`
	Cookies     []Cookie        `json:"cookies"`
	Headers     []NameValuePair `json:"headers"`
	Content     Content         `json:"content"`
	RedirectUrl string          `json:"redirectURL"`
	BodySize    int64           `json:"bodySize"`
	HeadersSize int64           `json:"headersSize"`
	Comment     string          `json:"comment,omitempty"`
}

func parseResponse(ctx *goproxy.ProxyCtx) *Response { _ = "STUB: not implemented"; return nil }

func parseRequest(ctx *goproxy.ProxyCtx) *Request { _ = "STUB: not implemented"; return nil }

func parseStringArrMap(stringArrMap map[string][]string) []NameValuePair {
	_ = "STUB: not implemented"
	return nil
}

// Use original key if unescaping fails

// Use original joined values if unescaping fails

func parseCookies(cookies []*http.Cookie) []Cookie { _ = "STUB: not implemented"; return nil }

type Cookie struct {
	Name     string     `json:"name"`
	Value    string     `json:"value"`
	Path     string     `json:"path,omitempty"`
	Domain   string     `json:"domain,omitempty"`
	Expires  *time.Time `json:"expires,omitempty"`
	HttpOnly bool       `json:"httpOnly,omitempty"`
	Secure   bool       `json:"secure,omitempty"`
}

type NameValuePair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type PostData struct {
	MimeType string          `json:"mimeType"`
	Params   []PostDataParam `json:"params,omitempty"`
	Text     string          `json:"text,omitempty"`
	Comment  string          `json:"comment,omitempty"`
}

type PostDataParam struct {
	Name        string `json:"name"`
	Value       string `json:"value,omitempty"`
	FileName    string `json:"fileName,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

type Content struct {
	Size        int    `json:"size"`
	Compression int    `json:"compression,omitempty"`
	MimeType    string `json:"mimeType"`
	Text        string `json:"text,omitempty"`
	Encoding    string `json:"encoding,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

type PageTimings struct {
	OnContentLoad int64  `json:"onContentLoad"`
	OnLoad        int64  `json:"onLoad"`
	Comment       string `json:"comment,omitempty"`
}

type Timings struct {
	Dns     int64  `json:"dns,omitempty"`
	Blocked int64  `json:"blocked,omitempty"`
	Connect int64  `json:"connect,omitempty"`
	Send    int64  `json:"send"`
	Wait    int64  `json:"wait"`
	Receive int64  `json:"receive"`
	Ssl     int64  `json:"ssl,omitempty"`
	Comment string `json:"comment,omitempty"`
}
