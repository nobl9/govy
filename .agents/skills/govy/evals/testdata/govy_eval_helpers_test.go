package validation

import (
	"errors"
	"testing"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonpath"
)

func govyEvalAssertErrors(t *testing.T, err error, want map[string]int) map[string][]*govy.RuleError {
	t.Helper()
	if len(want) == 0 {
		if err != nil {
			t.Fatalf("unexpected validation error: %v", err)
		}
		return nil
	}
	if err == nil {
		t.Fatalf("validation succeeded; want rule counts %v", want)
	}
	var validationErr *govy.ValidatorError
	if !errors.As(err, &validationErr) || validationErr == nil {
		t.Fatalf("error has type %T; want *govy.ValidatorError: %v", err, err)
	}
	if len(validationErr.Errors) != len(want) {
		t.Fatalf("got %d property errors; want %d for %v: %v", len(validationErr.Errors), len(want), want, err)
	}
	got := make(map[string][]*govy.RuleError, len(want))
	for _, propertyErr := range validationErr.Errors {
		if propertyErr == nil {
			t.Fatal("nil property error")
		}
		path := propertyErr.PropertyPath.String()
		count, exists := want[path]
		if !exists {
			t.Fatalf("unexpected property path %q; want %v", path, want)
		}
		if !propertyErr.PropertyPath.Equal(jsonpath.Parse(path)) {
			t.Fatalf("path %q does not have the expected segments", path)
		}
		if _, duplicate := got[path]; duplicate {
			t.Fatalf("duplicate property error at %q", path)
		}
		if len(propertyErr.Errors) != count {
			t.Fatalf("path %q has %d rule errors; want %d: %v", path, len(propertyErr.Errors), count, err)
		}
		for _, ruleErr := range propertyErr.Errors {
			if ruleErr == nil {
				t.Fatalf("nil rule error at %q", path)
			}
		}
		got[path] = propertyErr.Errors
	}
	for path := range want {
		if _, exists := got[path]; !exists {
			t.Fatalf("missing property error at %q", path)
		}
	}
	return got
}

func govyEvalPlanProperty(t *testing.T, plan *govy.ValidatorPlan, path string) *govy.PropertyPlan {
	t.Helper()
	if plan == nil {
		t.Fatal("validation plan is nil")
	}
	want := jsonpath.NewRoot().Join(jsonpath.Parse(path))
	for _, property := range plan.Properties {
		if property != nil && property.Path.Equal(want) {
			if len(property.Rules) == 0 {
				t.Fatalf("plan has no rules for %q", path)
			}
			return property
		}
	}
	t.Fatalf("plan has no property with path %q", want.String())
	return nil
}

func govyEvalPtr[T any](value T) *T {
	return &value
}
