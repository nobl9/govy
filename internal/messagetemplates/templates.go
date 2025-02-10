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
)

var rawMessageTemplates = map[templateKey]string{
	LengthTemplate:    "length must be between {{ .MinLength }} and {{ .MaxLength }}",
	MinLengthTemplate: "length must be greater than or equal to {{ .ComparedValue }}",
	MaxLengthTemplate: "length must be less than or equal to '{{ .ComparedValue }}'",
	EQTemplate:        "should be equal to '{{ .ComparedValue }}'",
	NEQTemplate:       "should be not equal to '{{ .ComparedValue }}'",
	GTTemplate:        "should be greater than '{{ .ComparedValue }}'",
	GTETemplate:       "should be greater than or equal to '{{ .ComparedValue }}'",
	LTTemplate:        "should be less than '{{ .ComparedValue }}'",
	LTETemplate:       "should be less than or equal to '{{ .ComparedValue }}'",
}

// commonTemplateSuffix is a suffix that is added to all message templates.
// It includes examples and details and handles their absence.
const commonTemplateSuffix = "{{- if .Examples }} {{ formatExamples .Examples }}{{- end }}" +
	"{{- if .Details }}; {{ .Details }}{{- end }}"
