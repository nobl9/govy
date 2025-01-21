package messagetemplates

// templateKey is a key that uniquely identifies a message template.
type templateKey int

const (
	StringLengthTemplate templateKey = iota + 1
)

var rawMessageTemplates = map[templateKey]string{
	StringLengthTemplate: "length must be between {{ .MinLength }} and {{ .MaxLength }}",
}

// commonTemplateSuffix is a suffix that is added to all message templates.
// It includes examples and details and handles their absence.
const commonTemplateSuffix = "{{- if .Examples }} {{ formatExamples .Examples }}{{- end }}" +
	"{{- if .Details }}; {{ .Details }}{{- end }}"
