package rules

import (
	"regexp"
	"sync"
)

// nolint: lll
// Define all regular expressions here:
var (
	validUUIDRegexp = lazyRegexCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	asciiRegexp     = lazyRegexCompile(`^[\x00-\x7F]*$`)
)

// lazyRegexCompile returns a function that compiles the regular expression
// once, when the function is called for the first time.
// If the function is never called, the regular expression is never compiled,
// thus saving on performance.
//
// All regular expression literals should be compiled using this function.
//
// Credits: https://github.com/go-playground/validator/commit/2e1df48b5ab876bdd461bdccc51d109389e7572f
func lazyRegexCompile(str string) func() *regexp.Regexp {
	var (
		regex *regexp.Regexp
		once  sync.Once
	)
	return func() *regexp.Regexp {
		once.Do(func() {
			regex = regexp.MustCompile(str)
		})
		return regex
	}
}
