package rules

import (
	"errors"
	"net"
	"net/url"
	"strconv"
	"strings"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// StringMongoDBObjectID ensures the property's value is a 24-character
// lowercase hexadecimal MongoDB ObjectID.
func StringMongoDBObjectID() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringMongoDBObjectIDTemplate)

	return govy.NewRule(func(s string) error {
		if !mongoDBObjectIDRegexp().MatchString(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMongoDBObjectID).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringMongoDBConnectionString ensures the property's value is a MongoDB
// connection string with a mongodb:// or mongodb+srv:// scheme and valid,
// non-empty host entries. The mongodb:// scheme also accepts URL-encoded
// Unix domain socket paths.
func StringMongoDBConnectionString() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringMongoDBConnectionStringTemplate)

	return govy.NewRule(func(s string) error {
		if err := validateMongoDBConnectionString(s); err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Error:         err.Error(),
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMongoDBConnectionString).
		WithMessageTemplate(tpl).
		WithDescription("MongoDB connection string must use mongodb:// or mongodb+srv:// and contain valid, non-empty hosts or Unix sockets")
}

func validateMongoDBConnectionString(s string) error {
	scheme, hostList, err := splitMongoDBConnectionString(s)
	if err != nil {
		return err
	}

	switch scheme {
	case "mongodb", "mongodb+srv":
	default:
		return errors.New("scheme must be mongodb:// or mongodb+srv://")
	}
	if hostList == "" {
		return errors.New("host must not be empty")
	}

	hosts := strings.Split(hostList, ",")
	if scheme == "mongodb+srv" && len(hosts) != 1 {
		return errors.New("mongodb+srv connection string must contain exactly one host")
	}
	for _, host := range hosts {
		if err := validateMongoDBHost(host, scheme == "mongodb+srv"); err != nil {
			return err
		}
	}
	return nil
}

func splitMongoDBConnectionString(s string) (scheme, hostList string, err error) {
	scheme, rest, ok := strings.Cut(s, "://")
	if !ok {
		return "", "", nil
	}
	scheme = strings.ToLower(scheme)

	authority, _, _ := strings.Cut(rest, "/")
	authority, _, _ = strings.Cut(authority, "?")
	authority, _, _ = strings.Cut(authority, "#")
	if strings.Count(authority, "@") > 1 {
		return "", "", errors.New("userinfo must not contain unescaped @")
	}
	if userInfoEnd := strings.LastIndexByte(authority, '@'); userInfoEnd >= 0 {
		if err = validateMongoDBUserInfo(authority[:userInfoEnd]); err != nil {
			return "", "", err
		}
		authority = authority[userInfoEnd+1:]
	}
	return scheme, authority, nil
}

func validateMongoDBUserInfo(userInfo string) error {
	if strings.Count(userInfo, ":") > 1 {
		return errors.New("userinfo must contain no more than one unescaped colon")
	}
	if strings.ContainsAny(userInfo, "$[]") {
		return errors.New("userinfo contains unescaped reserved character")
	}
	if _, err := url.PathUnescape(userInfo); err != nil {
		return errors.New("userinfo contains invalid percent-encoding")
	}
	return nil
}

func validateMongoDBHost(host string, srv bool) error {
	if host == "" {
		return errors.New("host must not be empty")
	}
	if strings.ContainsAny(host, " \t\r\n") {
		return errors.New("host must not contain whitespace")
	}
	if strings.HasPrefix(host, "[") {
		hostname, port, err := splitMongoDBBracketedHost(host)
		if err != nil {
			return err
		}
		if srv && port != "" {
			return errors.New("mongodb+srv host must not include a port")
		}
		if srv {
			return errors.New("mongodb+srv host must be a DNS name")
		}
		if net.ParseIP(hostname) == nil || !strings.Contains(hostname, ":") {
			return errors.New("host must be a valid IP address or DNS name")
		}
		return nil
	}
	if strings.ContainsAny(host, "[]") {
		return errors.New("host contains malformed IPv6 brackets")
	}
	if isMongoDBUnixSocketHost(host) {
		if srv {
			return errors.New("mongodb+srv host must be a DNS name")
		}
		return nil
	}
	if strings.Count(host, ":") > 1 {
		return errors.New("IPv6 hosts must be enclosed in brackets")
	}

	hostname, port, hasPort := strings.Cut(host, ":")
	if hostname == "" {
		return errors.New("host must not be empty")
	}
	if hasPort {
		if srv {
			return errors.New("mongodb+srv host must not include a port")
		}
		if err := validateMongoDBPort(port); err != nil {
			return err
		}
	}
	if srv && net.ParseIP(hostname) != nil {
		return errors.New("mongodb+srv host must be a DNS name")
	}
	if srv && !isValidSRVHostname(hostname) {
		return errors.New("mongodb+srv host must include hostname, domain, and top-level domain")
	}
	if !isValidMongoDBHostname(hostname) {
		return errors.New("host must be a valid IP address, DNS name, or URL-encoded Unix domain socket")
	}
	return nil
}

func splitMongoDBBracketedHost(host string) (hostname, port string, err error) {
	end := strings.IndexByte(host, ']')
	if end < 0 {
		return "", "", errors.New("host contains malformed IPv6 brackets")
	}

	hostname = host[1:end]
	if hostname == "" {
		return "", "", errors.New("host must not be empty")
	}
	rest := host[end+1:]
	if rest == "" {
		return hostname, "", nil
	}
	if !strings.HasPrefix(rest, ":") {
		return "", "", errors.New("host contains malformed IPv6 brackets")
	}

	port = rest[1:]
	if err := validateMongoDBPort(port); err != nil {
		return "", "", err
	}
	return hostname, port, nil
}

func validateMongoDBPort(port string) error {
	if port == "" {
		return errors.New("port must not be empty")
	}
	for _, r := range port {
		if r < '0' || r > '9' {
			return errors.New("port must contain only digits")
		}
	}
	n, err := strconv.Atoi(port)
	if err != nil || n < 1 || n > 65535 {
		return errors.New("port must be between 1 and 65535")
	}
	return nil
}

func isMongoDBUnixSocketHost(host string) bool {
	socketPath, err := url.PathUnescape(host)
	if err != nil {
		return false
	}
	return strings.HasPrefix(socketPath, "/") &&
		strings.HasSuffix(socketPath, ".sock") &&
		!strings.ContainsAny(socketPath, " \t\r\n")
}

func isValidMongoDBHostname(hostname string) bool {
	if hostname == "" || len(hostname) > 253 {
		return false
	}
	if ip := net.ParseIP(hostname); ip != nil {
		return true
	}

	hostname = strings.TrimSuffix(strings.ToLower(hostname), ".")
	if hostname == "" {
		return false
	}
	for _, label := range strings.Split(hostname, ".") {
		if !mongoDBHostLabelRegexp().MatchString(label) {
			return false
		}
	}
	return true
}

func isValidSRVHostname(hostname string) bool {
	if !isValidMongoDBHostname(hostname) {
		return false
	}
	hostname = strings.TrimSuffix(hostname, ".")
	return strings.Count(hostname, ".") >= 2
}
