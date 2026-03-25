package jsonpath_test

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonpath"
)

func TestPath(t *testing.T) {
	tests := map[string]struct {
		path     jsonpath.Path
		expected string
	}{
		"empty": {
			path:     jsonpath.NewPath(),
			expected: "",
		},
		"single name": {
			path:     jsonpath.NewPath().Name("metadata"),
			expected: "metadata",
		},
		"two names": {
			path:     jsonpath.NewPath().Name("metadata").Name("name"),
			expected: "metadata.name",
		},
		"name with dot": {
			path:     jsonpath.NewPath().Name("complex.key"),
			expected: "['complex.key']",
		},
		"name then index": {
			path:     jsonpath.NewPath().Name("metadata").Name("labels").Index(0),
			expected: "metadata.labels[0]",
		},
		"index only": {
			path:     jsonpath.NewPath().Index(2),
			expected: "[2]",
		},
		"nested index": {
			path:     jsonpath.NewPath().Index(0).Index(1),
			expected: "[0][1]",
		},
		"name after index": {
			path:     jsonpath.NewPath().Index(0).Name("field"),
			expected: "[0].field",
		},
		"complex path with special chars": {
			path:     jsonpath.NewPath().Name("parent").Name("foo.bar").Index(3),
			expected: "parent['foo.bar'][3]",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.path.String())
		})
	}
}
