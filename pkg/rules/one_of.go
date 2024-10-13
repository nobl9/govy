package rules

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"golang.org/x/exp/maps"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/pkg/govy"
)

// OneOf checks if the property's value matches one of the provided values.
// The values must be comparable.
func OneOf[T comparable](values ...T) govy.Rule[T] {
	return govy.NewRule(func(v T) error {
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

// OneOfProperties checks if at least one of the properties is set.
// Property is considered set if its value is not empty (non-zero).
func OneOfProperties[S any](getters map[string]func(s S) any) govy.Rule[S] {
	return govy.NewRule(func(s S) error {
		for _, getter := range getters {
			v := getter(s)
			if !internal.IsEmpty(v) {
				return nil
			}
		}
		keys := maps.Keys(getters)
		slices.Sort(keys)
		return fmt.Errorf(
			"one of %s properties must be set, none was provided",
			prettyOneOfList(keys))
	}).
		WithErrorCode(ErrorCodeOneOfProperties).
		WithDescription(func() string {
			keys := maps.Keys(getters)
			return fmt.Sprintf("at least one of the properties must be set: %s", strings.Join(keys, ", "))
		}())
}

// MutuallyExclusive checks if properties are mutually exclusive.
// This means, exactly one of the properties can be set.
// Property is considered set if its value is not empty (non-zero).
// If required is true, then a single non-empty property is required.
func MutuallyExclusive[S any](required bool, getters map[string]func(s S) any) govy.Rule[S] {
	return govy.NewRule(func(s S) error {
		var nonEmpty []string
		for name, getter := range getters {
			v := getter(s)
			if internal.IsEmpty(v) {
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
			return fmt.Errorf(
				"one of %s properties must be set, none was provided",
				prettyOneOfList(keys))
		case 1:
			return nil
		default:
			slices.Sort(nonEmpty)
			return fmt.Errorf(
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
