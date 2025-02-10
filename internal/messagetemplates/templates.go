package messagetemplates

// templateKey is a key that uniquely identifies a message template.
type templateKey int

const (
	LengthTemplate templateKey = iota + 1
	MinLengthTemplate
	MaxLengthTemplate
)

var rawMessageTemplates = map[templateKey]string{
	LengthTemplate:    "length must be between {{ .MinLength }} and {{ .MaxLength }}",
	MinLengthTemplate: "length must be greater than or equal to {{ .MinLength }}",
	MaxLengthTemplate: "length must be less than or equal to {{ .MaxLength }}",
}

// commonTemplateSuffix is a suffix that is added to all message templates.
// It includes examples and details and handles their absence.
const commonTemplateSuffix = "{{- if .Examples }} {{ formatExamples .Examples }}{{- end }}" +
	"{{- if .Details }}; {{ .Details }}{{- end }}"
