package govy

import (
	"fmt"
	"reflect"

	"github.com/nobl9/govy/pkg/jsonschema"
)

type JSONSchemaBuilder func(schema jsonschema.Schema) (jsonschema.Schema, error)

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
		return nil, fmt.Errorf("failed to generate JSON Schema: %w", err)
	}

	document := jsonschema.Document{
		Title: plan.Name,
		Type:  schemaType,
	}
	return &document, nil
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
