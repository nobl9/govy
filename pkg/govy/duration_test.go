package validation

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDurationPrecision(t *testing.T) {
	tests := []struct {
		name      string
		duration  time.Duration
		precision time.Duration
		expected  error
	}{
		{
			name:      "valid precision 1ns",
			duration:  time.Duration(123456),
			precision: time.Nanosecond,
			expected:  nil,
		},
		{
			name:      "valid precision 1m",
			duration:  time.Hour + time.Minute,
			precision: time.Minute,
			expected:  nil,
		},
		{
			name:      "invalid precision 1m1s",
			duration:  time.Minute + time.Second,
			precision: time.Minute,
			expected:  NewRuleError("duration must be defined with 1m0s precision", ErrorCodeDurationPrecision),
		},
		{
			name:      "invalid precision",
			duration:  time.Duration(123456),
			precision: 10 * time.Nanosecond,
			expected:  NewRuleError("duration must be defined with 10ns precision", ErrorCodeDurationPrecision),
		},
		{
			name:      "minute precision",
			duration:  time.Duration(123456),
			precision: time.Minute,
			expected:  NewRuleError("duration must be defined with 1m0s precision", ErrorCodeDurationPrecision),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := DurationPrecision(tt.precision)
			result := rule.Validate(tt.duration)
			assert.Equal(t, tt.expected, result)
		})
	}
}
