package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

func TestStringE164(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		in            string
		expectedError string
	}{
		"minimum length": {
			in: "+12",
		},
		"maximum length": {
			in: "+123456789012345",
		},
		"common US number": {
			in: "+14155552671",
		},
		"missing plus sign": {
			in:            "14155552671",
			expectedError: "string must be a valid E.164 phone number",
		},
		"starts with zero": {
			in:            "+0123456789",
			expectedError: "string must be a valid E.164 phone number",
		},
		"too short": {
			in:            "+1",
			expectedError: "string must be a valid E.164 phone number",
		},
		"too long": {
			in:            "+1234567890123456",
			expectedError: "string must be a valid E.164 phone number",
		},
		"contains spaces": {
			in:            "+1 4155552671",
			expectedError: "string must be a valid E.164 phone number",
		},
		"contains punctuation": {
			in:            "+1-415-555-2671",
			expectedError: "string must be a valid E.164 phone number",
		},
		"empty": {
			expectedError: "string must be a valid E.164 phone number",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := StringE164().Validate(tt.in)
			if tt.expectedError != "" {
				assert.Require(t, assert.Error(t, err))
				assert.EqualError(t, err, tt.expectedError)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringE164))
				return
			}
			assert.NoError(t, err)
		})
	}
}

func BenchmarkStringE164(b *testing.B) {
	tests := map[string]string{
		"valid":   "+14155552671",
		"invalid": "+1-415-555-2671",
	}

	for name, in := range tests {
		b.Run(name, func(b *testing.B) {
			rule := StringE164()
			for b.Loop() {
				_ = rule.Validate(in)
			}
		})
	}
}
