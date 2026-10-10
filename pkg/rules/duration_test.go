package rules

import (
	"fmt"
	"testing"
	"time"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonschematest"
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
	{
		name:      "valid precision 10ns",
		duration:  100 * time.Nanosecond,
		precision: 10 * time.Nanosecond,
	},
	{
		name:      "zero duration",
		precision: time.Minute,
	},
	{
		name:      "negative duration",
		duration:  -time.Minute,
		precision: time.Minute,
	},
	{
		name:          "negative duration with invalid precision",
		duration:      -time.Minute - time.Second,
		precision:     time.Minute,
		expectedError: "duration must be defined with 1m0s precision",
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

func TestDurationPrecision_JSONSchema(t *testing.T) {
	t.Parallel()

	groups := make(map[time.Duration][]jsonschematest.Case[time.Duration])
	for _, tc := range durationPrecisionTestCases {
		groups[tc.precision] = append(groups[tc.precision], jsonschematest.Case[time.Duration]{
			Name:  tc.name,
			Input: tc.duration,
			Valid: tc.expectedError == "",
		})
	}
	for precision, cases := range groups {
		t.Run(precision.String(), func(t *testing.T) {
			t.Parallel()
			schema, err := govy.JSONSchema(govy.New(
				govy.For(govy.GetSelf[time.Duration]()).Rules(DurationPrecision(precision)),
			))
			assert.Require(t, assert.NoError(t, err))
			fixture := fmt.Sprintf("testdata/jsonschema/expected_duration_precision_%d.json", precision)
			jsonschematest.Assert(t, schema, fixture, cases)
		})
	}
}

func BenchmarkDurationPrecision(b *testing.B) {
	for _, tc := range durationPrecisionTestCases {
		rule := DurationPrecision(tc.precision)
		for range b.N {
			_ = rule.Validate(tc.duration)
		}
	}
}
