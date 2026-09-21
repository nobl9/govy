package validation

import (
	"strings"
	"testing"

	"github.com/nobl9/govy/pkg/govy"
)

func TestGovyEvalConditionsTransformsAndPlans(t *testing.T) {
	var validator govy.Validator[Endpoint] = EndpointValidator
	for _, mode := range []string{"local", "remote"} {
		for _, timeout := range []string{"", "1s", "1ns", "0.5s", "1h30m"} {
			t.Run(mode+"_timeout_"+timeout, func(t *testing.T) {
				input := Endpoint{Name: "valid", Mode: mode, Address: "https://example.com", Timeout: timeout}
				govyEvalAssertErrors(t, validator.Validate(input), nil)
			})
		}
	}
	for _, address := range []string{"", "not a URL", "http://example.com", "https://"} {
		t.Run("local_ignores_"+address, func(t *testing.T) {
			govyEvalAssertErrors(t, validator.Validate(Endpoint{Name: "valid", Mode: "local", Address: address}), nil)
		})
	}
	for _, address := range []string{"https://example.com", "https://example.com/path?q=1", "https://example.com:8443", "https://localhost"} {
		t.Run("remote_accepts_"+address, func(t *testing.T) {
			govyEvalAssertErrors(t, validator.Validate(Endpoint{Name: "valid", Mode: "remote", Address: address}), nil)
		})
	}
	for _, address := range []string{
		"",
		"http://example.com",
		"ftp://example.com",
		"example.com",
		"https://",
		"https:///path",
		"https://example.com/%zz",
		"https://bad host",
		"%",
	} {
		t.Run("remote_rejects_"+address, func(t *testing.T) {
			input := Endpoint{Name: "valid", Mode: "remote", Address: address}
			govyEvalAssertErrors(t, validator.Validate(input), map[string]int{"address": 1})
			input.Name = ""
			govyEvalAssertErrors(t, validator.Validate(input), map[string]int{"name": 1, "address": 1})
		})
	}
	for _, timeout := range []string{"invalid", "1", " ", "0", "0s", "0ms", "-1ns", "-5s"} {
		for _, mode := range []string{"local", "remote"} {
			t.Run(mode+"_rejects_timeout_"+timeout, func(t *testing.T) {
				input := Endpoint{Name: "valid", Mode: mode, Address: "https://example.com", Timeout: timeout}
				govyEvalAssertErrors(t, validator.Validate(input), map[string]int{"timeout": 1})
			})
		}
	}
	for _, mode := range []string{"", "other", "REMOTE"} {
		t.Run("rejects_mode_"+mode, func(t *testing.T) {
			input := Endpoint{Name: "valid", Mode: mode, Address: "not a URL"}
			govyEvalAssertErrors(t, validator.Validate(input), map[string]int{"mode": 1})
		})
	}
	t.Run("name_required_for_local", func(t *testing.T) {
		govyEvalAssertErrors(t, validator.Validate(Endpoint{Mode: "local"}), map[string]int{"name": 1})
	})
	t.Run("independent_failures", func(t *testing.T) {
		input := Endpoint{Mode: "remote", Address: "https://", Timeout: "0ms"}
		govyEvalAssertErrors(t, validator.Validate(input), map[string]int{"name": 1, "address": 1, "timeout": 1})
	})
	t.Run("strict_plan_metadata", func(t *testing.T) {
		plan, err := govy.Plan(validator, govy.PlanStrictMode())
		if err != nil {
			t.Fatalf("generate strict plan: %v", err)
		}
		for _, path := range []string{"name", "mode", "address", "timeout"} {
			govyEvalPlanProperty(t, plan, path)
		}
		for _, property := range plan.Properties {
			if property == nil {
				t.Fatal("nil property plan")
			}
			for _, rule := range property.Rules {
				if strings.TrimSpace(rule.Description) == "" {
					t.Errorf("rule at %q has an empty description", property.Path.String())
				}
				for _, condition := range rule.Conditions {
					if strings.TrimSpace(condition) == "" {
						t.Errorf("rule at %q has an empty condition description", property.Path.String())
					}
				}
			}
		}
		address := govyEvalPlanProperty(t, plan, "address")
		hasAddressCondition := false
		for _, rule := range address.Rules {
			if len(rule.Conditions) > 0 {
				hasAddressCondition = true
			}
		}
		if !hasAddressCondition {
			t.Error("address plan has no condition descriptions")
		}
		for _, path := range []string{"address", "timeout"} {
			property := govyEvalPlanProperty(t, plan, path)
			hasContentRule := false
			for _, rule := range property.Rules {
				if rule.ErrorCode != "required" && rule.ErrorCode != "optional" {
					hasContentRule = true
				}
			}
			if !hasContentRule {
				t.Errorf("plan for %q describes only presence, not its content constraints", path)
			}
		}
	})
}
