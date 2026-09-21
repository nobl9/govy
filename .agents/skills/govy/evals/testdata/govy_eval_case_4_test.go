package validation

import (
	"testing"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonpath"
)

func TestGovyEvalPathsAndOptionalData(t *testing.T) {
	var validator govy.Validator[Profile] = ProfileValidator
	tests := []struct {
		name  string
		input Profile
		want  map[string]int
	}{
		{name: "nil_details", input: Profile{Code: "valid"}},
		{name: "valid_details", input: Profile{Code: "valid", Details: &Details{Name: "valid"}}},
		{name: "synthetic_path", input: Profile{}, want: map[string]int{"settings.code": 1}},
		{name: "nested_name", input: Profile{Code: "valid", Details: &Details{}}, want: map[string]int{"details.name": 1}},
		{name: "independent_failures", input: Profile{Details: &Details{}}, want: map[string]int{"settings.code": 1, "details.name": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			govyEvalAssertErrors(t, validator.Validate(tt.input), tt.want)
		})
	}
	t.Run("two_synthetic_segments", func(t *testing.T) {
		err := validator.Validate(Profile{})
		govyEvalAssertErrors(t, err, map[string]int{"settings.code": 1})
		validationErr := err.(*govy.ValidatorError)
		path := validationErr.Errors[0].PropertyPath
		want := jsonpath.New().Name("settings").Name("code")
		if !path.Equal(want) || path.Equal(jsonpath.New().Name("settings.code")) {
			t.Fatalf("synthetic path must contain separate settings and code segments: %q", path.String())
		}
	})
}
