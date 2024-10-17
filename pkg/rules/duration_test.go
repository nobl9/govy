package rules

import (
	"testing"
	"time"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

var durationPrecisionTestCases = []*struct {
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

func TestDurationPrecision(t *testing.T) {
	for _, tc := range durationPrecisionTestCases {
		t.Run(tc.name, func(t *testing.T) {
			rule := DurationPrecision(tc.precision)
			err := rule.Validate(tc.duration)
			if tc.expectedError != "" {
				assert.Require(t, assert.Error(t, err))
				assert.EqualError(t, err, tc.expectedError)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeDurationPrecision))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkDurationPrecision(b *testing.B) {
	for range b.N {
		for _, tc := range durationPrecisionTestCases {
			_ = DurationPrecision(tc.precision).Validate(tc.duration)
		}
	}
}
