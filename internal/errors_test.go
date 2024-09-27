package internal

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/nobl9/govy/internal/assert"
)

func TestJoinErrors(t *testing.T) {
	tests := []struct {
		in  []error
		out string
	}{
		{nil, ""},
		{[]error{nil, nil}, ""},
		// Incorrect formatting, this test case ensures the function does not panic.
		{[]error{nil, errors.New("some error"), nil}, " - some error\n"},
		{[]error{errors.New("- some error")}, " - some error"},
		{[]error{errors.New("- some error"), errors.New("some other error")}, " - some error\n - some other error"},
	}
	for _, tc := range tests {
		b := strings.Builder{}
		JoinErrors(&b, tc.in, " ")
		assert.Equal(t, tc.out, b.String())
	}
	t.Run("custom indent", func(t *testing.T) {
		b := strings.Builder{}
		JoinErrors(&b, []error{errors.New("some error")}, "   ")
		assert.Equal(t, "   - some error", b.String())
	})
}

func TestPropertyValueString(t *testing.T) {
	tests := []struct {
		in  any
		out string
	}{
		{nil, ""},
		{any(nil), ""},
		{false, "false"},
		{true, "true"},
		{any("this"), "this"},
		{func() {}, "func"},
		{ptr("this"), "this"},
		{struct{ This string }{This: "this"}, `{"This":"this"}`},
		{ptr(struct{ This string }{This: "this"}), `{"This":"this"}`},
		{struct {
			This string `json:"this"`
		}{This: "this"}, `{"this":"this"}`},
		{map[string]string{"this": "this"}, `{"this":"this"}`},
		{[]string{"this", "that"}, `["this","that"]`},
		{0, "0"},
		{0.0, "0"},
		{2, "2"},
		{0.123, "0.123"},
		{time.Second, "1s"},
		{time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), "2024-01-01T00:00:00Z"},
		{mockEmptyStringer{}, "mock"},
		{mockStringerWithTags{}, ""},
		{mockStringerWithTags{Mock: "mock"}, `{"mock":"mock"}`},
	}
	for _, tc := range tests {
		got := PropertyValueString(tc.in)
		assert.Equal(t, tc.out, got)
	}
}

type mockEmptyStringer struct{}

func (m mockEmptyStringer) String() string {
	return "mock"
}

type mockStringerWithTags struct {
	Mock string `json:"mock"`
}

func (m mockStringerWithTags) String() string {
	return "stringer"
}

func ptr[T any](v T) *T { return &v }
