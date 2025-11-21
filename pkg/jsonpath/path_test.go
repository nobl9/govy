package jsonpath_test

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/jsonpath"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{"simple property path", "$.spec.name", "$.spec.name"},
		{"path with index", "$.items[0].name", "$.items[0].name"},
		{"path with quoted property", "$.['property with spaces'].value", "$['property with spaces'].value"},
		{"complex path", "$.spec.['this is'][0].foo", "$.spec['this is'][0].foo"},
		{"partial path without root", "spec.name", "spec.name"},
		{"partial path with index", "items[0].name", "items[0].name"},
		{"partial path starting with bracket", "[0].name", "[0].name"},
		{"property name with spaces", "$.['property with spaces'].value", "$['property with spaces'].value"},
		{"property name with brackets", "$.['property[0]'].value", "$['property[0]'].value"},
		{"property with quotes", `$.['can\'t'].value`, `$['can\'t'].value`},
		{"property with dots", "$.['foo.bar'].value", "$['foo.bar'].value"},
		{"simple identifiers not escaped", "$.simple_property.another-one", "$.simple_property.another-one"},
		{"root only", "$", "$"},
		{"large index", "$[999999].foo", "$[999999].foo"},
		{"zero index", "$[0]", "$[0]"},
		{"multiple consecutive indices", "$[0][1][2]", "$[0][1][2]"},
		{"property with tab character", `$.['has\ttab'].value`, `$['has\ttab'].value`},
		{"property with newline character", `$.['has\nline'].value`, `$['has\nline'].value`},
		{"property with carriage return", `$.['has\rreturn'].value`, `$['has\rreturn'].value`},
		{"empty property name", "$.['']", "$['']"},
		{"numeric property names", "$.['123'].value", "$.123.value"},
		{"property name starting with underscore", "$._private._internal", "$._private._internal"},
		{"mixed bracket and dot notation", "$.foo['bar'].baz[0]['qux']", "$.foo.bar.baz[0].qux"},
		{"unicode in property name", "$.['日本語']", "$.日本語"},
		{"property with all special chars", `$.['a b.c[d]e\'f\ng\th\ri']`, `$['a b.c[d]e\'f\ng\th\ri']`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			path, err := jsonpath.Parse(tc.input)
			assert.NoError(t, err)
			assert.Equal(t, tc.output, path.String())
		})
	}

	t.Run("property with backslash", func(t *testing.T) {
		path, err := jsonpath.Parse(`$.['back\\slash'].value`)
		assert.NoError(t, err)
		segments := path.Segments()
		prop, ok := segments[1].(jsonpath.PropertySegment)
		assert.True(t, ok)
		assert.Equal(t, `back\slash`, prop.Name)
	})

	t.Run("property with multiple quotes", func(t *testing.T) {
		path, err := jsonpath.Parse(`$['\'\'\'test\'\'\''].value`)
		assert.NoError(t, err)
		segments := path.Segments()
		prop, ok := segments[1].(jsonpath.PropertySegment)
		assert.True(t, ok)
		assert.Equal(t, "'''test'''", prop.Name)
	})

	t.Run("root only has single segment", func(t *testing.T) {
		path, err := jsonpath.Parse("$")
		assert.NoError(t, err)
		segments := path.Segments()
		assert.Require(t, assert.Len(t, segments, 1))
		assert.Equal(t, jsonpath.TypeRoot, segments[0].Type())
	})
}

func TestParseErrors(t *testing.T) {
	errorCases := []struct {
		name        string
		input       string
		expectedErr string
	}{
		{"empty path", "", "empty path"},
		{"unclosed bracket", "$[0", "unclosed bracket"},
		{"unclosed quoted string", "$.['foo", "unclosed quoted string"},
		{"missing closing bracket after quote", "$.['foo'", "expected ] after closing quote"},
		{"invalid bracket content", "$[abc]", "invalid bracket content: abc"},
		{"bracket with unquoted string", "$[foo]", "invalid bracket content: foo"},
		{"float index", "$[1.5]", "invalid bracket content: 1.5"},
	}

	for _, tc := range errorCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := jsonpath.Parse(tc.input)
			assert.Error(t, err)
			assert.EqualError(t, err, tc.expectedErr)
		})
	}

	t.Run("negative index is parsed as index", func(t *testing.T) {
		path, err := jsonpath.Parse("$[-1]")
		assert.NoError(t, err)
		assert.Equal(t, "$[-1]", path.String())
		segments := path.Segments()
		idx, ok := segments[1].(jsonpath.IndexSegment)
		assert.True(t, ok)
		assert.Equal(t, -1, idx.Index)
	})

	t.Run("path ending with dot", func(t *testing.T) {
		path, err := jsonpath.Parse("$.foo.")
		assert.NoError(t, err)
		assert.Equal(t, "$.foo['']", path.String())
	})

	t.Run("double dot", func(t *testing.T) {
		path, err := jsonpath.Parse("$.foo..bar")
		assert.NoError(t, err)
		assert.Equal(t, "$.foo[''].bar", path.String())
	})
}

func TestParse_RoundTrip(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{"simple path", "$.foo.bar", "$.foo.bar"},
		{"with index", "$.foo[0].bar", "$.foo[0].bar"},
		{"with escaped property", "$.['foo bar'].baz", "$['foo bar'].baz"},
		{"complex", "$.a['b c'][0].d['e.f']", "$.a['b c'][0].d['e.f']"},
		{"root only", "$", "$"},
		{"partial path", "foo.bar", "foo.bar"},
		{"index only", "[0]", "[0]"},
		{"unicode", "$.['日本語'].value", "$.日本語.value"},
		{"escaped quotes", `$.['can\'t'].do`, `$['can\'t'].do`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			path1, err := jsonpath.Parse(tc.input)
			assert.NoError(t, err)
			str1 := path1.String()
			assert.Equal(t, tc.want, str1)

			path2, err := jsonpath.Parse(str1)
			assert.NoError(t, err)
			str2 := path2.String()
			assert.Equal(t, str1, str2)
		})
	}
}

func TestSegmentType(t *testing.T) {
	testCases := []struct {
		name         string
		segment      jsonpath.Segment
		expectedType jsonpath.Type
	}{
		{"RootSegment", jsonpath.RootSegment{}, jsonpath.TypeRoot},
		{"PropertySegment", jsonpath.PropertySegment{Name: "foo"}, jsonpath.TypeProperty},
		{"IndexSegment", jsonpath.IndexSegment{Index: 5}, jsonpath.TypeIndex},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedType, tc.segment.Type())
		})
	}

	t.Run("check types from parsed path", func(t *testing.T) {
		path, err := jsonpath.Parse("$.foo[0].bar")
		assert.NoError(t, err)

		segments := path.Segments()
		assert.Require(t, assert.Len(t, segments, 4))
		assert.Equal(t, jsonpath.TypeRoot, segments[0].Type())
		assert.Equal(t, jsonpath.TypeProperty, segments[1].Type())
		assert.Equal(t, jsonpath.TypeIndex, segments[2].Type())
		assert.Equal(t, jsonpath.TypeProperty, segments[3].Type())
	})
}

func TestSegmentString(t *testing.T) {
	t.Run("RootSegment", func(t *testing.T) {
		seg := jsonpath.RootSegment{}
		assert.Equal(t, "$", seg.String())
	})

	propertyTestCases := []struct {
		name     string
		propName string
		expected string
	}{
		{"simple", "foo", "foo"},
		{"with spaces", "foo bar", "['foo bar']"},
		{"with dot", "foo.bar", "['foo.bar']"},
		{"with brackets", "foo[0]", "['foo[0]']"},
		{"with quotes", "can't", `['can\'t']`},
		{"with newline", "foo\nbar", `['foo\nbar']`},
		{"with tab", "foo\tbar", `['foo\tbar']`},
		{"with carriage return", "foo\rbar", `['foo\rbar']`},
		{"empty", "", "['']"},
		{"with underscore", "_private", "_private"},
		{"with hyphen", "foo-bar", "foo-bar"},
	}

	for _, tc := range propertyTestCases {
		t.Run("PropertySegment "+tc.name, func(t *testing.T) {
			seg := jsonpath.PropertySegment{Name: tc.propName}
			assert.Equal(t, tc.expected, seg.String())
		})
	}

	indexTestCases := []struct {
		name     string
		index    int
		expected string
	}{
		{"zero", 0, "[0]"},
		{"positive", 42, "[42]"},
		{"large", 999999, "[999999]"},
	}

	for _, tc := range indexTestCases {
		t.Run("IndexSegment "+tc.name, func(t *testing.T) {
			seg := jsonpath.IndexSegment{Index: tc.index}
			assert.Equal(t, tc.expected, seg.String())
		})
	}
}

func TestSegments(t *testing.T) {
	t.Run("access individual segments", func(t *testing.T) {
		path, err := jsonpath.Parse("$.foo['bar'][0].baz")
		assert.NoError(t, err)

		segments := path.Segments()
		assert.Require(t, assert.Len(t, segments, 5))

		root, ok := segments[0].(jsonpath.RootSegment)
		assert.True(t, ok)
		assert.Equal(t, "$", root.String())

		prop1, ok := segments[1].(jsonpath.PropertySegment)
		assert.True(t, ok)
		assert.Equal(t, "foo", prop1.Name)

		prop2, ok := segments[2].(jsonpath.PropertySegment)
		assert.True(t, ok)
		assert.Equal(t, "bar", prop2.Name)

		idx, ok := segments[3].(jsonpath.IndexSegment)
		assert.True(t, ok)
		assert.Equal(t, 0, idx.Index)

		prop3, ok := segments[4].(jsonpath.PropertySegment)
		assert.True(t, ok)
		assert.Equal(t, "baz", prop3.Name)
	})

	t.Run("empty segments for root only", func(t *testing.T) {
		path, err := jsonpath.Parse("$")
		assert.NoError(t, err)

		segments := path.Segments()
		assert.Require(t, assert.Len(t, segments, 1))
		assert.Equal(t, jsonpath.TypeRoot, segments[0].Type())
	})
}

func TestAppend(t *testing.T) {
	t.Run("append single property segment", func(t *testing.T) {
		path, err := jsonpath.Parse("$.foo")
		assert.NoError(t, err)

		newPath := path.Append(jsonpath.PropertySegment{Name: "bar"})
		assert.Equal(t, "$.foo.bar", newPath.String())
	})

	t.Run("append single index segment", func(t *testing.T) {
		path, err := jsonpath.Parse("$.foo")
		assert.NoError(t, err)

		newPath := path.Append(jsonpath.IndexSegment{Index: 0})
		assert.Equal(t, "$.foo[0]", newPath.String())
	})

	t.Run("append multiple segments", func(t *testing.T) {
		path, err := jsonpath.Parse("$.foo")
		assert.NoError(t, err)

		newPath := path.Append(
			jsonpath.PropertySegment{Name: "bar"},
			jsonpath.IndexSegment{Index: 0},
			jsonpath.PropertySegment{Name: "baz"},
		)
		assert.Equal(t, "$.foo.bar[0].baz", newPath.String())
	})

	t.Run("append to root only", func(t *testing.T) {
		path, err := jsonpath.Parse("$")
		assert.NoError(t, err)

		newPath := path.Append(jsonpath.PropertySegment{Name: "foo"})
		assert.Equal(t, "$.foo", newPath.String())
	})

	t.Run("append to partial path", func(t *testing.T) {
		path, err := jsonpath.Parse("foo")
		assert.NoError(t, err)

		newPath := path.Append(jsonpath.PropertySegment{Name: "bar"})
		assert.Equal(t, "foo.bar", newPath.String())
	})

	t.Run("original path unchanged", func(t *testing.T) {
		path, err := jsonpath.Parse("$.foo")
		assert.NoError(t, err)
		original := path.String()

		newPath := path.Append(jsonpath.PropertySegment{Name: "bar"})
		assert.Equal(t, "$.foo", path.String())
		assert.Equal(t, original, path.String())
		assert.Equal(t, "$.foo.bar", newPath.String())
	})

	t.Run("append with special characters", func(t *testing.T) {
		path, err := jsonpath.Parse("$.foo")
		assert.NoError(t, err)

		newPath := path.Append(jsonpath.PropertySegment{Name: "bar baz"})
		assert.Equal(t, "$.foo['bar baz']", newPath.String())
	})

	t.Run("append no segments", func(t *testing.T) {
		path, err := jsonpath.Parse("$.foo")
		assert.NoError(t, err)

		newPath := path.Append()
		assert.Equal(t, "$.foo", newPath.String())
		assert.Require(t, assert.Len(t, newPath.Segments(), 2))
	})

	t.Run("chain multiple appends", func(t *testing.T) {
		path, err := jsonpath.Parse("$")
		assert.NoError(t, err)

		newPath := path.
			Append(jsonpath.PropertySegment{Name: "foo"}).
			Append(jsonpath.IndexSegment{Index: 0}).
			Append(jsonpath.PropertySegment{Name: "bar"})

		assert.Equal(t, "$.foo[0].bar", newPath.String())
	})
}
