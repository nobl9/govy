package rules

import (
	"strings"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// HashFunction accepts a value and returns a comparable hash.
type HashFunction[V any, H comparable] func(v V) H

// HashFuncSelf returns a HashFunction which returns its input value as a hash itself.
// The value must be comparable.
func HashFuncSelf[H comparable]() HashFunction[H, H] {
	return func(v H) H { return v }
}

type sliceUniqueTemplateVars struct {
	Constraints   []string
	FirstOrdinal  string
	SecondOrdinal string
}

// SliceUnique ensures that a slice contains unique elements based on a provided HashFunction.
// You can optionally specify constraints which will be included in the error message to further
// clarify the reason for breaking uniqueness.
func SliceUnique[S []V, V any, H comparable](hashFunc HashFunction[V, H], constraints ...string) govy.Rule[S] {
	tpl := messagetemplates.Get(messagetemplates.SliceUniqueTemplate)

	return govy.NewRule(func(slice S) error {
		unique := make(map[H]int)
		for i := range slice {
			hash := hashFunc(slice[i])
			if j, ok := unique[hash]; ok {
				return govy.NewRuleErrorTemplate(govy.TemplateVars{
					PropertyValue: slice,
					Custom: sliceUniqueTemplateVars{
						Constraints:   constraints,
						FirstOrdinal:  ordinalString(j + 1),
						SecondOrdinal: ordinalString(i + 1),
					},
				})
			}
			unique[hash] = i
		}
		return nil
	}).
		WithErrorCode(ErrorCodeSliceUnique).
		WithMessageTemplate(tpl).
		WithDescription(func() string {
			msg := "elements must be unique"
			if len(constraints) > 0 {
				msg += " according to the following constraints: " + strings.Join(constraints, ", ")
			}
			return msg
		}())
}
