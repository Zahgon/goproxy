package regretable

import (
	"io"
)

// Reader in regretable package will allow you to read from a reader,
// and then to "regret" reading it, and push back everything you've read.
// For example:
//
//	rb := NewRegretableReader(bytes.NewBuffer([]byte{1,2,3}))
//	var b = make([]byte,1)
//	rb.Read(b) // b[0] = 1
//	rb.Regret()
//	ioutil.ReadAll(rb.Read) // returns []byte{1,2,3},nil
type Reader struct {
	reader   io.Reader
	overflow bool
	r, w     int
	buf      []byte
}

const _defaultBufferSize = 500

// The next read from the RegretableReader will be as if the underlying reader
// was never read (or from the last point forget is called).
func (rb *Reader) Regret() { _ = "STUB: not implemented"; return }

// Will "forget" everything read so far.
//
//	rb := NewRegretableReader(bytes.NewBuffer([]byte{1,2,3}))
//	var b = make([]byte,1)
//	rb.Read(b) // b[0] = 1
//	rb.Forget()
//	rb.Read(b) // b[0] = 2
//	rb.Regret()
//	ioutil.ReadAll(rb.Read) // returns []byte{2,3},nil
func (rb *Reader) Forget() { _ = "STUB: not implemented"; return }

// initialize a RegretableReader with underlying reader r, whose buffer is size bytes long.
func NewRegretableReaderSize(r io.Reader, size int) *Reader { _ = "STUB: not implemented"; return nil }

// initialize a RegretableReader with underlying reader r.
func NewRegretableReader(r io.Reader) *Reader { _ = "STUB: not implemented"; return nil }

// reads from the underlying reader. Will buffer all input until Regret is called.
func (rb *Reader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// ReaderCloser is the same as Reader, but allows closing the underlying reader.
type ReaderCloser struct {
	Reader
	c io.Closer
}

// initialize a RegretableReaderCloser with underlying readCloser rc.
func NewRegretableReaderCloser(rc io.ReadCloser) *ReaderCloser {
	_ = "STUB: not implemented"
	return nil
}

// initialize a RegretableReaderCloser with underlying readCloser rc.
func NewRegretableReaderCloserSize(rc io.ReadCloser, size int) *ReaderCloser {
	_ = "STUB: not implemented"
	return nil
}

// Closes the underlying readCloser, you cannot regret after closing the stream.
func (rbc *ReaderCloser) Close() error { _ = "STUB: not implemented"; return nil }
