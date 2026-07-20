package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

var validEINTestCases = map[string]string{
	"valid prefix":          "12-3456789",
	"valid internet prefix": "99-3456789",
}

var invalidEINTestCases = map[string]string{
	"zero prefix":          "00-0000000",
	"unassigned prefix":    "07-3456789",
	"missing dash":         "123456789",
	"short prefix":         "1-23456789",
	"short serial":         "12-345678",
	"long serial":          "12-34567890",
	"letters":              "AB-3456789",
	"underscore separator": "12_3456789",
}

func TestStringEIN(t *testing.T) {
	rule := StringEIN()
	t.Run("valid EINs", func(t *testing.T) {
		for name, in := range validEINTestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(in))
			})
		}
	})
	t.Run("invalid EINs", func(t *testing.T) {
		err := rule.Validate(invalidEINTestCases["zero prefix"])
		assert.EqualError(t, err, "string must be a valid Employer Identification Number (EIN)")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringEIN))

		for name, in := range invalidEINTestCases {
			t.Run(name, func(t *testing.T) {
				assert.Error(t, rule.Validate(in))
			})
		}
	})
}

func BenchmarkStringEIN(b *testing.B) {
	rule := StringEIN()
	for b.Loop() {
		_ = rule.Validate("12-3456789")
	}
}

var validSSNTestCases = map[string]string{
	"valid":            "123-45-6789",
	"valid high area":  "899-99-9999",
	"valid low groups": "001-01-0001",
}

var invalidSSNTestCases = map[string]string{
	"zero area":      "000-45-6789",
	"666 area":       "666-45-6789",
	"900 area":       "900-45-6789",
	"999 area":       "999-45-6789",
	"zero group":     "123-00-6789",
	"zero serial":    "123-45-0000",
	"missing dashes": "123456789",
	"letters":        "12A-45-6789",
}

func TestStringSSN(t *testing.T) {
	rule := StringSSN()
	t.Run("valid SSNs", func(t *testing.T) {
		for name, in := range validSSNTestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(in))
			})
		}
	})
	t.Run("invalid SSNs", func(t *testing.T) {
		err := rule.Validate(invalidSSNTestCases["zero area"])
		assert.EqualError(t, err, "string must be a valid Social Security Number (SSN)")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringSSN))

		for name, in := range invalidSSNTestCases {
			t.Run(name, func(t *testing.T) {
				assert.Error(t, rule.Validate(in))
			})
		}
	})
}

func BenchmarkStringSSN(b *testing.B) {
	rule := StringSSN()
	for b.Loop() {
		_ = rule.Validate("123-45-6789")
	}
}
