package validation

import (
	"testing"

	"github.com/nobl9/govy/pkg/govy"
)

func TestGovyEvalPointerPresence(t *testing.T) {
	var validator govy.Validator[Sample] = SampleValidator
	valid := Sample{Enabled: govyEvalPtr(false), Target: govyEvalPtr(0.0)}
	tests := []struct {
		name   string
		change func(*Sample)
		want   map[string]int
	}{
		{name: "present_false_and_zero"},
		{name: "present_true", change: func(v *Sample) { v.Enabled = govyEvalPtr(true) }},
		{name: "target_maximum", change: func(v *Sample) { v.Target = govyEvalPtr(100.0) }},
		{name: "target_fraction", change: func(v *Sample) { v.Target = govyEvalPtr(50.25) }},
		{name: "target_negative", change: func(v *Sample) { v.Target = govyEvalPtr(-1.0) }, want: map[string]int{"target": 1}},
		{name: "target_too_high", change: func(v *Sample) { v.Target = govyEvalPtr(101.0) }, want: map[string]int{"target": 1}},
		{name: "target_just_below_zero", change: func(v *Sample) { v.Target = govyEvalPtr(-0.001) }, want: map[string]int{"target": 1}},
		{name: "target_just_above_maximum", change: func(v *Sample) { v.Target = govyEvalPtr(100.001) }, want: map[string]int{"target": 1}},
		{name: "enabled_missing", change: func(v *Sample) { v.Enabled = nil }, want: map[string]int{"enabled": 1}},
		{name: "target_missing", change: func(v *Sample) { v.Target = nil }, want: map[string]int{"target": 1}},
		{name: "both_required_values_missing", change: func(v *Sample) { v.Enabled = nil; v.Target = nil }, want: map[string]int{"enabled": 1, "target": 1}},
		{name: "note_present", change: func(v *Sample) { v.Note = govyEvalPtr("text") }},
		{name: "note_present_empty", change: func(v *Sample) { v.Note = govyEvalPtr("") }, want: map[string]int{"note": 1}},
		{
			name: "independent_failures",
			change: func(v *Sample) {
				v.Enabled = nil
				v.Target = govyEvalPtr(101.0)
				v.Note = govyEvalPtr("")
			},
			want: map[string]int{"enabled": 1, "target": 1, "note": 1},
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
