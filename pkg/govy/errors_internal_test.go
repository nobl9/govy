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
		InputName     jsonpath.Path
		ExpectedName  string
	}{
		{
			PropertyError: &PropertyError{},
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.Parse("test")},
			ExpectedName:  "test",
		},
		{
			PropertyError: &PropertyError{},
			InputName:     jsonpath.Parse("new"),
			ExpectedName:  "new",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.Parse("original")},
			InputName:     jsonpath.Parse("added"),
			ExpectedName:  "added.original",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.Parse("bar"), IsSliceElementError: true},
			InputName:     jsonpath.Parse("foo[1]"),
			ExpectedName:  "foo[1].bar",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.Parse("[2]"), IsSliceElementError: true},
			InputName:     jsonpath.Parse("foo"),
			ExpectedName:  "foo[2]",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.Parse("foo"), IsSliceElementError: true},
			InputName:     jsonpath.Parse("[0]"),
			ExpectedName:  "[0].foo",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.Parse("[1]"), IsSliceElementError: true},
			InputName:     jsonpath.Parse("[0]"),
			ExpectedName:  "[0][1]",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.Parse("['foo.bar']")},
			InputName:     jsonpath.Parse("parent"),
			ExpectedName:  "parent['foo.bar']",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.Parse("child")},
			InputName:     jsonpath.Parse("['complex.parent']"),
			ExpectedName:  "['complex.parent'].child",
		},
	}

	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := tc.PropertyError.prependParentPropertyPath(tc.InputName)
			assert.Equal(t, tc.ExpectedName, result.PropertyPath.String())
		})
	}
}
