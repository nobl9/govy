package typeinfo

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
)

const packageName = "github.com/nobl9/govy/internal/typeinfo"

type customString string

type customStruct struct{}

type customMap map[string]int

type customList []customMap

type customNestedMap map[customString]customList

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
			expected: TypeInfo{Name: "int", Package: "", Kind: "int"},
		},
		{
			name:     "pointer to int",
			typeFunc: func() TypeInfo { return Get[*int]() },
			expected: TypeInfo{Name: "int", Package: "", Kind: "int"},
		},
		{
			name:     "slice of int",
			typeFunc: func() TypeInfo { return Get[[]int]() },
			expected: TypeInfo{Name: "[]int", Package: "", Kind: "[]int"},
		},
		{
			name:     "slice of customString",
			typeFunc: func() TypeInfo { return Get[[]customString]() },
			expected: TypeInfo{Name: "[]customString", Package: packageName, Kind: "[]string"},
		},
		{
			name:     "map of string to int",
			typeFunc: func() TypeInfo { return Get[map[string]int]() },
			expected: TypeInfo{Name: "map[string]int", Package: "", Kind: "map[string]int"},
		},
		{
			name:     "custom string",
			typeFunc: func() TypeInfo { return Get[customString]() },
			expected: TypeInfo{Name: "customString", Package: packageName, Kind: "string"},
		},
		{
			name:     "custom struct",
			typeFunc: func() TypeInfo { return Get[customStruct]() },
			expected: TypeInfo{Name: "customStruct", Package: packageName, Kind: "struct"},
		},
		{
			name:     "pointer to custom struct",
			typeFunc: func() TypeInfo { return Get[*customStruct]() },
			expected: TypeInfo{Name: "customStruct", Package: packageName, Kind: "struct"},
		},
		{
			name:     "custom map",
			typeFunc: func() TypeInfo { return Get[customMap]() },
			expected: TypeInfo{Name: "customMap", Package: packageName, Kind: "map[string]int"},
		},
		{
			name:     "custom nested map",
			typeFunc: func() TypeInfo { return Get[customNestedMap]() },
			expected: TypeInfo{Name: "customNestedMap", Package: packageName, Kind: "map[string][]map[string]int"},
		},
		{
			name:     "custom list",
			typeFunc: func() TypeInfo { return Get[customList]() },
			expected: TypeInfo{Name: "customList", Package: packageName, Kind: "[]map[string]int"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.typeFunc()
			assert.Equal(t, tc.expected, actual)
		})
	}
}
