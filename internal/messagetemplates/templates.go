package messagetemplates

import "github.com/nobl9/govy/internal"

// templateKey is a key that uniquely identifies a message template.
type templateKey int

const (
	LengthTemplate templateKey = iota + 1
	MinLengthTemplate
	MaxLengthTemplate
	EQTemplate
	NEQTemplate
	GTTemplate
	GTETemplate
	LTTemplate
	LTETemplate
	EqualPropertiesTemplate
	DurationPrecisionTemplate
	ForbiddenTemplate
	OneOfTemplate
	OneOfPropertiesTemplate
	MutuallyExclusiveTemplate
	RequiredTemplate
	StringNonEmptyTemplate
	StringMatchRegexpTemplate
	StringDenyRegexpTemplate
	StringEmailTemplate
	StringMACTemplate
	StringIPTemplate
	StringIPv4Template
	StringIPv6Template
	StringCIDRTemplate
	StringCIDRv4Template
	StringCIDRv6Template
	StringJSONTemplate
	StringContainsTemplate
	StringExcludesTemplate
	StringStartsWithTemplate
	StringEndsWithTemplate
	StringTitleTemplate
	StringGitRefTemplate
	StringFileSystemPathTemplate
	StringFilePathTemplate
	StringDirPathTemplate
	StringMatchFileSystemPathTemplate
	StringRegexpTemplate
	StringCrontabTemplate
	StringDateTimeTemplate
	StringTimeZoneTemplate
	URLTemplate
	SliceUniqueTemplate
)

var rawMessageTemplates = map[templateKey]string{
	LengthTemplate:            "length must be between {{ .MinLength }} and {{ .MaxLength }}",
	MinLengthTemplate:         "length must be greater than or equal to {{ .ComparisonValue }}",
	MaxLengthTemplate:         "length must be less than or equal to {{ .ComparisonValue }}",
	EQTemplate:                "should be equal to '{{ .ComparisonValue }}'",
	NEQTemplate:               "should be not equal to '{{ .ComparisonValue }}'",
	GTTemplate:                "should be greater than '{{ .ComparisonValue }}'",
	GTETemplate:               "should be greater than or equal to '{{ .ComparisonValue }}'",
	LTTemplate:                "should be less than '{{ .ComparisonValue }}'",
	LTETemplate:               "should be less than or equal to '{{ .ComparisonValue }}'",
	EqualPropertiesTemplate:   `all of [{{ joinStringSlice .ComparisonValue "" }}] properties must be equal, but '{{ .Custom.FirstNotEqual }}' is not equal to '{{ .Custom.SecondNotEqual }}'`,
	DurationPrecisionTemplate: "duration must be defined with {{ .ComparisonValue }} precision",
	ForbiddenTemplate:         "property is forbidden",
	OneOfTemplate:             `must be one of [{{ joinStringSlice .ComparisonValue "" }}]`,
	OneOfPropertiesTemplate:   `one of [{{ joinStringSlice .ComparisonValue "" }}] properties must be set, none was provided`,
	MutuallyExclusiveTemplate: `
{{- if .Custom.NoProperties -}}
one of [{{ joinStringSlice .ComparisonValue "" }}] properties must be set, none was provided
{{- else -}}
[{{ joinStringSlice .ComparisonValue "" }}] properties are mutually exclusive, provide only one of them
{{- end }}`,
	RequiredTemplate:          internal.RequiredErrorMessage,
	StringNonEmptyTemplate:    "string cannot be empty",
	StringMatchRegexpTemplate: "string must match regular expression: '{{ .ComparisonValue }}'",
	StringDenyRegexpTemplate:  "string must not match regular expression: '{{ .ComparisonValue }}'",
	StringEmailTemplate:       "string must be a valid email address: {{ .Error }}",
	StringMACTemplate:         "string must be a valid MAC address",
	StringIPTemplate:          "string must be a valid IP address",
	StringIPv4Template:        "string must be a valid IPv4 address",
	StringIPv6Template:        "string must be a valid IPv6 address",
	StringCIDRTemplate:        "string must be a valid CIDR notation IP address",
	StringCIDRv4Template:      "string must be a valid CIDR notation IPv4 address",
	StringCIDRv6Template:      "string must be a valid CIDR notation IPv6 address",
	StringJSONTemplate:        "string must be a valid JSON",
	StringContainsTemplate:    `string must contain the following substrings: {{ joinStringSlice .ComparisonValue "'" }}`,
	StringExcludesTemplate:    `string must not contain any of the following substrings: {{ joinStringSlice .ComparisonValue "'" }}`,
	StringStartsWithTemplate: `
{{- if eq (len .ComparisonValue) 1 -}}
string must start with '{{ index .ComparisonValue 0 }}' prefix
{{- else -}}
string must start with one of the following prefixes: {{ joinStringSlice .ComparisonValue "'" }}
{{- end }}
`,
	StringEndsWithTemplate: `
{{- if eq (len .ComparisonValue) 1 -}}
string must end with '{{ index .ComparisonValue 0 }}' suffix
{{- else -}}
string must end with one of the following suffixes: {{ joinStringSlice .ComparisonValue "'" }}
{{- end }}
`,
	StringTitleTemplate: "each word in a string must start with a capital letter",
	StringGitRefTemplate: `
{{- if eq .Custom.GitRefEmpty true -}}
git reference cannot be empty
{{- else if eq .Custom.GitRefEndsWithDot true -}}
git reference must not end with a '.'
{{- else if eq .Custom.GitRefAtLeastOneSlash true -}}
git reference must contain at least one '/'
{{- else if eq .Custom.GitRefEmptyPart true -}}
git reference must not have empty parts
{{- else if eq .Custom.GitRefStartsWithDash true -}}
git branch and tag references must not start with '-'
{{- else if eq .Custom.GitRefForbiddenChars true -}}
git reference contains forbidden characters
{{- else -}}
string must be a valid git reference
{{- end }}
`,
	StringFileSystemPathTemplate:      "string must be an existing file system path{{- if .Error }}: {{ .Error }}{{- end }}",
	StringFilePathTemplate:            "string must be a file system path to an existing file{{- if .Error }}: {{ .Error }}{{- end }}",
	StringDirPathTemplate:             "string must be a file system path to an existing directory{{- if .Error }}: {{ .Error }}{{- end }}",
	StringMatchFileSystemPathTemplate: "string must match file path pattern: '{{ .ComparisonValue }}'{{- if .Error }}: {{ .Error }}{{- end }}",
	StringRegexpTemplate:              "string must be a valid regular expression{{- if .Error }}: {{ .Error }}{{- end }}",
	StringCrontabTemplate:             "string must be a valid cron schedule expression",
	StringDateTimeTemplate:            "string must be a valid date and time in '{{ .ComparisonValue }}' format{{- if .Error }}: {{ .Error }}{{- end }}",
	StringTimeZoneTemplate:            "string must be a valid IANA Time Zone Database code{{- if .Error }}: {{ .Error }}{{- end }}",
	SliceUniqueTemplate: `elements are not unique, {{ .Custom.FirstOrdinal }} and {{ .Custom.SecondOrdinal }} elements collide
{{- if gt (len .Custom.Constraints) 0 }} based on constraints: {{ joinStringSlice .Custom.Constraints "" }}{{- end }}`,
	URLTemplate: "{{ .Error }}",
}

// commonTemplateSuffix is a suffix that is added to all message templates.
// It includes examples and details and handles their absence.
const commonTemplateSuffix = "{{- if .Examples }} {{ formatExamples .Examples }}{{- end }}" +
	"{{- if .Details }}; {{ .Details }}{{- end }}"
