package main

import (
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/transport"
)

type FileStream struct {
	path string
	f    *os.File
}

func NewFileStream(path string) *FileStream { _ = "STUB: not implemented"; return nil }

func (fs *FileStream) Write(b []byte) (nr int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (fs *FileStream) Close() error { _ = "STUB: not implemented"; return nil }

type Meta struct {
	req      *http.Request
	resp     *http.Response
	err      error
	t        time.Time
	sess     int64
	bodyPath string
	from     string
}

func fprintf(nr *int64, err *error, w io.Writer, pat string, a ...any) {
	_ = "STUB: not implemented"
	return
}

func write(nr *int64, err *error, w io.Writer, b []byte) { _ = "STUB: not implemented"; return }

func (m *Meta) WriteTo(w io.Writer) (nr int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// note the empty response

// HttpLogger is an asynchronous HTTP request/response logger. It traces
// requests and responses headers in a "log" file in logger directory and dumps
// their bodies in files prefixed with the session identifiers.
// Close it to ensure pending items are correctly logged.
type HttpLogger struct {
	path  string
	c     chan *Meta
	errch chan error
}

func NewLogger(basepath string) (*HttpLogger, error) { _ = "STUB: not implemented"; return nil, nil }

func (logger *HttpLogger) LogResp(resp *http.Response, ctx *goproxy.ProxyCtx) {
	_ = "STUB: not implemented"
	return
}

var emptyResp = &http.Response{}
var emptyReq = &http.Request{}

func (logger *HttpLogger) LogReq(req *http.Request, ctx *goproxy.ProxyCtx) {
	_ = "STUB: not implemented"
	return
}

func (logger *HttpLogger) LogMeta(m *Meta) { _ = "STUB: not implemented"; return }

func (logger *HttpLogger) Close() error { _ = "STUB: not implemented"; return nil }

// TeeReadCloser extends io.TeeReader by allowing reader and writer to be
// closed.
type TeeReadCloser struct {
	r io.Reader
	w io.WriteCloser
	c io.Closer
}

func NewTeeReadCloser(r io.ReadCloser, w io.WriteCloser) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

func (t *TeeReadCloser) Read(b []byte) (int, error) {
	_ = "STUB: not implemented"

	// Close attempts to close the reader and write. It returns an error if both
	// failed to Close.
	return 0, nil
}

func (t *TeeReadCloser) Close() error { _ = "STUB: not implemented"; return nil }

// stoppableListener serves stoppableConn and tracks their lifetime to notify
// when it is safe to terminate the application.
type stoppableListener struct {
	net.Listener
	sync.WaitGroup
}

type stoppableConn struct {
	net.Conn
	wg *sync.WaitGroup
}

func newStoppableListener(l net.Listener) *stoppableListener { _ = "STUB: not implemented"; return nil }

func (sl *stoppableListener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (sc *stoppableConn) Close() error { _ = "STUB: not implemented"; return nil }

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("l", ":8080", "on which address should the proxy listen")
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose
	if err := os.MkdirAll("db", 0755); err != nil {
		log.Fatal("Can't create dir", err)
	}
	logger, err := NewLogger("db")
	if err != nil {
		log.Fatal("can't open log file", err)
	}
	tr := transport.Transport{Proxy: transport.ProxyFromEnvironment}
	// For every incoming request, override the RoundTripper to extract
	// connection information. Store it is session context log it after
	// handling the response.
	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		ctx.RoundTripper = goproxy.RoundTripperFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (resp *http.Response, err error) {
			ctx.UserData, resp, err = tr.DetailedRoundTrip(req)
			return
		})
		logger.LogReq(req, ctx)
		return req, nil
	})
	proxy.OnResponse().DoFunc(func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		logger.LogResp(resp, ctx)
		return resp
	})
	l, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal("listen:", err)
	}
	sl := newStoppableListener(l)
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	go func() {
		<-ch
		log.Println("Got SIGINT exiting")
		sl.Add(1)
		sl.Close()
		logger.Close()
		sl.Done()
	}()
	log.Println("Starting Proxy")
	http.Serve(sl, proxy)
	sl.Wait()
	log.Println("All connections closed - exit")
}
