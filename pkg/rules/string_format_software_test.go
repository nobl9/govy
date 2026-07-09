package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

func TestStringSemver(t *testing.T) {
	tests := map[string]struct {
		in            string
		expectedError string
	}{
		"zero major": {
			in: "0.1.0",
		},
		"release": {
			in: "1.0.0",
		},
		"prerelease": {
			in: "1.0.0-alpha",
		},
		"prerelease numeric identifier": {
			in: "1.0.0-alpha.1",
		},
		"build metadata": {
			in: "1.0.0+20130313144700",
		},
		"prerelease with build metadata": {
			in: "1.0.0-beta+exp.sha.5114f85",
		},
		"complex prerelease and build": {
			in: "2.7.3-rc.1+build.11.e0f985a",
		},
		"empty": {
			in:            "",
			expectedError: "string must be a valid semantic version",
		},
		"major only": {
			in:            "1",
			expectedError: "string must be a valid semantic version",
		},
		"major minor only": {
			in:            "1.2",
			expectedError: "string must be a valid semantic version",
		},
		"too many numeric components": {
			in:            "1.2.3.4",
			expectedError: "string must be a valid semantic version",
		},
		"leading zero major": {
			in:            "01.2.3",
			expectedError: "string must be a valid semantic version",
		},
		"leading zero minor": {
			in:            "1.02.3",
			expectedError: "string must be a valid semantic version",
		},
		"leading zero patch": {
			in:            "1.2.03",
			expectedError: "string must be a valid semantic version",
		},
		"empty prerelease": {
			in:            "1.2.3-",
			expectedError: "string must be a valid semantic version",
		},
		"leading zero prerelease number": {
			in:            "1.2.3-01",
			expectedError: "string must be a valid semantic version",
		},
		"v prefix": {
			in:            "v1.2.3",
			expectedError: "string must be a valid semantic version",
		},
		"empty build identifier": {
			in:            "1.2.3+build..1",
			expectedError: "string must be a valid semantic version",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := StringSemver().Validate(tt.in)
			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringSemver))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkStringSemver(b *testing.B) {
	rule := StringSemver()
	for b.Loop() {
		_ = rule.Validate("2.7.3-rc.1+build.11.e0f985a")
	}
}

func TestStringCVE(t *testing.T) {
	tests := map[string]struct {
		in            string
		expectedError string
	}{
		"four digit sequence": {
			in: "CVE-1999-0001",
		},
		"four digit year before 1999": {
			in: "CVE-1998-0001",
		},
		"zero sequence": {
			in: "CVE-2021-0000",
		},
		"sequence with leading zero": {
			in: "CVE-2014-0160",
		},
		"five digit sequence with leading zero": {
			in: "CVE-2021-00001",
		},
		"sequence with two leading zeroes": {
			in: "CVE-2021-0990",
		},
		"five digit sequence": {
			in: "CVE-2021-44228",
		},
		"long sequence": {
			in: "CVE-2024-12345",
		},
		"nineteen digit sequence": {
			in: "CVE-2024-1234567890123456789",
		},
		"empty": {
			in:            "",
			expectedError: "string must be a valid CVE ID",
		},
		"lowercase prefix": {
			in:            "cve-2021-44228",
			expectedError: "string must be a valid CVE ID",
		},
		"short sequence": {
			in:            "CVE-2021-123",
			expectedError: "string must be a valid CVE ID",
		},
		"letters in sequence": {
			in:            "CVE-2021-ABCD",
			expectedError: "string must be a valid CVE ID",
		},
		"five digit year": {
			in:            "CVE-10000-0001",
			expectedError: "string must be a valid CVE ID",
		},
		"twenty digit sequence": {
			in:            "CVE-2024-12345678901234567890",
			expectedError: "string must be a valid CVE ID",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := StringCVE().Validate(tt.in)
			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCVE))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkStringCVE(b *testing.B) {
	rule := StringCVE()
	for b.Loop() {
		_ = rule.Validate("CVE-2021-44228")
	}
}
