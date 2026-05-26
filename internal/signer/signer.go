package signer

import (
	"crypto/tls"
)

const _goproxySignerVersion = ":goproxy2"

func hashSorted(lst []string) []byte { _ = "STUB: not implemented"; return nil }

func SignHost(ca tls.Certificate, hosts []string) (cert *tls.Certificate, err error) {
	_ = "STUB: not implemented"
	// Use the provided CA for certificate generation.
	// Use already parsed Leaf certificate when present.
	return nil, nil
}

// -30 days
// 365 days

// Always generate a positive int value
// (Two complement is not enabled when the first bit is 0)

// Save an already parsed leaf certificate to use less CPU
// when it will be used
