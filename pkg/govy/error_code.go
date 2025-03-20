package govy

import "strings"

// ErrorCode is a unique string that represents a specific [RuleError].
// It can be used to precisely identify the error without inspecting its message.
type ErrorCode string

const (
	ErrorCodeTransform ErrorCode = "transform"
)

// Add extends the error code with a new error code.
// Codes are prepended, the last code in chain is always the first one set.
// Example:
//
//	ErrorCode("first").Add("another").Add("last") --> ErrorCode("last:another:first")
func (e ErrorCode) Add(code ErrorCode) ErrorCode {
	switch {
	case e == "":
		return code
	case code == "":
		return e
	default:
		return code + ErrorCode(ErrorCodeSeparator) + e
	}
}

// Has reports whether given error code is in the examined error code's chain.
// Example:
//
//	ErrorCode("foo:bar").Has("foo") --> true
//	ErrorCode("foo:bar").Has("bar") --> true
//	ErrorCode("foo:bar").Has("baz") --> false
func (e ErrorCode) Has(code ErrorCode) bool {
	if e == "" || code == "" {
		return false
	}
	if e == code {
		return true
	}
	i := 0
	for {
		if i >= len(e) {
			return false
		}
		if e[i:] == code {
			return true
		}
		next := strings.Index(string(e[i:]), ErrorCodeSeparator)
		switch {
		case next == -1:
			return false
		case e[i:i+next] == code:
			return true
		}
		i += next + 1
	}
}
