package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

func TestStringEIN(t *testing.T) {
	tests := map[string]struct {
		in            string
		expectedError string
	}{
		"valid prefix": {
			in: "12-3456789",
		},
		"valid internet prefix": {
			in: "99-3456789",
		},
		"zero prefix": {
			in:            "00-0000000",
			expectedError: "string must be a valid EIN",
		},
		"unassigned prefix": {
			in:            "07-3456789",
			expectedError: "string must be a valid EIN",
		},
		"missing dash": {
			in:            "123456789",
			expectedError: "string must be a valid EIN",
		},
		"short prefix": {
			in:            "1-23456789",
			expectedError: "string must be a valid EIN",
		},
		"short serial": {
			in:            "12-345678",
			expectedError: "string must be a valid EIN",
		},
		"long serial": {
			in:            "12-34567890",
			expectedError: "string must be a valid EIN",
		},
		"letters": {
			in:            "AB-3456789",
			expectedError: "string must be a valid EIN",
		},
		"underscore separator": {
			in:            "12_3456789",
			expectedError: "string must be a valid EIN",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := StringEIN().Validate(tt.in)
			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringEIN))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkStringEIN(b *testing.B) {
	rule := StringEIN()
	for b.Loop() {
		_ = rule.Validate("12-3456789")
	}
}

func TestStringSSN(t *testing.T) {
	tests := map[string]struct {
		in            string
		expectedError string
	}{
		"valid": {
			in: "123-45-6789",
		},
		"valid high area": {
			in: "899-99-9999",
		},
		"valid low groups": {
			in: "001-01-0001",
		},
		"zero area": {
			in:            "000-45-6789",
			expectedError: "string must be a valid SSN",
		},
		"666 area": {
			in:            "666-45-6789",
			expectedError: "string must be a valid SSN",
		},
		"900 area": {
			in:            "900-45-6789",
			expectedError: "string must be a valid SSN",
		},
		"999 area": {
			in:            "999-45-6789",
			expectedError: "string must be a valid SSN",
		},
		"zero group": {
			in:            "123-00-6789",
			expectedError: "string must be a valid SSN",
		},
		"zero serial": {
			in:            "123-45-0000",
			expectedError: "string must be a valid SSN",
		},
		"missing dashes": {
			in:            "123456789",
			expectedError: "string must be a valid SSN",
		},
		"letters": {
			in:            "12A-45-6789",
			expectedError: "string must be a valid SSN",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := StringSSN().Validate(tt.in)
			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringSSN))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkStringSSN(b *testing.B) {
	rule := StringSSN()
	for b.Loop() {
		_ = rule.Validate("123-45-6789")
	}
}
