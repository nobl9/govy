package internal

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
)

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		in  any
		out bool
	}{
		{nil, true},
		{any(nil), true},
		{any(""), true},
		{"", true},
		{0, true},
		{0.0, true},
		{false, true},
		{struct{}{}, true},
		{map[int]string{}, false},
		{[]int{}, false},
		{ptr(struct{}{}), false},
		{ptr(""), false},
		{make(chan int), false},
		{any("this"), false},
		{0.123, false},
		{true, false},
		{struct{ This string }{This: "this"}, false},
		{map[int]string{0: ""}, false},
		{ptr(struct{ This string }{This: "this"}), false},
		{[]int{0}, false},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.out, IsEmpty(tc.in))
	}
}