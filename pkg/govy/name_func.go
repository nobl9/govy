package govy

import (
	"fmt"
	"strings"
)

type NameFunc[T any] func(value T) string

// NameFuncFromTypeName returns a
func NameFuncFromTypeName[T any]() func(value T) string {
	return func(value T) string {
		split := strings.Split(fmt.Sprintf("%T", *new(T)), ".")
		if len(split) == 0 {
			return ""
		}
		return split[len(split)-1]
	}
}
