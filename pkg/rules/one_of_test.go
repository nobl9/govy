package rules

import (
	"fmt"
	"math"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonschematest"
	"github.com/nobl9/govy/pkg/govy"
)

var oneOfTestCases = []*struct {
	in            string
	options       []string
	expectedError string
}{
	{"that", []string{"this", "that"}, ""},
	{"those", []string{"this", "that"}, "must be one of: this, that"},
}

func TestOneOf(t *testing.T) {
	for _, tc := range oneOfTestCases {
		err := OneOf(tc.options...).Validate(tc.in)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeOneOf))
		} else {
			assert.NoError(t, err)
		}
	}
	t.Run("panic if values are empty", func(t *testing.T) {
		assert.Panic(t, func() { OneOf[string]() }, "values must not be empty")
	})
}

func TestOneOf_JSONSchema(t *testing.T) {
	t.Parallel()
	options := oneOfTestCases[0].options
	validator := govy.New(govy.For(govy.GetSelf[string]()).Rules(OneOf(options...)))
	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))
	cases := make([]jsonschematest.Case[string], len(oneOfTestCases))
	for i, tc := range oneOfTestCases {
		assert.Require(t, assert.Equal(t, options, tc.options))
		cases[i] = jsonschematest.Case[string]{
			Name:  tc.in,
			Input: tc.in,
			Valid: tc.expectedError == "",
		}
	}
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_one_of.json", cases)
}

func BenchmarkOneOf(b *testing.B) {
	for _, tc := range oneOfTestCases {
		rule := OneOf(tc.options...)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var notOneOfTestCases = []*struct {
	in            string
	options       []string
	expectedError string
}{
	{"that", []string{"this", "that"}, "must not be one of: this, that"},
	{"those", []string{"this", "that"}, ""},
}

func TestNotOneOf(t *testing.T) {
	for _, tc := range notOneOfTestCases {
		err := NotOneOf(tc.options...).Validate(tc.in)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeNotOneOf))
		} else {
			assert.NoError(t, err)
		}
	}
	t.Run("panic if values are empty", func(t *testing.T) {
		assert.Panic(t, func() { NotOneOf[string]() }, "values must not be empty")
	})
}

func TestNotOneOf_JSONSchema(t *testing.T) {
	t.Parallel()
	options := notOneOfTestCases[0].options
	validator := govy.New(govy.For(govy.GetSelf[string]()).Rules(NotOneOf(options...)))
	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))
	cases := make([]jsonschematest.Case[string], len(notOneOfTestCases))
	for i, tc := range notOneOfTestCases {
		assert.Require(t, assert.Equal(t, options, tc.options))
		cases[i] = jsonschematest.Case[string]{
			Name:  tc.in,
			Input: tc.in,
			Valid: tc.expectedError == "",
		}
	}
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_not_one_of.json", cases)
}

func BenchmarkNotOneOf(b *testing.B) {
	for _, tc := range notOneOfTestCases {
		rule := NotOneOf(tc.options...)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

type paymentMethod struct {
	Cash     *string
	Card     *string
	Transfer *string
}

var paymentMethodGetters = map[string]func(p paymentMethod) any{
	"Cash":     func(p paymentMethod) any { return p.Cash },
	"Card":     func(p paymentMethod) any { return p.Card },
	"Transfer": func(p paymentMethod) any { return p.Transfer },
}

var oneOfPropertiesTestCases = []*struct {
	paymentMethod        paymentMethod
	expectedError        string
	jsonSchemaDifference string
}{
	{
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     ptr("2$"),
			Transfer: nil,
		},
	},
	{
		paymentMethod: paymentMethod{
			Cash:     ptr("1$"),
			Card:     ptr("2$"),
			Transfer: nil,
		},
	},
	{
		paymentMethod: paymentMethod{
			Cash:     ptr("1$"),
			Card:     ptr("2$"),
			Transfer: ptr("3$"),
		},
	},
	{
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     nil,
			Transfer: nil,
		},
		expectedError:        "one of [Card, Cash, Transfer] properties must be set, none was provided",
		jsonSchemaDifference: "JSON includes all three properties as null; JSON Schema checks presence, not non-zero values.",
	},
}

func TestOneOfProperties(t *testing.T) {
	for _, tc := range oneOfPropertiesTestCases {
		err := OneOfProperties(paymentMethodGetters).Validate(tc.paymentMethod)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeOneOfProperties))
		} else {
			assert.NoError(t, err)
		}
	}
	t.Run("panic if getters are empty", func(t *testing.T) {
		assert.Panic(t,
			func() { OneOfProperties(map[string]func(paymentMethod) any{}) },
			"getters must not be empty")
	})
}

func TestOneOfProperties_JSONSchema(t *testing.T) {
	t.Parallel()
	validator := govy.New(govy.For(govy.GetSelf[paymentMethod]()).Rules(OneOfProperties(paymentMethodGetters)))
	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))
	cases := make([]jsonschematest.Case[paymentMethod], len(oneOfPropertiesTestCases))
	for i, tc := range oneOfPropertiesTestCases {
		cases[i] = jsonschematest.Case[paymentMethod]{
			Name:                 fmt.Sprint(i),
			Input:                tc.paymentMethod,
			Valid:                tc.expectedError == "",
			JSONSchemaDifference: tc.jsonSchemaDifference,
		}
	}
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_one_of_properties.json", cases)
}

func BenchmarkOneOfProperties(b *testing.B) {
	for _, tc := range oneOfPropertiesTestCases {
		rule := OneOfProperties(paymentMethodGetters)
		for range b.N {
			_ = rule.Validate(tc.paymentMethod)
		}
	}
}

var mutuallyExclusiveTestCases = []*struct {
	required             bool
	paymentMethod        paymentMethod
	expectedError        string
	jsonSchemaDifference string
}{
	{
		required: true,
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     ptr("2$"),
			Transfer: nil,
		},
		jsonSchemaDifference: "JSON includes both nil pointers as present null properties, so more than one property is present.",
	},
	{
		required: false,
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     nil,
			Transfer: nil,
		},
		jsonSchemaDifference: "JSON includes all three nil pointers as present null properties, so more than one property is present.",
	},
	{
		required: true,
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     ptr("2$"),
			Transfer: ptr("2$"),
		},
		expectedError: "[Card, Transfer] properties are mutually exclusive, provide only one of them",
	},
	{
		required: false,
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     ptr("2$"),
			Transfer: ptr("2$"),
		},
		expectedError: "[Card, Transfer] properties are mutually exclusive, provide only one of them",
	},
	{
		required: true,
		paymentMethod: paymentMethod{
			Cash:     ptr("2$"),
			Card:     ptr("2$"),
			Transfer: ptr("2$"),
		},
		expectedError: "[Card, Cash, Transfer] properties are mutually exclusive, provide only one of them",
	},
	{
		required: false,
		paymentMethod: paymentMethod{
			Cash:     ptr("2$"),
			Card:     ptr("2$"),
			Transfer: ptr("2$"),
		},
		expectedError: "[Card, Cash, Transfer] properties are mutually exclusive, provide only one of them",
	},
	{
		required: true,
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     nil,
			Transfer: nil,
		},
		expectedError: "one of [Card, Cash, Transfer] properties must be set, none was provided",
	},
}

func TestMutuallyExclusive(t *testing.T) {
	for _, tc := range mutuallyExclusiveTestCases {
		err := MutuallyExclusive(tc.required, paymentMethodGetters).Validate(tc.paymentMethod)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeMutuallyExclusive))
		} else {
			assert.NoError(t, err)
		}
	}
	invalidGetters := map[string]map[string]func(paymentMethod) any{
		"no properties": {},
		"one property": {
			"Cash": func(p paymentMethod) any { return p.Cash },
		},
	}
	for name, getters := range invalidGetters {
		t.Run("panic if getters contain "+name, func(t *testing.T) {
			assert.Panic(t,
				func() { MutuallyExclusive(false, getters) },
				"getters must contain at least two properties")
		})
	}
}

func TestMutuallyExclusive_JSONSchema(t *testing.T) {
	t.Parallel()
	for _, required := range []bool{false, true} {
		t.Run(fmt.Sprintf("required=%t", required), func(t *testing.T) {
			t.Parallel()
			validator := govy.New(
				govy.For(govy.GetSelf[paymentMethod]()).Rules(MutuallyExclusive(required, paymentMethodGetters)),
			)
			schema, err := govy.JSONSchema(validator)
			assert.Require(t, assert.NoError(t, err))
			var cases []jsonschematest.Case[paymentMethod]
			for i, tc := range mutuallyExclusiveTestCases {
				if tc.required != required {
					continue
				}
				cases = append(cases, jsonschematest.Case[paymentMethod]{
					Name:                 fmt.Sprint(i),
					Input:                tc.paymentMethod,
					Valid:                tc.expectedError == "",
					JSONSchemaDifference: tc.jsonSchemaDifference,
				})
			}
			fixture := fmt.Sprintf("testdata/jsonschema/expected_mutually_exclusive_%t.json", required)
			jsonschematest.Assert(t, schema, fixture, cases)
		})
	}
}

func BenchmarkMutuallyExclusive(b *testing.B) {
	for _, tc := range mutuallyExclusiveTestCases {
		rule := MutuallyExclusive(tc.required, paymentMethodGetters)
		for range b.N {
			_ = rule.Validate(tc.paymentMethod)
		}
	}
}

var mutuallyDependentTestCases = []*struct {
	paymentMethod        paymentMethod
	expectedError        string
	jsonSchemaDifference string
}{
	{
		paymentMethod: paymentMethod{
			Cash:     ptr("2$"),
			Card:     ptr("2$"),
			Transfer: ptr("2$"),
		},
	},
	{
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     nil,
			Transfer: nil,
		},
	},
	{
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     ptr("2$"),
			Transfer: nil,
		},
		expectedError: "[Card, Cash, Transfer] properties are mutually dependent," +
			" since [Card] is provided, [Cash, Transfer] properties must also be set",
		jsonSchemaDifference: "JSON includes the nil pointers as null, so all dependent properties are present.",
	},
	{
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     ptr("2$"),
			Transfer: ptr("2$"),
		},
		expectedError: "[Card, Cash, Transfer] properties are mutually dependent," +
			" since [Card, Transfer] are provided, [Cash] property must also be set",
		jsonSchemaDifference: "JSON includes the nil pointer as null, so all dependent properties are present.",
	},
}

func TestMutuallyDependent(t *testing.T) {
	for _, tc := range mutuallyDependentTestCases {
		err := MutuallyDependent(paymentMethodGetters).Validate(tc.paymentMethod)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeMutuallyDependent))
		} else {
			assert.NoError(t, err)
		}
	}
	t.Run("panic if getters contain fewer than two properties", func(t *testing.T) {
		assert.Panic(t,
			func() {
				MutuallyDependent(map[string]func(paymentMethod) any{
					"Cash": func(p paymentMethod) any { return p.Cash },
				})
			},
			"getters must contain at least two properties")
	})
}

func TestMutuallyDependent_JSONSchema(t *testing.T) {
	t.Parallel()
	validator := govy.New(govy.For(govy.GetSelf[paymentMethod]()).Rules(MutuallyDependent(paymentMethodGetters)))
	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))
	cases := make([]jsonschematest.Case[paymentMethod], len(mutuallyDependentTestCases))
	for i, tc := range mutuallyDependentTestCases {
		cases[i] = jsonschematest.Case[paymentMethod]{
			Name:                 fmt.Sprint(i),
			Input:                tc.paymentMethod,
			Valid:                tc.expectedError == "",
			JSONSchemaDifference: tc.jsonSchemaDifference,
		}
	}
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_mutually_dependent.json", cases)
}

func BenchmarkMutuallyDependent(b *testing.B) {
	for _, tc := range mutuallyDependentTestCases {
		rule := MutuallyDependent(paymentMethodGetters)
		for range b.N {
			_ = rule.Validate(tc.paymentMethod)
		}
	}
}

func TestOneOfAndNotOneOf_JSONSchemaRepresentations(t *testing.T) {
	t.Parallel()
	t.Run("duplicate strings", func(t *testing.T) {
		t.Parallel()
		testEnumJSONSchemas(t, "strings", []string{"this", "this", "that"}, []jsonschematest.Case[string]{
			{Name: "first", Input: "this", Valid: true},
			{Name: "second", Input: "that", Valid: true},
			{Name: "not listed", Input: "those", Valid: false},
		})
	})
	t.Run("integers", func(t *testing.T) {
		t.Parallel()
		testEnumJSONSchemas(t, "integers", []int{-1, 0, 2}, []jsonschematest.Case[int]{
			{Name: "negative", Input: -1, Valid: true},
			{Name: "zero", Input: 0, Valid: true},
			{Name: "positive", Input: 2, Valid: true},
			{Name: "not listed", Input: 1, Valid: false},
		})
	})
	t.Run("booleans", func(t *testing.T) {
		t.Parallel()
		testEnumJSONSchemas(t, "booleans", []bool{true}, []jsonschematest.Case[bool]{
			{Name: "listed", Input: true, Valid: true},
			{Name: "not listed", Input: false, Valid: false},
		})
	})
	t.Run("arrays", func(t *testing.T) {
		t.Parallel()
		testEnumJSONSchemas(t, "arrays", [][2]int{{1, 2}, {2, 1}}, []jsonschematest.Case[[2]int]{
			{Name: "first", Input: [2]int{1, 2}, Valid: true},
			{Name: "second", Input: [2]int{2, 1}, Valid: true},
			{Name: "not listed", Input: [2]int{1, 1}, Valid: false},
		})
	})
	t.Run("structs", func(t *testing.T) {
		t.Parallel()
		type choice struct {
			Name     string `json:"name"`
			Revision int    `json:"-"`
		}
		testEnumJSONSchemas(t, "structs", []choice{
			{Name: "first", Revision: 1},
			{Name: "first", Revision: 2},
			{Name: "second", Revision: 1},
		}, []jsonschematest.Case[choice]{
			{Name: "first", Input: choice{Name: "first", Revision: 1}, Valid: true},
			{Name: "duplicate JSON value", Input: choice{Name: "first", Revision: 2}, Valid: true},
			{Name: "second", Input: choice{Name: "second", Revision: 1}, Valid: true},
			{Name: "not listed", Input: choice{Name: "third", Revision: 1}, Valid: false},
			{
				Name:                 "ignored field differs",
				Input:                choice{Name: "first", Revision: 3},
				Valid:                false,
				JSONSchemaDifference: "Go compares the ignored Revision field; JSON Schema compares only the serialized object.",
			},
		})
	})
	t.Run("pointers", func(t *testing.T) {
		t.Parallel()
		listed := ptr("first")
		testEnumJSONSchemas(t, "pointers", []*string{listed}, []jsonschematest.Case[*string]{
			{Name: "same pointer", Input: listed, Valid: true},
			{
				Name:                 "same value at different address",
				Input:                ptr("first"),
				Valid:                false,
				JSONSchemaDifference: "Go compares pointer identity; JSON Schema compares the pointed-to value.",
			},
			{Name: "different value", Input: ptr("second"), Valid: false},
		})
	})
}

func TestOneOfAndNotOneOf_JSONSchemaUnsupportedValues(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name       string
		value      float64
		errorValue string
	}{
		{name: "NaN", value: math.NaN(), errorValue: "NaN"},
		{name: "positive infinity", value: math.Inf(1), errorValue: "+Inf"},
		{name: "negative infinity", value: math.Inf(-1), errorValue: "-Inf"},
	} {
		for code, rule := range map[govy.ErrorCode]govy.Rule[float64]{
			ErrorCodeOneOf:    OneOf(tc.value),
			ErrorCodeNotOneOf: NotOneOf(tc.value),
		} {
			t.Run(fmt.Sprintf("%s/%s", code, tc.name), func(t *testing.T) {
				t.Parallel()
				schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[float64]()).Rules(rule)))
				assert.True(t, schema == nil)
				assert.EqualError(t, err, fmt.Sprintf(
					`failed to build JSON Schema for "$" property and %q rule: convert value at index 0: marshal value: json: unsupported value: %s`,
					code,
					tc.errorValue,
				))
			})
		}
	}
	t.Run("non-JSON struct field", func(t *testing.T) {
		t.Parallel()
		type value struct {
			Channel chan struct{}
		}
		for code, rule := range map[govy.ErrorCode]govy.Rule[value]{
			ErrorCodeOneOf:    OneOf(value{}),
			ErrorCodeNotOneOf: NotOneOf(value{}),
		} {
			t.Run(string(code), func(t *testing.T) {
				t.Parallel()
				schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[value]()).Rules(rule)))
				assert.True(t, schema == nil)
				assert.EqualError(t, err, fmt.Sprintf(
					`failed to build JSON Schema for "$" property and %q rule: convert value at index 0: marshal value: json: unsupported type: chan struct {}`,
					code,
				))
			})
		}
	})
}

func testEnumJSONSchemas[T comparable](t *testing.T, fixture string, values []T, cases []jsonschematest.Case[T]) {
	t.Helper()
	for code, rule := range map[govy.ErrorCode]govy.Rule[T]{
		ErrorCodeOneOf:    OneOf(values...),
		ErrorCodeNotOneOf: NotOneOf(values...),
	} {
		t.Run(string(code), func(t *testing.T) {
			t.Parallel()
			schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[T]()).Rules(rule)))
			assert.Require(t, assert.NoError(t, err))
			inputs := make([]jsonschematest.Case[T], len(cases))
			for i, tc := range cases {
				if code == ErrorCodeNotOneOf {
					tc.Valid = !tc.Valid
				}
				inputs[i] = tc
				assert.Equal(t, tc.Valid, rule.Validate(tc.Input) == nil)
			}
			path := fmt.Sprintf("testdata/jsonschema/expected_%s_%s.json", code, fixture)
			jsonschematest.Assert(t, schema, path, inputs)
		})
	}
}

func TestPresenceRules_JSONSchemaMissingProperties(t *testing.T) {
	t.Parallel()
	type presenceCase struct {
		name       string
		input      map[string]string
		atLeastOne bool
		exactlyOne bool
		atMostOne  bool
		allOrNone  bool
	}
	cases := []presenceCase{
		{
			name:       "none",
			input:      map[string]string{},
			atLeastOne: false,
			exactlyOne: false,
			atMostOne:  true,
			allOrNone:  true,
		},
		{
			name:       "cash",
			input:      map[string]string{"Cash": "1$"},
			atLeastOne: true,
			exactlyOne: true,
			atMostOne:  true,
			allOrNone:  false,
		},
		{
			name:       "card",
			input:      map[string]string{"Card": "1$"},
			atLeastOne: true,
			exactlyOne: true,
			atMostOne:  true,
			allOrNone:  false,
		},
		{
			name:       "transfer",
			input:      map[string]string{"Transfer": "1$"},
			atLeastOne: true,
			exactlyOne: true,
			atMostOne:  true,
			allOrNone:  false,
		},
		{
			name:       "cash and card",
			input:      map[string]string{"Cash": "1$", "Card": "1$"},
			atLeastOne: true,
			exactlyOne: false,
			atMostOne:  false,
			allOrNone:  false,
		},
		{
			name:       "cash and transfer",
			input:      map[string]string{"Cash": "1$", "Transfer": "1$"},
			atLeastOne: true,
			exactlyOne: false,
			atMostOne:  false,
			allOrNone:  false,
		},
		{
			name:       "card and transfer",
			input:      map[string]string{"Card": "1$", "Transfer": "1$"},
			atLeastOne: true,
			exactlyOne: false,
			atMostOne:  false,
			allOrNone:  false,
		},
		{
			name:       "all",
			input:      map[string]string{"Cash": "1$", "Card": "1$", "Transfer": "1$"},
			atLeastOne: true,
			exactlyOne: false,
			atMostOne:  false,
			allOrNone:  true,
		},
	}
	getters := map[string]func(map[string]string) any{
		"Cash":     func(v map[string]string) any { return v["Cash"] },
		"Card":     func(v map[string]string) any { return v["Card"] },
		"Transfer": func(v map[string]string) any { return v["Transfer"] },
	}
	for _, ruleCase := range []struct {
		name  string
		rule  govy.Rule[map[string]string]
		valid func(presenceCase) bool
	}{
		{name: "one_of_properties", rule: OneOfProperties(getters), valid: func(tc presenceCase) bool { return tc.atLeastOne }},
		{name: "mutually_exclusive_true", rule: MutuallyExclusive(true, getters), valid: func(tc presenceCase) bool { return tc.exactlyOne }},
		{name: "mutually_exclusive_false", rule: MutuallyExclusive(false, getters), valid: func(tc presenceCase) bool { return tc.atMostOne }},
		{name: "mutually_dependent", rule: MutuallyDependent(getters), valid: func(tc presenceCase) bool { return tc.allOrNone }},
	} {
		t.Run(ruleCase.name, func(t *testing.T) {
			t.Parallel()
			validator := govy.New(govy.For(govy.GetSelf[map[string]string]()).Rules(ruleCase.rule))
			schema, err := govy.JSONSchema(validator)
			assert.Require(t, assert.NoError(t, err))
			inputs := make([]jsonschematest.Case[map[string]string], len(cases))
			for i, tc := range cases {
				inputs[i] = jsonschematest.Case[map[string]string]{
					Name:  tc.name,
					Input: tc.input,
					Valid: ruleCase.valid(tc),
				}
				assert.Equal(t, inputs[i].Valid, validator.Validate(tc.input) == nil)
			}
			fixture := "testdata/jsonschema/expected_" + ruleCase.name + ".json"
			jsonschematest.Assert(t, schema, fixture, inputs)
		})
	}
}
