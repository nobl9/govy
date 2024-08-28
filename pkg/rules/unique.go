package rules

import (
	"errors"
	"fmt"
	"strings"

	"github.com/nobl9/govy/pkg/govy"
)

// HashFunction accepts a value and returns a comparable hash.
type HashFunction[V any, H comparable] func(v V) H

// HashFuncSelf returns a HashFunction which returns its input value as a hash itself.
// The value must be comparable.
func HashFuncSelf[H comparable]() HashFunction[H, H] {
	return func(v H) H { return v }
}

// SliceUnique ensures that a slice contains unique elements based on a provided HashFunction.
// You can optionally specify constraints which will be included in the error message to further
// clarify the reason for breaking uniqueness.
func SliceUnique[S []V, V any, H comparable](hashFunc HashFunction[V, H], constraints ...string) govy.Rule[S] {
	return govy.NewRule(func(slice S) error {
		unique := make(map[H]int)
		for i := range slice {
			hash := hashFunc(slice[i])
			if j, ok := unique[hash]; ok {
				errMsg := fmt.Sprintf("elements are not unique, index %d collides with index %d", j, i)
				if len(constraints) > 0 {
					errMsg += " based on constraints: " + strings.Join(constraints, ", ")
				}
				return errors.New(errMsg)
			}
			unique[hash] = i
		}
		return nil
	}).
		WithErrorCode(ErrorCodeSliceUnique).
		WithDescription(func() string {
			msg := "elements must be unique"
			if len(constraints) > 0 {
				msg += " according to the following constraints: " + strings.Join(constraints, ", ")
			}
			return msg
		}())
}
