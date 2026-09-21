package validation

import (
	"errors"
	"slices"
	"strings"
	"testing"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonpath"
)

func TestGovyEvalPlainLanguageComposition(t *testing.T) {
	var validate func(Registration) error = ValidateRegistration
	valid := Registration{Username: "Ab3", DisplayName: "Xy9"}
	tests := []struct {
		name   string
		change func(*Registration)
		want   []string
	}{
		{name: "nil_optional_address"},
		{
			name: "maximum_names",
			change: func(v *Registration) {
				v.Username = strings.Repeat("a", 20)
				v.DisplayName = strings.Repeat("9", 20)
			},
		},
		{name: "username_required", change: func(v *Registration) { v.Username = "" }, want: []string{"username"}},
		{
			name:   "display_name_required",
			change: func(v *Registration) { v.DisplayName = "" },
			want:   []string{"displayName"},
		},
		{name: "username_too_short", change: func(v *Registration) { v.Username = "ab" }, want: []string{"username"}},
		{
			name:   "display_name_too_short",
			change: func(v *Registration) { v.DisplayName = "ab" },
			want:   []string{"displayName"},
		},
		{
			name:   "username_too_long",
			change: func(v *Registration) { v.Username = strings.Repeat("a", 21) },
			want:   []string{"username"},
		},
		{
			name:   "display_name_too_long",
			change: func(v *Registration) { v.DisplayName = strings.Repeat("a", 21) },
			want:   []string{"displayName"},
		},
		{name: "username_non_ascii", change: func(v *Registration) { v.Username = "abé" }, want: []string{"username"}},
		{
			name:   "display_name_non_ascii",
			change: func(v *Registration) { v.DisplayName = "abé" },
			want:   []string{"displayName"},
		},
		{
			name:   "username_unicode_digit",
			change: func(v *Registration) { v.Username = "ab٣" },
			want:   []string{"username"},
		},
		{
			name:   "display_name_unicode_digit",
			change: func(v *Registration) { v.DisplayName = "ab٣" },
			want:   []string{"displayName"},
		},
		{name: "username_space", change: func(v *Registration) { v.Username = "a b" }, want: []string{"username"}},
		{
			name:   "display_name_punctuation",
			change: func(v *Registration) { v.DisplayName = "ab_" },
			want:   []string{"displayName"},
		},
		{
			name:   "username_trailing_newline",
			change: func(v *Registration) { v.Username = "abc\n" },
			want:   []string{"username"},
		},
		{name: "address_present", change: func(v *Registration) { v.Address = &Address{City: "Warsaw"} }},
		{
			name:   "address_city_required",
			change: func(v *Registration) { v.Address = &Address{} },
			want:   []string{"address.city"},
		},
		{
			name: "independent_failures",
			change: func(v *Registration) {
				v.Username = "ab"
				v.DisplayName = "ab_"
				v.Address = &Address{}
			},
			want: []string{"address.city", "displayName", "username"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := valid
			if tt.change != nil {
				tt.change(&input)
			}
			govyEvalAssertPropertyPaths(t, validate(input), tt.want)
		})
	}
}

func govyEvalAssertPropertyPaths(t *testing.T, err error, want []string) {
	t.Helper()
	if len(want) == 0 {
		govyEvalAssertErrors(t, err, nil)
		return
	}
	var validationErr *govy.ValidatorError
	if !errors.As(err, &validationErr) || validationErr == nil {
		t.Fatalf("error has type %T; want *govy.ValidatorError: %v", err, err)
	}
	got := make([]string, 0, len(validationErr.Errors))
	for _, propertyErr := range validationErr.Errors {
		if propertyErr == nil || len(propertyErr.Errors) == 0 {
			t.Fatal("property error has no rule failures")
		}
		path := propertyErr.PropertyPath.String()
		if !propertyErr.PropertyPath.Equal(jsonpath.Parse(path)) {
			t.Fatalf("path %q does not have the expected segments", path)
		}
		got = append(got, path)
	}
	slices.Sort(got)
	if !slices.Equal(got, want) {
		t.Fatalf("got property paths %v; want %v", got, want)
	}
}
