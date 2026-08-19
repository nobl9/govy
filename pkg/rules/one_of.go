package rules

import (
	"fmt"
	"slices"
	"strings"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/collections"
	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonschema"
)

// OneOf checks if the property's value matches one of the provided values.
// The values must be comparable.
//
// For reversed rule see [NotOneOf].
// It panics if no values are provided.
func OneOf[T comparable](values ...T) govy.Rule[T] {
	if len(values) == 0 {
		panic("values must not be empty")
	}
	tpl := messagetemplates.Get(messagetemplates.OneOfTemplate)

	return govy.NewRule(func(v T) error {
		if slices.Contains(values, v) {
			return nil
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
		WithPlanModifiers(govy.RulePlanModifierValidValues(values...)).
		WithJSONSchema(func(ctx govy.JSONSchemaBuilderContext) error {
			converted, err := jsonSchemaValues(values)
			if err != nil {
				return err
			}
			addJSONSchemaEnum(ctx.Schema, converted)
			return nil
		})
}

// NotOneOf checks if the property's value does not match any of the provided values.
// The values must be comparable.
//
// For reversed rule see [OneOf].
// It panics if no values are provided.
func NotOneOf[T comparable](values ...T) govy.Rule[T] {
	if len(values) == 0 {
		panic("values must not be empty")
	}
	tpl := messagetemplates.Get(messagetemplates.NotOneOfTemplate)

	return govy.NewRule(func(v T) error {
		if slices.Contains(values, v) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   v,
				ComparisonValue: values,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeNotOneOf).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: values,
		})).
		WithJSONSchema(func(ctx govy.JSONSchemaBuilderContext) error {
			converted, err := jsonSchemaValues(values)
			if err != nil {
				return err
			}
			addJSONSchemaNot(ctx.Schema, &jsonschema.Schema{Enum: converted})
			return nil
		})
}

// OneOfProperties checks if at least one of the properties is set.
// Property is considered set if its value is not empty (non-zero).
// It panics if getters is empty.
func OneOfProperties[T any](getters map[string]func(parent T) any) govy.Rule[T] {
	if len(getters) == 0 {
		panic("getters must not be empty")
	}
	tpl := messagetemplates.Get(messagetemplates.OneOfPropertiesTemplate)
	sortedKeys := collections.SortedKeys(getters)

	return govy.NewRule(func(parent T) error {
		for _, getter := range getters {
			v := getter(parent)
			if !internal.IsEmpty(v) {
				return nil
			}
		}
		return govy.NewRuleErrorTemplate(govy.TemplateVars{
			PropertyValue:   parent,
			ComparisonValue: collections.SortedKeys(getters),
		})
	}).
		WithErrorCode(ErrorCodeOneOfProperties).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: sortedKeys,
		})).
		WithJSONSchema(func(ctx govy.JSONSchemaBuilderContext) error {
			addJSONSchemaAnyOf(ctx.Schema, jsonSchemaRequiredAlternatives(sortedKeys))
			return nil
		})
}

type mutuallyExclusiveTemplateVars struct {
	// NoProperties is set to true if no properties were set and exactly one was required.
	NoProperties bool
}

// MutuallyExclusive checks if properties are mutually exclusive.
// This means, exactly one of the properties can be set.
// Property is considered set if its value is not empty (non-zero).
// If required is true, then a single non-empty property is required.
// It panics if getters contains fewer than two properties.
func MutuallyExclusive[T any](required bool, getters map[string]func(parent T) any) govy.Rule[T] {
	if len(getters) < 2 {
		panic("getters must contain at least two properties")
	}
	tpl := messagetemplates.Get(messagetemplates.MutuallyExclusiveTemplate)
	sortedKeys := collections.SortedKeys(getters)

	return govy.NewRule(func(parent T) error {
		var nonEmpty []string
		for name, getter := range getters {
			v := getter(parent)
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
				PropertyValue:   parent,
				ComparisonValue: collections.SortedKeys(getters),
				Custom:          mutuallyExclusiveTemplateVars{NoProperties: true},
			})
		case 1:
			return nil
		default:
			slices.Sort(nonEmpty)
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   parent,
				ComparisonValue: nonEmpty,
				Custom:          mutuallyExclusiveTemplateVars{NoProperties: false},
			})
		}
	}).
		WithErrorCode(ErrorCodeMutuallyExclusive).
		WithMessageTemplate(tpl).
		WithDescription(func() string {
			return fmt.Sprintf("properties are mutually exclusive: %s",
				strings.Join(sortedKeys, ", "))
		}()).
		WithJSONSchema(func(ctx govy.JSONSchemaBuilderContext) error {
			if required {
				addJSONSchemaOneOf(ctx.Schema, jsonSchemaRequiredAlternatives(sortedKeys))
				return nil
			}
			pairs := make([]*jsonschema.Schema, 0, len(sortedKeys)*(len(sortedKeys)-1)/2)
			for i := range sortedKeys {
				for j := i + 1; j < len(sortedKeys); j++ {
					pairs = append(pairs, &jsonschema.Schema{
						Required: []string{sortedKeys[i], sortedKeys[j]},
					})
				}
			}
			addJSONSchemaNot(ctx.Schema, &jsonschema.Schema{AnyOf: pairs})
			return nil
		})
}

type mutuallyDependentTemplateVars struct {
	NonEmptyProperties []string
	EmptyProperties    []string
}

// MutuallyDependent checks if properties are mutually dependent.
// This means, if any of the properties is set, the rest must be also set.
// Property is considered set if its value is not empty (non-zero).
// It panics if getters contains fewer than two properties.
func MutuallyDependent[T any](getters map[string]func(parent T) any) govy.Rule[T] {
	if len(getters) < 2 {
		panic("getters must contain at least two properties")
	}
	tpl := messagetemplates.Get(messagetemplates.MutuallyDependentTemplate)
	sortedKeys := collections.SortedKeys(getters)

	return govy.NewRule(func(parent T) error {
		emptyIndexes := make([]bool, len(getters))
		emptyCtr := 0
		for i, name := range sortedKeys {
			v := getters[name](parent)
			if internal.IsEmpty(v) {
				emptyIndexes[i] = true
				emptyCtr++
				continue
			}
		}
		if emptyCtr == 0 || emptyCtr == len(getters) {
			return nil
		}
		empty := make([]string, 0, emptyCtr)
		nonEmpty := make([]string, 0, len(getters)-emptyCtr)
		for i, isEmpty := range emptyIndexes {
			switch isEmpty {
			case true:
				empty = append(empty, sortedKeys[i])
			case false:
				nonEmpty = append(nonEmpty, sortedKeys[i])
			}
		}
		return govy.NewRuleErrorTemplate(govy.TemplateVars{
			PropertyValue:   parent,
			ComparisonValue: sortedKeys,
			Custom: mutuallyDependentTemplateVars{
				NonEmptyProperties: nonEmpty,
				EmptyProperties:    empty,
			},
		})
	}).
		WithErrorCode(ErrorCodeMutuallyDependent).
		WithMessageTemplate(tpl).
		WithDescription(func() string {
			return fmt.Sprintf("properties are mutually dependent: %s", strings.Join(sortedKeys, ", "))
		}()).
		WithJSONSchema(func(ctx govy.JSONSchemaBuilderContext) error {
			addJSONSchemaAnyOf(ctx.Schema, []*jsonschema.Schema{
				{Required: sortedKeys},
				{Not: &jsonschema.Schema{AnyOf: jsonSchemaRequiredAlternatives(sortedKeys)}},
			})
			return nil
		})
}
