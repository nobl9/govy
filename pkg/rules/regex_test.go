package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLazyRegexCompile(t *testing.T) {
	lazyRegexp := lazyRegexCompile("^test$")

	re1 := lazyRegexp()
	assert.True(t, re1.MatchString("test"))
	re2 := lazyRegexp()
	assert.True(t, re2.MatchString("test"))

	assert.True(t, re1 == re2, "both regular expression must be the same pointer")
}
