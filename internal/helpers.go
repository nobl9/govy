package internal

import (
	"os"
	"path/filepath"
	"reflect"
	"sync"
)

const RequiredErrorMessage = "property is required but was empty"

const RequiredErrorCodeString = "required"

// IsEmptyFunc verifies if the value is zero value of its type.
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
