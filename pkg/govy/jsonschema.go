package govy

import (
	"fmt"
	"math"
	"reflect"

	"github.com/nobl9/govy/pkg/jsonpath"
	"github.com/nobl9/govy/pkg/jsonschema"
)

type JSONSchemaBuilder func(schema *jsonschema.Schema) error

const maxJSONSchemaPrefixItemsIndex = math.MaxInt - 1

// JSONSchema creates a JSON Schema document for the provided [Validator].
// It uses exclusively [Draft 2020-12] version.
// It returns an error for Go kinds without a default JSON representation.
//
// [Draft 2020-12]: https://json-schema.org/draft/2020-12/schema
func JSONSchema[T any](v Validator[T]) (*jsonschema.Document, error) {
	plan, err := Plan(v, PlanStrictMode(), planRecordJSONSchema())
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
		schemaSelector := schema
		var segment jsonpath.Segment
		for _, segment = range prop.Path.Segments() {
			if segment.Kind() == jsonpath.SegmentIndex && segment.Index() > maxJSONSchemaPrefixItemsIndex {
				return nil, fmt.Errorf(
					"failed to generate JSON Schema for %q property: array index %d exceeds maximum supported index %d",
					prop.Path,
					segment.Index(),
					maxJSONSchemaPrefixItemsIndex,
				)
			}
			schemaSelector = getJSONSchemaForSegment(schemaSelector, segment)
		}
		schemaSelector.Type, err = jsonSchemaTypeFromTypeInfo(prop.TypeInfo)
		if err != nil {
			return nil, fmt.Errorf("failed to generate JSON Schema type info for %q property: %w", prop.Path, err)
		}
		for _, rule := range prop.Rules {
			for _, builder := range rule.jsonSchemaBuilders {
				if err = builder(schemaSelector); err != nil {
					return nil, fmt.Errorf(
						"failed to build JSON Schema for %q property and %q rule: %w",
						prop.Path,
						rule.ErrorCode,
						err,
					)
				}
			}
		}
		switch segment.Kind() {
		case jsonpath.SegmentName:
		case jsonpath.SegmentRoot:
		case jsonpath.SegmentUnknownIndex, jsonpath.SegmentIndex, jsonpath.SegmentValueWildcard:
		case jsonpath.SegmentKeyWildcard:
		}
	}
	normalizeJSONSchemaWildcardApplicators(schema)

	document := jsonschema.Document(*schema)
	return &document, nil
}

// normalizeJSONSchemaWildcardApplicators separates wildcard and specific applicators
// so both apply to matching children.
func normalizeJSONSchemaWildcardApplicators(schema *jsonschema.Schema) {
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

	for _, child := range schema.AllOf {
		normalizeJSONSchemaWildcardApplicators(child)
	}
	for _, child := range schema.AnyOf {
		normalizeJSONSchemaWildcardApplicators(child)
	}
	for _, child := range schema.OneOf {
		normalizeJSONSchemaWildcardApplicators(child)
	}
	normalizeJSONSchemaWildcardApplicators(schema.Not)
	for _, child := range schema.Properties {
		normalizeJSONSchemaWildcardApplicators(child)
	}
	normalizeJSONSchemaWildcardApplicators(schema.AdditionalProperties)
	normalizeJSONSchemaWildcardApplicators(schema.PropertyNames)
	for _, child := range schema.PrefixItems {
		normalizeJSONSchemaWildcardApplicators(child)
	}
	normalizeJSONSchemaWildcardApplicators(schema.Items)
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
		// Do nothing, schema is already initalized.
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
