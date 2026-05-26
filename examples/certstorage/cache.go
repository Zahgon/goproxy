package main

import (
	"crypto/tls"
	"sync"
)

// CertStorage is a simple certificate cache that keeps
// everything in memory.
type CertStorage struct {
	certs map[string]*tls.Certificate
	mtx   sync.RWMutex
}

func (cs *CertStorage) Fetch(hostname string, gen func() (*tls.Certificate, error)) (*tls.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewCertStorage() *CertStorage { _ = "STUB: not implemented"; return nil }
