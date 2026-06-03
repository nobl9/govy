package govy

import (
	"fmt"
	"strings"
)

// NameFunc is a function that produces a descriptive name for a given value.
type NameFunc[T any] func(value T) string

// NameFuncFromTypeName returns a [NameFunc] that produces the name of type T
// (without its package path) for any provided value.
func NameFuncFromTypeName[T any]() NameFunc[T] {
	return func(value T) string {
		split := strings.Split(fmt.Sprintf("%T", *new(T)), ".")
		if len(split) == 0 {
			return ""
		}
		return split[len(split)-1]
	}
}
