package validation

import (
	"strings"
	"testing"

	"github.com/nobl9/govy/pkg/govy"
)

func TestGovyEvalCustomRuleMetadata(t *testing.T) {
	var validator govy.Validator[Contact] = ContactValidator
	for _, email := range []string{"person@example.com", "person@EXAMPLE.COM", "person@ExAmPlE.CoM", "other+tag@example.com"} {
		t.Run("accept_"+email, func(t *testing.T) {
			govyEvalAssertErrors(t, validator.Validate(Contact{Email: email}), nil)
		})
	}
	for _, email := range []string{
		"person@example.org",
		"person@sub.example.com",
		"person@notexample.com",
		"person@example.com.invalid",
		"person@example.com.",
		"example.com",
	} {
		t.Run("reject_"+email, func(t *testing.T) {
			failures := govyEvalAssertErrors(t, validator.Validate(Contact{Email: email}), map[string]int{"email": 1})
			failure := failures["email"][0]
			if !failure.Code.Has("email_domain") {
				t.Errorf("error code = %q; want email_domain", failure.Code)
			}
			if !strings.Contains(failure.Message, email) {
				t.Errorf("runtime message %q does not contain rejected address %q", failure.Message, email)
			}
		})
	}
	t.Run("required_owns_empty_failure", func(t *testing.T) {
		failures := govyEvalAssertErrors(t, validator.Validate(Contact{}), map[string]int{"email": 1})
		failure := failures["email"][0]
		if !failure.Code.Has("required") || failure.Code.Has("email_domain") {
			t.Errorf("empty email error code = %q; want only the required failure", failure.Code)
		}
	})
	t.Run("custom_rule_plan", func(t *testing.T) {
		plan, err := govy.Plan(validator)
		if err != nil {
			t.Fatalf("generate plan: %v", err)
		}
		emailPlan := govyEvalPlanProperty(t, plan, "email")
		found := false
		for _, rule := range emailPlan.Rules {
			if !rule.ErrorCode.Has("email_domain") {
				continue
			}
			found = true
			if strings.TrimSpace(rule.Description) == "" {
				t.Error("email_domain plan description is empty")
			}
			validExample := false
			for _, example := range rule.Examples {
				local, domain, ok := strings.Cut(example, "@")
				if ok && local != "" && strings.EqualFold(domain, "example.com") &&
					validator.Validate(Contact{Email: example}) == nil {
					validExample = true
				}
			}
			if !validExample {
				t.Errorf("email_domain plan has no valid example: %v", rule.Examples)
			}
		}
		if !found {
			t.Fatal("email plan has no rule with error code email_domain")
		}
	})
}
