package messagetemplates

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
	DurationPrecisionTemplate
	ForbiddenTemplate
)

var rawMessageTemplates = map[templateKey]string{
	LengthTemplate:            "length must be between {{ .MinLength }} and {{ .MaxLength }}",
	MinLengthTemplate:         "length must be greater than or equal to {{ .ComparisonValue }}",
	MaxLengthTemplate:         "length must be less than or equal to '{{ .ComparisonValue }}'",
	EQTemplate:                "should be equal to '{{ .ComparisonValue }}'",
	NEQTemplate:               "should be not equal to '{{ .ComparisonValue }}'",
	GTTemplate:                "should be greater than '{{ .ComparisonValue }}'",
	GTETemplate:               "should be greater than or equal to '{{ .ComparisonValue }}'",
	LTTemplate:                "should be less than '{{ .ComparisonValue }}'",
	LTETemplate:               "should be less than or equal to '{{ .ComparisonValue }}'",
	DurationPrecisionTemplate: "duration must be defined with {{ .ComparisonValue }} precision",
	ForbiddenTemplate:         "property is forbidden",
}

// commonTemplateSuffix is a suffix that is added to all message templates.
// It includes examples and details and handles their absence.
const commonTemplateSuffix = "{{- if .Examples }} {{ formatExamples .Examples }}{{- end }}" +
	"{{- if .Details }}; {{ .Details }}{{- end }}"
