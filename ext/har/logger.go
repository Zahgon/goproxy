package har

import (
	"net/http"
	"time"

	"github.com/elazarl/goproxy"
)

// ExportFunc is a function type that users can implement to handle exported entries
type ExportFunc func([]Entry)

// Logger implements a HAR logging extension for goproxy
type Logger struct {
	exportFunc      ExportFunc
	exportInterval  time.Duration
	exportThreshold int
	dataCh          chan Entry
}

// LoggerOption is a function type for configuring the Logger
type LoggerOption func(*Logger)

// WithExportInterval sets the interval for automatic exports
func WithExportInterval(d time.Duration) LoggerOption {
	_ = "STUB: not implemented"
	return *new(LoggerOption)
}

// WithExportCount sets the number of requests after which to export entries
func WithExportThreshold(threshold int) LoggerOption {
	_ = "STUB: not implemented"
	return *new(LoggerOption)
}

// NewLogger creates a new HAR logger instance
func NewLogger(exportFunc ExportFunc, opts ...LoggerOption) *Logger {
	_ = "STUB: not implemented"
	return nil
}

// Default threshold
// Default no interval

// Apply options

// OnRequest handles incoming HTTP requests
func (l *Logger) OnRequest(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OnResponse handles HTTP responses
func (l *Logger) OnResponse(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	_ = "STUB: not implemented"
	return nil
}

func (l *Logger) exportLoop() { _ = "STUB: not implemented"; return }

func (l *Logger) Stop() { _ = "STUB: not implemented"; return }
