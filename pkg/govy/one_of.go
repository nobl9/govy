package validation

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func OneOf[T comparable](values ...T) SingleRule[T] {
	return NewSingleRule(func(v T) error {
		for i := range values {
			if v == values[i] {
				return nil
			}
		}
		return errors.New("must be one of " + prettyOneOfList(values))
	}).
		WithErrorCode(ErrorCodeOneOf).
		WithDescription(func() string {
			b := strings.Builder{}
			prettyStringListBuilder(&b, values, false)
			return "must be one of: " + b.String()
		}())
}

// MutuallyExclusive checks if properties are mutually exclusive.
// This means, exactly one of the properties can be provided.
// If required is true, then a single non-empty property is required.
func MutuallyExclusive[S any](required bool, getters map[string]func(s S) any) SingleRule[S] {
	return NewSingleRule(func(s S) error {
		var nonEmpty []string
		for name, getter := range getters {
			v := getter(s)
			if isEmptyFunc(v) {
				continue
			}
			nonEmpty = append(nonEmpty, name)
		}
		switch len(nonEmpty) {
		case 0:
			if !required {
				return nil
			}
			keys := maps.Keys(getters)
			slices.Sort(keys)
			return errors.Errorf(
				"one of %s properties must be set, none was provided",
				prettyOneOfList(keys))
		case 1:
			return nil
		default:
			slices.Sort(nonEmpty)
			return errors.Errorf(
				"%s properties are mutually exclusive, provide only one of them",
				prettyOneOfList(nonEmpty))
		}
	}).
		WithErrorCode(ErrorCodeMutuallyExclusive).
		WithDescription(func() string {
			keys := maps.Keys(getters)
			return fmt.Sprintf("properties are mutually exclusive: %s", strings.Join(keys, ", "))
		}())
}

func prettyOneOfList[T any](values []T) string {
	b := strings.Builder{}
	b.Grow(2 + len(values))
	b.WriteString("[")
	prettyStringListBuilder(&b, values, false)
	b.WriteString("]")
	return b.String()
}
