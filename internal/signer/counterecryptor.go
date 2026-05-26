package signer

import (
	"crypto/cipher"
)

type CounterEncryptorRand struct {
	cipher  cipher.Block
	counter []byte
	rand    []byte
	ix      int
}

func NewCounterEncryptorRandFromKey(key any, seed []byte) (r CounterEncryptorRand, err error) {
	_ = "STUB: not implemented"
	return *new(CounterEncryptorRand), nil
}

func (c *CounterEncryptorRand) Seed(b []byte) { _ = "STUB: not implemented"; return }

func (c *CounterEncryptorRand) refill() { _ = "STUB: not implemented"; return }

func (c *CounterEncryptorRand) Read(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
