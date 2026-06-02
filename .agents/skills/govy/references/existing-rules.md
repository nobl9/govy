# Existing Rules

Generated reference of exported predefined rule constructors from `pkg/rules`.
Use this when choosing an existing rule before writing a custom one.

[//]: # (docs: pkg/rules?kind=func&returns=govy.Rule,govy.RuleSet)

## Comparable

### `EQ`

```go
// EQ ensures the property's value is equal to the compared value.
func EQ[T comparable](compared T) govy.Rule[T]
```

### `NEQ`

```go
// NEQ ensures the property's value is not equal to the compared value.
func NEQ[T comparable](compared T) govy.Rule[T]
```

### `GT`

```go
// GT ensures the property's value is greater than the compared value.
func GT[T cmp.Ordered](compared T) govy.Rule[T]
```

### `GTE`

```go
// GTE ensures the property's value is greater than or equal to the compared value.
func GTE[T cmp.Ordered](compared T) govy.Rule[T]
```

### `LT`

```go
// LT ensures the property's value is less than the compared value.
func LT[T cmp.Ordered](compared T) govy.Rule[T]
```

### `LTE`

```go
// LTE ensures the property's value is less than or equal to the compared value.
func LTE[T cmp.Ordered](compared T) govy.Rule[T]
```

### `EqualProperties`

```go
// EqualProperties checks if all the specified properties are equal.
// It uses the provided [ComparisonFunc] to compare the values.
// The following built-in comparison functions are available:
//   - [CompareFunc]
//   - [CompareDeepEqualFunc]
//
// If builtin [ComparisonFunc] is not enough, a custom function can be used.
func EqualProperties[T, P any](compare ComparisonFunc[T], getters map[string]func(parent P) T) govy.Rule[P]
```

### `GTProperties`

```go
// GTProperties ensures the first property's value is greater than the second property's value.
// It works with [cmp.Ordered] types (int, float64, string, etc.).
func GTProperties[T cmp.Ordered, P any](
	firstName string,
	firstGetter func(parent P) T,
	secondName string,
	secondGetter func(parent P) T,
) govy.Rule[P]
```

### `GTEProperties`

```go
// GTEProperties ensures the first property's value is greater than or equal to the second property's value.
// It works with [cmp.Ordered] types (int, float64, string, etc.).
func GTEProperties[T cmp.Ordered, P any](
	firstName string,
	firstGetter func(parent P) T,
	secondName string,
	secondGetter func(parent P) T,
) govy.Rule[P]
```

### `LTProperties`

```go
// LTProperties ensures the first property's value is less than the second property's value.
// It works with [cmp.Ordered] types (int, float64, string, etc.).
func LTProperties[T cmp.Ordered, P any](
	firstName string,
	firstGetter func(parent P) T,
	secondName string,
	secondGetter func(parent P) T,
) govy.Rule[P]
```

### `LTEProperties`

```go
// LTEProperties ensures the first property's value is less than or equal to the second property's value.
// It works with [cmp.Ordered] types (int, float64, string, etc.).
func LTEProperties[T cmp.Ordered, P any](
	firstName string,
	firstGetter func(parent P) T,
	secondName string,
	secondGetter func(parent P) T,
) govy.Rule[P]
```

### `GTComparableProperties`

```go
// GTComparableProperties ensures the first property's value is greater than the second property's value.
// It works with types that implement [Comparable], such as [time.Time].
func GTComparableProperties[T Comparable[T], P any](
	firstName string,
	firstGetter func(parent P) T,
	secondName string,
	secondGetter func(parent P) T,
) govy.Rule[P]
```

### `GTEComparableProperties`

```go
// GTEComparableProperties ensures the first property's value is greater than or equal to the second property's value.
// It works with types that implement [Comparable], such as [time.Time].
func GTEComparableProperties[T Comparable[T], P any](
	firstName string,
	firstGetter func(parent P) T,
	secondName string,
	secondGetter func(parent P) T,
) govy.Rule[P]
```

### `LTComparableProperties`

```go
// LTComparableProperties ensures the first property's value is less than the second property's value.
// It works with types that implement [Comparable], such as [time.Time].
func LTComparableProperties[T Comparable[T], P any](
	firstName string,
	firstGetter func(parent P) T,
	secondName string,
	secondGetter func(parent P) T,
) govy.Rule[P]
```

### `LTEComparableProperties`

```go
// LTEComparableProperties ensures the first property's value is less than or equal to the second property's value.
// It works with types that implement [Comparable], such as [time.Time].
func LTEComparableProperties[T Comparable[T], P any](
	firstName string,
	firstGetter func(parent P) T,
	secondName string,
	secondGetter func(parent P) T,
) govy.Rule[P]
```

## Duration

### `DurationPrecision`

```go
// DurationPrecision ensures the duration is defined with the specified precision.
func DurationPrecision(precision time.Duration) govy.Rule[time.Duration]
```

## Forbidden

### `Forbidden`

```go
// Forbidden ensures the property's value is its type's zero value, i.e. it's empty.
func Forbidden[T any]() govy.Rule[T]
```

## Length

### `StringLength`

```go
// StringLength ensures the string's length is between min and max (closed interval).
//
// The following, additional template variables are supported:
//   - [govy.TemplateVars.MinLength]
//   - [govy.TemplateVars.MaxLength]
func StringLength(minLen, maxLen int) govy.Rule[string]
```

### `StringMinLength`

```go
// StringMinLength ensures the string's length is greater than or equal to the limit.
func StringMinLength(limit int) govy.Rule[string]
```

### `StringMaxLength`

```go
// StringMaxLength ensures the string's length is less than or equal to the limit.
func StringMaxLength(limit int) govy.Rule[string]
```

### `SliceLength`

```go
// SliceLength ensures the slice's length is between min and max (closed interval).
//
// The following, additional template variables are supported:
//   - [govy.TemplateVars.MinLength]
//   - [govy.TemplateVars.MaxLength]
func SliceLength[S ~[]E, E any](minLen, maxLen int) govy.Rule[S]
```

### `SliceMinLength`

```go
// SliceMinLength ensures the slice's length is greater than or equal to the limit.
func SliceMinLength[S ~[]E, E any](limit int) govy.Rule[S]
```

### `SliceMaxLength`

```go
// SliceMaxLength ensures the slice's length is less than or equal to the limit.
func SliceMaxLength[S ~[]E, E any](limit int) govy.Rule[S]
```

### `MapLength`

```go
// MapLength ensures the map's length is between min and max (closed interval).
//
// The following, additional template variables are supported:
//   - [govy.TemplateVars.MinLength]
//   - [govy.TemplateVars.MaxLength]
func MapLength[M ~map[K]V, K comparable, V any](minLen, maxLen int) govy.Rule[M]
```

### `MapMinLength`

```go
// MapMinLength ensures the map's length is greater than or equal to the limit.
func MapMinLength[M ~map[K]V, K comparable, V any](limit int) govy.Rule[M]
```

### `MapMaxLength`

```go
// MapMaxLength ensures the map's length is less than or equal to the limit.
func MapMaxLength[M ~map[K]V, K comparable, V any](limit int) govy.Rule[M]
```

## One Of

### `OneOf`

```go
// OneOf checks if the property's value matches one of the provided values.
// The values must be comparable.
//
// For reversed rule see [NotOneOf].
func OneOf[T comparable](values ...T) govy.Rule[T]
```

### `NotOneOf`

```go
// NotOneOf checks if the property's value does not match any of the provided values.
// The values must be comparable.
//
// For reversed rule see [OneOf].
func NotOneOf[T comparable](values ...T) govy.Rule[T]
```

### `OneOfProperties`

```go
// OneOfProperties checks if at least one of the properties is set.
// Property is considered set if its value is not empty (non-zero).
func OneOfProperties[T any](getters map[string]func(parent T) any) govy.Rule[T]
```

### `MutuallyExclusive`

```go
// MutuallyExclusive checks if properties are mutually exclusive.
// This means, exactly one of the properties can be set.
// Property is considered set if its value is not empty (non-zero).
// If required is true, then a single non-empty property is required.
func MutuallyExclusive[T any](required bool, getters map[string]func(parent T) any) govy.Rule[T]
```

### `MutuallyDependent`

```go
// MutuallyDependent checks if properties are mutually dependent.
// This means, if any of the properties is set, the rest must be also set.
// Property is considered set if its value is not empty (non-zero).
func MutuallyDependent[T any](getters map[string]func(parent T) any) govy.Rule[T]
```

## Required

### `Required`

```go
// Required ensures the property's value is not empty (i.e. it's not its type's zero value).
func Required[T any]() govy.Rule[T]
```

## String

### `StringNotEmpty`

```go
// StringNotEmpty ensures the property's value is not empty.
// The string is considered empty if it contains only whitespace characters.
func StringNotEmpty() govy.Rule[string]
```

### `StringMatchRegexp`

```go
// StringMatchRegexp ensures the property's value matches the regular expression.
// The error message can be enhanced with examples of valid values.
func StringMatchRegexp(re *regexp.Regexp) govy.Rule[string]
```

### `StringDenyRegexp`

```go
// StringDenyRegexp ensures the property's value does not match the regular expression.
// The error message can be enhanced with examples of invalid values.
func StringDenyRegexp(re *regexp.Regexp) govy.Rule[string]
```

### `StringDNSLabel`

```go
// StringDNSLabel ensures the property's value is a valid DNS label as defined by [RFC 1123].
//
// [RFC 1123]: https://www.ietf.org/rfc/rfc1123.txt
func StringDNSLabel() govy.RuleSet[string]
```

### `StringDNSSubdomain`

```go
// StringDNSSubdomain ensures the property's value is a valid DNS subdomain as defined by [RFC 1123].
//
// [RFC 1123]: https://www.ietf.org/rfc/rfc1123.txt
func StringDNSSubdomain() govy.RuleSet[string]
```

### `StringEmail`

```go
// StringEmail ensures the property's value is a valid email address.
// It follows [RFC 5322] specification which is more permissive in regards to domain names.
//
// [RFC 5322]: https://www.ietf.org/rfc/rfc5322.txt
func StringEmail() govy.Rule[string]
```

### `StringURL`

```go
// StringURL ensures property's value is a valid URL as defined by [url.Parse] function.
// Unlike [URL] it does not impose any additional rules upon parsed [url.URL].
func StringURL() govy.Rule[string]
```

### `StringMAC`

```go
// StringMAC ensures property's value is a valid MAC address.
func StringMAC() govy.Rule[string]
```

### `StringIP`

```go
// StringIP ensures property's value is a valid IP address.
func StringIP() govy.Rule[string]
```

### `StringIPv4`

```go
// StringIPv4 ensures property's value is a valid IPv4 address.
func StringIPv4() govy.Rule[string]
```

### `StringIPv6`

```go
// StringIPv6 ensures property's value is a valid IPv6 address.
func StringIPv6() govy.Rule[string]
```

### `StringCIDR`

```go
// StringCIDR ensures property's value is a valid CIDR notation IP address.
func StringCIDR() govy.Rule[string]
```

### `StringCIDRv4`

```go
// StringCIDRv4 ensures property's value is a valid CIDR notation IPv4 address.
func StringCIDRv4() govy.Rule[string]
```

### `StringCIDRv6`

```go
// StringCIDRv6 ensures property's value is a valid CIDR notation IPv6 address.
func StringCIDRv6() govy.Rule[string]
```

### `StringUUID`

```go
// StringUUID ensures property's value is a valid UUID string as defined by [RFC 4122].
// It does not enforce a specific UUID version.
//
// [RFC 4122]: https://www.ietf.org/rfc/rfc4122.txt
func StringUUID() govy.Rule[string]
```

### `StringASCII`

```go
// StringASCII ensures property's value contains only ASCII characters.
func StringASCII() govy.Rule[string]
```

### `StringJSON`

```go
// StringJSON ensures property's value is a valid JSON literal.
func StringJSON() govy.Rule[string]
```

### `StringContains`

```go
// StringContains ensures the property's value contains all the provided substrings.
func StringContains(substrings ...string) govy.Rule[string]
```

### `StringExcludes`

```go
// StringExcludes ensures the property's value does not contain any of the provided substrings.
func StringExcludes(substrings ...string) govy.Rule[string]
```

### `StringStartsWith`

```go
// StringStartsWith ensures the property's value starts with one of the provided prefixes.
func StringStartsWith(prefixes ...string) govy.Rule[string]
```

### `StringEndsWith`

```go
// StringEndsWith ensures the property's value ends with one of the provided suffixes.
func StringEndsWith(suffixes ...string) govy.Rule[string]
```

### `StringTitle`

```go
// StringTitle ensures each word in a string starts with a capital letter.
func StringTitle() govy.Rule[string]
```

### `StringGitRef`

```go
// StringGitRef ensures a git reference name follows the [git-check-ref-format] rules.
//
// It is important to note that this function does not check if the reference exists in the repository.
// It only checks if the reference name is valid.
// This functions does not support the '--refspec-pattern', '--normalize', and '--allow-onelevel' options.
//
// Git imposes the following rules on how references are named:
//
//  1. They can include slash '/' for hierarchical (directory) grouping, but no
//     slash-separated component can begin with a dot '.' or end with the
//     sequence '.lock'.
//  2. They must contain at least one '/'. This enforces the presence of a
//     category (e.g. 'heads/', 'tags/'), but the actual names are not restricted.
//  3. They cannot have ASCII control characters (i.e. bytes whose values are
//     lower than '\040', or '\177' DEL).
//  4. They cannot have '?', '*', '[', ' ', '~', '^', ', '\t', '\n', '@{', '\\' and '..',
//  5. They cannot begin or end with a slash '/'.
//  6. They cannot end with a '.'.
//  7. They cannot be the single character '@'.
//  8. 'HEAD' is an allowed special name.
//
// Slightly modified version of [go-git] implementation, kudos to the authors!
//
// [git-check-ref-format] :https://git-scm.com/docs/git-check-ref-format
// [go-git]: https://github.com/go-git/go-git/blob/95afe7e1cdf71c59ee8a71971fac71880020a744/plumbing/reference.go#L167
func StringGitRef() govy.Rule[string]
```

### `StringFileSystemPath`

```go
// StringFileSystemPath ensures the property's value is an existing file system path.
func StringFileSystemPath() govy.Rule[string]
```

### `StringFilePath`

```go
// StringFilePath ensures the property's value is a file system path pointing to an existing file.
func StringFilePath() govy.Rule[string]
```

### `StringDirPath`

```go
// StringDirPath ensures the property's value is a file system path pointing to an existing directory.
func StringDirPath() govy.Rule[string]
```

### `StringMatchFileSystemPath`

```go
// StringMatchFileSystemPath ensures the property's value matches the provided file path pattern.
// It uses [filepath.Match] to match the pattern. The native function comes with some limitations,
// most notably it does not support '**' recursive expansion.
// It does not check if the file path exists on the file system.
func StringMatchFileSystemPath(pattern string) govy.Rule[string]
```

### `StringRegexp`

```go
// StringRegexp ensures the property's value is a valid regular expression.
// The accepted regular expression syntax must comply to RE2.
// It is described at https://golang.org/s/re2syntax, except for \C.
// For an overview of the syntax, see [regexp/syntax] package.
//
// [regexp/syntax]: https://pkg.go.dev/regexp/syntax
func StringRegexp() govy.Rule[string]
```

### `StringCrontab`

```go
// StringCrontab ensures the property's value is a valid crontab schedule expression.
// For more details on cron expressions read [crontab manual] and visit [crontab.guru].
//
// [crontab manual]: https://www.man7.org/linux/man-pages/man5/crontab.5.html
// [crontab.guru]: https://crontab.guru
func StringCrontab() govy.Rule[string]
```

### `StringDateTime`

```go
// StringDateTime ensures the property's value is a valid date and time in the specified layout.
//
// The layout must be a valid time format string as defined by [time.Parse],
// an example of which is [time.RFC3339].
func StringDateTime(layout string) govy.Rule[string]
```

### `StringTimeZone`

```go
// StringTimeZone ensures the property's value is a valid time zone name which
// uniquely identifies a time zone in the IANA Time Zone database.
// Example: "America/New_York", "Europe/London".
//
// Under the hood [time.LoadLocation] is called to parse the zone.
// The native function allows empty string and 'Local' keyword to be supplied.
// However, these two options are explicitly forbidden by [StringTimeZone].
//
// Furthermore, the time zone data is not readily available in one predefined place.
// [time.LoadLocation] looks for the IANA Time Zone database in specific places,
// please refer to its documentation for more information.
func StringTimeZone() govy.Rule[string]
```

### `StringAlpha`

```go
// StringAlpha ensures the property's value consists only of ASCII letters.
func StringAlpha() govy.Rule[string]
```

### `StringAlphanumeric`

```go
// StringAlphanumeric ensures the property's value consists only of ASCII letters and numbers.
func StringAlphanumeric() govy.Rule[string]
```

### `StringAlphaUnicode`

```go
// StringAlphaUnicode ensures the property's value consists only of Unicode letters.
func StringAlphaUnicode() govy.Rule[string]
```

### `StringAlphanumericUnicode`

```go
// StringAlphanumericUnicode ensures the property's value consists only of Unicode letters and numbers.
func StringAlphanumericUnicode() govy.Rule[string]
```

### `StringFQDN`

```go
// StringFQDN ensures the property's value is a fully qualified domain name (FQDN).
func StringFQDN() govy.Rule[string]
```

### `StringKubernetesQualifiedName`

```go
// StringKubernetesQualifiedName ensures the property's value is a valid "qualified name"
// as defined by [Kubernetes validation].
// The qualified name is used in various parts of the Kubernetes system, examples:
//   - annotation names
//   - label names
//
// [Kubernetes validation]: https://github.com/kubernetes/kubernetes/blob/55573a0739785292e62b32a748c0b0735ff963ba/staging/src/k8s.io/apimachinery/pkg/util/validation/validation.go#L41
func StringKubernetesQualifiedName() govy.RuleSet[string]
```

## Unique

### `SliceUnique`

```go
// SliceUnique ensures that a slice contains unique elements based on a provided [HashFunction].
// The following built-in hashing functions are available:
//   - [HashFuncSelf]
//
// You can optionally specify constraints which will be included in the error message to further
// clarify the reason for breaking uniqueness.
func SliceUnique[S []V, V any, H comparable](hashFunc HashFunction[V, H], constraints ...string) govy.Rule[S]
```

### `UniqueProperties`

```go
// UniqueProperties ensures each property is unique based on a provided [HashFunction].
// The following built-in hashing functions are available:
//   - [HashFuncSelf]
//
// You can optionally specify constraints which will be included in the error message to further
// clarify the reason for breaking uniqueness.
func UniqueProperties[V, P any, H comparable](
	hashFunc HashFunction[V, H],
	getters map[string]func(parent P) V,
	constraints ...string,
) govy.Rule[P]
```

## Url

### `URL`

```go
// URL ensures the URL is valid.
// The URL must have a scheme (e.g. https://) and contain either host, fragment or opaque data.
func URL() govy.Rule[*url.URL]
```

[//]: # (end-docs)
