package govy

import (
	"strconv"
	"testing"

	"github.com/nobl9/govy/internal/assert"
)

func TestPropertyError_prependPropertyPath(t *testing.T) {
	tests := []struct {
		PropertyError *PropertyError
		InputName     string
		ExpectedName  string
	}{
		{
			PropertyError: &PropertyError{},
		},
		{
			PropertyError: &PropertyError{PropertyPath: "test"},
			ExpectedName:  "test",
		},
		{
			PropertyError: &PropertyError{},
			InputName:     "new",
			ExpectedName:  "new",
		},
		{
			PropertyError: &PropertyError{PropertyPath: "original"},
			InputName:     "added",
			ExpectedName:  "added.original",
		},
		{
			PropertyError: &PropertyError{PropertyPath: "bar", IsSliceElementError: true},
			InputName:     "foo[1]",
			ExpectedName:  "foo[1].bar",
		},
		{
			PropertyError: &PropertyError{PropertyPath: "[2]", IsSliceElementError: true},
			InputName:     "foo",
			ExpectedName:  "foo[2]",
		},
		{
			PropertyError: &PropertyError{PropertyPath: "foo", IsSliceElementError: true},
			InputName:     "[0]",
			ExpectedName:  "[0].foo",
		},
		{
			PropertyError: &PropertyError{PropertyPath: "[1]", IsSliceElementError: true},
			InputName:     "[0]",
			ExpectedName:  "[0][1]",
		},
	}

	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tc.ExpectedName, tc.PropertyError.prependParentPropertyPath(tc.InputName).PropertyPath)
		})
	}
}
