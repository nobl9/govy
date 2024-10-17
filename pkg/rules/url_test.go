package rules

import (
	"net/url"
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

var urlTestCases = []*struct {
	url        string
	shouldFail bool
}{
	{"http://foo.bar#com", false},
	{"http://foobar.com", false},
	{"https://foobar.com", false},
	{"http://foobar.coffee/", false},
	{"http://foobar.中文网/", false},
	{"http://foobar.org/", false},
	{"http://foobar.org:8080/", false},
	{"ftp://foobar.ua/", false},
	{"http://user:pass@www.foobar.com/", false},
	{"http://127.0.0.1/", false},
	{"http://duckduckgo.com/?q=%2F", false},
	{"http://localhost:3000/", false},
	{"http://foobar.com/?foo=bar#baz=qux", false},
	{"http://foobar.com?foo=bar", false},
	{"http://www.xn--froschgrn-x9a.net/", false},
	{"xyz://foobar.com", false},
	{"rtmp://foobar.com", false},
	{"http://www.foo_bar.com/", false},
	{"http://localhost:3000/", false},
	{"http://foobar.com/#baz", false},
	{"http://foobar.com#baz=qux", false},
	{"http://foobar.com/t$-_.+!*\\'(),", false},
	{"http://www.foobar.com/~foobar", false},
	{"http://www.-foobar.com/", false},
	{"http://www.foo---bar.com/", false},
	{"mailto:someone@example.com", false},
	{"irc://irc.server.org/channel", false},
	{"irc://#channel@network", false},
	{"foobar.com", true},
	{"", true},
	{"invalid.", true},
	{".com", true},
	{"/abs/test/dir", true},
	{"./rel/test/dir", true},
	{"irc:", true},
	{"http://", true},
}

func TestURL(t *testing.T) {
	for _, tc := range urlTestCases {
		u, err := url.Parse(tc.url)
		assert.Require(t, assert.NoError(t, err))
		err = URL().Validate(u)
		if tc.shouldFail {
			assert.Require(t, assert.Error(t, err))
			assert.True(t, govy.HasErrorCode(err, ErrorCodeURL))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkURL(b *testing.B) {
	parsedURLs := make([]*url.URL, 0, len(urlTestCases))
	for _, tc := range urlTestCases {
		u, err := url.Parse(tc.url)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
		parsedURLs = append(parsedURLs, u)
	}
	b.ResetTimer()

	for range b.N {
		for _, u := range parsedURLs {
			_ = URL().Validate(u)
		}
	}
}
