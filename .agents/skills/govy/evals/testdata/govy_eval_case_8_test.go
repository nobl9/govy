package validation

import (
	"path/filepath"
	"testing"

	"github.com/nobl9/govy/pkg/govy"
)

func TestGovyEvalStrictFormatSelection(t *testing.T) {
	var validator govy.Validator[ImportRecord] = ImportRecordValidator
	valid := ImportRecord{
		ID:         "f47ac10b-58cc-4372-a567-0e02b2c3d479",
		Language:   "en-US",
		Payload:    "Zg",
		OutputPath: filepath.Join(t.TempDir(), "not-created"),
	}
	tests := []struct {
		name   string
		change func(*ImportRecord)
		want   map[string]int
	}{
		{name: "valid_formats_and_nonexistent_absolute_path"},
		{name: "url_alphabet", change: func(v *ImportRecord) { v.Payload = "-_8" }},
		{name: "no_padding_needed", change: func(v *ImportRecord) { v.Payload = "Zm9v" }},
		{name: "preferred_language", change: func(v *ImportRecord) { v.Language = "he" }},
		{name: "explicit_script", change: func(v *ImportRecord) { v.Language = "en-Latn" }},
		{
			name:   "uuid_version",
			change: func(v *ImportRecord) { v.ID = "f47ac10b-58cc-5372-a567-0e02b2c3d479" },
			want:   map[string]int{"id": 1},
		},
		{
			name:   "uuid_variant",
			change: func(v *ImportRecord) { v.ID = "f47ac10b-58cc-4372-7567-0e02b2c3d479" },
			want:   map[string]int{"id": 1},
		},
		{
			name:   "nil_uuid",
			change: func(v *ImportRecord) { v.ID = "00000000-0000-0000-0000-000000000000" },
			want:   map[string]int{"id": 1},
		},
		{name: "uuid_required", change: func(v *ImportRecord) { v.ID = "" }, want: map[string]int{"id": 1}},
		{
			name:   "language_case",
			change: func(v *ImportRecord) { v.Language = "en-us" },
			want:   map[string]int{"language": 1},
		},
		{
			name:   "language_alias",
			change: func(v *ImportRecord) { v.Language = "iw" },
			want:   map[string]int{"language": 1},
		},
		{
			name:   "language_underscore",
			change: func(v *ImportRecord) { v.Language = "en_US" },
			want:   map[string]int{"language": 1},
		},
		{
			name:   "language_required",
			change: func(v *ImportRecord) { v.Language = "" },
			want:   map[string]int{"language": 1},
		},
		{
			name:   "padded_payload",
			change: func(v *ImportRecord) { v.Payload = "Zg==" },
			want:   map[string]int{"payload": 1},
		},
		{
			name:   "padded_url_payload",
			change: func(v *ImportRecord) { v.Payload = "-_8=" },
			want:   map[string]int{"payload": 1},
		},
		{
			name:   "standard_alphabet",
			change: func(v *ImportRecord) { v.Payload = "+/8" },
			want:   map[string]int{"payload": 1},
		},
		{name: "trailing_bits", change: func(v *ImportRecord) { v.Payload = "Zh" }, want: map[string]int{"payload": 1}},
		{
			name:   "payload_newline",
			change: func(v *ImportRecord) { v.Payload = "Zg\n" },
			want:   map[string]int{"payload": 1},
		},
		{
			name:   "payload_space",
			change: func(v *ImportRecord) { v.Payload = " Zg" },
			want:   map[string]int{"payload": 1},
		},
		{
			name:   "payload_required",
			change: func(v *ImportRecord) { v.Payload = "" },
			want:   map[string]int{"payload": 1},
		},
		{
			name:   "relative_path",
			change: func(v *ImportRecord) { v.OutputPath = "relative/path" },
			want:   map[string]int{"outputPath": 1},
		},
		{
			name:   "path_required",
			change: func(v *ImportRecord) { v.OutputPath = "" },
			want:   map[string]int{"outputPath": 1},
		},
		{
			name: "independent_failures",
			change: func(v *ImportRecord) {
				v.ID = "f47ac10b-58cc-5372-a567-0e02b2c3d479"
				v.Language = "en-us"
				v.Payload = "Zg=="
				v.OutputPath = "relative/path"
			},
			want: map[string]int{"id": 1, "language": 1, "payload": 1, "outputPath": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := valid
			if tt.change != nil {
				tt.change(&input)
			}
			govyEvalAssertErrors(t, validator.Validate(input), tt.want)
		})
	}
}
