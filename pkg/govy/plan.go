package govy

import (
	"cmp"
	"maps"
	"slices"
	"strings"

	"github.com/nobl9/govy/internal/collections"
	"github.com/nobl9/govy/internal/compare"
)

// ValidatorPlan is a validation plan for a single [Validator].
type ValidatorPlan struct {
	// Name is the value provided to [Validator.WithName].
	Name string `json:"name,omitempty"`
	// Properties which this [Validator] defines.
	Properties []*PropertyPlan `json:"properties"`
}

// PropertyPlan is a validation plan for a single [PropertyRules].
type PropertyPlan struct {
	// Path is a JSON path to the property.
	Path string `json:"path"`
	// TypeInfo contains the type information of the property.
	TypeInfo TypeInfo `json:"typeInfo"`
	// IsHidden indicates if the property was marked with [PropertyRules.HideValue].
	IsHidden bool `json:"isHidden,omitempty"`
	// Examples lists example, valid values for this property.
	// These values are not exhaustive, for an exhaustive list of valid values see [Values].
	Examples []string `json:"examples,omitempty"`
	// Values unlike [Examples] should list ALL valid values for this property.
	// These values are constructed as an intersection of all [RulePlan] values
	// for this property.
	Values []string `json:"values,omitempty"`
	// Rules which apply to this property.
	Rules []RulePlan `json:"rules,omitempty"`
}

// TypeInfo contains the type information of a property.
type TypeInfo struct {
	// Name is a Go type name.
	// Example: "Pod", "string", "int", "bool", etc.
	Name string `json:"name"`
	// Kind is a Go type kind.
	// Example: "string", "int", "bool", "struct", "slice", etc.
	Kind string `json:"kind"`
	// Package is the full package path of the type.
	// It's empty for builtin types.
	// Example: "github.com/nobl9/govy/pkg/govy", "time", etc.
	Package string `json:"package,omitempty"`
}

// RulePlan is a validation plan for a single [Rule].
type RulePlan struct {
	// Description is the value provided to [Rule.WithDescription].
	Description string `json:"description"`
	// Details is the value provided to [Rule.WithDetails].
	Details string `json:"details,omitempty"`
	// ErrorCode is the value provided to [Rule.WithErrorCode].
	ErrorCode ErrorCode `json:"errorCode,omitempty"`
	// Conditions are all the predicates set through [PropertyRules.When] and [Validator.When]
	// which had [WhenDescription] added to the [WhenOptions].
	Conditions []string `json:"conditions,omitempty"`
	// Examples is the value provided to [Rule.WithExamples].
	// These values are not exhaustive, for an exhaustive list of valid values see [PropertyPlan.ValidValues].
	Examples []string `json:"examples,omitempty"`

	// values unlike [Examples] should list ALL valid values which meet this rule.
	// It is not exported as it is only here to contribute to the [PropertyPlan.ValidValues].
	values []string
}

func (r RulePlan) isEmpty() bool {
	return r.Description == "" && r.Details == "" && r.ErrorCode == ""
}

// Plan creates a validation plan for the provided [Validator].
// Each property is represented by a [PropertyPlan] which aggregates its every [RulePlan].
// If a property does not have any rules, it won't be included in the result.
func Plan[S any](v Validator[S]) *ValidatorPlan {
	builders := make([]planBuilder, 0)
	v.plan(planBuilder{path: "$", children: &builders})
	properties := aggregatePropertyPlans(builders)
	// Post-processing after the properties have been aggregated.
	for _, prop := range properties {
		prop.collectValidValuesFromRules()
		prop.removeDeduplicatedRules()
	}
	name := v.name
	// Best effort name function setting.
	if v.nameFunc != nil {
		name = v.nameFunc(*new(S))
	}
	return &ValidatorPlan{
		Name:       name,
		Properties: properties,
	}
}

func aggregatePropertyPlans(builders []planBuilder) []*PropertyPlan {
	propertiesMap := make(map[string]*PropertyPlan)
	for _, b := range builders {
		entry, ok := propertiesMap[b.path]
		if !ok {
			entry = &PropertyPlan{
				Path:     b.path,
				TypeInfo: b.propertyPlan.TypeInfo,
				Examples: b.propertyPlan.Examples,
				IsHidden: b.propertyPlan.IsHidden,
			}
		}
		if !b.rulePlan.isEmpty() {
			entry.Rules = append(entry.Rules, b.rulePlan)
		}
		propertiesMap[b.path] = entry
	}
	return slices.SortedFunc(
		maps.Values(propertiesMap),
		func(a, b *PropertyPlan) int { return cmp.Compare(a.Path, b.Path) },
	)
}

func (p *PropertyPlan) collectValidValuesFromRules() {
	validValuesSlices := make([][]string, 0)
	for _, rule := range p.Rules {
		if len(rule.values) > 0 {
			validValuesSlices = append(validValuesSlices, rule.values)
		}
	}
	// TODO: If there are indeed conflicting elements, we might want to drop an error?
	p.Values = collections.Intersection(validValuesSlices...)
}

// removeDeduplicatedRules removes duplicate rules from the [PropertyPlan].
func (p *PropertyPlan) removeDeduplicatedRules() {
	if len(p.Rules) == 0 {
		return
	}
	uniqueRules := make([]RulePlan, 0, len(p.Rules))
	for _, rule := range p.Rules {
		isDuplicate := false
		for _, uniqueRule := range uniqueRules {
			if compare.EqualExportedFields(rule, uniqueRule) {
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			uniqueRules = append(uniqueRules, rule)
		}
	}
	p.Rules = uniqueRules
}

// planner is an interface for types that can create a [PropertyPlan] or [RulePlan].
type planner interface {
	plan(builder planBuilder)
}

// planBuilder is used to traverse the validation rules and build a slice of [PropertyPlan].
type planBuilder struct {
	path         string
	rulePlan     RulePlan
	propertyPlan PropertyPlan
	// children stores every rule for the current property.
	// It's not safe for concurrent usage.
	children *[]planBuilder
}

func (p planBuilder) appendPath(path string) planBuilder {
	builder := planBuilder{
		children:     p.children,
		rulePlan:     p.rulePlan,
		propertyPlan: p.propertyPlan,
	}
	switch {
	case p.path == "" && path != "":
		builder.path = path
	case p.path != "" && path != "":
		if strings.HasPrefix(path, "[") {
			builder.path = p.path + path
		} else {
			builder.path = p.path + "." + path
		}
	default:
		builder.path = p.path
	}
	return builder
}

func (p planBuilder) setExamples(examples ...string) planBuilder {
	p.propertyPlan.Examples = examples
	return p
}

func appendPredicatesToPlanBuilder[T any](builder planBuilder, predicates []predicateContainer[T]) planBuilder {
	for _, predicate := range predicates {
		if predicate.description == "" {
			continue
		}
		builder.rulePlan.Conditions = append(builder.rulePlan.Conditions, predicate.description)
	}
	return builder
}
