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
