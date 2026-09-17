package validation

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/nobl9/govy/pkg/govy"
)

func TestGovyEvalHiddenValuesAndPlanTypes(t *testing.T) {
	var validator govy.Validator[Account] = AccountValidator
	var callback func(Account) error = ValidateAccount
	checks := []struct {
		name     string
		validate func(Account) error
	}{
		{"validator", func(value Account) error { return validator.Validate(value) }},
		{"callback", callback},
	}
	tests := []struct {
		name  string
		input Account
		want  map[string]int
	}{
		{name: "empty_collection", input: Account{Name: "active"}},
		{
			name:  "token_minimum",
			input: Account{Name: "active", Credentials: []Credential{{Token: strings.Repeat("a", 16)}}},
		},
		{
			name:  "token_minimum_bytes",
			input: Account{Name: "active", Credentials: []Credential{{Token: strings.Repeat("é", 8)}}},
		},
		{
			name:  "long_token",
			input: Account{Name: "active", Credentials: []Credential{{Token: strings.Repeat("a", 64)}}},
		},
		{name: "name_maximum_bytes", input: Account{Name: "éééé"}},
		{name: "name_too_many_bytes", input: Account{Name: "ééééa"}, want: map[string]int{"name": 1}},
		{name: "name_required", input: Account{}, want: map[string]int{"name": 1}},
		{
			name:  "token_required",
			input: Account{Name: "active", Credentials: []Credential{{}}},
			want:  map[string]int{"credentials[0].token": 1},
		},
		{
			name:  "short_token",
			input: Account{Name: "active", Credentials: []Credential{{Token: "private-token"}}},
			want:  map[string]int{"credentials[0].token": 1},
		},
		{
			name:  "token_just_below_minimum",
			input: Account{Name: "active", Credentials: []Credential{{Token: strings.Repeat("a", 15)}}},
			want:  map[string]int{"credentials[0].token": 1},
		},
		{
			name: "independent_failures",
			input: Account{
				Name: "visible-name-too-long",
				Credentials: []Credential{
					{Token: "private-token"},
					{Token: strings.Repeat("a", 16)},
					{Token: "second-secret"},
				},
			},
			want: map[string]int{"name": 1, "credentials[0].token": 1, "credentials[2].token": 1},
		},
	}
	for _, check := range checks {
		t.Run(check.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					err := check.validate(tt.input)
					ruleErrors := govyEvalAssertErrors(t, err, tt.want)
					if err == nil {
						return
					}
					var validatorErr *govy.ValidatorError
					if !errors.As(err, &validatorErr) {
						t.Fatalf("error has type %T; want *govy.ValidatorError", err)
					}
					encoded, encodeErr := json.Marshal(err)
					if encodeErr != nil {
						t.Fatal(encodeErr)
					}
					for _, propertyErr := range validatorErr.Errors {
						if strings.HasPrefix(propertyErr.PropertyPath.String(), "credentials[") &&
							propertyErr.PropertyValue != "" {
							t.Fatalf("hidden property retains its value: %#v", propertyErr.PropertyValue)
						}
						if propertyErr.PropertyPath.String() == "name" && propertyErr.PropertyValue != tt.input.Name {
							t.Fatalf("visible name changed: %#v", propertyErr.PropertyValue)
						}
					}
					for _, credential := range tt.input.Credentials {
						if credential.Token == "" || len(credential.Token) >= 16 {
							continue
						}
						if strings.Contains(err.Error(), credential.Token) ||
							strings.Contains(string(encoded), credential.Token) {
							t.Fatalf("hidden token appears in the error: %v; JSON: %s", err, encoded)
						}
					}
					for path, failures := range ruleErrors {
						if !strings.HasPrefix(path, "credentials[") || tt.name == "token_required" {
							continue
						}
						if failures[0].Code != "token_length" || failures[0].Message != `invalid token "[hidden]"` {
							t.Fatalf("unexpected hidden token error at %s: %#v", path, failures[0])
						}
					}
				})
			}
		})
	}

	plan, err := govy.Plan(validator)
	if err != nil {
		t.Fatal(err)
	}
	if plan.Name != "account" || plan.TypeInfo.Name != "Account" || plan.TypeInfo.Kind != "struct" ||
		plan.TypeInfo.Package == "" {
		t.Fatalf("plan loses the display name or root Go type: %#v", plan)
	}
	encoded, err := json.Marshal(plan)
	if err != nil {
		t.Fatal(err)
	}
	var decoded struct {
		TypeInfo govy.TypeInfo `json:"typeInfo"`
	}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.TypeInfo != plan.TypeInfo {
		t.Fatalf("JSON loses root type information: %s", encoded)
	}
}
