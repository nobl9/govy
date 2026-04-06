package govy

import (
	"strconv"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonpath"
)

func TestPropertyError_prependPropertyPath(t *testing.T) {
	tests := []struct {
		PropertyError *PropertyError
		InputName     Path
		ExpectedName  string
	}{
		{
			PropertyError: &PropertyError{},
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.ParsePath("test")},
			ExpectedName:  "test",
		},
		{
			PropertyError: &PropertyError{},
			InputName:     jsonpath.ParsePath("new"),
			ExpectedName:  "new",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.ParsePath("original")},
			InputName:     jsonpath.ParsePath("added"),
			ExpectedName:  "added.original",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.ParsePath("bar"), IsSliceElementError: true},
			InputName:     jsonpath.ParsePath("foo[1]"),
			ExpectedName:  "foo[1].bar",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.ParsePath("[2]"), IsSliceElementError: true},
			InputName:     jsonpath.ParsePath("foo"),
			ExpectedName:  "foo[2]",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.ParsePath("foo"), IsSliceElementError: true},
			InputName:     jsonpath.ParsePath("[0]"),
			ExpectedName:  "[0].foo",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.ParsePath("[1]"), IsSliceElementError: true},
			InputName:     jsonpath.ParsePath("[0]"),
			ExpectedName:  "[0][1]",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.ParsePath("['foo.bar']")},
			InputName:     jsonpath.ParsePath("parent"),
			ExpectedName:  "parent['foo.bar']",
		},
		{
			PropertyError: &PropertyError{PropertyPath: jsonpath.ParsePath("child")},
			InputName:     jsonpath.ParsePath("['complex.parent']"),
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
