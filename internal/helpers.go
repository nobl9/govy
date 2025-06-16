package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
)

const (
	RequiredErrorMessage    = "property is required but was empty"
	RequiredDescription     = "property is required"
	RequiredErrorCodeString = "required"
)

// IsEmpty verifies if the value is zero value of its type.
func IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	return rv.Kind() == 0 || rv.IsZero()
}

var (
	moduleRoot string
	once       sync.Once
)

// FindModuleRoot finds the root of the current module.
// It does so by looking for a go.mod file in the current working directory.
func FindModuleRoot() string {
	once.Do(func() {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		dir = filepath.Clean(dir)
		for {
			if fi, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil && !fi.IsDir() {
				moduleRoot = dir
				return
			}
			d := filepath.Dir(dir)
			if d == dir {
				break
			}
			dir = d
		}
	})
	return moduleRoot
}

// PrettyStringListBuilder writes a list of arbitrary values to the provided [strings.Builder].
// It produces a human-readable comma-separated list.
// Example:
//
//	PrettyStringListBuilder(b, []string{"foo", "bar"}, "") -> "foo, bar"
//	PrettyStringListBuilder(b, []string{"foo", "bar"}, "'") -> "'foo', 'bar'"
func PrettyStringListBuilder[T any](b *strings.Builder, values []T, surroundingStr string) {
	b.Grow(len(values))
	for i := range values {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(surroundingStr)
		fmt.Fprint(b, values[i])
		b.WriteString(surroundingStr)
	}
}
