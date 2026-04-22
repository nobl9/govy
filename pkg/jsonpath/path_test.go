package jsonpath_test

import (
	"encoding/json"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/jsonpath"
)

func TestPath(t *testing.T) {
	tests := map[string]struct {
		path     jsonpath.Path
		expected string
	}{
		"empty": {
			path:     jsonpath.New(),
			expected: "",
		},
		"single name": {
			path:     jsonpath.New().Name("metadata"),
			expected: "metadata",
		},
		"two names": {
			path:     jsonpath.New().Name("metadata").Name("name"),
			expected: "metadata.name",
		},
		"name with dot": {
			path:     jsonpath.New().Name("complex.key"),
			expected: "['complex.key']",
		},
		"name then index": {
			path:     jsonpath.New().Name("metadata").Name("labels").Index(0),
			expected: "metadata.labels[0]",
		},
		"name then wildcard index": {
			path:     jsonpath.New().Name("metadata").Name("labels").IndexWildcard(),
			expected: "metadata.labels[*]",
		},
		"name then wildcard value": {
			path:     jsonpath.New().Name("metadata").ValueWildcard(),
			expected: "metadata.*",
		},
		"name then wildcard key": {
			path:     jsonpath.New().Name("metadata").KeyWildcard(),
			expected: "metadata.*",
		},
		"index only": {
			path:     jsonpath.New().Index(2),
			expected: "[2]",
		},
		"nested index": {
			path:     jsonpath.New().Index(0).Index(1),
			expected: "[0][1]",
		},
		"name after index": {
			path:     jsonpath.New().Index(0).Name("field"),
			expected: "[0].field",
		},
		"complex path with special chars": {
			path:     jsonpath.New().Name("parent").Name("foo.bar").Index(3),
			expected: "parent['foo.bar'][3]",
		},
		"root only": {
			path:     jsonpath.NewRoot(),
			expected: "$",
		},
		"root then name": {
			path:     jsonpath.NewRoot().Name("metadata"),
			expected: "$.metadata",
		},
		"name with dollar": {
			path:     jsonpath.New().Name("$"),
			expected: "['$']",
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
		"root only": {
			input:    "$",
			expected: "$",
		},
		"absolute path": {
			input:    "$.items[0]",
			expected: "$.items[0]",
		},
		"dollar in middle is a name": {
			input:    "items.$.name",
			expected: "items['$'].name",
		},
		"wildcard index": {
			input:    "items[*]",
			expected: "items[*]",
		},
		"legacy key wildcard is a literal name": {
			input:    "items.*~",
			expected: "items['*~']",
		},
		"legacy key wildcard at root is a literal name": {
			input:    "*~",
			expected: "['*~']",
		},
		"standard wildcard": {
			input:    "items.*",
			expected: "items.*",
		},
		"tilde is a literal name": {
			input:    "~",
			expected: "~",
		},
		"tilde child segment is a literal name": {
			input:    "items.~",
			expected: "items.~",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			p := jsonpath.Parse(tc.input)
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
			path:     jsonpath.New().Name("map").Key("myKey"),
			expected: "map.myKey",
		},
		"key with dot": {
			path:     jsonpath.New().Name("map").Key("complex.key"),
			expected: "map['complex.key']",
		},
		"integer key": {
			path:     jsonpath.New().Name("map").Key(42),
			expected: "map.42",
		},
		"key only": {
			path:     jsonpath.New().Key("solo"),
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
			base:     jsonpath.New(),
			other:    jsonpath.New(),
			expected: "",
		},
		"empty base": {
			base:     jsonpath.New(),
			other:    jsonpath.New().Name("child"),
			expected: "child",
		},
		"empty other": {
			base:     jsonpath.New().Name("parent"),
			other:    jsonpath.New(),
			expected: "parent",
		},
		"two simple paths": {
			base:     jsonpath.New().Name("parent"),
			other:    jsonpath.New().Name("child"),
			expected: "parent.child",
		},
		"other starts with bracket": {
			base:     jsonpath.New().Name("parent"),
			other:    jsonpath.Parse("['complex.child']"),
			expected: "parent['complex.child']",
		},
		"other starts with array index": {
			base:     jsonpath.New().Name("items"),
			other:    jsonpath.New().Index(0),
			expected: "items[0]",
		},
		"starts with root": {
			base:     jsonpath.NewRoot().Name("items"),
			other:    jsonpath.New().Name("name"),
			expected: "$.items.name",
		},
		"other starts with root": {
			base:     jsonpath.New().Name("items"),
			other:    jsonpath.NewRoot().Name("name"),
			expected: "$.name.items",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.base.Join(tc.other).String())
		})
	}
}

func TestPath_UnknownIndex(t *testing.T) {
	assert.Equal(t, "[]", jsonpath.New().UnknownIndex().String())
	assert.Equal(t, "items[]", jsonpath.New().Name("items").UnknownIndex().String())
	assert.Equal(t, "items[].name", jsonpath.New().Name("items").UnknownIndex().Name("name").String())
}

func TestPath_Compare(t *testing.T) {
	a1 := jsonpath.New().Name("a")
	a2 := jsonpath.New().Name("a")
	b := jsonpath.New().Name("b")
	empty1 := jsonpath.New()
	empty2 := jsonpath.New()

	assert.Equal(t, 0, a1.Compare(a2))
	assert.Equal(t, -1, a1.Compare(b))
	assert.Equal(t, 1, b.Compare(a1))
	assert.Equal(t, 0, empty1.Compare(empty2))
	assert.Equal(t, -1, empty1.Compare(a1))
	assert.Equal(t, 1, a1.Compare(empty1))
}

func TestPath_IsEmpty(t *testing.T) {
	assert.True(t, jsonpath.New().IsEmpty())
	assert.False(t, jsonpath.New().Name("x").IsEmpty())
	assert.True(t, jsonpath.Parse("").IsEmpty())
	assert.False(t, jsonpath.Parse("foo").IsEmpty())
}

func TestPath_MarshalText(t *testing.T) {
	type wrapper struct {
		Path jsonpath.Path `json:"path"`
	}

	t.Run("marshal", func(t *testing.T) {
		w := wrapper{Path: jsonpath.New().Name("foo").Name("bar")}
		data, err := json.Marshal(w)
		assert.NoError(t, err)
		assert.Equal(t, `{"path":"foo.bar"}`, string(data))
	})

	t.Run("marshal empty", func(t *testing.T) {
		w := wrapper{Path: jsonpath.New()}
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
		original := wrapper{Path: jsonpath.New().Name("parent").Name("foo.bar").Index(3)}
		data, err := json.Marshal(original)
		assert.NoError(t, err)

		var decoded wrapper
		err = json.Unmarshal(data, &decoded)
		assert.NoError(t, err)
		assert.Equal(t, original.Path.String(), decoded.Path.String())
	})

	t.Run("marshal bracket notation", func(t *testing.T) {
		w := wrapper{Path: jsonpath.New().Name("complex.key")}
		data, err := json.Marshal(w)
		assert.NoError(t, err)
		assert.Equal(t, `{"path":"['complex.key']"}`, string(data))
	})
}
