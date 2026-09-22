package rules

import (
	"net/url"
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

var urlTestCases = []*struct {
	url                  string
	shouldFail           bool
	jsonSchemaDifference string
}{
	{url: "http://foo.bar#com"},
	{url: "http://foobar.com"},
	{url: "https://foobar.com"},
	{url: "http://foobar.coffee/"},
	{
		url:                  "http://foobar.中文网/",
		jsonSchemaDifference: "Ajv's URI format requires non-ASCII hostname characters to be encoded.",
	},
	{url: "http://foobar.org/"},
	{url: "http://foobar.org:8080/"},
	{url: "ftp://foobar.ua/"},
	{url: "http://user:pass@www.foobar.com/"}, //nolint:gosec // Dummy credentials test URL user information.
	{url: "http://127.0.0.1/"},
	{url: "http://duckduckgo.com/?q=%2F"},
	{url: "http://localhost:3000/"},
	{url: "http://foobar.com/?foo=bar#baz=qux"},
	{url: "http://foobar.com?foo=bar"},
	{url: "http://www.xn--froschgrn-x9a.net/"},
	{url: "xyz://foobar.com"},
	{url: "rtmp://foobar.com"},
	{url: "http://www.foo_bar.com/"},
	{url: "http://localhost:3000/"},
	{url: "http://foobar.com/#baz"},
	{url: "http://foobar.com#baz=qux"},
	{
		url:                  "http://foobar.com/t$-_.+!*\\'(),",
		jsonSchemaDifference: "Ajv's URI format rejects unescaped backslashes accepted by Go's URL parser.",
	},
	{url: "http://www.foobar.com/~foobar"},
	{url: "http://www.-foobar.com/"},
	{url: "http://www.foo---bar.com/"},
	{url: "mailto:someone@example.com"},
	{url: "irc://irc.server.org/channel"},
	{url: "irc://#channel@network"},
	{url: "foobar.com", shouldFail: true},
	{url: "", shouldFail: true},
	{url: "invalid.", shouldFail: true},
	{url: ".com", shouldFail: true},
	{url: "/abs/test/dir", shouldFail: true},
	{url: "./rel/test/dir", shouldFail: true},
	{url: "irc:", shouldFail: true},
	{
		url:                  "http://",
		shouldFail:           true,
		jsonSchemaDifference: "The URI format permits an empty authority, while Govy requires content after the scheme.",
	},
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

func TestURL_WithOptions(t *testing.T) {
	tests := map[string]struct {
		url         string
		options     []URLOption
		expectedErr string
		shouldFail  bool
	}{
		"scheme allowed": {
			url:     "https://foobar.com",
			options: []URLOption{URLSchemes("https")},
		},
		"scheme rejected": {
			url:         "http://foobar.com",
			options:     []URLOption{URLSchemes("https")},
			expectedErr: "valid URL must use one of the following schemes: 'https'",
			shouldFail:  true,
		},
		"host required allows host": {
			url:     "https://foobar.com",
			options: []URLOption{URLHostRequired()},
		},
		"host required rejects opaque": {
			url:         "mailto:someone@example.com",
			options:     []URLOption{URLHostRequired()},
			expectedErr: "valid URL must have a host",
			shouldFail:  true,
		},
		"host required rejects fragment only": {
			url:         "https:#fragment",
			options:     []URLOption{URLHostRequired()},
			expectedErr: "valid URL must have a host",
			shouldFail:  true,
		},
		"user info forbidden rejects user info": {
			url:         "https://user@foobar.com",
			options:     []URLOption{URLUserInfoForbidden()},
			expectedErr: "valid URL must not contain user information",
			shouldFail:  true,
		},
		"host allow list accepts exact hostname": {
			url:     "https://foobar.com",
			options: []URLOption{URLHostAllowList("foobar.com")},
		},
		"host allow list accepts hostname with different case": {
			url:     "https://FooBar.COM",
			options: []URLOption{URLHostAllowList("foobar.com")},
		},
		"host allow list rejects unlisted hostname": {
			url:         "https://barfoo.com",
			options:     []URLOption{URLHostAllowList("foobar.com")},
			expectedErr: "valid URL must use one of the following hostnames: 'foobar.com'",
			shouldFail:  true,
		},
		"host allow list ignores port": {
			url:     "https://foobar.com:8443",
			options: []URLOption{URLHostAllowList("foobar.com")},
		},
		"host deny list rejects exact hostname": {
			url:         "https://foobar.com",
			options:     []URLOption{URLHostDenyList("foobar.com")},
			expectedErr: "valid URL must not use any of the following hostnames: 'foobar.com'",
			shouldFail:  true,
		},
		"host deny list rejects hostname with different case": {
			url:         "https://FooBar.COM",
			options:     []URLOption{URLHostDenyList("foobar.com")},
			expectedErr: "valid URL must not use any of the following hostnames: 'foobar.com'",
			shouldFail:  true,
		},
		"host deny list accepts other hostnames": {
			url:     "https://barfoo.com",
			options: []URLOption{URLHostDenyList("foobar.com")},
		},
		"host deny list wins over allow list": {
			url: "https://foobar.com",
			options: []URLOption{
				URLHostAllowList("foobar.com"),
				URLHostDenyList("foobar.com"),
			},
			expectedErr: "valid URL must not use any of the following hostnames: 'foobar.com'",
			shouldFail:  true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			u, err := url.Parse(tc.url)
			assert.Require(t, assert.NoError(t, err))

			err = URL(tc.options...).Validate(u)
			if tc.shouldFail {
				assert.Require(t, assert.Error(t, err))
				assert.EqualError(t, err, tc.expectedErr)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeURL))
			} else {
				assert.NoError(t, err)
			}
		})
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

	for _, u := range parsedURLs {
		for range b.N {
			rule := URL()
			_ = rule.Validate(u)
		}
	}
}
