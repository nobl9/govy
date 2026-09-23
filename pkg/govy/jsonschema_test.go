package govy_test

import (
	"encoding/json"
	"errors"
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

func TestJSONSchema_UnlabelledBuilder(t *testing.T) {
	t.Parallel()

	rule := govy.NewRule(func(value int) error {
		if value < 3 {
			return fmt.Errorf("value must be at least 3")
		}
		return nil
	}).WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
		return &jsonschema.Schema{Minimum: "3"}, nil
	})
	validator := govy.New(govy.For(govy.GetSelf[int]()).Rules(rule))
	schema, err := govy.JSONSchema(validator, govy.JSONSchemaIncludeOmittedRules())
	assert.Require(t, assert.NoError(t, err))
	cases := []jsonschematest.Case[int]{
		{Name: "below minimum", Input: 2},
		{Name: "at minimum", Input: 3, Valid: true},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
	}
	jsonschematest.Assert(t, schema, "test_data/expected_unlabelled_builder_json_schema.json", cases)
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

func TestJSONSchema_WhenNestedScopes(t *testing.T) {
	t.Parallel()

	type child struct {
		Enabled bool   `json:"enabled"`
		Kind    string `json:"kind,omitempty"`
		Value   string `json:"value,omitempty"`
	}
	type document struct {
		Enabled bool  `json:"enabled"`
		Strict  bool  `json:"strict"`
		Child   child `json:"child"`
	}
	enabled, requiredKind := any(true), any("required")
	whenEnabled := govy.WhenJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
		return &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{"enabled": {Const: &enabled}},
			Required:   []string{"enabled"},
		}, nil
	})
	childValidator := govy.New(
		govy.For(func(v child) string { return v.Value }).
			WithName("value").
			Required().
			When(func(v child) bool { return v.Kind == "required" },
				govy.WhenJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
					return &jsonschema.Schema{
						Properties: map[string]*jsonschema.Schema{"kind": {Const: &requiredKind}},
						Required:   []string{"kind"},
					}, nil
				})),
	).When(func(v child) bool { return v.Enabled }, whenEnabled)
	validator := govy.New(
		govy.For(func(v document) child { return v.Child }).
			WithName("child").
			Include(childValidator).
			When(func(v document) bool { return v.Strict },
				govy.WhenJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
					return &jsonschema.Schema{
						Properties: map[string]*jsonschema.Schema{"strict": {Const: &enabled}},
						Required:   []string{"strict"},
					}, nil
				})),
	).When(func(v document) bool { return v.Enabled }, whenEnabled)

	valid := document{Enabled: true, Strict: true, Child: child{Enabled: true, Kind: "required", Value: "present"}}
	missing := valid
	missing.Child.Value = ""
	rootDisabled, propertyDisabled, childDisabled, optionalKind := missing, missing, missing, missing
	rootDisabled.Enabled = false
	propertyDisabled.Strict = false
	childDisabled.Child.Enabled = false
	optionalKind.Child.Kind = "optional"
	cases := []jsonschematest.Case[document]{
		{Name: "all conditions match with value", Input: valid, Valid: true},
		{Name: "all conditions match without value", Input: missing},
		{Name: "root condition does not match", Input: rootDisabled, Valid: true},
		{Name: "parent property condition does not match", Input: propertyDisabled, Valid: true},
		{Name: "included validator condition does not match", Input: childDisabled, Valid: true},
		{Name: "child property condition does not match", Input: optionalKind, Valid: true},
	}
	for _, tc := range cases {
		t.Run("govy/"+tc.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
		})
	}
	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "test_data/expected_when_nested_scopes_json_schema.json", cases)
}

func TestJSONSchema_RepeatedConditionTypes(t *testing.T) {
	t.Parallel()

	type document struct {
		Name string
	}
	tests := []struct {
		name  string
		path  string
		build func(govy.Validator[string]) (*jsonschema.Document, error)
	}{
		{name: "root", path: "$", build: func(v govy.Validator[string]) (*jsonschema.Document, error) {
			return govy.JSONSchema(v)
		}},
		{name: "named property", path: "$.name", build: func(v govy.Validator[string]) (*jsonschema.Document, error) {
			return govy.JSONSchema(govy.New(govy.For(func(d document) string { return d.Name }).
				WithName("name").Include(v)))
		}},
		{name: "array item", path: "$[*]", build: func(v govy.Validator[string]) (*jsonschema.Document, error) {
			return govy.JSONSchema(govy.New(govy.ForSlice(govy.GetSelf[[]string]()).IncludeForEach(v)))
		}},
		{name: "map value", path: "$.*", build: func(v govy.Validator[string]) (*jsonschema.Document, error) {
			return govy.JSONSchema(govy.New(govy.ForMap(govy.GetSelf[map[string]string]()).IncludeForValues(v)))
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var contexts []govy.JSONSchemaBuilderContext
			condition := govy.WhenJSONSchema(func(ctx govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
				contexts = append(contexts, ctx)
				return &jsonschema.Schema{}, nil
			})
			v := govy.New(govy.For(govy.GetSelf[string]()).Rules(rules.StringMinLength(1))).
				When(func(string) bool { return true }, condition).
				When(func(string) bool { return true }, condition)
			_, err := tt.build(v)
			assert.Require(t, assert.NoError(t, err))
			assert.Require(t, assert.Len(t, contexts, 2))
			for _, ctx := range contexts {
				assert.Equal(t, jsonschema.TypeString, ctx.Type)
				assert.Equal(t, tt.path, ctx.Path.String())
			}
		})
	}
}

func TestJSONSchema_WhenIndependentBranches(t *testing.T) {
	t.Parallel()

	type document struct {
		First  bool   `json:"first"`
		Second bool   `json:"second"`
		Value  string `json:"value,omitempty"`
	}
	enabled := any(true)
	value := govy.For(func(v document) string { return v.Value }).WithName("value").Required()
	validator := govy.New(
		value.When(func(v document) bool { return v.First },
			govy.WhenDescription("branch is enabled"),
			govy.WhenJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
				return &jsonschema.Schema{
					Properties: map[string]*jsonschema.Schema{"first": {Const: &enabled}},
					Required:   []string{"first"},
				}, nil
			})),
		value.When(func(v document) bool { return v.Second },
			govy.WhenDescription("branch is enabled"),
			govy.WhenJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
				return &jsonschema.Schema{
					Properties: map[string]*jsonschema.Schema{"second": {Const: &enabled}},
					Required:   []string{"second"},
				}, nil
			})),
	)
	cases := []jsonschematest.Case[document]{
		{Name: "neither branch matches", Input: document{}, Valid: true},
		{Name: "first branch requires value", Input: document{First: true}},
		{Name: "second branch requires value", Input: document{Second: true}},
		{Name: "both branches require value", Input: document{First: true, Second: true}},
		{Name: "first branch satisfied", Input: document{First: true, Value: "present"}, Valid: true},
		{Name: "second branch satisfied", Input: document{Second: true, Value: "present"}, Valid: true},
		{Name: "both branches satisfied", Input: document{First: true, Second: true, Value: "present"}, Valid: true},
	}
	for _, tc := range cases {
		t.Run("govy/"+tc.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
		})
	}
	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "test_data/expected_when_independent_branches_json_schema.json", cases)
}

func TestJSONSchema_WhenCollections(t *testing.T) {
	t.Parallel()

	type item struct {
		Kind  string `json:"kind,omitempty"`
		Value string `json:"value,omitempty"`
	}
	type document struct {
		Enabled bool            `json:"enabled"`
		Items   []item          `json:"items"`
		Values  map[string]item `json:"values"`
	}
	enabled, requiredKind := any(true), any("required")
	whenEnabled := govy.WhenJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
		return &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{"enabled": {Const: &enabled}},
			Required:   []string{"enabled"},
		}, nil
	})
	itemValidator := govy.New(
		govy.For(func(v item) string { return v.Value }).WithName("value").Required(),
	).When(func(v item) bool { return v.Kind == "required" },
		govy.WhenJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return &jsonschema.Schema{
				Properties: map[string]*jsonschema.Schema{"kind": {Const: &requiredKind}},
				Required:   []string{"kind"},
			}, nil
		}))
	validator := govy.New(
		govy.ForSlice(func(v document) []item { return v.Items }).
			WithName("items").
			IncludeForEach(itemValidator).
			When(func(v document) bool { return v.Enabled }, whenEnabled),
		govy.ForMap(func(v document) map[string]item { return v.Values }).
			WithName("values").
			IncludeForValues(itemValidator).
			When(func(v document) bool { return v.Enabled }, whenEnabled),
	)
	cases := []jsonschematest.Case[document]{
		{
			Name:  "empty collections",
			Input: document{Enabled: true, Items: []item{}, Values: map[string]item{}},
			Valid: true,
		},
		{
			Name: "mixed optional and required children",
			Input: document{
				Enabled: true,
				Items:   []item{{}, {Kind: "required", Value: "present"}},
				Values:  map[string]item{"optional": {}, "required": {Kind: "required", Value: "present"}},
			},
			Valid: true,
		},
		{
			Name: "condition applies to later array items",
			Input: document{
				Enabled: true,
				Items:   []item{{}, {Kind: "required"}},
				Values:  map[string]item{},
			},
		},
		{
			Name: "condition applies to each map value",
			Input: document{
				Enabled: true,
				Items:   []item{},
				Values:  map[string]item{"optional": {}, "required": {Kind: "required"}},
			},
		},
		{
			Name: "parent condition disables both collections",
			Input: document{
				Items:  []item{{Kind: "required"}},
				Values: map[string]item{"required": {Kind: "required"}},
			},
			Valid: true,
		},
	}
	for _, tc := range cases {
		t.Run("govy/"+tc.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
		})
	}
	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "test_data/expected_when_collections_json_schema.json", cases)
}

func TestJSONSchema_WhenMapKeysAndValues(t *testing.T) {
	t.Parallel()

	keyValidator := govy.New(
		govy.For(govy.GetSelf[string]()).Rules(rules.StringMinLength(6)),
	).When(func(v string) bool { return strings.HasPrefix(v, "app/") },
		govy.WhenJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return &jsonschema.Schema{Pattern: "^app/"}, nil
		}))
	valueValidator := govy.New(
		govy.For(govy.GetSelf[int]()).Rules(rules.GTE(10)),
	).When(func(v int) bool { return v > 0 },
		govy.WhenJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return &jsonschema.Schema{ExclusiveMinimum: "0"}, nil
		}))
	validator := govy.New(
		govy.ForMap(govy.GetSelf[map[string]int]()).
			IncludeForKeys(keyValidator).
			IncludeForValues(valueValidator),
	)
	cases := []jsonschematest.Case[map[string]int]{
		{Name: "empty map", Input: map[string]int{}, Valid: true},
		{Name: "both conditions match", Input: map[string]int{"app/ok": 10}, Valid: true},
		{Name: "both conditions skip", Input: map[string]int{"x": -1}, Valid: true},
		{Name: "zero skips value condition", Input: map[string]int{"app/ok": 0}, Valid: true},
		{Name: "matching key fails length", Input: map[string]int{"app/a": 10}},
		{Name: "matching value fails minimum", Input: map[string]int{"x": 1}},
		{Name: "both selected values fail", Input: map[string]int{"app/a": 1}},
		{Name: "valid sibling does not hide invalid key", Input: map[string]int{"app/ok": 10, "app/a": -1}},
		{Name: "valid sibling does not hide invalid value", Input: map[string]int{"app/ok": 10, "x": 1}},
	}
	for _, tc := range cases {
		t.Run("govy/"+tc.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
		})
	}
	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "test_data/expected_when_map_keys_and_values_json_schema.json", cases)
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

func TestJSONSchema_CustomBuilderComposition(t *testing.T) {
	t.Parallel()

	prefix := rules.StringStartsWith("pre", "alt").
		WithErrorCode("custom").
		WithDescription("custom contribution").
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return &jsonschema.Schema{AnyOf: []*jsonschema.Schema{
				{Pattern: "^pre"},
				{Pattern: "^alt"},
			}}, nil
		})
	suffix := rules.StringEndsWith("end", "stop").
		WithErrorCode("custom").
		WithDescription("custom contribution").
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return &jsonschema.Schema{AnyOf: []*jsonschema.Schema{
				{Pattern: "end$"},
				{Pattern: "stop$"},
			}}, nil
		})
	exclude := rules.StringExcludes("bad").
		WithErrorCode("custom").
		WithDescription("custom contribution").
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return &jsonschema.Schema{AllOf: []*jsonschema.Schema{
				{Not: &jsonschema.Schema{Pattern: "bad"}},
			}}, nil
		})
	noOp := govy.NewRule(func(string) error { return nil }).WithDescription("no constraint")
	validator := govy.New(
		govy.For(govy.GetSelf[string]()).Rules(
			prefix,
			govy.NewRuleSet(suffix, exclude),
			noOp.WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) { return nil, nil }),
			noOp.WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
				return &jsonschema.Schema{}, nil
			}),
		),
	)
	cases := []jsonschematest.Case[string]{
		{Name: "first prefix and suffix", Input: "preend", Valid: true},
		{Name: "second prefix and suffix", Input: "altstop", Valid: true},
		{Name: "first prefix and second suffix", Input: "prestop", Valid: true},
		{Name: "second prefix and first suffix", Input: "altend", Valid: true},
		{Name: "first anyOf still applies", Input: "otherend"},
		{Name: "second anyOf still applies", Input: "preother"},
		{Name: "explicit allOf still applies", Input: "prebadend"},
		{Name: "neither anyOf matches", Input: "other"},
		{Name: "empty input", Input: ""},
	}
	for _, tc := range cases {
		t.Run("govy/"+tc.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
		})
	}
	schema, err := govy.JSONSchema(validator, govy.JSONSchemaIncludeOmittedRules())
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "test_data/expected_custom_builder_composition_json_schema.json", cases)
}

func TestJSONSchema_CustomObjectApplicators(t *testing.T) {
	t.Parallel()

	contribution := func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
		return &jsonschema.Schema{
			Properties:           map[string]*jsonschema.Schema{"name": {Type: jsonschema.TypeString}},
			AdditionalProperties: &jsonschema.Schema{Type: jsonschema.TypeInteger},
		}, nil
	}
	matches := func(value map[string]any) bool {
		for key, item := range value {
			if key == "name" {
				if _, ok := item.(string); !ok {
					return false
				}
			} else if _, ok := item.(int); !ok {
				return false
			}
		}
		return true
	}
	rule := govy.NewRule(func(value map[string]any) error {
		if !matches(value) {
			return fmt.Errorf("invalid object")
		}
		return nil
	}).WithDescription("custom object").WithJSONSchema(contribution)
	base := govy.New(govy.For(govy.GetSelf[map[string]any]()).Rules(rule))
	for _, tt := range []struct {
		name      string
		validator govy.Validator[map[string]any]
		fixture   string
		cases     []jsonschematest.Case[map[string]any]
	}{
		{
			name: "custom contribution", validator: base, fixture: "object",
			cases: []jsonschematest.Case[map[string]any]{
				{Name: "named string and integer extra", Input: map[string]any{"name": "ok", "extra": 1}, Valid: true},
				{Name: "invalid extra", Input: map[string]any{"name": "ok", "extra": "wrong"}},
			},
		},
		{
			name: "custom condition",
			validator: govy.New(govy.For(govy.GetSelf[map[string]any]()).Rules(rules.MapMaxLength[map[string]any](2))).
				When(matches, govy.WhenJSONSchema(contribution)),
			fixture: "condition",
			cases: []jsonschematest.Case[map[string]any]{
				{Name: "condition matches", Input: map[string]any{"name": "ok", "a": 1}, Valid: true},
				{Name: "condition enables length check", Input: map[string]any{"name": "ok", "a": 1, "b": 2}},
				{Name: "condition disables length check", Input: map[string]any{"name": 3, "a": 1, "b": 2}, Valid: true},
			},
		},
		{
			name: "custom and generated wildcards",
			validator: govy.New(
				govy.For(govy.GetSelf[map[string]any]()).Rules(rule),
				govy.ForMap(govy.GetSelf[map[string]any]()).RulesForValues(rules.NEQ[any](1)),
			),
			fixture: "mixed",
			cases: []jsonschematest.Case[map[string]any]{
				{Name: "named string remains valid", Input: map[string]any{"name": "ok", "extra": 2}, Valid: true},
				{Name: "wildcard rejects integer", Input: map[string]any{"name": "ok", "extra": 1}},
				{Name: "invalid extra", Input: map[string]any{"name": "ok", "extra": "wrong"}},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			schema, err := govy.JSONSchema(tt.validator)
			assert.Require(t, assert.NoError(t, err))
			for _, tc := range tt.cases {
				assert.Equal(t, tc.Valid, tt.validator.Validate(tc.Input) == nil)
			}
			jsonschematest.Assert(t, schema, "test_data/expected_custom_applicator_"+tt.fixture+".json", tt.cases)
		})
	}
}

func TestJSONSchema_CustomArrayApplicators(t *testing.T) {
	t.Parallel()

	rule := govy.NewRule(func(value []any) error {
		for i, item := range value {
			if i == 0 {
				if _, ok := item.(string); !ok {
					return fmt.Errorf("invalid prefix item")
				}
			} else if _, ok := item.(int); !ok {
				return fmt.Errorf("invalid remaining item")
			}
		}
		return nil
	}).WithDescription("custom tuple").
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return &jsonschema.Schema{
				PrefixItems: []*jsonschema.Schema{{Type: jsonschema.TypeString}},
				Items:       &jsonschema.Schema{Type: jsonschema.TypeInteger},
			}, nil
		})
	validator := govy.New(govy.For(govy.GetSelf[[]any]()).Rules(rule))
	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))
	cases := []jsonschematest.Case[[]any]{
		{Name: "string prefix and integer tail", Input: []any{"ok", 1}, Valid: true},
		{Name: "invalid tail", Input: []any{"ok", "wrong"}},
		{Name: "invalid prefix", Input: []any{1, 2}},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
	}
	jsonschematest.Assert(t, schema, "test_data/expected_custom_applicator_array.json", cases)
}

func TestJSONSchema_CustomApplicatorsWithExplicitPaths(t *testing.T) {
	t.Parallel()
	t.Run("named property", func(t *testing.T) {
		t.Parallel()
		rule := govy.NewRule(func(value map[string]int) error {
			for _, item := range value {
				if item < 2 {
					return fmt.Errorf("all values must be at least 2")
				}
			}
			return nil
		}).WithDescription("minimum for all values").
			WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
				return &jsonschema.Schema{AdditionalProperties: &jsonschema.Schema{Minimum: "2"}}, nil
			})
		v := govy.New(
			govy.For(govy.GetSelf[map[string]int]()).Rules(rule),
			govy.For(func(value map[string]int) int { return value["name"] }).WithName("name").Rules(rules.GTE(1)),
		)
		schema, err := govy.JSONSchema(v)
		assert.Require(t, assert.NoError(t, err))
		cases := []jsonschematest.Case[map[string]int]{
			{Name: "both constraints pass", Input: map[string]int{"name": 2, "extra": 2}, Valid: true},
			{Name: "custom rule rejects named value", Input: map[string]int{"name": 1, "extra": 2}},
			{Name: "custom rule rejects extra value", Input: map[string]int{"name": 2, "extra": 1}},
		}
		for _, tc := range cases {
			assert.Equal(t, tc.Valid, v.Validate(tc.Input) == nil)
		}
		jsonschematest.Assert(t, schema, "test_data/expected_custom_applicator_named.json", cases)
	})
	t.Run("indexed item", func(t *testing.T) {
		t.Parallel()
		rule := govy.NewRule(func(value []int) error {
			for _, item := range value {
				if item < 2 {
					return fmt.Errorf("all items must be at least 2")
				}
			}
			return nil
		}).WithDescription("minimum for all items").
			WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
				return &jsonschema.Schema{Items: &jsonschema.Schema{Minimum: "2"}}, nil
			})
		v := govy.New(
			govy.For(govy.GetSelf[[]int]()).Rules(rule),
			govy.For(func(value []int) int { return value[0] }).WithPath(jsonpath.New().Index(0)).Rules(rules.GTE(1)),
		)
		schema, err := govy.JSONSchema(v)
		assert.Require(t, assert.NoError(t, err))
		cases := []jsonschematest.Case[[]int]{
			{Name: "both constraints pass", Input: []int{2, 2}, Valid: true},
			{Name: "custom rule rejects indexed value", Input: []int{1, 2}},
			{Name: "custom rule rejects remaining value", Input: []int{2, 1}},
		}
		for _, tc := range cases {
			assert.Equal(t, tc.Valid, v.Validate(tc.Input) == nil)
		}
		jsonschematest.Assert(t, schema, "test_data/expected_custom_applicator_indexed.json", cases)
	})
}

func TestJSONSchema_RequiredErrorCodes(t *testing.T) {
	t.Parallel()

	type document struct {
		Name string `json:"name,omitempty"`
	}
	for name, rule := range map[string]govy.RulesInterface[string]{
		"original": rules.Required[string](),
		"grouped":  govy.NewRuleSet(rules.Required[string]()).WithErrorCode("group"),
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := govy.New(govy.For(func(d document) string { return d.Name }).
				WithName("name").Rules(rule))
			schema, err := govy.JSONSchema(validator)
			assert.Require(t, assert.NoError(t, err))
			cases := []jsonschematest.Case[document]{
				{Name: "missing", Input: document{}},
				{Name: "present", Input: document{Name: "ok"}, Valid: true},
			}
			for _, tc := range cases {
				assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
			}
			jsonschematest.Assert(t, schema, "test_data/expected_required_codes_json_schema.json", cases)
		})
	}
}

func TestJSONSchema_IncludedRequiredProperties(t *testing.T) {
	t.Parallel()

	type child struct {
		Name string `json:"name,omitempty"`
	}
	type document struct {
		Optional *child           `json:"optional,omitempty"`
		Required *child           `json:"required,omitempty"`
		Items    []child          `json:"items"`
		Values   map[string]child `json:"values"`
	}
	childValidator := govy.New(
		govy.For(func(v child) string { return v.Name }).WithName("name").Required(),
	)
	validator := govy.New(
		govy.ForPointer(func(v document) *child { return v.Optional }).WithName("optional").Include(childValidator),
		govy.ForPointer(func(v document) *child { return v.Required }).
			WithName("required").Required().Include(childValidator),
		govy.ForSlice(func(v document) []child { return v.Items }).WithName("items").IncludeForEach(childValidator),
		govy.ForMap(func(v document) map[string]child { return v.Values }).
			WithName("values").
			IncludeForValues(childValidator),
	)
	valid := document{Required: &child{Name: "present"}, Items: []child{}, Values: map[string]child{}}
	allPresent := valid
	allPresent.Optional = &child{Name: "present"}
	allPresent.Items = []child{{Name: "first"}, {Name: "second"}}
	allPresent.Values = map[string]child{"first": {Name: "first"}, "second": {Name: "second"}}
	missingParent, missingRequiredName, missingOptionalName, missingItemName, missingMapValueName := valid, valid, valid, valid, valid
	missingParent.Required = nil
	missingRequiredName.Required = &child{}
	missingOptionalName.Optional = &child{}
	missingItemName.Items = []child{{Name: "present"}, {}}
	missingMapValueName.Values = map[string]child{"first": {Name: "present"}, "second": {}}
	cases := []jsonschematest.Case[document]{
		{Name: "optional parent absent and collections empty", Input: valid, Valid: true},
		{Name: "all included children satisfy required property", Input: allPresent, Valid: true},
		{Name: "required parent absent", Input: missingParent},
		{Name: "required parent lacks required child property", Input: missingRequiredName},
		{Name: "present optional parent lacks required child property", Input: missingOptionalName},
		{Name: "later array item lacks required property", Input: missingItemName},
		{Name: "map value lacks required property", Input: missingMapValueName},
	}
	for _, tc := range cases {
		t.Run("govy/"+tc.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
		})
	}
	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "test_data/expected_included_required_properties_json_schema.json", cases)
}

func TestJSONSchema_RemovePropertiesByID(t *testing.T) {
	t.Parallel()

	type child struct {
		Name   string `json:"name,omitempty"`
		Legacy string `json:"legacy,omitempty"`
	}
	type document struct {
		Child child   `json:"child"`
		Items []child `json:"items"`
	}
	childValidator := govy.New(
		govy.For(func(v child) string { return v.Name }).WithName("name").WithID("name").Required(),
		govy.For(func(v child) string { return v.Legacy }).
			WithName("legacy").WithID("legacy").Required().Rules(rules.EQ("fixed")),
	)
	validator := govy.New(
		govy.For(func(v document) child { return v.Child }).WithName("child").Include(&childValidator),
		govy.ForSlice(func(v document) []child { return v.Items }).WithName("items").IncludeForEach(&childValidator),
	)
	valid := document{
		Child: child{Name: "present", Legacy: "fixed"},
		Items: []child{{Name: "present", Legacy: "fixed"}},
	}
	missingLegacy, wrongLegacy, missingItemLegacy, missingName, missingItemName := valid, valid, valid, valid, valid
	missingLegacy.Child.Legacy = ""
	wrongLegacy.Child.Legacy = "other"
	missingItemLegacy.Items = []child{{Name: "present"}}
	missingName.Child.Name = ""
	missingItemName.Items = []child{{Legacy: "fixed"}}
	tests := []struct {
		name      string
		validator govy.Validator[document]
		removed   bool
		fixture   string
	}{
		{
			name:      "removed from both included validators",
			validator: validator.RemovePropertiesByID("legacy"),
			removed:   true,
			fixture:   "expected_removed_properties_json_schema.json",
		},
		{
			name:      "original remains unchanged",
			validator: validator,
			fixture:   "expected_original_properties_json_schema.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			cases := []jsonschematest.Case[document]{
				{Name: "all original constraints satisfied", Input: valid, Valid: true},
				{Name: "nested required property removed", Input: missingLegacy, Valid: tt.removed},
				{Name: "nested rule removed", Input: wrongLegacy, Valid: tt.removed},
				{Name: "array item required property removed", Input: missingItemLegacy, Valid: tt.removed},
				{Name: "other nested required property retained", Input: missingName},
				{Name: "other array item required property retained", Input: missingItemName},
			}
			for _, tc := range cases {
				t.Run("govy/"+tc.Name, func(t *testing.T) {
					t.Parallel()
					assert.Equal(t, tc.Valid, tt.validator.Validate(tc.Input) == nil)
				})
			}
			schema, err := govy.JSONSchema(tt.validator)
			assert.Require(t, assert.NoError(t, err))
			jsonschematest.Assert(t, schema, "test_data/"+tt.fixture, cases)
		})
	}
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

	builderError := errors.New("builder failed")
	tests := map[string]struct {
		validator     govy.Validator[struct{}]
		expectedError string
		expectedCause error
	}{
		"builder without descriptive metadata": {
			validator: govy.New(
				govy.For(func(struct{}) int { return 0 }).Rules(
					govy.NewRule(func(int) error { return nil }).
						WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
							return nil, builderError
						}),
				),
			),
			expectedError: `failed to build JSON Schema for "$" property and "" rule: builder failed`,
			expectedCause: builderError,
		},
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
					if tc.expectedCause != nil {
						assert.True(t, errors.Is(err, tc.expectedCause))
					}
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

func TestJSONSchema_NumericMapKeyRules(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		run  func(*testing.T) error
		kind string
	}{
		{
			name: "signed inequality", kind: "int",
			run: func(t *testing.T) error {
				v := govy.New(govy.ForMap(govy.GetSelf[map[int]string]()).RulesForKeys(rules.GTE(1)))
				assert.NoError(t, v.Validate(map[int]string{1: "ok"}))
				_, err := govy.JSONSchema(v)
				return err
			},
		},
		{
			name: "unsigned constant", kind: "uint",
			run: func(t *testing.T) error {
				v := govy.New(govy.ForMap(govy.GetSelf[map[uint]string]()).RulesForKeys(rules.EQ(uint(1))))
				assert.NoError(t, v.Validate(map[uint]string{1: "ok"}))
				_, err := govy.JSONSchema(v)
				return err
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualError(
				t,
				tt.run(t),
				`cannot generate JSON Schema for "$.*~" property: map key rules require string keys, got Go kind "`+tt.kind+`"`,
			)
		})
	}
	t.Run("map without key rules", func(t *testing.T) {
		t.Parallel()
		_, err := govy.JSONSchema(govy.New(govy.ForMap(govy.GetSelf[map[int]string]()).
			RulesForValues(rules.StringMinLength(1))))
		assert.NoError(t, err)
	})
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
