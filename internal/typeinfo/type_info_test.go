package typeinfo

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/nobl9/govy/internal/assert"
)

const packageName = "github.com/nobl9/govy/internal/typeinfo"

type customString string

type customStruct struct{}

type customMap map[string]int

type customSlice []customMap

type customNestedMap map[customString]customSlice

type customStringSlice []string

type testCase struct {
	name     string
	typeFunc func() TypeInfo
	expected TypeInfo
}

func TestGet(t *testing.T) {
	tests := []testCase{
		{
			name:     "any",
			typeFunc: func() TypeInfo { return Get[any]() },
			expected: TypeInfo{},
		},
		{
			name:     "int",
			typeFunc: func() TypeInfo { return Get[int]() },
			expected: TypeInfo{Name: "int", Package: "", Kind: "int", ReflectKind: reflect.Int},
		},
		{
			name:     "pointer to int",
			typeFunc: func() TypeInfo { return Get[*int]() },
			expected: TypeInfo{Name: "int", Package: "", Kind: "int", ReflectKind: reflect.Int},
		},
		{
			name:     "slice of int",
			typeFunc: func() TypeInfo { return Get[[]int]() },
			expected: TypeInfo{Name: "[]int", Package: "", Kind: "[]int", ReflectKind: reflect.Slice},
		},
		{
			name:     "slice of customString",
			typeFunc: func() TypeInfo { return Get[[]customString]() },
			expected: TypeInfo{
				Name: "[]customString", Package: packageName, Kind: "[]string", ReflectKind: reflect.Slice,
			},
		},
		{
			name:     "map of string to int",
			typeFunc: func() TypeInfo { return Get[map[string]int]() },
			expected: TypeInfo{
				Name: "map[string]int", Package: "", Kind: "map[string]int", ReflectKind: reflect.Map,
			},
		},
		{
			name:     "custom string",
			typeFunc: func() TypeInfo { return Get[customString]() },
			expected: TypeInfo{
				Name: "customString", Package: packageName, Kind: "string", ReflectKind: reflect.String,
			},
		},
		{
			name:     "custom struct",
			typeFunc: func() TypeInfo { return Get[customStruct]() },
			expected: TypeInfo{
				Name: "customStruct", Package: packageName, Kind: "struct", ReflectKind: reflect.Struct,
			},
		},
		{
			name:     "pointer to custom struct",
			typeFunc: func() TypeInfo { return Get[*customStruct]() },
			expected: TypeInfo{
				Name: "customStruct", Package: packageName, Kind: "struct", ReflectKind: reflect.Struct,
			},
		},
		{
			name:     "custom map",
			typeFunc: func() TypeInfo { return Get[customMap]() },
			expected: TypeInfo{
				Name: "customMap", Package: packageName, Kind: "map[string]int", ReflectKind: reflect.Map,
			},
		},
		{
			name:     "custom nested map",
			typeFunc: func() TypeInfo { return Get[customNestedMap]() },
			expected: TypeInfo{
				Name: "customNestedMap", Package: packageName, Kind: "map[string][]map[string]int",
				ReflectKind: reflect.Map,
			},
		},
		{
			name:     "custom slice",
			typeFunc: func() TypeInfo { return Get[customSlice]() },
			expected: TypeInfo{
				Name: "customSlice", Package: packageName, Kind: "[]map[string]int", ReflectKind: reflect.Slice,
			},
		},
		{
			name:     "custom string slice",
			typeFunc: func() TypeInfo { return Get[customStringSlice]() },
			expected: TypeInfo{
				Name: "customStringSlice", Package: packageName, Kind: "[]string", ReflectKind: reflect.Slice,
			},
		},
		{
			name:     "json.Number",
			typeFunc: func() TypeInfo { return Get[json.Number]() },
			expected: TypeInfo{
				Name: "Number", Package: "encoding/json", Kind: "string",
				ReflectKind: reflect.String, JSONKind: reflect.Float64,
			},
		},
		{
			name:     "pointer to json.Number",
			typeFunc: func() TypeInfo { return Get[*json.Number]() },
			expected: TypeInfo{
				Name: "Number", Package: "encoding/json", Kind: "string",
				ReflectKind: reflect.String, JSONKind: reflect.Float64,
			},
		},
		{
			name:     "time.Time",
			typeFunc: func() TypeInfo { return Get[time.Time]() },
			expected: TypeInfo{
				Name: "Time", Package: "time", Kind: "struct",
				ReflectKind: reflect.Struct, JSONKind: reflect.String,
			},
		},
		{
			name:     "pointer to time.Time",
			typeFunc: func() TypeInfo { return Get[*time.Time]() },
			expected: TypeInfo{
				Name: "Time", Package: "time", Kind: "struct",
				ReflectKind: reflect.Struct, JSONKind: reflect.String,
			},
		},
		{
			name:     "JSON marshaler",
			typeFunc: func() TypeInfo { return Get[json.RawMessage]() },
			expected: TypeInfo{
				Name: "RawMessage", Package: "encoding/json", Kind: "[]uint8",
				ReflectKind: reflect.Slice, JSONKind: reflect.Interface,
			},
		},
		{
			name:     "value with pointer JSON method",
			typeFunc: func() TypeInfo { return Get[customJSONByte]() },
			expected: TypeInfo{
				Name: "customJSONByte", Package: packageName, Kind: "uint8",
				ReflectKind: reflect.Uint8, JSONKind: reflect.Interface,
			},
		},
		{
			name:     "pointer JSON marshaler",
			typeFunc: func() TypeInfo { return Get[*customJSONByte]() },
			expected: TypeInfo{
				Name: "customJSONByte", Package: packageName, Kind: "uint8",
				ReflectKind: reflect.Uint8, JSONKind: reflect.Interface,
			},
		},
		{
			name:     "value with pointer text method",
			typeFunc: func() TypeInfo { return Get[customTextByte]() },
			expected: TypeInfo{
				Name: "customTextByte", Package: packageName, Kind: "uint8",
				ReflectKind: reflect.Uint8, JSONKind: reflect.Interface,
			},
		},
		{
			name:     "pointer text marshaler",
			typeFunc: func() TypeInfo { return Get[*customTextByte]() },
			expected: TypeInfo{
				Name: "customTextByte", Package: packageName, Kind: "uint8",
				ReflectKind: reflect.Uint8, JSONKind: reflect.String,
			},
		},
		{
			name:     "byte slice",
			typeFunc: func() TypeInfo { return Get[[]byte]() },
			expected: TypeInfo{
				Name: "[]uint8", Kind: "[]uint8", ReflectKind: reflect.Slice, JSONKind: reflect.String,
			},
		},
		{
			name:     "pointer to byte slice",
			typeFunc: func() TypeInfo { return Get[*[]byte]() },
			expected: TypeInfo{
				Name: "[]uint8", Kind: "[]uint8", ReflectKind: reflect.Slice, JSONKind: reflect.String,
			},
		},
		{
			name:     "byte array",
			typeFunc: func() TypeInfo { return Get[[2]byte]() },
			expected: TypeInfo{
				Name: "[2]uint8", Kind: "array", ReflectKind: reflect.Array,
			},
		},
		{
			name:     "byte slice with JSON element methods",
			typeFunc: func() TypeInfo { return Get[[]customJSONByte]() },
			expected: TypeInfo{
				Name: "[]customJSONByte", Package: packageName, Kind: "[]uint8", ReflectKind: reflect.Slice,
			},
		},
		{
			name:     "byte slice with text element methods",
			typeFunc: func() TypeInfo { return Get[[]customTextByte]() },
			expected: TypeInfo{
				Name: "[]customTextByte", Package: packageName, Kind: "[]uint8", ReflectKind: reflect.Slice,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.typeFunc()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

type customJSONByte byte

func (b *customJSONByte) MarshalJSON() ([]byte, error) { return json.Marshal(string(rune(*b))) }

type customTextByte byte

func (b *customTextByte) MarshalText() ([]byte, error) { return []byte{byte(*b)}, nil }
