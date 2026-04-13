package jsonpath_test

import (
	"encoding/json"
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

func TestParsePath(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"empty": {
			input:    "",
			expected: "",
		},
		"simple path": {
			input:    "foo.bar",
			expected: "foo.bar",
		},
		"bracket notation": {
			input:    "parent['foo.bar']",
			expected: "parent['foo.bar']",
		},
		"array index": {
			input:    "items[0]",
			expected: "items[0]",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			p := jsonpath.ParsePath(tc.input)
			assert.Equal(t, tc.expected, p.String())
		})
	}
}

func TestPath_Key(t *testing.T) {
	tests := map[string]struct {
		path     jsonpath.Path
		expected string
	}{
		"simple key": {
			path:     jsonpath.NewPath().Name("map").Key("myKey"),
			expected: "map.myKey",
		},
		"key with dot": {
			path:     jsonpath.NewPath().Name("map").Key("complex.key"),
			expected: "map['complex.key']",
		},
		"integer key": {
			path:     jsonpath.NewPath().Name("map").Key(42),
			expected: "map.42",
		},
		"key only": {
			path:     jsonpath.NewPath().Key("solo"),
			expected: "solo",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.path.String())
		})
	}
}

func TestPath_JoinPath(t *testing.T) {
	tests := map[string]struct {
		base     jsonpath.Path
		other    jsonpath.Path
		expected string
	}{
		"both empty": {
			base:     jsonpath.NewPath(),
			other:    jsonpath.NewPath(),
			expected: "",
		},
		"empty base": {
			base:     jsonpath.NewPath(),
			other:    jsonpath.NewPath().Name("child"),
			expected: "child",
		},
		"empty other": {
			base:     jsonpath.NewPath().Name("parent"),
			other:    jsonpath.NewPath(),
			expected: "parent",
		},
		"two simple paths": {
			base:     jsonpath.NewPath().Name("parent"),
			other:    jsonpath.NewPath().Name("child"),
			expected: "parent.child",
		},
		"other starts with bracket": {
			base:     jsonpath.NewPath().Name("parent"),
			other:    jsonpath.ParsePath("['complex.child']"),
			expected: "parent['complex.child']",
		},
		"other starts with array index": {
			base:     jsonpath.NewPath().Name("items"),
			other:    jsonpath.NewPath().Index(0),
			expected: "items[0]",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.base.JoinPath(tc.other).String())
		})
	}
}

func TestPath_UnknownIndex(t *testing.T) {
	assert.Equal(t, "[]", jsonpath.NewPath().UnknownIndex().String())
	assert.Equal(t, "items[]", jsonpath.NewPath().Name("items").UnknownIndex().String())
	assert.Equal(t, "items[].name", jsonpath.NewPath().Name("items").UnknownIndex().Name("name").String())
}

func TestPath_Compare(t *testing.T) {
	a1 := jsonpath.NewPath().Name("a")
	a2 := jsonpath.NewPath().Name("a")
	b := jsonpath.NewPath().Name("b")
	empty1 := jsonpath.NewPath()
	empty2 := jsonpath.NewPath()

	assert.Equal(t, 0, a1.Compare(a2))
	assert.Equal(t, -1, a1.Compare(b))
	assert.Equal(t, 1, b.Compare(a1))
	assert.Equal(t, 0, empty1.Compare(empty2))
	assert.Equal(t, -1, empty1.Compare(a1))
	assert.Equal(t, 1, a1.Compare(empty1))
}

func TestPath_IsEmpty(t *testing.T) {
	assert.True(t, jsonpath.NewPath().IsEmpty())
	assert.False(t, jsonpath.NewPath().Name("x").IsEmpty())
	assert.True(t, jsonpath.ParsePath("").IsEmpty())
	assert.False(t, jsonpath.ParsePath("foo").IsEmpty())
}

func TestPath_MarshalText(t *testing.T) {
	type wrapper struct {
		Path jsonpath.Path `json:"path"`
	}

	t.Run("marshal", func(t *testing.T) {
		w := wrapper{Path: jsonpath.NewPath().Name("foo").Name("bar")}
		data, err := json.Marshal(w)
		assert.NoError(t, err)
		assert.Equal(t, `{"path":"foo.bar"}`, string(data))
	})

	t.Run("marshal empty", func(t *testing.T) {
		w := wrapper{Path: jsonpath.NewPath()}
		data, err := json.Marshal(w)
		assert.NoError(t, err)
		assert.Equal(t, `{"path":""}`, string(data))
	})

	t.Run("unmarshal", func(t *testing.T) {
		var w wrapper
		err := json.Unmarshal([]byte(`{"path":"foo.bar[0]"}`), &w)
		assert.NoError(t, err)
		assert.Equal(t, "foo.bar[0]", w.Path.String())
	})

	t.Run("round trip", func(t *testing.T) {
		original := wrapper{Path: jsonpath.NewPath().Name("parent").Name("foo.bar").Index(3)}
		data, err := json.Marshal(original)
		assert.NoError(t, err)

		var decoded wrapper
		err = json.Unmarshal(data, &decoded)
		assert.NoError(t, err)
		assert.Equal(t, original.Path.String(), decoded.Path.String())
	})

	t.Run("marshal bracket notation", func(t *testing.T) {
		w := wrapper{Path: jsonpath.NewPath().Name("complex.key")}
		data, err := json.Marshal(w)
		assert.NoError(t, err)
		assert.Equal(t, `{"path":"['complex.key']"}`, string(data))
	})
}
