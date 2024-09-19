package rules

import (
	"net/url"
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

var validURLs = []string{
	"http://foo.bar#com",
	"http://foobar.com",
	"https://foobar.com",
	"http://foobar.coffee/",
	"http://foobar.中文网/",
	"http://foobar.org/",
	"http://foobar.org:8080/",
	"ftp://foobar.ua/",
	"http://user:pass@www.foobar.com/",
	"http://127.0.0.1/",
	"http://duckduckgo.com/?q=%2F",
	"http://localhost:3000/",
	"http://foobar.com/?foo=bar#baz=qux",
	"http://foobar.com?foo=bar",
	"http://www.xn--froschgrn-x9a.net/",
	"xyz://foobar.com",
	"rtmp://foobar.com",
	"http://www.foo_bar.com/",
	"http://localhost:3000/",
	"http://foobar.com/#baz",
	"http://foobar.com#baz=qux",
	"http://foobar.com/t$-_.+!*\\'(),",
	"http://www.foobar.com/~foobar",
	"http://www.-foobar.com/",
	"http://www.foo---bar.com/",
	"mailto:someone@example.com",
	"irc://irc.server.org/channel",
	"irc://#channel@network",
}

var invalidURLs = []string{
	"foobar.com",
	"",
	"invalid.",
	".com",
	"/abs/test/dir",
	"./rel/test/dir",
	"irc:",
	"http://",
}

func TestURL(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		for _, input := range validURLs {
			u, err := url.Parse(input)
			assert.Require(t, assert.NoError(t, err))
			err = URL().Validate(u)
			assert.NoError(t, err)
		}
	})
	t.Run("fails", func(t *testing.T) {
		for _, input := range invalidURLs {
			u, err := url.Parse(input)
			assert.Require(t, assert.NoError(t, err))
			err = URL().Validate(u)
			assert.Require(t, assert.Error(t, err))
			assert.True(t, govy.HasErrorCode(err, ErrorCodeURL))
		}
	})
}
