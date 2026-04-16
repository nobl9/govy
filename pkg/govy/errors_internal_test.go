package govy

import (
	"strconv"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/jsonpath"
)

func TestPropertyError_prependPropertyPath(t *testing.T) {
	tests := []struct {
		PropertyError *PropertyError
		InputPath     jsonpath.Path
		ExpectedPath  string
	}{
		{
			PropertyError: &PropertyError{},
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.Parse("test")},
			ExpectedPath:  "test",
		},
		{
			PropertyError: &PropertyError{},
			InputPath:     jsonpath.Parse("new"),
			ExpectedPath:  "new",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.Parse("original")},
			InputPath:     jsonpath.Parse("added"),
			ExpectedPath:  "added.original",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.Parse("bar"), IsSliceElementError: true},
			InputPath:     jsonpath.Parse("foo[1]"),
			ExpectedPath:  "foo[1].bar",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.Parse("[2]"), IsSliceElementError: true},
			InputPath:     jsonpath.Parse("foo"),
			ExpectedPath:  "foo[2]",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.Parse("foo"), IsSliceElementError: true},
			InputPath:     jsonpath.Parse("[0]"),
			ExpectedPath:  "[0].foo",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.Parse("[1]"), IsSliceElementError: true},
			InputPath:     jsonpath.Parse("[0]"),
			ExpectedPath:  "[0][1]",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.Parse("['foo.bar']")},
			InputPath:     jsonpath.Parse("parent"),
			ExpectedPath:  "parent['foo.bar']",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.Parse("child")},
			InputPath:     jsonpath.Parse("['complex.parent']"),
			ExpectedPath:  "['complex.parent'].child",
		},
	}

	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := tc.PropertyError.prependParentPropertyPath(tc.InputPath)
			assert.Equal(t, tc.ExpectedPath, result.PropertyPath.String())
		})
	}
}
