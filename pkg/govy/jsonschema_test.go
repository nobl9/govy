package govy_test

import (
	"encoding/json"
	"fmt"
	"math"
	"regexp"
	"strings"
	"testing"
	"unsafe"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonschematest"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonpath"
	"github.com/nobl9/govy/pkg/jsonschema"
	"github.com/nobl9/govy/pkg/rules"
)

func TestValidatorPlan_JSONSchema(t *testing.T) {
	t.Parallel()

	validator := newPodValidator()
	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))

	valid := Pod{
		APIVersion: "v1", Kind: "Pod",
		Metadata: PodMetadata{
			Name: "pod", Namespace: "default", Labels: Labels{"app": "web"}, Annotations: Annotations{"owner": "team"},
		},
		Spec: PodSpec{
			DNSPolicy: "Default", Containers: []Container{{Name: "web", Image: "image", Env: []EnvVar{}}},
		},
	}
	invalidKind, invalidLabel, invalidContainer := valid, valid, valid
	invalidKind.Kind = "Deployment"
	invalidLabel.Metadata.Labels = Labels{"INVALID": "web"}
	invalidContainer.Spec.Containers = []Container{{Name: "INVALID", Image: "image", Env: []EnvVar{}}}
	duplicateContainers, matchingAnnotation, emptyPolicy, nilLabels := valid, valid, valid, valid
	duplicateContainers.Spec.Containers = []Container{valid.Spec.Containers[0], valid.Spec.Containers[0]}
	matchingAnnotation.Metadata.Annotations = Annotations{"owner": "owner"}
	emptyPolicy.Spec.DNSPolicy = ""
	nilLabels.Metadata.Labels = nil
	cases := []jsonschematest.Case[Pod]{
		{Name: "valid nested document", Input: valid, Valid: true},
		{Name: "root property rule", Input: invalidKind},
		{Name: "map key rule", Input: invalidLabel},
		{Name: "included array item rule", Input: invalidContainer},
		{
			Name:                 "custom uniqueness rule omitted",
			Input:                duplicateContainers,
			JSONSchemaDifference: "SliceUnique compares selected container names and has no JSON Schema builder.",
		},
		{
			Name:                 "custom map item rule omitted",
			Input:                matchingAnnotation,
			JSONSchemaDifference: "The custom rule relating an annotation key to its value has no JSON Schema builder.",
		},
		{
			Name:                 "empty optional enum",
			Input:                emptyPolicy,
			Valid:                true,
			JSONSchemaDifference: "OmitEmpty skips the empty Go string, but a present JSON string still must match the enum.",
		},
		{
			Name:                 "nil map serializes as null",
			Input:                nilLabels,
			Valid:                true,
			JSONSchemaDifference: "JSON Schema emits only the object type, while a nil Go map serializes as null.",
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
	}
	jsonschematest.Assert(t, schema, "test_data/expected_pod_json_schema.json", cases)
}

func TestJSONSchema_BuilderContext(t *testing.T) {
	t.Parallel()

	type document struct {
		Custom           string `json:"custom"`
		PropertyRequired string `json:"propertyRequired,omitempty"`
		RuleRequired     bool   `json:"ruleRequired"`
	}
	rootRule := govy.NewRule(func(document) error { return nil }).
		WithDescription("customize root schema").
		WithJSONSchema(func(ctx govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			if !ctx.Path.Equal(jsonpath.NewRoot()) {
				return nil, fmt.Errorf("expected root builder path, got %q", ctx.Path)
			}
			if ctx.Type != jsonschema.TypeObject {
				return nil, fmt.Errorf("expected object type, got %q", ctx.Type)
			}
			return &jsonschema.Schema{Title: "Builder Context"}, nil
		})
	customRule := govy.NewRule(func(v string) error {
		if !strings.Contains(v, "custom") {
			return fmt.Errorf("value must contain custom")
		}
		return nil
	}).
		WithDescription("customize property schema").
		WithJSONSchema(func(ctx govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			if !ctx.Path.Equal(jsonpath.Parse("$.custom")) {
				return nil, fmt.Errorf("expected custom property path, got %q", ctx.Path)
			}
			if ctx.Type != jsonschema.TypeString {
				return nil, fmt.Errorf("expected string type, got %q", ctx.Type)
			}
			return &jsonschema.Schema{Pattern: "custom"}, nil
		})
	validator := govy.New(
		govy.For(govy.GetSelf[document]()).Rules(rootRule),
		govy.For(func(v document) string { return v.Custom }).
			WithName("custom").
			Rules(customRule),
		govy.For(func(v document) string { return v.PropertyRequired }).
			WithName("propertyRequired").
			Required(),
		govy.For(func(v document) bool { return v.RuleRequired }).
			WithName("ruleRequired").
			Required().
			Rules(rules.Required[bool]()),
	)

	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))

	cases := []jsonschematest.Case[document]{
		{
			Name:  "custom builder and required properties",
			Input: document{Custom: "custom value", PropertyRequired: "present", RuleRequired: true},
			Valid: true,
		},
		{
			Name:  "custom builder rejects unmatched value",
			Input: document{Custom: "other", PropertyRequired: "present", RuleRequired: true},
		},
		{
			Name:  "missing required property",
			Input: document{Custom: "custom", RuleRequired: true},
		},
		{
			Name:                 "present false boolean",
			Input:                document{Custom: "custom", PropertyRequired: "present"},
			JSONSchemaDifference: "JSON Schema required checks presence, not Go zero values.",
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
	}
	jsonschematest.Assert(t, schema, "test_data/expected_builder_context_json_schema.json", cases)
}

func TestJSONSchema_When(t *testing.T) {
	t.Parallel()

	type item struct {
		Kind  string `json:"kind,omitempty"`
		Value string `json:"value,omitempty"`
	}
	type document struct {
		Enabled bool   `json:"enabled"`
		Ignored string `json:"ignored,omitempty"`
		Items   []item `json:"items"`
	}
	enabled := any(true)
	kind := any("required")
	itemValidator := govy.New(
		govy.For(func(v item) string { return v.Value }).
			WithName("value").
			Required().
			When(
				func(v item) bool { return v.Kind == "required" },
				govy.WhenJSONSchema(func(ctx govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
					if !ctx.Path.Equal(jsonpath.Parse("$.items[*]")) {
						return nil, fmt.Errorf("expected item condition path, got %q", ctx.Path)
					}
					return &jsonschema.Schema{
						Properties: map[string]*jsonschema.Schema{
							"kind": {Const: &kind},
						},
						Required: []string{"kind"},
					}, nil
				}),
			),
	)
	validator := govy.New(
		govy.For(func(v document) string { return v.Ignored }).
			WithName("ignored").
			Required().
			When(func(document) bool { return true }),
		govy.ForSlice(func(v document) []item { return v.Items }).
			WithName("items").
			IncludeForEach(itemValidator),
	).
		WithName("Conditions").
		When(
			func(v document) bool { return v.Enabled },
			govy.WhenJSONSchema(func(ctx govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
				if !ctx.Path.Equal(jsonpath.NewRoot()) {
					return nil, fmt.Errorf("expected root condition path, got %q", ctx.Path)
				}
				if ctx.Type != jsonschema.TypeObject {
					return nil, fmt.Errorf("expected object type, got %q", ctx.Type)
				}
				return &jsonschema.Schema{
					Properties: map[string]*jsonschema.Schema{
						"enabled": {Const: &enabled},
					},
					Required: []string{"enabled"},
				}, nil
			}),
		)

	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))

	cases := []jsonschematest.Case[document]{
		{
			Name:  "disabled validator skips required item value",
			Input: document{Items: []item{{Kind: "required"}}},
			Valid: true,
		},
		{
			Name:  "enabled required item has value",
			Input: document{Enabled: true, Ignored: "present", Items: []item{{Kind: "required", Value: "present"}}},
			Valid: true,
		},
		{
			Name:  "enabled optional item omits value",
			Input: document{Enabled: true, Ignored: "present", Items: []item{{Kind: "optional"}, {}}},
			Valid: true,
		},
		{
			Name:  "required item omits value",
			Input: document{Enabled: true, Ignored: "present", Items: []item{{Kind: "required"}}},
		},
		{
			Name:  "condition is evaluated for each item",
			Input: document{Enabled: true, Ignored: "present", Items: []item{{Kind: "optional"}, {Kind: "required"}}},
		},
		{
			Name:                 "unmapped condition omits required rule",
			Input:                document{Enabled: true, Items: []item{}},
			JSONSchemaDifference: "Rules guarded by a condition without WhenJSONSchema are omitted.",
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
	}
	jsonschematest.Assert(t, schema, "test_data/expected_when_json_schema.json", cases)
}

func TestJSONSchema_RuleComposition(t *testing.T) {
	t.Parallel()

	type document struct {
		Text     string  `json:"text"`
		Count    int     `json:"count"`
		Optional *string `json:"optional,omitempty"`
		Required bool    `json:"required,omitempty"`
	}
	validator := govy.New(
		govy.For(func(v document) string { return v.Text }).WithName("text").Rules(
			rules.StringStartsWith("pre"),
			rules.StringEndsWith("post"),
			rules.StringMatchRegexp(regexp.MustCompile("^[a-z]+$")),
		),
		govy.For(func(v document) int { return v.Count }).WithName("count").Rules(
			rules.GTE(2), rules.GTE(5), rules.LTE(10), rules.LTE(8),
		),
		govy.For(func(v document) *string { return v.Optional }).WithName("optional").
			Rules(govy.RuleToPointer(rules.OneOf("first", "second"))),
		govy.For(func(v document) bool { return v.Required }).WithName("required").Rules(rules.Required[bool]()),
	)
	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))

	allowed, disallowed := "first", "third"
	cases := []jsonschematest.Case[document]{
		{
			Name:  "all rules match",
			Input: document{Text: "prepost", Count: 5, Optional: &allowed, Required: true},
			Valid: true,
		},
		{
			Name:  "optional pointer absent",
			Input: document{Text: "prepost", Count: 8, Required: true},
			Valid: true,
		},
		{
			Name:  "first pattern fails",
			Input: document{Text: "otherpost", Count: 5, Required: true},
		},
		{
			Name:  "second pattern fails",
			Input: document{Text: "preother", Count: 5, Required: true},
		},
		{
			Name:  "third pattern fails",
			Input: document{Text: "pre-post", Count: 5, Required: true},
		},
		{
			Name:  "stricter minimum fails",
			Input: document{Text: "prepost", Count: 4, Required: true},
		},
		{
			Name:  "stricter maximum fails",
			Input: document{Text: "prepost", Count: 9, Required: true},
		},
		{
			Name:  "pointer rule fails",
			Input: document{Text: "prepost", Count: 5, Optional: &disallowed, Required: true},
		},
		{
			Name:  "required rule lifts to parent",
			Input: document{Text: "prepost", Count: 5},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
	}
	jsonschematest.Assert(t, schema, "test_data/expected_rule_composition_json_schema.json", cases)
}

func TestJSONSchema_ZeroLengthLimits(t *testing.T) {
	t.Parallel()

	type document struct {
		String string            `json:"string"`
		Slice  []string          `json:"slice"`
		Map    map[string]string `json:"map"`
	}
	validator := govy.New(
		govy.For(func(v document) string { return v.String }).
			WithName("string").
			Rules(rules.StringLength(0, 0), rules.StringMinLength(0), rules.StringMaxLength(0)),
		govy.For(func(v document) []string { return v.Slice }).
			WithName("slice").
			Rules(rules.SliceLength[[]string](0, 0), rules.SliceMinLength[[]string](0), rules.SliceMaxLength[[]string](0)),
		govy.For(func(v document) map[string]string { return v.Map }).
			WithName("map").
			Rules(
				rules.MapLength[map[string]string](0, 0),
				rules.MapMinLength[map[string]string](0),
				rules.MapMaxLength[map[string]string](0),
			),
	)

	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))

	jsonschematest.Assert(
		t,
		schema,
		"test_data/expected_zero_length_limits_json_schema.json",
		[]jsonschematest.Case[document]{
			{
				Name:  "empty values",
				Input: document{Slice: []string{}, Map: map[string]string{}},
				Valid: true,
			},
			{
				Name:  "nonempty string",
				Input: document{String: "a", Slice: []string{}, Map: map[string]string{}},
			},
			{
				Name:  "nonempty slice",
				Input: document{Slice: []string{"a"}, Map: map[string]string{}},
			},
			{
				Name:  "nonempty map",
				Input: document{Slice: []string{}, Map: map[string]string{"a": "b"}},
			},
		},
	)
}

func TestJSONSchema_RuleBuilderError(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		validator     govy.Validator[struct{}]
		expectedError string
	}{
		"value without a JSON representation": {
			validator: govy.New(
				govy.For(func(struct{}) any { return nil }).
					WithName("value").
					Rules(rules.EQ[any](make(chan int))),
			),
			expectedError: `failed to build JSON Schema for "$.value" property and "equal_to" rule: ` +
				`marshal value: json: unsupported type: chan int`,
		},
		"non-finite number": {
			validator: govy.New(
				govy.For(func(struct{}) float64 { return 0 }).
					WithName("value").
					Rules(rules.GT(math.Inf(1))),
			),
			expectedError: `failed to build JSON Schema for "$.value" property and "greater_than" rule: ` +
				`marshal value: json: unsupported value: +Inf`,
		},
		"unsupported regular expression": {
			validator: govy.New(
				govy.For(func(struct{}) string { return "" }).
					WithName("value").
					Rules(rules.StringMatchRegexp(regexp.MustCompile(`(?m)^value$`))),
			),
			expectedError: `failed to build JSON Schema for "$.value" property and "string_match_regexp" rule: ` +
				`translate Go regular expression "(?m)^value$": ` +
				`unsupported Go regular expression construct: multiline beginning anchor`,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			for name, options := range map[string][]govy.JSONSchemaOption{
				"default":               nil,
				"include omitted rules": {govy.JSONSchemaIncludeOmittedRules()},
			} {
				t.Run(name, func(t *testing.T) {
					_, err := govy.JSONSchema(tc.validator, options...)
					assert.EqualError(t, err, tc.expectedError)
				})
			}
		})
	}
}

func TestJSONSchema_FormatComposition(t *testing.T) {
	t.Parallel()

	validator := govy.New(govy.For(govy.GetSelf[string]()).Rules(rules.StringIPv4(), rules.StringIPv6()))
	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))

	cases := []jsonschematest.Case[string]{
		{Name: "IPv4 fails IPv6 constraint", Input: "127.0.0.1"},
		{Name: "IPv6 fails IPv4 constraint", Input: "::1"},
		{Name: "neither format", Input: "not an IP address"},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
	}
	jsonschematest.Assert(t, schema, "test_data/expected_format_composition_json_schema.json", cases)
}

func TestJSONSchema_UnsupportedType(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		run           func() error
		expectedError string
	}{
		"channel": {
			run:           runJSONSchema[chan struct{}],
			expectedError: `failed to generate JSON Schema type info for validator: unsupported Go kind "chan"`,
		},
		"complex64": {
			run:           runJSONSchema[complex64],
			expectedError: `failed to generate JSON Schema type info for validator: unsupported Go kind "complex64"`,
		},
		"complex128": {
			run:           runJSONSchema[complex128],
			expectedError: `failed to generate JSON Schema type info for validator: unsupported Go kind "complex128"`,
		},
		"function": {
			run:           runJSONSchema[func()],
			expectedError: `failed to generate JSON Schema type info for validator: unsupported Go kind "func"`,
		},
		"unsafe pointer": {
			run:           runJSONSchema[unsafe.Pointer],
			expectedError: `failed to generate JSON Schema type info for validator: unsupported Go kind "unsafe.Pointer"`,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.EqualError(t, tc.run(), tc.expectedError)
		})
	}
}

func TestJSONSchema_MapItemUsesValueType(t *testing.T) {
	t.Parallel()

	type annotations map[string]string
	type document struct {
		Annotations annotations `json:"annotations"`
	}
	validator := govy.New(
		govy.ForMap(func(v document) annotations { return v.Annotations }).
			WithName("annotations").
			RulesForItems(
				govy.NewRule(func(v govy.MapItem[string, string]) error {
					if v.Value == "" {
						return fmt.Errorf("value must not be empty")
					}
					return nil
				}).WithDescription("map value must not be empty").
					WithJSONSchema(func(ctx govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
						if !ctx.Path.Equal(jsonpath.Parse("$.annotations.*")) || ctx.Type != jsonschema.TypeString {
							return nil, fmt.Errorf("builder must describe the map value")
						}
						minimum := uint64(1)
						return &jsonschema.Schema{MinLength: &minimum}, nil
					}),
			),
	)

	schema, err := govy.JSONSchema(validator, govy.JSONSchemaIncludeOmittedRules())
	assert.Require(t, assert.NoError(t, err))

	cases := []jsonschematest.Case[document]{
		{Name: "empty map", Input: document{Annotations: annotations{}}, Valid: true},
		{
			Name:  "string values",
			Input: document{Annotations: annotations{"first": "one", "second": "two"}},
			Valid: true,
		},
		{Name: "empty value", Input: document{Annotations: annotations{"first": ""}}},
		{Name: "one empty value", Input: document{Annotations: annotations{"first": "one", "second": ""}}},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
	}
	jsonschematest.Assert(t, schema, "test_data/expected_map_item_json_schema.json", cases)
}

func TestJSONSchema_FixedArrayIndexes(t *testing.T) {
	t.Parallel()

	type document struct{}
	validator := govy.New(
		govy.For(func(document) string { return "" }).
			WithPath(jsonpath.New().Name("tuple").Index(2).Name("a")),
		govy.For(func(document) string { return "" }).
			WithPath(jsonpath.New().Name("tuple").Index(2).Name("b")),
		govy.For(func(document) string { return "" }).
			WithPath(jsonpath.New().Name("tuple").Index(10).Name("c")),
	)

	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))

	jsonschematest.Assert(
		t,
		schema,
		"test_data/expected_fixed_array_indexes_json_schema.json",
		[]jsonschematest.Case[json.RawMessage]{
			{Name: "empty array", Input: json.RawMessage(`{"tuple":[]}`), Valid: true},
			{
				Name:  "both fixed indexes",
				Input: json.RawMessage(`{"tuple":[0,1,{"a":"first","b":"second"},3,4,5,6,7,8,9,{"c":"third"}]}`),
				Valid: true,
			},
			{
				Name:  "unconstrained positions and tail",
				Input: json.RawMessage(`{"tuple":[{"a":1},false,{"a":"first"},3,4,5,6,7,8,9,{}, {"c":1}]}`),
				Valid: true,
			},
			{Name: "first property at index two", Input: json.RawMessage(`{"tuple":[0,1,{"a":2}]}`)},
			{Name: "second property at index two", Input: json.RawMessage(`{"tuple":[0,1,{"b":2}]}`)},
			{
				Name:  "property at index ten",
				Input: json.RawMessage(`{"tuple":[0,1,{},3,4,5,6,7,8,9,{"c":2}]}`),
			},
		},
	)
}

func TestJSONSchema_ArrayWildcardWithFixedIndex(t *testing.T) {
	t.Parallel()

	type document struct{}
	validator := govy.New(
		govy.For(func(document) string { return "" }).
			WithPath(jsonpath.New().Name("values").IndexWildcard()),
		govy.For(func(document) int { return 0 }).
			WithPath(jsonpath.New().Name("values").Index(0)),
	)

	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))

	jsonschematest.Assert(
		t,
		schema,
		"test_data/expected_array_wildcard_with_fixed_index_json_schema.json",
		[]jsonschematest.Case[json.RawMessage]{
			{Name: "empty array", Input: json.RawMessage(`{"values":[]}`), Valid: true},
			{Name: "fixed index still requires integer", Input: json.RawMessage(`{"values":["text"]}`)},
			{Name: "wildcard also applies to fixed index", Input: json.RawMessage(`{"values":[1]}`)},
			{
				Name:  "valid tail does not bypass wildcard at fixed index",
				Input: json.RawMessage(`{"values":[1,"text"]}`),
			},
		},
	)
}

func TestJSONSchema_ValueWildcardWithNamedProperty(t *testing.T) {
	t.Parallel()

	type document struct{}
	validator := govy.New(
		govy.For(func(document) string { return "" }).
			WithPath(jsonpath.New().Name("values").ValueWildcard()),
		govy.For(func(document) int { return 0 }).
			WithPath(jsonpath.New().Name("values").Name("foo")),
	)

	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))

	jsonschematest.Assert(
		t,
		schema,
		"test_data/expected_value_wildcard_with_named_property_json_schema.json",
		[]jsonschematest.Case[json.RawMessage]{
			{Name: "empty object", Input: json.RawMessage(`{"values":{}}`), Valid: true},
			{Name: "unnamed property", Input: json.RawMessage(`{"values":{"bar":"text"}}`), Valid: true},
			{Name: "named property still requires integer", Input: json.RawMessage(`{"values":{"foo":"text"}}`)},
			{Name: "wildcard also applies to named property", Input: json.RawMessage(`{"values":{"foo":1}}`)},
			{Name: "wildcard applies to unnamed property", Input: json.RawMessage(`{"values":{"bar":1}}`)},
		},
	)
}

func TestJSONSchema_ArrayIndexOutsideIntRange(t *testing.T) {
	t.Parallel()

	type document struct{}
	index := uint(math.MaxUint)
	validator := govy.New(
		govy.For(func(document) string { return "" }).
			WithPath(jsonpath.New().Name("tuple").Index(index)),
	)

	_, err := govy.JSONSchema(validator)
	assert.EqualError(t, err, fmt.Sprintf(
		`failed to generate JSON Schema for %q property: array index %d exceeds maximum supported index %d`,
		jsonpath.NewRoot().Name("tuple").Index(index),
		index,
		math.MaxInt-1,
	))
}

func runJSONSchema[T any]() error {
	_, err := govy.JSONSchema(govy.New[T]())
	return err
}
