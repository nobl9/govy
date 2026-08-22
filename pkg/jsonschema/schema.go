package jsonschema

import (
	"encoding/json"
)

const draft2020 = "https://json-schema.org/draft/2020-12/schema"

// Type is a JSON Schema primitive type name.
type Type string

// JSON Schema primitive types.
const (
	TypeNull    Type = "null"
	TypeBoolean Type = "boolean"
	TypeObject  Type = "object"
	TypeArray   Type = "array"
	TypeNumber  Type = "number"
	TypeString  Type = "string"
	TypeInteger Type = "integer"
)

// Schema represents the subset of a JSON Schema Draft 2020-12 schema object
// required by Govy. Convert a root Schema to [Document] before marshaling it.
type Schema struct {
	// Title describes the validated value.
	Title string `json:"title,omitempty"`

	// Type sets the type keyword to one primitive type.
	Type Type `json:"type,omitempty"`

	// Enum limits a value to the listed constants.
	Enum []any `json:"enum,omitempty"`
	// Const limits a value to one constant. A pointer distinguishes an absent
	// keyword from a constant whose value is JSON null.
	Const *any `json:"const,omitempty"`

	// AllOf requires every child schema to match.
	AllOf []*Schema `json:"allOf,omitempty"`
	// AnyOf requires at least one child schema to match.
	AnyOf []*Schema `json:"anyOf,omitempty"`
	// OneOf requires exactly one child schema to match.
	OneOf []*Schema `json:"oneOf,omitempty"`
	// Not requires the child schema not to match.
	Not *Schema `json:"not,omitempty"`
	// If determines whether [Schema.Then] applies.
	If *Schema `json:"if,omitempty"`
	// Then applies when [Schema.If] matches.
	Then *Schema `json:"then,omitempty"`

	// Properties maps object property names to their schemas.
	Properties map[string]*Schema `json:"properties,omitempty"`
	// Required lists object properties that must be present.
	Required []string `json:"required,omitempty"`
	// AdditionalProperties constrains object properties not named in Properties.
	AdditionalProperties *Schema `json:"additionalProperties,omitempty"`
	// PropertyNames constrains every property name in an object.
	PropertyNames *Schema `json:"propertyNames,omitempty"`
	// MinProperties is the inclusive lower bound for an object's property count.
	MinProperties *uint64 `json:"minProperties,omitempty"`
	// MaxProperties is the inclusive upper bound for an object's property count.
	MaxProperties *uint64 `json:"maxProperties,omitempty"`

	// PrefixItems constrains array items at the corresponding positions.
	PrefixItems []*Schema `json:"prefixItems,omitempty"`
	// Items constrains array items not covered by [Schema.PrefixItems].
	Items *Schema `json:"items,omitempty"`
	// MinItems is the inclusive lower bound for an array's length.
	MinItems *uint64 `json:"minItems,omitempty"`
	// MaxItems is the inclusive upper bound for an array's length.
	MaxItems *uint64 `json:"maxItems,omitempty"`

	// MinLength is the inclusive lower bound for a string's length.
	MinLength *uint64 `json:"minLength,omitempty"`
	// MaxLength is the inclusive upper bound for a string's length.
	MaxLength *uint64 `json:"maxLength,omitempty"`
	// Pattern is an ECMA-262 regular expression that a string must match.
	Pattern string `json:"pattern,omitempty"`

	// MultipleOf requires a number to be a multiple of this value.
	MultipleOf json.Number `json:"multipleOf,omitempty"`
	// Minimum is the inclusive lower bound for a number.
	Minimum json.Number `json:"minimum,omitempty"`
	// Maximum is the inclusive upper bound for a number.
	Maximum json.Number `json:"maximum,omitempty"`
	// ExclusiveMinimum is the exclusive lower bound for a number.
	ExclusiveMinimum json.Number `json:"exclusiveMinimum,omitempty"`
	// ExclusiveMaximum is the exclusive upper bound for a number.
	ExclusiveMaximum json.Number `json:"exclusiveMaximum,omitempty"`
}

// Document represents a top-level JSON Schema document. Its JSON
// representation always declares the [draft2020] schema version.
type Document Schema

// MarshalJSON declares the Draft 2020-12 dialect and encodes the root schema.
func (d Document) MarshalJSON() ([]byte, error) {
	type documentWithoutMethods Document
	return json.Marshal(struct {
		Version string `json:"$schema"`
		documentWithoutMethods
	}{
		Version:                draft2020,
		documentWithoutMethods: documentWithoutMethods(d),
	})
}
