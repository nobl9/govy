package validation

import (
	"fmt"
	"testing"

	"github.com/nobl9/govy/pkg/govy"
)

func TestGovyEvalNestedCollections(t *testing.T) {
	var validator govy.Validator[Environment] = EnvironmentValidator
	projects := func(count int) []Project {
		values := make([]Project, count)
		for i := range values {
			values[i] = Project{Name: fmt.Sprintf("project-%d", i)}
		}
		return values
	}
	withObjective := func(objective float64) Environment {
		return Environment{Projects: []Project{{Name: "valid", SLOs: []SLO{{Objective: objective}}}}}
	}
	tests := []struct {
		name  string
		input Environment
		want  map[string]int
	}{
		{name: "nil_projects"},
		{name: "empty_projects", input: Environment{Projects: []Project{}}},
		{name: "objective_zero", input: withObjective(0)},
		{name: "objective_maximum", input: withObjective(100)},
		{name: "objective_fraction", input: withObjective(50.25)},
		{name: "objective_negative", input: withObjective(-1), want: map[string]int{"projects[0].slos[0].objective": 1}},
		{name: "objective_too_high", input: withObjective(101), want: map[string]int{"projects[0].slos[0].objective": 1}},
		{name: "objective_just_below_zero", input: withObjective(-0.001), want: map[string]int{"projects[0].slos[0].objective": 1}},
		{name: "objective_just_above_maximum", input: withObjective(100.001), want: map[string]int{"projects[0].slos[0].objective": 1}},
		{name: "projects_maximum", input: Environment{Projects: projects(20)}},
		{name: "projects_too_many", input: Environment{Projects: projects(21)}, want: map[string]int{"projects": 1}},
		{name: "slos_maximum", input: Environment{Projects: []Project{{Name: "valid", SLOs: make([]SLO, 100)}}}},
		{name: "slos_too_many", input: Environment{Projects: []Project{{Name: "valid", SLOs: make([]SLO, 101)}}}, want: map[string]int{"projects[0].slos": 1}},
		{name: "project_name_required", input: Environment{Projects: []Project{{}}}, want: map[string]int{"projects[0].name": 1}},
		{
			name:  "second_slo_path",
			input: Environment{Projects: []Project{{Name: "valid", SLOs: []SLO{{Objective: 0}, {Objective: 101}}}}},
			want:  map[string]int{"projects[0].slos[1].objective": 1},
		},
		{
			name: "independent_failures_with_valid_neighbors",
			input: Environment{Projects: []Project{
				{Name: "", SLOs: []SLO{{Objective: 0}, {Objective: 101}}},
				{Name: "valid", SLOs: []SLO{{Objective: 100}}},
				{Name: "also-valid", SLOs: []SLO{{Objective: -1}}},
			}},
			want: map[string]int{"projects[0].name": 1, "projects[0].slos[1].objective": 1, "projects[2].slos[0].objective": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			govyEvalAssertErrors(t, validator.Validate(tt.input), tt.want)
		})
	}
}
