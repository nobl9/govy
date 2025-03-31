package rules

import (
	"fmt"
	"slices"
	"strings"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/collections"
	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// OneOf checks if the property's value matches one of the provided values.
// The values must be comparable.
func OneOf[T comparable](values ...T) govy.Rule[T] {
	tpl := messagetemplates.Get(messagetemplates.OneOfTemplate)

	return govy.NewRule(func(v T) error {
		for i := range values {
			if v == values[i] {
				return nil
			}
		}
		return govy.NewRuleErrorTemplate(govy.TemplateVars{
			PropertyValue:   v,
			ComparisonValue: values,
		})
	}).
		WithErrorCode(ErrorCodeOneOf).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: values,
		})).
		WithPlanModifiers(govy.RulePlanModifierValidValues(values...))
}

// OneOfProperties checks if at least one of the properties is set.
// Property is considered set if its value is not empty (non-zero).
func OneOfProperties[S any](getters map[string]func(s S) any) govy.Rule[S] {
	tpl := messagetemplates.Get(messagetemplates.OneOfPropertiesTemplate)

	return govy.NewRule(func(s S) error {
		for _, getter := range getters {
			v := getter(s)
			if !internal.IsEmpty(v) {
				return nil
			}
		}
		return govy.NewRuleErrorTemplate(govy.TemplateVars{
			PropertyValue:   s,
			ComparisonValue: collections.SortedKeys(getters),
		})
	}).
		WithErrorCode(ErrorCodeOneOfProperties).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: collections.SortedKeys(getters),
		}))
}

type mutuallyExclusiveTemplateVars struct {
	// NoProperties is set to true if no properties were set and exactly one was required.
	NoProperties bool
}

// MutuallyExclusive checks if properties are mutually exclusive.
// This means, exactly one of the properties can be set.
// Property is considered set if its value is not empty (non-zero).
// If required is true, then a single non-empty property is required.
func MutuallyExclusive[S any](required bool, getters map[string]func(s S) any) govy.Rule[S] {
	tpl := messagetemplates.Get(messagetemplates.MutuallyExclusiveTemplate)

	return govy.NewRule(func(s S) error {
		var nonEmpty []string
		for name, getter := range getters {
			v := getter(s)
			if internal.IsEmpty(v) {
				continue
			}
			nonEmpty = append(nonEmpty, name)
		}
		switch len(nonEmpty) {
		case 0:
			if !required {
				return nil
			}
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   s,
				ComparisonValue: collections.SortedKeys(getters),
				Custom:          mutuallyExclusiveTemplateVars{NoProperties: true},
			})
		case 1:
			return nil
		default:
			slices.Sort(nonEmpty)
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   s,
				ComparisonValue: nonEmpty,
				Custom:          mutuallyExclusiveTemplateVars{NoProperties: false},
			})
		}
	}).
		WithErrorCode(ErrorCodeMutuallyExclusive).
		WithMessageTemplate(tpl).
		WithDescription(func() string {
			return fmt.Sprintf("properties are mutually exclusive: %s",
				strings.Join(collections.SortedKeys(getters), ", "))
		}())
}
