package govy

import (
	"reflect"
	"sort"
	"strings"

	"golang.org/x/exp/maps"
)

// ValidatorPlan is a validation plan for a single [Validator].
type ValidatorPlan struct {
	Name       string         `json:"name,omitempty"`
	Properties []PropertyPlan `json:"properties"`
}

// PropertyPlan is a validation plan for a single [PropertyRules].
type PropertyPlan struct {
	// Path is a JSON path to the property.
	Path string `json:"path"`
	// Type is a Go type name of the property.
	Type string `json:"type"`
	// Package is the full package path of the Type.
	Package string `json:"package,omitempty"`
	// IsOptional indicates if the property was marked with [PropertyRules.OmitEmpty].
	IsOptional bool `json:"isOptional,omitempty"`
	// IsHidden indicates if the property was marked with [PropertyRules.HideValue].
	IsHidden bool       `json:"isHidden,omitempty"`
	Examples []string   `json:"examples,omitempty"`
	Rules    []RulePlan `json:"rules,omitempty"`
}

// RulePlan is a validation plan for a single [Rule].
type RulePlan struct {
	Description string    `json:"description"`
	Details     string    `json:"details,omitempty"`
	ErrorCode   ErrorCode `json:"errorCode,omitempty"`
	// Conditions are all the predicates set through [PropertyRules.When] and [Validator.When]
	// which had [WhenDescription] added to the [WhenOptions].
	Conditions []string `json:"conditions,omitempty"`
	Examples   []string `json:"examples,omitempty"`
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
				Type:       p.propertyPlan.Type,
				Package:    p.propertyPlan.Package,
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

// typeInfo stores the type name and its package if it's available.
type typeInfo struct {
	Name    string
	Package string
}

// getTypeInfo returns the information for the type T.
// It returns the type name without package path or name.
// It strips the pointer '*' from the type name.
// Package is only available if the type is not a built-in type.
func getTypeInfo[T any]() typeInfo {
	typ := reflect.TypeOf(*new(T))
	if typ == nil {
		return typeInfo{}
	}
	var result typeInfo
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Slice {
		typ = typ.Elem()
		result.Name = "[]"
	}
	if typ.PkgPath() == "" {
		result.Name += typ.String()
	} else {
		result.Name += typ.Name()
		result.Package = typ.PkgPath()
	}
	return result
}
