package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
)

func TestOrdinalString(t *testing.T) {
	tests := []struct {
		in  int
		out string
	}{
		{0, "0th"},
		{1, "1st"},
		{2, "2nd"},
		{3, "3rd"},
		{4, "4th"},
		{10, "10th"},
		{11, "11th"},
		{12, "12th"},
		{13, "13th"},
		{21, "21st"},
		{32, "32nd"},
		{43, "43rd"},
		{101, "101st"},
		{102, "102nd"},
		{103, "103rd"},
		{211, "211th"},
		{212, "212th"},
		{213, "213th"},
	}
	for _, tc := range tests {
		t.Run(tc.out, func(t *testing.T) {
			assert.Equal(t, tc.out, ordinalString(tc.in))
		})
	}
}
