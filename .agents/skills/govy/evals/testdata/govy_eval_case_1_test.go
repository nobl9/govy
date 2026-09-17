package validation

import (
	"fmt"
	"strings"
	"testing"

	"github.com/nobl9/govy/pkg/govy"
)

func TestGovyEvalFieldBoundaries(t *testing.T) {
	var validator govy.Validator[AlertPolicyDraft] = AlertPolicyValidator
	valid := AlertPolicyDraft{Name: "abc", Severity: "low"}
	labels := func(count int) []string {
		values := make([]string, count)
		for i := range values {
			values[i] = fmt.Sprintf("label-%d", i)
		}
		return values
	}
	tests := []struct {
		name   string
		change func(*AlertPolicyDraft)
		want   map[string]int
	}{
		{name: "nil_optional_values"},
		{name: "name_minimum", change: func(v *AlertPolicyDraft) { v.Name = strings.Repeat("a", 3) }},
		{name: "name_maximum", change: func(v *AlertPolicyDraft) { v.Name = strings.Repeat("a", 63) }},
		{name: "name_required", change: func(v *AlertPolicyDraft) { v.Name = "" }, want: map[string]int{"name": 1}},
		{name: "name_too_short", change: func(v *AlertPolicyDraft) { v.Name = "ab" }, want: map[string]int{"name": 1}},
		{name: "name_too_long", change: func(v *AlertPolicyDraft) { v.Name = strings.Repeat("a", 64) }, want: map[string]int{"name": 1}},
		{name: "name_non_ascii", change: func(v *AlertPolicyDraft) { v.Name = "abé" }, want: map[string]int{"name": 1}},
		{name: "description_empty", change: func(v *AlertPolicyDraft) { v.Description = govyEvalPtr("") }},
		{name: "description_maximum", change: func(v *AlertPolicyDraft) { v.Description = govyEvalPtr(strings.Repeat("d", 200)) }},
		{name: "description_too_long", change: func(v *AlertPolicyDraft) { v.Description = govyEvalPtr(strings.Repeat("d", 201)) }, want: map[string]int{"description": 1}},
		{name: "description_non_ascii", change: func(v *AlertPolicyDraft) { v.Description = govyEvalPtr("é") }, want: map[string]int{"description": 1}},
		{name: "severity_medium", change: func(v *AlertPolicyDraft) { v.Severity = "medium" }},
		{name: "severity_high", change: func(v *AlertPolicyDraft) { v.Severity = "high" }},
		{name: "severity_critical", change: func(v *AlertPolicyDraft) { v.Severity = "critical" }},
		{name: "severity_empty", change: func(v *AlertPolicyDraft) { v.Severity = "" }, want: map[string]int{"severity": 1}},
		{name: "severity_unknown", change: func(v *AlertPolicyDraft) { v.Severity = "urgent" }, want: map[string]int{"severity": 1}},
		{name: "severity_case_sensitive", change: func(v *AlertPolicyDraft) { v.Severity = "LOW" }, want: map[string]int{"severity": 1}},
		{name: "labels_empty", change: func(v *AlertPolicyDraft) { v.Labels = []string{} }},
		{name: "labels_maximum", change: func(v *AlertPolicyDraft) { v.Labels = labels(10) }},
		{name: "labels_too_many", change: func(v *AlertPolicyDraft) { v.Labels = labels(11) }, want: map[string]int{"labels": 1}},
		{name: "labels_duplicate", change: func(v *AlertPolicyDraft) { v.Labels = []string{"x", "y", "x"} }, want: map[string]int{"labels": 1}},
		{
			name: "independent_failures",
			change: func(v *AlertPolicyDraft) {
				v.Name = "ab"
				v.Description = govyEvalPtr(strings.Repeat("d", 201))
				v.Severity = "urgent"
				v.Labels = []string{"x", "x"}
			},
			want: map[string]int{"name": 1, "description": 1, "severity": 1, "labels": 1},
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
