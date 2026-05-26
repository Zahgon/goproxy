package transport

type badStringError struct {
	what string
	str  string
}

func (e *badStringError) Error() string { _ = "STUB: not implemented"; return "" }

func hasPort(s string) bool { _ = "STUB: not implemented"; return false }
