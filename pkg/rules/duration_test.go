package rules

import (
	"testing"
	"time"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

func TestDurationPrecision(t *testing.T) {
	tests := []struct {
		name          string
		duration      time.Duration
		precision     time.Duration
		expectedError string
	}{
		{
			name:      "valid precision 1ns",
			duration:  time.Duration(123456),
			precision: time.Nanosecond,
		},
		{
			name:      "valid precision 1m",
			duration:  time.Hour + time.Minute,
			precision: time.Minute,
		},
		{
			name:          "invalid precision 1m1s",
			duration:      time.Minute + time.Second,
			precision:     time.Minute,
			expectedError: "duration must be defined with 1m0s precision",
		},
		{
			name:          "invalid precision",
			duration:      101 * time.Nanosecond,
			precision:     10 * time.Nanosecond,
			expectedError: "duration must be defined with 10ns precision",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := DurationPrecision(tt.precision)
			err := rule.Validate(tt.duration)
			if tt.expectedError != "" {
				assert.Require(t, assert.Error(t, err))
				assert.EqualError(t, err, tt.expectedError)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeDurationPrecision))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
