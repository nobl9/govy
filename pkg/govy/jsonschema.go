package govy

import (
	"fmt"
	"maps"
	"math"
	"reflect"
	"slices"

	"github.com/nobl9/govy/pkg/jsonpath"
	"github.com/nobl9/govy/pkg/jsonschema"
)

// JSONSchemaBuilderContext describes the JSON Schema node selected for a builder.
type JSONSchemaBuilderContext struct {
	// Root is the document root.
	Root *jsonschema.Schema
	// Path is the absolute JSON path represented by Schema.
	Path jsonpath.Path
	// Schema is the node selected by Segment.
	Schema *jsonschema.Schema
	// Parent is the node that contains jsonschema.Schema, or nil when Schema is the document root.
	Parent *jsonschema.Schema
	// Segment is the final path segment that selected Schema.
	Segment jsonpath.Segment
}

// JSONSchemaBuilder modifies a selected JSON Schema node.
type JSONSchemaBuilder func(ctx JSONSchemaBuilderContext) error

type jsonSchemaCondition struct {
	scope   jsonpath.Path
	builder JSONSchemaBuilder
}

const maxJSONSchemaPrefixItemsIndex = math.MaxInt - 1

// JSONSchema creates a JSON Schema document for the provided [Validator].
// It uses exclusively [Draft 2020-12] version.
// It returns an error for Go kinds without a default JSON representation.
//
// [Draft 2020-12]: https://json-schema.org/draft/2020-12/schema
func JSONSchema[T any](v Validator[T]) (*jsonschema.Document, error) {
	plan, err := Plan(v, planRecordJSONSchema())
	if err != nil {
		return nil, fmt.Errorf("failed to generate %T: %w", plan, err)
	}
	schemaType, err := jsonSchemaTypeFromTypeInfo(plan.TypeInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to generate JSON Schema type info for validator: %w", err)
	}

	schema := &jsonschema.Schema{
		Title: plan.Name,
		Type:  schemaType,
	}

	for _, prop := range plan.Properties {
		ctx, err := newJSONSchemaBuilderContext(schema, schema, prop.Path)
		if err != nil {
			return nil, fmt.Errorf("failed to generate JSON Schema for %q property: %w", prop.Path, err)
		}
		ctx.Schema.Type, err = jsonSchemaTypeFromTypeInfo(prop.TypeInfo)
		if err != nil {
			return nil, fmt.Errorf("failed to generate JSON Schema type info for %q property: %w", prop.Path, err)
		}
		for _, rule := range prop.Rules {
			for _, builder := range rule.jsonSchemaBuilders {
				if err = builder(ctx); err != nil {
					return nil, fmt.Errorf(
						"failed to build JSON Schema for %q property and %q rule: %w",
						prop.Path,
						rule.ErrorCode,
						err,
					)
				}
			}
		}
	}
	ensureWildcardApplicatorsApplyToAllChildren(schema)

	document := jsonschema.Document(*schema)
	return &document, nil
}

type jsonSchemaPlanBuilder struct {
	conditions  []jsonSchemaCondition
	ruleBuilder JSONSchemaBuilder
}

func newJSONSchemaPlanBuilder(
	conditions []jsonSchemaCondition,
	ruleBuilder JSONSchemaBuilder,
) *jsonSchemaPlanBuilder {
	conditions = slices.Clone(conditions)
	return &jsonSchemaPlanBuilder{
		conditions:  conditions,
		ruleBuilder: ruleBuilder,
	}
}

func (b *jsonSchemaPlanBuilder) Build(ctx JSONSchemaBuilderContext) error {
	propertyPath := ctx.Path
	currentSchema := ctx.Root
	rootScopeDepth := jsonpath.NewRoot().Len()
	currentScopeDepth := rootScopeDepth
	for _, condition := range b.conditions {
		conditionScopeDepth := condition.scope.Len()
		if conditionScopeDepth < currentScopeDepth ||
			conditionScopeDepth > propertyPath.Len() {
			return fmt.Errorf(
				"JSON Schema path %q is not nested under condition scope %q",
				propertyPath,
				condition.scope,
			)
		}
		expectedScope := propertyPath.Slice(0, conditionScopeDepth)
		if !condition.scope.Equal(expectedScope) {
			return fmt.Errorf(
				"JSON Schema path %q is not nested under condition scope %q",
				propertyPath,
				condition.scope,
			)
		}
		relativeScope := jsonpath.NewRoot().Join(
			propertyPath.Slice(currentScopeDepth, conditionScopeDepth),
		)
		scopeContext, err := newJSONSchemaBuilderContext(ctx.Root, currentSchema, relativeScope)
		if err != nil {
			return fmt.Errorf("select JSON Schema condition scope %q: %w", condition.scope, err)
		}
		conditionSchema := new(jsonschema.Schema)
		conditionContext, err := newJSONSchemaBuilderContext(
			ctx.Root,
			conditionSchema,
			jsonpath.NewRoot(),
		)
		if err != nil {
			return err
		}
		conditionContext.Path = condition.scope
		if err = condition.builder(conditionContext); err != nil {
			return fmt.Errorf("build JSON Schema condition at %q: %w", condition.scope, err)
		}
		conditionalSchema := &jsonschema.Schema{
			If:   conditionSchema,
			Then: new(jsonschema.Schema),
		}
		scopeContext.Schema.AllOf = append(scopeContext.Schema.AllOf, conditionalSchema)
		currentSchema = conditionalSchema.Then
		currentScopeDepth = conditionScopeDepth
	}

	relativeProperty := jsonpath.NewRoot().Join(
		propertyPath.Slice(currentScopeDepth, propertyPath.Len()),
	)
	ruleContext, err := newJSONSchemaBuilderContext(ctx.Root, currentSchema, relativeProperty)
	if err != nil {
		return fmt.Errorf("select JSON Schema property %q: %w", propertyPath, err)
	}
	ruleContext.Path = propertyPath
	return b.ruleBuilder(ruleContext)
}

func newJSONSchemaBuilderContext(
	root *jsonschema.Schema,
	schema *jsonschema.Schema,
	path jsonpath.Path,
) (JSONSchemaBuilderContext, error) {
	ctx := JSONSchemaBuilderContext{Root: root, Path: path, Schema: schema}
	for _, segment := range path.Segments() {
		if segment.Kind() == jsonpath.SegmentIndex && segment.Index() > maxJSONSchemaPrefixItemsIndex {
			return JSONSchemaBuilderContext{}, fmt.Errorf(
				"array index %d exceeds maximum supported index %d",
				segment.Index(),
				maxJSONSchemaPrefixItemsIndex,
			)
		}
		if segment.Kind() != jsonpath.SegmentRoot {
			ctx.Parent = ctx.Schema
		}
		ctx.Schema = getJSONSchemaForSegment(ctx.Schema, segment)
		ctx.Segment = segment
	}
	return ctx, nil
}

// ensureWildcardApplicatorsApplyToAllChildren rewrites mixed wildcard and
// explicit child schemas so wildcard constraints also apply to explicit children.
// For example:
//
//	prefixItems: [specific], items: wildcard
//
// becomes:
//
//	prefixItems: [specific], allOf: [{items: wildcard}]
//
// Likewise:
//
//	properties: {name: specific}, additionalProperties: wildcard
//
// becomes:
//
//	properties: {name: specific}, allOf: [{additionalProperties: wildcard}]
//
// The allOf schema has no properties keyword, so its additionalProperties
// constraint applies to name and every other property.
func ensureWildcardApplicatorsApplyToAllChildren(schema *jsonschema.Schema) {
	if schema == nil {
		return
	}
	if schema.Items != nil && len(schema.PrefixItems) > 0 {
		items := schema.Items
		schema.Items = nil
		schema.AllOf = append(schema.AllOf, &jsonschema.Schema{Items: items})
	}
	if schema.AdditionalProperties != nil && len(schema.Properties) > 0 {
		additionalProperties := schema.AdditionalProperties
		schema.AdditionalProperties = nil
		schema.AllOf = append(schema.AllOf, &jsonschema.Schema{
			AdditionalProperties: additionalProperties,
		})
	}

	applicators := []*jsonschema.Schema{
		schema.Not,
		schema.If,
		schema.Then,
		schema.Items,
		schema.AdditionalProperties,
		schema.PropertyNames,
	}
	applicators = append(applicators, schema.AllOf...)
	applicators = append(applicators, schema.AnyOf...)
	applicators = append(applicators, schema.OneOf...)
	applicators = append(applicators, schema.PrefixItems...)
	applicators = slices.AppendSeq(applicators, maps.Values(schema.Properties))

	for _, applicator := range applicators {
		ensureWildcardApplicatorsApplyToAllChildren(applicator)
	}
}

func getJSONSchemaForSegment(schema *jsonschema.Schema, segment jsonpath.Segment) *jsonschema.Schema {
	switch segment.Kind() {
	case jsonpath.SegmentName:
		if existing, ok := schema.Properties[segment.Name()]; ok {
			return existing
		}
		if schema.Properties == nil {
			schema.Properties = make(map[string]*jsonschema.Schema)
		}
		newSchema := new(jsonschema.Schema)
		schema.Properties[segment.Name()] = newSchema
		return newSchema
	case jsonpath.SegmentRoot:
		// Do nothing, schema is already initialized.
	case jsonpath.SegmentIndex:
		idx := int(segment.Index())
		if idx < len(schema.PrefixItems) && schema.PrefixItems[idx] != nil {
			return schema.PrefixItems[idx]
		}
		for len(schema.PrefixItems) <= idx {
			schema.PrefixItems = append(schema.PrefixItems, new(jsonschema.Schema))
		}
		if schema.PrefixItems[idx] == nil {
			schema.PrefixItems[idx] = new(jsonschema.Schema)
		}
		return schema.PrefixItems[idx]
	case jsonpath.SegmentIndexWildcard, jsonpath.SegmentUnknownIndex:
		if schema.Items != nil {
			return schema.Items
		}
		newSchema := new(jsonschema.Schema)
		schema.Items = newSchema
		return newSchema
	case jsonpath.SegmentValueWildcard:
		if schema.AdditionalProperties != nil {
			return schema.AdditionalProperties
		}
		newSchema := new(jsonschema.Schema)
		schema.AdditionalProperties = newSchema
		return newSchema
	case jsonpath.SegmentKeyWildcard:
		if schema.PropertyNames != nil {
			return schema.PropertyNames
		}
		newSchema := new(jsonschema.Schema)
		schema.PropertyNames = newSchema
		return newSchema
	}
	return schema
}

func jsonSchemaTypeFromTypeInfo(info TypeInfo) (jsonschema.Type, error) {
	switch info.reflectKind {
	case reflect.Invalid, reflect.Interface, reflect.Pointer:
		return "", nil
	case reflect.Bool:
		return jsonschema.TypeBoolean, nil
	case reflect.String:
		return jsonschema.TypeString, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return jsonschema.TypeInteger, nil
	case reflect.Float32, reflect.Float64:
		return jsonschema.TypeNumber, nil
	case reflect.Struct, reflect.Map:
		return jsonschema.TypeObject, nil
	case reflect.Array, reflect.Slice:
		return jsonschema.TypeArray, nil
	default:
		return "", fmt.Errorf("unsupported Go kind %q", info.reflectKind)
	}
}
