package govy_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonpath"
	"github.com/nobl9/govy/pkg/rules"
)

func TestRule_HideValue(t *testing.T) {
	const secret = "secret"

	t.Run("plain error", func(t *testing.T) {
		rule := govy.NewRule(func(v string) error {
			return fmt.Errorf("invalid value %q", v)
		})
		hiddenRule := rule.HideValue()

		assert.EqualError(t, rule.Validate(secret), `invalid value "secret"`)
		assert.EqualError(t, hiddenRule.Validate(secret), `invalid value "[hidden]"`)
		assert.EqualError(t, rule.Validate(secret), `invalid value "secret"`)
	})

	t.Run("long repeated value", func(t *testing.T) {
		longSecret := strings.Repeat("sensitive-value-", 8)
		rule := govy.NewRule(func(v string) error {
			return fmt.Errorf("invalid values %q and %q", v, v)
		}).HideValue()

		assert.EqualError(t, rule.Validate(longSecret), `invalid values "[hidden]" and "[hidden]"`)
	})

	t.Run("truncated long value", func(t *testing.T) {
		longSecret := strings.Repeat("sensitive-value-", 8)
		rule := govy.NewRule(func(v string) error {
			return fmt.Errorf("invalid value %q", v[:100]+"...")
		}).HideValue()

		assert.EqualError(t, rule.Validate(longSecret), `invalid value "[hidden]"`)
	})

	t.Run("non-string value", func(t *testing.T) {
		rule := govy.NewRule(func(v int) error {
			return fmt.Errorf("invalid value %d", v)
		}).HideValue()

		assert.EqualError(t, rule.Validate(42), "invalid value [hidden]")
	})

	t.Run("message template", func(t *testing.T) {
		rule := govy.NewRule(func(v string) error {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				Error: fmt.Sprintf("nested error for %q", v),
			})
		}).
			WithMessageTemplateString("{{ .PropertyValue }}: {{ .Error }}").
			HideValue()

		assert.EqualError(t, rule.Validate(secret), `[hidden]: nested error for "[hidden]"`)
	})

	t.Run("message template metadata", func(t *testing.T) {
		rule := govy.NewRule(func(v string) error {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				Error:  fmt.Sprintf("nested error for %q", v),
				Custom: fmt.Sprintf("custom value %q", v),
			})
		}).
			WithDetails(fmt.Sprintf("details for %q", secret)).
			WithExamples(secret).
			WithMessageTemplateString(
				"{{ .PropertyValue }}: {{ .Error }}; {{ .Details }}; {{ index .Examples 0 }}; {{ .Custom }}",
			).
			HideValue()

		assert.EqualError(
			t,
			rule.Validate(secret),
			`[hidden]: nested error for "[hidden]"; details for "[hidden]"; [hidden]; custom value "[hidden]"`,
		)
	})

	t.Run("configured message", func(t *testing.T) {
		rule := govy.NewRule(func(string) error {
			return errors.New("original error")
		}).
			WithMessage(fmt.Sprintf("invalid value %q", secret)).
			WithDetails(fmt.Sprintf("details for %q", secret)).
			WithExamples(secret).
			HideValue()

		assert.EqualError(
			t,
			rule.Validate(secret),
			`invalid value "[hidden]" (e.g. '[hidden]'); details for "[hidden]"`,
		)
	})

	t.Run("structured rule error", func(t *testing.T) {
		rule := govy.NewRule(func(v string) error {
			return govy.NewRuleError(fmt.Sprintf("invalid value %q", v), "inner")
		}).
			WithErrorCode("outer").
			WithDescription("validation description").
			HideValue()

		assert.Equal(t, &govy.RuleError{
			Message:     `invalid value "[hidden]"`,
			Code:        "outer:inner",
			Description: "validation description",
		}, rule.Validate(secret))
	})

	t.Run("structured rule error with configured message", func(t *testing.T) {
		rule := govy.NewRule(func(string) error {
			return govy.NewRuleError("underlying error", "inner")
		}).
			WithMessage(fmt.Sprintf("invalid value %q", secret)).
			WithDetails(fmt.Sprintf("details for %q", secret)).
			WithExamples(secret).
			WithErrorCode("outer").
			WithDescription("validation description").
			HideValue()

		assert.Equal(t, &govy.RuleError{
			Message:     `invalid value "[hidden]" (e.g. '[hidden]'); details for "[hidden]"`,
			Code:        "outer:inner",
			Description: "validation description",
		}, rule.Validate(secret))
	})

	t.Run("structured property error without property value", func(t *testing.T) {
		rule := govy.NewRule(func(v string) error {
			return govy.NewPropertyError(
				jsonpath.New().Name("nested"),
				nil,
				govy.NewRuleError(fmt.Sprintf("invalid value %q", v)),
			)
		}).HideValue()

		err := mustErrorType[*govy.PropertyError](t, rule.Validate(secret))
		assert.Equal(t, &govy.PropertyError{
			PropertyPath: jsonpath.Parse("nested"),
			Errors:       []*govy.RuleError{{Message: `invalid value "[hidden]"`}},
		}, err)
		assertDoesNotContainSecret(t, err, secret)
	})

	t.Run("pointer conversion", func(t *testing.T) {
		rule := govy.NewRule(func(v string) error {
			return fmt.Errorf("invalid value %q", v)
		}).HideValue()

		assert.EqualError(t, govy.RuleToPointer(rule).Validate(ptr(secret)), `invalid value "[hidden]"`)
	})

	t.Run("rule set pointer conversion", func(t *testing.T) {
		ruleSet := govy.NewRuleSet(govy.NewRule(func(v string) error {
			return fmt.Errorf("invalid value %q", v)
		}).HideValue())

		err := govy.RuleSetToPointer(ruleSet).Validate(ptr(secret))
		errs := mustErrorType[govy.RuleSetError](t, err)
		assert.Require(t, assert.Len(t, errs, 1))
		assert.EqualError(t, errs[0], `invalid value "[hidden]"`)
	})
}

func TestPropertyRules_HideValue(t *testing.T) {
	const secret = "secret"

	t.Run("plain error", func(t *testing.T) {
		expectedErr := errors.New("oh no! here's the value: 'secret'")
		propertyRules := govy.For(govy.GetSelf[string]()).
			WithPath(jsonpath.New().Name("test").Name("path")).
			Rules(govy.NewRule(func(string) error { return expectedErr }))
		hiddenPropertyRules := propertyRules.HideValue()

		assert.EqualError(
			t,
			propertyRules.Validate(secret),
			"- 'test.path' with value 'secret':\n  - oh no! here's the value: 'secret'",
		)
		errs := mustPropertyErrors(t, hiddenPropertyRules.Validate(secret))
		assert.Require(t, assert.Len(t, errs, 1))
		assert.Equal(t, &govy.PropertyError{
			PropertyPath:  jsonpath.Parse("test.path"),
			PropertyValue: "",
			Errors:        []*govy.RuleError{{Message: "oh no! here's the value: '[hidden]'"}},
		}, errs[0])
		assert.EqualError(
			t,
			propertyRules.Validate(secret),
			"- 'test.path' with value 'secret':\n  - oh no! here's the value: 'secret'",
		)
	})

	t.Run("empty value", func(t *testing.T) {
		propertyRules := govy.For(govy.GetSelf[string]()).
			WithName("password").
			Required().
			HideValue()

		err := propertyRules.Validate("")
		assert.EqualError(t, err, "- 'password':\n  - property is required but was empty")
	})

	t.Run("whitespace-only value", func(t *testing.T) {
		tests := map[string]string{
			"ASCII whitespace":   " \t\n",
			"Unicode whitespace": "\u3000",
		}
		for name, value := range tests {
			t.Run(name, func(t *testing.T) {
				propertyRules := govy.For(govy.GetSelf[string]()).
					WithName("password").
					HideValue().
					Rules(rules.StringNotEmpty())

				err := propertyRules.Validate(value)
				assert.EqualError(t, err, "- 'password':\n  - string must not be empty")
			})
		}
	})

	t.Run("rule-only scope", func(t *testing.T) {
		propertyRules := govy.For(govy.GetSelf[string]()).
			WithName("password").
			Rules(govy.NewRule(func(v string) error {
				return fmt.Errorf("invalid value %q", v)
			}).HideValue())

		errs := mustPropertyErrors(t, propertyRules.Validate(secret))
		assert.Require(t, assert.Len(t, errs, 1))
		assert.Equal(t, &govy.PropertyError{
			PropertyPath:  jsonpath.Parse("password"),
			PropertyValue: secret,
			Errors:        []*govy.RuleError{{Message: `invalid value "[hidden]"`}},
		}, errs[0])
	})

	t.Run("sibling property scope", func(t *testing.T) {
		type credentials struct {
			Password string
			Username string
		}
		invalidValueRule := func() govy.Rule[string] {
			return govy.NewRule(func(v string) error {
				return fmt.Errorf("invalid value %q", v)
			})
		}
		validator := govy.New(
			govy.For(func(v credentials) string { return v.Password }).
				WithName("password").
				HideValue().
				Rules(invalidValueRule()),
			govy.For(func(v credentials) string { return v.Username }).
				WithName("username").
				Rules(invalidValueRule()),
		)

		err := validator.Validate(credentials{Password: secret, Username: "public"})
		assert.Equal(t, &govy.ValidatorError{Errors: govy.PropertyErrors{
			{
				PropertyPath: jsonpath.Parse("password"),
				Errors:       []*govy.RuleError{{Message: `invalid value "[hidden]"`}},
			},
			{
				PropertyPath:  jsonpath.Parse("username"),
				PropertyValue: "public",
				Errors:        []*govy.RuleError{{Message: `invalid value "public"`}},
			},
		}}, err)
	})

	t.Run("structured property error", func(t *testing.T) {
		const nestedSecret = "nested-secret"
		rule := govy.NewRule(func(string) error {
			return govy.NewPropertyError(
				jsonpath.New().Name("nested"),
				nestedSecret,
				govy.NewRuleError(fmt.Sprintf("invalid value %q", nestedSecret), "inner"),
			)
		}).WithErrorCode("outer")
		propertyRules := govy.For(govy.GetSelf[string]()).
			WithName("password").
			HideValue().
			Rules(rule)

		errs := mustPropertyErrors(t, propertyRules.Validate(secret))
		assert.Require(t, assert.Len(t, errs, 1))
		assert.Equal(t, &govy.PropertyError{
			PropertyPath:  jsonpath.Parse("password.nested"),
			PropertyValue: "",
			Errors: []*govy.RuleError{{
				Message: `invalid value "[hidden]"`,
				Code:    "outer:inner",
			}},
		}, errs[0])
		assertDoesNotContainSecret(t, errs[0], nestedSecret)
	})

	t.Run("rule set", func(t *testing.T) {
		propertyRules := govy.For(govy.GetSelf[string]()).
			WithName("password").
			HideValue().
			Rules(govy.NewRuleSet(govy.NewRule(func(v string) error {
				return fmt.Errorf("invalid value %q", v)
			})))

		propErr := assertHiddenPropertyError(
			t,
			propertyRules.Validate(secret),
			"password",
			`invalid value "[hidden]"`,
		)
		assertDoesNotContainSecret(t, propErr, secret)
	})

	t.Run("included validator", func(t *testing.T) {
		included := govy.New(
			govy.For(govy.GetSelf[string]()).
				WithName("value").
				Rules(govy.NewRule(func(v string) error {
					return fmt.Errorf("invalid value %q", v)
				})),
		)
		propertyRules := govy.For(govy.GetSelf[string]()).
			WithName("password").
			HideValue().
			Include(included)

		propErr := assertHiddenPropertyError(
			t,
			propertyRules.Validate(secret),
			"password.value",
			`invalid value "[hidden]"`,
		)
		assertDoesNotContainSecret(t, propErr, secret)
	})

	t.Run("included whole-slice rules", func(t *testing.T) {
		included := govy.New(
			govy.ForSlice(govy.GetSelf[[]string]()).
				Rules(govy.NewRule(func(v []string) error {
					return fmt.Errorf("invalid value [%q]", v[0])
				})),
		)
		propertyRules := govy.For(govy.GetSelf[[]string]()).
			WithName("passwords").
			HideValue().
			Include(included)

		propErr := assertHiddenPropertyError(
			t,
			propertyRules.Validate([]string{secret}),
			"passwords",
			"invalid value [hidden]",
		)
		assert.False(t, propErr.IsSliceElementError)
		assertDoesNotContainSecret(t, propErr, secret)
	})

	t.Run("included slice element rules", func(t *testing.T) {
		included := govy.New(
			govy.ForSlice(govy.GetSelf[[]string]()).
				RulesForEach(govy.NewRule(func(v string) error {
					return fmt.Errorf("invalid value %q", v)
				})),
		)
		propertyRules := govy.For(govy.GetSelf[[]string]()).
			WithName("passwords").
			HideValue().
			Include(included)

		propErr := assertHiddenPropertyError(
			t,
			propertyRules.Validate([]string{secret}),
			"passwords[0]",
			`invalid value "[hidden]"`,
		)
		assert.True(t, propErr.IsSliceElementError)
		assertDoesNotContainSecret(t, propErr, secret)
	})

	t.Run("included whole-map rules", func(t *testing.T) {
		included := govy.New(
			govy.ForMap(govy.GetSelf[map[string]string]()).
				Rules(govy.NewRule(func(v map[string]string) error {
					return fmt.Errorf(`invalid value {"primary":%q}`, v["primary"])
				})),
		)
		propertyRules := govy.For(govy.GetSelf[map[string]string]()).
			WithName("passwords").
			HideValue().
			Include(included)

		propErr := assertHiddenPropertyError(
			t,
			propertyRules.Validate(map[string]string{"primary": secret}),
			"passwords",
			"invalid value [hidden]",
		)
		assertDoesNotContainSecret(t, propErr, secret)
	})

	t.Run("included map key rules", func(t *testing.T) {
		const key = "primary"
		included := govy.New(
			govy.ForMap(govy.GetSelf[map[string]string]()).
				RulesForKeys(govy.NewRule(func(v string) error {
					return fmt.Errorf("invalid value %q", v)
				})),
		)
		propertyRules := govy.For(govy.GetSelf[map[string]string]()).
			WithName("passwords").
			HideValue().
			Include(included)

		propErr := assertHiddenPropertyError(
			t,
			propertyRules.Validate(map[string]string{key: "public"}),
			"passwords.primary",
			`invalid value "[hidden]"`,
		)
		assert.True(t, propErr.IsKeyError)
	})

	t.Run("included map value rules", func(t *testing.T) {
		included := govy.New(
			govy.ForMap(govy.GetSelf[map[string]string]()).
				RulesForValues(govy.NewRule(func(v string) error {
					return fmt.Errorf("invalid value %q", v)
				})),
		)
		propertyRules := govy.For(govy.GetSelf[map[string]string]()).
			WithName("passwords").
			HideValue().
			Include(included)

		propErr := assertHiddenPropertyError(
			t,
			propertyRules.Validate(map[string]string{"primary": secret}),
			"passwords.primary",
			`invalid value "[hidden]"`,
		)
		assert.False(t, propErr.IsKeyError)
		assertDoesNotContainSecret(t, propErr, secret)
	})

	t.Run("included map item rules", func(t *testing.T) {
		included := govy.New(
			govy.ForMap(govy.GetSelf[map[string]string]()).
				RulesForItems(govy.NewRule(func(item govy.MapItem[string, string]) error {
					return fmt.Errorf("invalid value %q", item.Value)
				})),
		)
		propertyRules := govy.For(govy.GetSelf[map[string]string]()).
			WithName("passwords").
			HideValue().
			Include(included)

		propErr := assertHiddenPropertyError(
			t,
			propertyRules.Validate(map[string]string{"primary": secret}),
			"passwords.primary",
			`invalid value "[hidden]"`,
		)
		assertDoesNotContainSecret(t, propErr, secret)
	})
}

func TestTransform_HideValue(t *testing.T) {
	t.Run("transform failure", func(t *testing.T) {
		transformed := govy.Transform(govy.GetSelf[string](), strconv.Atoi).
			WithName("prop").
			HideValue().
			Rules(rules.GT(123))

		errs := mustPropertyErrors(t, transformed.Validate("secret!"))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, expectedErrorOutput(t, "property_error_transform_with_hidden_value.txt"))
		assert.True(t, govy.HasErrorCode(errs, govy.ErrorCodeTransform))
	})

	t.Run("rule failure", func(t *testing.T) {
		transformed := govy.Transform(
			govy.GetSelf[string](),
			func(v string) (string, error) { return "transformed-" + v, nil },
		).
			WithName("prop").
			HideValue().
			Rules(govy.NewRule(func(v string) error {
				return fmt.Errorf("invalid transformed value %q", v)
			}))

		propErr := assertHiddenPropertyError(
			t,
			transformed.Validate("secret"),
			"prop",
			`invalid transformed value "[hidden]"`,
		)
		assertDoesNotContainSecret(t, propErr, "secret")
	})
}

func TestPlan_HiddenProperty(t *testing.T) {
	propertyRules := govy.For(govy.GetSelf[string]()).
		WithName("password").
		Rules(rules.StringNotEmpty())
	visiblePlan, err := govy.Plan(govy.New(propertyRules))
	assert.Require(t, assert.NoError(t, err))
	assert.Require(t, assert.Len(t, visiblePlan.Properties, 1))
	assert.False(t, visiblePlan.Properties[0].IsHidden)

	hiddenPlan, err := govy.Plan(govy.New(propertyRules.HideValue()))
	assert.Require(t, assert.NoError(t, err))
	assert.Require(t, assert.Len(t, hiddenPlan.Properties, 1))
	assert.True(t, hiddenPlan.Properties[0].IsHidden)
}

func assertHiddenPropertyError(t *testing.T, err error, path, message string) *govy.PropertyError {
	t.Helper()

	errs := mustPropertyErrors(t, err)
	assert.Require(t, assert.Len(t, errs, 1))
	assert.Equal(t, jsonpath.Parse(path), errs[0].PropertyPath)
	assert.Equal(t, "", errs[0].PropertyValue)
	assert.Require(t, assert.Len(t, errs[0].Errors, 1))
	assert.Equal(t, message, errs[0].Errors[0].Message)
	return errs[0]
}

func assertDoesNotContainSecret(t *testing.T, value any, secret string) {
	t.Helper()

	if rendered := fmt.Sprint(value); strings.Contains(rendered, secret) {
		assert.Fail(t, "Rendered error contains secret %q: %s", secret, rendered)
	}
	if errValue, ok := value.(error); ok && strings.Contains(errValue.Error(), secret) {
		assert.Fail(t, "Rendered error contains secret %q: %s", secret, errValue.Error())
	}
	data, err := json.Marshal(value)
	assert.Require(t, assert.NoError(t, err))
	if strings.Contains(string(data), secret) {
		assert.Fail(t, "JSON error contains secret %q: %s", secret, data)
	}
}
