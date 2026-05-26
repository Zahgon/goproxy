package limitation

import (
	"github.com/elazarl/goproxy"
)

// ConcurrentRequests implements a mechanism to limit the number of
// concurrently handled HTTP requests, configurable by the user.
// The ReqHandler can simply be added to the server with OnRequest().
func ConcurrentRequests(limit int) goproxy.ReqHandler {
	_ = "STUB: not implemented"
	// Do nothing when the specified limit is invalid
	return *new(goproxy.ReqHandler)
}

// Release semaphore when request finishes
