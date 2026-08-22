package govy_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"regexp"
	"testing"
	"time"
	"unsafe"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonpath"
	"github.com/nobl9/govy/pkg/jsonschema"
	"github.com/nobl9/govy/pkg/rules"
)

func TestValidatorPlan_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(newPodValidator())
	assert.Require(t, assert.NoError(t, err))

	expected := readTestData(t, "expected_pod_json_schema.json")
	var actual bytes.Buffer
	encoder := json.NewEncoder(&actual)
	encoder.SetIndent("", "  ")
	assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
	assert.Equal(t, expected, actual.String())
}

func TestJSONSchema_BuilderContext(t *testing.T) {
	t.Parallel()

	type document struct {
		Custom           string
		PropertyRequired string
		RuleRequired     bool
	}
	rootRule := govy.NewRule(func(document) error { return nil }).
		WithDescription("customize root schema").
		WithJSONSchema(func(ctx govy.JSONSchemaBuilderContext) error {
			if !ctx.Path.Equal(jsonpath.NewRoot()) {
				return fmt.Errorf("expected root builder path, got %q", ctx.Path)
			}
			if ctx.Root != ctx.Schema {
				return fmt.Errorf("expected root builder schema to be the document root")
			}
			if ctx.Parent != nil {
				return fmt.Errorf("expected root builder parent to be nil")
			}
			if ctx.Segment.Kind() != jsonpath.SegmentRoot {
				return fmt.Errorf("expected root segment, got %v", ctx.Segment.Kind())
			}
			ctx.Schema.Title = "Builder Context"
			return nil
		})
	customRule := govy.NewRule(func(string) error { return nil }).
		WithDescription("customize property schema").
		WithJSONSchema(func(ctx govy.JSONSchemaBuilderContext) error {
			if !ctx.Path.Equal(jsonpath.Parse("$.custom")) {
				return fmt.Errorf("expected custom property path, got %q", ctx.Path)
			}
			if ctx.Root == nil || ctx.Root != ctx.Parent {
				return fmt.Errorf("expected property builder document root")
			}
			if ctx.Parent == nil {
				return fmt.Errorf("expected property builder parent")
			}
			if ctx.Segment.Kind() != jsonpath.SegmentName || ctx.Segment.Name() != "custom" {
				return fmt.Errorf("expected custom property segment")
			}
			ctx.Schema.Pattern = "custom"
			ctx.Parent.Required = append(ctx.Parent.Required, ctx.Segment.Name())
			return nil
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

	expected := readTestData(t, "expected_builder_context_json_schema.json")
	var actual bytes.Buffer
	encoder := json.NewEncoder(&actual)
	encoder.SetIndent("", "  ")
	assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
	assert.Equal(t, expected, actual.String())
}

func TestJSONSchema_When(t *testing.T) {
	t.Parallel()

	type item struct {
		Kind  string
		Value string
	}
	type document struct {
		Enabled bool
		Ignored string
		Items   []item
	}
	enabled := any(true)
	kind := any("required")
	itemValidator := govy.New(
		govy.For(func(v item) string { return v.Value }).
			WithName("value").
			Required().
			When(
				func(v item) bool { return v.Kind == "required" },
				govy.WhenJSONSchema(func(ctx govy.JSONSchemaBuilderContext) error {
					if !ctx.Path.Equal(jsonpath.Parse("$.items[*]")) {
						return fmt.Errorf("expected item condition path, got %q", ctx.Path)
					}
					if ctx.Root == nil || ctx.Root == ctx.Schema {
						return fmt.Errorf("expected condition builder document root")
					}
					if ctx.Parent != nil || ctx.Segment.Kind() != jsonpath.SegmentRoot {
						return fmt.Errorf("expected condition root builder context")
					}
					ctx.Schema.Properties = map[string]*jsonschema.Schema{
						"kind": {Const: &kind},
					}
					ctx.Schema.Required = []string{"kind"}
					return nil
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
			govy.WhenJSONSchema(func(ctx govy.JSONSchemaBuilderContext) error {
				if !ctx.Path.Equal(jsonpath.NewRoot()) {
					return fmt.Errorf("expected root condition path, got %q", ctx.Path)
				}
				if ctx.Root == nil || ctx.Root == ctx.Schema {
					return fmt.Errorf("expected condition builder document root")
				}
				if ctx.Parent != nil || ctx.Segment.Kind() != jsonpath.SegmentRoot {
					return fmt.Errorf("expected condition root builder context")
				}
				ctx.Schema.Properties = map[string]*jsonschema.Schema{
					"enabled": {Const: &enabled},
				}
				ctx.Schema.Required = []string{"enabled"}
				return nil
			}),
		)

	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))

	expected := readTestData(t, "expected_when_json_schema.json")
	var actual bytes.Buffer
	encoder := json.NewEncoder(&actual)
	encoder.SetIndent("", "  ")
	assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
	assert.Equal(t, expected, actual.String())
}

func TestJSONSchema_Rules(t *testing.T) {
	t.Parallel()

	type choice struct {
		Name    string `json:"name"`
		Enabled bool   `json:"enabled"`
	}
	type presence struct {
		A string `json:"a"`
		B string `json:"b"`
		C string `json:"c"`
	}
	presenceGetters := map[string]func(presence) any{
		"a": func(v presence) any { return v.A },
		"b": func(v presence) any { return v.B },
		"c": func(v presence) any { return v.C },
	}
	type document struct {
		ArrayLength  []string
		AtLeastOne   presence
		AtMostOne    presence
		Bounds       int
		Const        [2]int
		ContainsAll  string
		DenyPattern  string
		Dependent    presence
		Duration     time.Duration
		EndsWith     string
		Enum         choice
		ExactlyOne   presence
		Excludes     string
		Forbidden    int
		MatchPattern string
		NotConst     string
		NotEmpty     string
		NotEnum      int
		ObjectLength map[string]string
		Ordered      string
		Pointer      *string
		Required     bool
		StartsWith   string
	}
	validator := govy.New(
		govy.For(func(v document) []string { return v.ArrayLength }).
			WithName("arrayLength").
			Rules(rules.SliceLength[[]string](1, 3)),
		govy.For(func(v document) presence { return v.AtLeastOne }).
			WithName("atLeastOne").
			Rules(rules.OneOfProperties(presenceGetters)),
		govy.For(func(v document) presence { return v.AtMostOne }).
			WithName("atMostOne").
			Rules(rules.MutuallyExclusive(false, presenceGetters)),
		govy.For(func(v document) int { return v.Bounds }).
			WithName("bounds").
			Rules(rules.GT(1), rules.GTE(2), rules.LT(10), rules.LTE(9)),
		govy.For(func(v document) [2]int { return v.Const }).
			WithName("const").
			Rules(rules.EQ([2]int{1, 2})),
		govy.For(func(v document) string { return v.ContainsAll }).
			WithName("containsAll").
			Rules(rules.StringContains("alpha", "omega")),
		govy.For(func(v document) string { return v.DenyPattern }).
			WithName("denyPattern").
			Rules(rules.StringDenyRegexp(regexp.MustCompile("secret"))),
		govy.For(func(v document) presence { return v.Dependent }).
			WithName("dependent").
			Rules(rules.MutuallyDependent(presenceGetters)),
		govy.For(func(v document) time.Duration { return v.Duration }).
			WithName("duration").
			Rules(rules.DurationPrecision(250*time.Millisecond)),
		govy.For(func(v document) string { return v.EndsWith }).
			WithName("endsWith").
			Rules(rules.StringEndsWith(".json")),
		govy.For(func(v document) choice { return v.Enum }).
			WithName("enum").
			Rules(rules.OneOf(
				choice{Name: "first", Enabled: true},
				choice{Name: "second"},
			)),
		govy.For(func(v document) presence { return v.ExactlyOne }).
			WithName("exactlyOne").
			Rules(rules.MutuallyExclusive(true, presenceGetters)),
		govy.For(func(v document) string { return v.Excludes }).
			WithName("excludes").
			Rules(rules.StringExcludes("secret")),
		govy.For(func(v document) int { return v.Forbidden }).
			WithName("forbidden").
			Rules(rules.Forbidden[int]()),
		govy.For(func(v document) string { return v.MatchPattern }).
			WithName("matchPattern").
			Rules(rules.StringMatchRegexp(regexp.MustCompile("[a-z]+"))),
		govy.For(func(v document) string { return v.NotConst }).
			WithName("notConst").
			Rules(rules.NEQ("forbidden")),
		govy.For(func(v document) string { return v.NotEmpty }).
			WithName("notEmpty").
			Rules(rules.StringNotEmpty()),
		govy.For(func(v document) int { return v.NotEnum }).
			WithName("notEnum").
			Rules(rules.NotOneOf(1, 2)),
		govy.For(func(v document) map[string]string { return v.ObjectLength }).
			WithName("objectLength").
			Rules(rules.MapLength[map[string]string](1, 3)),
		govy.For(func(v document) string { return v.Ordered }).
			WithName("ordered").
			Rules(rules.GT("middle")),
		govy.For(func(v document) *string { return v.Pointer }).
			WithName("pointer").
			Rules(govy.RuleToPointer(rules.OneOf("first", "second"))),
		govy.For(func(v document) bool { return v.Required }).
			WithName("required").
			Rules(rules.Required[bool]()),
		govy.For(func(v document) string { return v.StartsWith }).
			WithName("startsWith").
			Rules(rules.StringStartsWith("prefix")),
	).
		WithName("Rules")

	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))

	expected := readTestData(t, "expected_rules_json_schema.json")
	var actual bytes.Buffer
	encoder := json.NewEncoder(&actual)
	encoder.SetIndent("", "  ")
	assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
	assert.Equal(t, expected, actual.String())
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
			_, err := govy.JSONSchema(tc.validator)
			assert.EqualError(t, err, tc.expectedError)
		})
	}
}

func TestJSONSchema_StringPatternRules(t *testing.T) {
	t.Parallel()

	type document struct {
		CVE           string
		DNSLabel      string
		E164          string
		Hexadecimal   string
		MD5           string
		MongoObjectID string
		Semver        string
		SHA256        string
		ULID          string
		UUID          string
		UUIDv4        string
	}
	validator := govy.New(
		govy.For(func(v document) string { return v.CVE }).
			WithName("cve").
			Rules(rules.StringCVE()),
		govy.For(func(v document) string { return v.DNSLabel }).
			WithName("dnsLabel").
			Rules(rules.StringDNSLabel()),
		govy.For(func(v document) string { return v.E164 }).
			WithName("e164").
			Rules(rules.StringE164()),
		govy.For(func(v document) string { return v.Hexadecimal }).
			WithName("hexadecimal").
			Rules(rules.StringHexadecimal()),
		govy.For(func(v document) string { return v.MD5 }).
			WithName("md5").
			Rules(rules.StringMD5()),
		govy.For(func(v document) string { return v.MongoObjectID }).
			WithName("mongoObjectID").
			Rules(rules.StringMongoDBObjectID()),
		govy.For(func(v document) string { return v.Semver }).
			WithName("semver").
			Rules(rules.StringSemver()),
		govy.For(func(v document) string { return v.SHA256 }).
			WithName("sha256").
			Rules(rules.StringSHA256()),
		govy.For(func(v document) string { return v.ULID }).
			WithName("ulid").
			Rules(rules.StringULID()),
		govy.For(func(v document) string { return v.UUID }).
			WithName("uuid").
			Rules(rules.StringUUID()),
		govy.For(func(v document) string { return v.UUIDv4 }).
			WithName("uuidV4").
			Rules(rules.StringUUIDv4()),
	).
		WithName("StringPatternRules")

	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))

	expected := readTestData(t, "expected_string_pattern_rules_json_schema.json")
	var actual bytes.Buffer
	encoder := json.NewEncoder(&actual)
	encoder.SetIndent("", "  ")
	assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
	assert.Equal(t, expected, actual.String())
}

func TestJSONSchema_StringFormatRules(t *testing.T) {
	t.Parallel()

	type document struct {
		CustomDateTime string
		DateTime       string
		DateTimeNano   string
		Email          string
		IP             string
		IPv4           string
		IPv6           string
		Multiple       string
		URL            string
	}
	validator := govy.New(
		govy.For(func(v document) string { return v.CustomDateTime }).
			WithName("customDateTime").
			Rules(rules.StringDateTime("2006-01-02")),
		govy.For(func(v document) string { return v.DateTime }).
			WithName("dateTime").
			Rules(rules.StringDateTime(time.RFC3339)),
		govy.For(func(v document) string { return v.DateTimeNano }).
			WithName("dateTimeNano").
			Rules(rules.StringDateTime(time.RFC3339Nano)),
		govy.For(func(v document) string { return v.Email }).
			WithName("email").
			Rules(rules.StringEmail()),
		govy.For(func(v document) string { return v.IP }).
			WithName("ip").
			Rules(rules.StringIP()),
		govy.For(func(v document) string { return v.IPv4 }).
			WithName("ipv4").
			Rules(rules.StringIPv4()),
		govy.For(func(v document) string { return v.IPv6 }).
			WithName("ipv6").
			Rules(rules.StringIPv6()),
		govy.For(func(v document) string { return v.Multiple }).
			WithName("multiple").
			Rules(rules.StringIPv4(), rules.StringIPv6()),
		govy.For(func(v document) string { return v.URL }).
			WithName("url").
			Rules(rules.StringURL()),
	).
		WithName("StringFormatRules")

	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))

	expected := readTestData(t, "expected_string_format_rules_json_schema.json")
	var actual bytes.Buffer
	encoder := json.NewEncoder(&actual)
	encoder.SetIndent("", "  ")
	assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
	assert.Equal(t, expected, actual.String())
}

func TestJSONSchema_UnsupportedType(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		run           func() error
		expectedError string
	}{
		"channel": {
			run:           runJSONSchema[chan struct{}],
			expectedError: `failed to generate JSON Schema: unsupported Go kind "chan"`,
		},
		"complex64": {
			run:           runJSONSchema[complex64],
			expectedError: `failed to generate JSON Schema: unsupported Go kind "complex64"`,
		},
		"complex128": {
			run:           runJSONSchema[complex128],
			expectedError: `failed to generate JSON Schema: unsupported Go kind "complex128"`,
		},
		"function": {
			run:           runJSONSchema[func()],
			expectedError: `failed to generate JSON Schema: unsupported Go kind "func"`,
		},
		"unsafe pointer": {
			run:           runJSONSchema[unsafe.Pointer],
			expectedError: `failed to generate JSON Schema: unsupported Go kind "unsafe.Pointer"`,
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
	type document struct{}
	validator := govy.New(
		govy.ForMap(func(document) annotations { return nil }).
			WithName("annotations").
			RulesForItems(
				govy.NewRule(func(govy.MapItem[string, string]) error { return nil }).
					WithDescription("key and value must differ"),
			),
	)

	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))

	expected := readTestData(t, "expected_map_item_json_schema.json")
	var actual bytes.Buffer
	encoder := json.NewEncoder(&actual)
	encoder.SetIndent("", "  ")
	assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
	assert.Equal(t, expected, actual.String())
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

	expected := readTestData(t, "expected_fixed_array_indexes_json_schema.json")
	var actual bytes.Buffer
	encoder := json.NewEncoder(&actual)
	encoder.SetIndent("", "  ")
	assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
	assert.Equal(t, expected, actual.String())
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

	expected := readTestData(t, "expected_array_wildcard_with_fixed_index_json_schema.json")
	var actual bytes.Buffer
	encoder := json.NewEncoder(&actual)
	encoder.SetIndent("", "  ")
	assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
	assert.Equal(t, expected, actual.String())
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

	expected := readTestData(t, "expected_value_wildcard_with_named_property_json_schema.json")
	var actual bytes.Buffer
	encoder := json.NewEncoder(&actual)
	encoder.SetIndent("", "  ")
	assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
	assert.Equal(t, expected, actual.String())
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
