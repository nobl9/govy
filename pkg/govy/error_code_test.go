package govy

import (
	"fmt"
	"testing"

	"github.com/nobl9/govy/internal/assert"
)

func TestErrorCode_Add(t *testing.T) {
	tests := []struct {
		in       ErrorCode
		current  ErrorCode
		expected ErrorCode
	}{
		{in: "", current: "", expected: ""},
		{in: "code", current: "", expected: "code"},
		{in: "", current: "code", expected: "code"},
		{in: "foo", current: "bar", expected: "foo:bar"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("'%s' + '%s'", tc.current, tc.in), func(t *testing.T) {
			actual := tc.current.Add(tc.in)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestErrorCode_Has(t *testing.T) {
	tests := []struct {
		in       ErrorCode
		current  ErrorCode
		expected bool
	}{
		{in: "", current: "", expected: false},
		{in: "code", current: "", expected: false},
		{in: "", current: "code", expected: false},
		{in: "foo", current: "baz", expected: false},
		{in: "foo", current: "baz:bar", expected: false},
		{in: "bar", current: "bar", expected: true},
		{in: "bar", current: ":", expected: false},
		{in: "bar", current: "::", expected: false},
		{in: "bar", current: ":bar:", expected: true},
		{in: "bar", current: ":bar", expected: true},
		{in: "bar", current: "bar:", expected: true},
		{in: "bar", current: ":baz", expected: false},
		{in: "bar", current: "baz:", expected: false},
		{in: "foo", current: "foo:bar", expected: true},
		{in: "foo", current: "foo!:bar", expected: false},
		{in: "bar", current: "foo:bar", expected: true},
		{in: "bar", current: "foo:barr", expected: false},
		{in: "baz", current: "foo:baz:bar", expected: true},
		{in: "baz", current: "foo:baz!:bar", expected: false},
		{in: "baz", current: "foo:!baz:bar", expected: false},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("'%s' has '%s'", tc.current, tc.in), func(t *testing.T) {
			actual := tc.current.Has(tc.in)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
