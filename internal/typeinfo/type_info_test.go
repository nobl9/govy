package typeinfo

import (
	"reflect"
	"testing"

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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.typeFunc()
			assert.Equal(t, tc.expected, actual)
		})
	}
}
