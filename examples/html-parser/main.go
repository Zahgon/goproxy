package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/elazarl/goproxy"
)

var (
	// who said we can't parse HTML with regexp?
	scriptMatcher  = regexp.MustCompile(`(?i:<script\s+)`)
	srcAttrMatcher = regexp.MustCompile(`^(?i:[^>]*\ssrc=["']([^"']*)["'])`)
)

// findScripts returns all sources of HTML script tags found in input text.
func findScriptSrc(html string) []string { _ = "STUB: not implemented"; return nil }

// -1 to capture the whitespace at the end of the script tag

// NewJQueryVersionProxy creates a proxy checking responses HTML content, looks
// for scripts referencing jQuery library and emits warnings if different
// versions of the library are being used for a given host.
func NewJQueryVersionProxy() *goproxy.ProxyHttpServer { _ = "STUB: not implemented"; return nil }

func main() {
	proxy := NewJQueryVersionProxy()
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
