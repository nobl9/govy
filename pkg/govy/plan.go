package govy

import (
	"sort"
	"strings"

	"github.com/nobl9/govy/internal/collections"
	"golang.org/x/exp/maps"
)

// ValidatorPlan is a validation plan for a single [Validator].
type ValidatorPlan struct {
	// Name is the value provided to [Validator.WithName].
	Name string `json:"name,omitempty"`
	// Properties which this [Validator] defines.
	Properties []PropertyPlan `json:"properties"`
}

// PropertyPlan is a validation plan for a single [PropertyRules].
type PropertyPlan struct {
	// Path is a JSON path to the property.
	Path string `json:"path"`
	// TypeInfo contains the type information of the property.
	TypeInfo TypeInfo `json:"typeInfo"`
	// IsOptional indicates if the property was marked with [PropertyRules.OmitEmpty].
	IsOptional bool `json:"isOptional,omitempty"`
	// IsHidden indicates if the property was marked with [PropertyRules.HideValue].
	IsHidden bool `json:"isHidden,omitempty"`
	// Examples lists example, valid values for this property.
	// These values are not exhaustive, for an exhaustive list of valid values see [ValidValues].
	Examples []string `json:"examples,omitempty"`
	// Values unlike [Examples] should list ALL valid values for this property.
	// These values are constructed as an intersection of all [RulePlan.Values]
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
	// These values are not exhaustive, for an exhaustive list of valid values see [ValidValues].
	Examples []string `json:"examples,omitempty"`

	// values unlike [Examples] should list ALL valid values which meet this rule.
	// It is not exported as it is only here to contribute to the [PropertyPlan.ValidValues].
	values []string
}

func (r RulePlan) isEmpty() bool {
	return r.Description == "" && r.Details == "" && r.ErrorCode == "" && len(r.Conditions) == 0
}

// Plan creates a validation plan for the provided [Validator].
// Each property is represented by a [PropertyPlan] which aggregates its every [RulePlan].
// If a property does not have any rules, it won't be included in the result.
func Plan[S any](v Validator[S]) *ValidatorPlan {
	all := make([]planBuilder, 0)
	v.plan(planBuilder{path: "$", children: &all})
	propertiesMap := make(map[string]PropertyPlan)
	for _, p := range all {
		entry, ok := propertiesMap[p.path]
		if ok {
			entry.Rules = append(entry.Rules, p.rulePlan)
			propertiesMap[p.path] = entry
		} else {
			entry = PropertyPlan{
				Path:       p.path,
				TypeInfo:   p.propertyPlan.TypeInfo,
				Examples:   p.propertyPlan.Examples,
				IsOptional: p.propertyPlan.IsOptional,
				IsHidden:   p.propertyPlan.IsHidden,
			}
			if !p.rulePlan.isEmpty() {
				entry.Rules = append(entry.Rules, p.rulePlan)
			}
			propertiesMap[p.path] = entry
		}
	}
	properties := maps.Values(propertiesMap)
	sort.Slice(properties, func(i, j int) bool { return properties[i].Path < properties[j].Path })

	for i, prop := range properties {
		validValuesSlices := make([]string, 0)
		for _, rule := range prop.Rules {
			validValuesSlices = append(validValuesSlices, rule.values...)
		}
		// TODO: If there are indeed conflicting elements, we might want to drop an error!
		prop.Values = collections.Intersection(validValuesSlices)
		properties[i] = prop
	}

	return &ValidatorPlan{
		Name:       v.name,
		Properties: properties,
	}
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
