package govy

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/nobl9/govy/internal/collections"
	"github.com/nobl9/govy/pkg/jsonpath"
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
	Path jsonpath.Path `json:"path"`
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

func (r RulePlan) equal(r2 RulePlan) bool {
	return r.Description == r2.Description &&
		r.Details == r2.Details &&
		r.ErrorCode == r2.ErrorCode &&
		collections.EqualSlices(r.Conditions, r2.Conditions) &&
		collections.EqualSlices(r.Examples, r2.Examples)
}

// planOptions contains options for configuring the behavior of the [Plan] function.
type planOptions struct {
	requirePredicateDescriptions bool
}

type PlanOption func(options planOptions) planOptions

// PlanRequirePredicateDescription returns a [PlanOption] that will cause [Plan] to return an error
// if any [Predicate] set through [Validator.When] or [PropertyRules.When] does not have
// a description provided via [WhenDescription].
func PlanRequirePredicateDescription() PlanOption {
	return func(options planOptions) planOptions {
		options.requirePredicateDescriptions = true
		return options
	}
}

// PlanStrictMode bundles all [Plan] validations into a single [PlanOption].
// These include:
//   - [PlanRequirePredicateDescription]
func PlanStrictMode() PlanOption {
	return func(options planOptions) planOptions {
		options = PlanRequirePredicateDescription()(options)
		return options
	}
}

// Plan creates a validation plan for the provided [Validator].
// Each property is represented by a [PropertyPlan] which aggregates its every [RulePlan].
// If a property does not have any rules, it won't be included in the result.
// You can customize the behavior of the plan generation by providing [PlanOption].
func Plan[T any](v Validator[T], opts ...PlanOption) (*ValidatorPlan, error) {
	builders := make([]planBuilder, 0)
	rootBuilder := planBuilder{
		propertyPath:        jsonpath.NewRoot(),
		path:                &builders,
		missingDescriptions: ptr(make([]predicateLocation, 0)),
	}
	for _, opt := range opts {
		rootBuilder.options = opt(rootBuilder.options)
	}
	v.plan(rootBuilder)

	if err := rootBuilder.validate(); err != nil {
		return nil, err
	}

	properties := aggregatePropertyPlans(builders)
	// Post-processing after the properties have been aggregated.
	for _, prop := range properties {
		prop.collectValidValuesFromRules()
		prop.removeDeduplicatedRules()
	}
	name := v.name
	// Best effort name function setting.
	if v.nameFunc != nil {
		name = v.nameFunc(*new(T))
	}
	return &ValidatorPlan{
		Name:       name,
		Properties: properties,
	}, nil
}

type predicateLocation struct {
	propertyPath jsonpath.Path
}

// missingPredicateDescriptionsError is returned when [Plan] is called with
// [PlanRequirePredicateDescription] option and there are any [Predicate] without description.
type missingPredicateDescriptionsError struct {
	locations []predicateLocation
}

func newMissingPredicateDescriptionsError(locations []predicateLocation) *missingPredicateDescriptionsError {
	return &missingPredicateDescriptionsError{locations: locations}
}

func (e *missingPredicateDescriptionsError) Error() string {
	var paths []string
	for _, loc := range e.locations {
		if loc.propertyPath.IsEmpty() || loc.propertyPath.IsRoot() {
			paths = append(paths, "validator level")
		} else {
			paths = append(paths, loc.propertyPath.String())
		}
	}
	return fmt.Sprintf("predicates without description found at: %s", strings.Join(paths, ", "))
}

func aggregatePropertyPlans(builders []planBuilder) []*PropertyPlan {
	propertiesMap := make(map[string]*PropertyPlan)
	for _, b := range builders {
		key := b.propertyPath.String()
		entry, ok := propertiesMap[key]
		if !ok {
			entry = &PropertyPlan{
				Path:     b.propertyPath,
				TypeInfo: b.propertyPlan.TypeInfo,
				Examples: b.propertyPlan.Examples,
				IsHidden: b.propertyPlan.IsHidden,
			}
		}
		if !b.rulePlan.isEmpty() {
			entry.Rules = append(entry.Rules, b.rulePlan)
		}
		propertiesMap[key] = entry
	}
	return slices.SortedFunc(
		maps.Values(propertiesMap),
		func(a, b *PropertyPlan) int { return a.Path.Compare(b.Path) },
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
		isDuplicate := slices.ContainsFunc(uniqueRules, rule.equal)
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

type planBuilder struct {
	propertyPath        jsonpath.Path
	rulePlan            RulePlan
	propertyPlan        PropertyPlan
	path                *[]planBuilder
	missingDescriptions *[]predicateLocation
	options             planOptions
}

func (p planBuilder) appendPath(path jsonpath.Path) planBuilder {
	builder := planBuilder{
		path:                p.path,
		missingDescriptions: p.missingDescriptions,
		options:             p.options,
		rulePlan:            p.rulePlan,
		propertyPlan:        p.propertyPlan,
		propertyPath:        p.propertyPath.Join(path),
	}
	return builder
}

func (p planBuilder) setExamples(examples ...string) planBuilder {
	p.propertyPlan.Examples = examples
	return p
}

func (p planBuilder) validate() error {
	if p.options.requirePredicateDescriptions && len(*p.missingDescriptions) > 0 {
		return newMissingPredicateDescriptionsError(*p.missingDescriptions)
	}
	return nil
}

func appendPredicatesToPlanBuilder[T any](builder planBuilder, predicates []predicateContainer[T]) planBuilder {
	for _, predicate := range predicates {
		if predicate.description == "" {
			if builder.options.requirePredicateDescriptions && builder.missingDescriptions != nil {
				*builder.missingDescriptions = append(*builder.missingDescriptions, predicateLocation{
					propertyPath: builder.propertyPath,
				})
			}
			continue
		}
		builder.rulePlan.Conditions = append(builder.rulePlan.Conditions, predicate.description)
	}
	return builder
}

func ptr[T any](v T) *T { return &v }
