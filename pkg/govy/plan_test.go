package govy_test

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonpath"
	"github.com/nobl9/govy/pkg/rules"
)

type Kind string

type Pod struct {
	APIVersion string      `json:"apiVersion"`
	Kind       Kind        `json:"kind"`
	Metadata   PodMetadata `json:"metadata"`
	Spec       PodSpec     `json:"spec"`
	Status     *PodStatus  `json:"status,omitempty"`
}

type PodMetadata struct {
	Name        string      `json:"name"`
	Namespace   string      `json:"namespace"`
	Labels      Labels      `json:"labels"`
	Annotations Annotations `json:"annotations"`
}

type Labels map[string]string

type Annotations map[string]string

type PodSpec struct {
	DNSPolicy  string      `json:"dnsPolicy"`
	Containers []Container `json:"containers"`
}

type Container struct {
	Name  string   `json:"name"`
	Image string   `json:"image"`
	Env   []EnvVar `json:"env"`
}

type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type PodStatus struct {
	HostIP string `json:"hostIP"`
}

func TestPlan(t *testing.T) {
	metadataValidator := govy.New(
		govy.For(func(p PodMetadata) string { return p.Name }).
			WithName("name").
			Required().
			Rules(rules.StringNotEmpty()),
		govy.For(func(p PodMetadata) string { return p.Namespace }).
			WithName("namespace").
			Required(),
		govy.ForMap(func(p PodMetadata) Labels { return p.Labels }).
			WithName("labels").
			Rules(rules.MapMaxLength[Labels](10)).
			RulesForKeys(rules.StringDNSLabel()).
			RulesForValues(rules.StringMaxLength(120)),
		govy.ForMap(func(p PodMetadata) Annotations { return p.Annotations }).
			WithName("annotations").
			Rules(rules.MapMaxLength[Annotations](10)).
			RulesForItems(
				govy.NewRule(func(a govy.MapItem[string, string]) error {
					if a.Key == a.Value {
						return errors.New("key and value must not be equal")
					}
					return nil
				}).WithDescription("key and value must not be equal"),
			),
	)

	specValidator := govy.New(
		govy.For(func(p PodSpec) string { return p.DNSPolicy }).
			WithName("dnsPolicy").
			OmitEmpty().
			Rules(rules.OneOf("ClusterFirst", "Default")),
		govy.ForSlice(func(p PodSpec) []Container { return p.Containers }).
			WithName("containers").
			Rules(
				rules.SliceMaxLength[[]Container](10),
				rules.SliceUnique(func(c Container) string { return c.Name }),
			).
			IncludeForEach(govy.New(
				govy.For(func(c Container) string { return c.Name }).
					WithName("name").
					Required().
					Rules(rules.StringDNSLabel()),
				govy.For(func(c Container) string { return c.Image }).
					WithName("image").
					Required().
					Rules(rules.StringNotEmpty()),
				govy.ForSlice(func(c Container) []EnvVar { return c.Env }).
					WithName("env").
					RulesForEach(
						govy.NewRule(func(e EnvVar) error {
							return nil
						}).WithDescription("custom error!"),
					),
			)),
	)

	validator := govy.New(
		govy.For(func(p Pod) string { return p.APIVersion }).
			WithName("apiVersion").
			Required().
			Rules(rules.OneOf("v1", "v2")),
		govy.For(func(p Pod) Kind { return p.Kind }).
			WithName("kind").
			Required().
			Rules(rules.EQ[Kind]("Pod")),
		govy.For(func(p Pod) PodMetadata { return p.Metadata }).
			WithName("metadata").
			Required().
			Include(metadataValidator),
		govy.For(func(p Pod) PodSpec { return p.Spec }).
			WithName("spec").
			Required().
			Include(specValidator),
	).
		WithName("Pod")

	plan, err := govy.Plan(validator)
	assert.Require(t, assert.NoError(t, err))

	actual := requireJSON(t, plan)
	assert.Equal(t, readExpectedPlan(t, "expected_pod_plan.json"), actual)
}

func TestPlan_mapKeyRulesUseStandardWildcardPath(t *testing.T) {
	type Labels map[string]string
	type PodMetadata struct {
		Labels Labels `json:"labels"`
	}

	validator := govy.New(
		govy.ForMap(func(p PodMetadata) Labels { return p.Labels }).
			WithName("labels").
			RulesForKeys(
				govy.NewRule(func(v string) error { return nil }).
					WithDescription("key rule"),
			).
			RulesForValues(
				govy.NewRule(func(v string) error { return nil }).
					WithDescription("value rule"),
			),
	)

	plan, err := govy.Plan(validator)
	assert.Require(t, assert.NoError(t, err))

	var keyPlan, valuePlan *govy.PropertyPlan
	for _, property := range plan.Properties {
		if property.Path.String() != "$.labels.*" {
			continue
		}
		if property.IsKey {
			keyPlan = property
			continue
		}
		valuePlan = property
	}

	if keyPlan == nil {
		t.Fatal("expected a map key property plan for $.labels.*")
	}
	if valuePlan == nil {
		t.Fatal("expected a map value property plan for $.labels.*")
	}

	assert.True(t, keyPlan.IsKey)
	assert.False(t, valuePlan.IsKey)
	assert.Equal(t, govy.TypeInfo{Name: "string", Kind: "string"}, keyPlan.TypeInfo)
	assert.Equal(t, govy.TypeInfo{Name: "string", Kind: "string"}, valuePlan.TypeInfo)
	assert.Equal(t, []govy.RulePlan{{Description: "key rule"}}, keyPlan.Rules)
	assert.Equal(t, []govy.RulePlan{{Description: "value rule"}}, valuePlan.Rules)
}

func TestPropertyPlan_Compare(t *testing.T) {
	tests := map[string]struct {
		left     *govy.PropertyPlan
		right    *govy.PropertyPlan
		expected int
	}{
		"path comes first": {
			left:     &govy.PropertyPlan{Path: jsonpath.Parse("$.a")},
			right:    &govy.PropertyPlan{Path: jsonpath.Parse("$.b")},
			expected: -1,
		},
		"key entries sort before value entries": {
			left: &govy.PropertyPlan{
				Path:     jsonpath.Parse("$.labels.*"),
				IsKey:    true,
				TypeInfo: govy.TypeInfo{Name: "string", Kind: "string"},
			},
			right: &govy.PropertyPlan{
				Path:     jsonpath.Parse("$.labels.*"),
				TypeInfo: govy.TypeInfo{Name: "string", Kind: "string"},
			},
			expected: -1,
		},
		"type package breaks ties": {
			left: &govy.PropertyPlan{
				Path:     jsonpath.Parse("$.labels.*"),
				TypeInfo: govy.TypeInfo{Name: "Alpha", Kind: "string", Package: "a/pkg"},
			},
			right: &govy.PropertyPlan{
				Path:     jsonpath.Parse("$.labels.*"),
				TypeInfo: govy.TypeInfo{Name: "Alpha", Kind: "string", Package: "b/pkg"},
			},
			expected: -1,
		},
		"type name breaks ties": {
			left: &govy.PropertyPlan{
				Path:     jsonpath.Parse("$.labels.*"),
				TypeInfo: govy.TypeInfo{Name: "Alpha", Kind: "string"},
			},
			right: &govy.PropertyPlan{
				Path:     jsonpath.Parse("$.labels.*"),
				TypeInfo: govy.TypeInfo{Name: "Beta", Kind: "string"},
			},
			expected: -1,
		},
		"type kind breaks final ties": {
			left: &govy.PropertyPlan{
				Path:     jsonpath.Parse("$.labels.*"),
				TypeInfo: govy.TypeInfo{Name: "Alpha", Kind: "map"},
			},
			right: &govy.PropertyPlan{
				Path:     jsonpath.Parse("$.labels.*"),
				TypeInfo: govy.TypeInfo{Name: "Alpha", Kind: "string"},
			},
			expected: -1,
		},
		"equal plans compare equal": {
			left: &govy.PropertyPlan{
				Path:     jsonpath.Parse("$.labels.*"),
				IsKey:    true,
				TypeInfo: govy.TypeInfo{Name: "Alpha", Kind: "string", Package: "pkg"},
			},
			right: &govy.PropertyPlan{
				Path:     jsonpath.Parse("$.labels.*"),
				IsKey:    true,
				TypeInfo: govy.TypeInfo{Name: "Alpha", Kind: "string", Package: "pkg"},
			},
			expected: 0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.left.Compare(tc.right))
			assert.Equal(t, -tc.expected, tc.right.Compare(tc.left))
		})
	}
}

func TestPlan_validValuesIntersection(t *testing.T) {
	validator := govy.New(
		govy.For(func(p PodMetadata) string { return p.Name }).
			Rules(rules.NEQ("baz")),
		govy.For(func(p PodMetadata) string { return p.Name }).
			Rules(rules.OneOf("foo", "bar")),
		govy.For(func(p PodMetadata) string { return p.Name }).
			Rules(rules.EQ("foo")),
		govy.For(func(p PodMetadata) string { return p.Name }).
			Rules(rules.OneOf("bar", "baz", "foo")),
	)

	plan, err := govy.Plan(validator)
	assert.Require(t, assert.NoError(t, err))

	actual := requireJSON(t, plan)
	assert.Equal(t, readExpectedPlan(t, "expected_values_intersection_plan.json"), actual)
}

func TestPlan_customSliceType(t *testing.T) {
	type Slice []string
	type Foo struct {
		Slice Slice
	}

	validator := govy.New(
		govy.For(func(f Foo) Slice { return f.Slice }).
			Rules(rules.SliceMaxLength[Slice](1)),
	)

	plan, err := govy.Plan(validator)
	assert.Require(t, assert.NoError(t, err))

	actual := requireJSON(t, plan)
	assert.Equal(t, readExpectedPlan(t, "expected_custom_slice_type_plan.json"), actual)
}

func TestPlan_conditionsWithoutRules(t *testing.T) {
	type Slice []string
	type Foo struct {
		Slice Slice
	}

	validator := govy.New(
		govy.For(func(f Foo) Slice { return f.Slice }).
			WithName("Slice").
			Include(govy.New(
				govy.For(govy.GetSelf[Slice]()).
					Rules(rules.SliceMinLength[Slice](2)),
			)).
			When(func(f Foo) bool { return true }, govy.WhenDescription("when true")),
		govy.For(func(f Foo) Slice { return f.Slice }).
			WithName("Slice").
			Include(govy.New(
				govy.For(govy.GetSelf[Slice]()).
					Rules(rules.SliceMaxLength[Slice](1)),
			)).
			When(func(f Foo) bool { return true }),
		govy.For(func(f Foo) Slice { return f.Slice }).
			WithName("Slice").
			When(func(f Foo) bool { return true }, govy.WhenDescription("when true")),
		govy.For(func(f Foo) Slice { return f.Slice }).
			WithName("Slice").
			When(func(f Foo) bool { return true }),
	)

	plan, err := govy.Plan(validator)
	assert.Require(t, assert.NoError(t, err))

	actual := requireJSON(t, plan)
	assert.Equal(t, readExpectedPlan(t, "expected_conditions_without_rules_plan.json"), actual)
}

func TestPlan_removeDuplicateRules(t *testing.T) {
	validator := govy.New(
		govy.For(func(s string) string { return s }).
			Required().
			Rules(rules.StringASCII()).
			WithName("String"),
		govy.For(func(s string) string { return s }).
			Required().
			WithName("String"),
	)

	plan, err := govy.Plan(validator)
	assert.Require(t, assert.NoError(t, err))

	actual := requireJSON(t, plan)
	assert.Equal(t, readExpectedPlan(t, "expected_remove_duplicate_rules_plan.json"), actual)
}

func TestPlan_optionalProperties(t *testing.T) {
	validator := govy.New(
		govy.For(func(s string) string { return s }).
			OmitEmpty().
			Rules(rules.StringASCII()).
			WithName("String2"),
		govy.ForPointer(func(s string) *string { return &s }).
			WithName("String1").
			When(func(s string) bool { return true }, govy.WhenDescription("when true")),
	)

	plan, err := govy.Plan(validator)
	assert.Require(t, assert.NoError(t, err))

	actual := requireJSON(t, plan)
	assert.Equal(t, readExpectedPlan(t, "expected_optional_properties_plan.json"), actual)
}

func requireJSON(t *testing.T, plan *govy.ValidatorPlan) string {
	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	err := enc.Encode(plan)
	assert.Require(t, assert.NoError(t, err))
	return buf.String()
}

func TestPlanRequirePredicateDescriptions(t *testing.T) {
	t.Run("validator level predicate without description", func(t *testing.T) {
		validator := govy.New(
			govy.For(func(s string) string { return s }).
				Required().
				WithName("String"),
		).When(func(s string) bool { return true })

		_, err := govy.Plan(validator, govy.PlanRequirePredicateDescription())
		assert.Require(t, assert.Error(t, err))

		assert.Equal(t, "predicates without description found at: validator level", err.Error())
	})

	t.Run("property level predicate without description", func(t *testing.T) {
		validator := govy.New(
			govy.For(func(s string) string { return s }).
				Required().
				WithName("String").
				When(func(s string) bool { return true }),
		)

		_, err := govy.Plan(validator, govy.PlanRequirePredicateDescription())
		assert.Require(t, assert.Error(t, err))

		assert.Equal(t, "predicates without description found at: $.String", err.Error())
	})

	t.Run("multiple predicates without description", func(t *testing.T) {
		type Foo struct {
			Name string
			Age  int
		}
		validator := govy.New(
			govy.For(func(f Foo) string { return f.Name }).
				Required().
				WithName("name").
				When(func(f Foo) bool { return true }),
			govy.For(func(f Foo) int { return f.Age }).
				Required().
				WithName("age").
				When(func(f Foo) bool { return true }),
		).When(func(f Foo) bool { return true })

		_, err := govy.Plan(validator, govy.PlanRequirePredicateDescription())
		assert.Require(t, assert.Error(t, err))

		assert.Equal(t, "predicates without description found at: validator level, $.name, $.age", err.Error())
	})

	t.Run("all predicates have descriptions", func(t *testing.T) {
		validator := govy.New(
			govy.For(func(s string) string { return s }).
				Required().
				WithName("String").
				When(func(s string) bool { return true }, govy.WhenDescription("when true")),
		).When(func(s string) bool { return true }, govy.WhenDescription("validator when true"))

		_, err := govy.Plan(validator, govy.PlanRequirePredicateDescription())
		assert.Require(t, assert.NoError(t, err))
	})

	t.Run("disabled by default", func(t *testing.T) {
		validator := govy.New(
			govy.For(func(s string) string { return s }).
				Required().
				WithName("String").
				When(func(s string) bool { return true }),
		).When(func(s string) bool { return true })

		_, err := govy.Plan(validator)
		assert.Require(t, assert.NoError(t, err))
	})
}

func TestPlanStrictMode(t *testing.T) {
	t.Run("invalid - all validations fail", func(t *testing.T) {
		validator := govy.New(
			govy.For(func(s string) string { return s }).
				Required().
				WithName("String"),
		).When(func(s string) bool { return true })

		_, err := govy.Plan(validator, govy.PlanStrictMode())
		assert.Require(t, assert.Error(t, err))

		assert.Equal(t, "predicates without description found at: validator level", err.Error())
	})

	t.Run("valid", func(t *testing.T) {
		validator := govy.New(
			govy.For(func(s string) string { return s }).
				Required().
				WithName("String").
				When(func(s string) bool { return true }, govy.WhenDescription("when true")),
		).When(func(s string) bool { return true }, govy.WhenDescription("validator when true"))

		_, err := govy.Plan(validator, govy.PlanStrictMode())
		assert.Require(t, assert.NoError(t, err))
	})
}

func readExpectedPlan(t *testing.T, name string) string {
	filename := filepath.Join(internal.FindModuleRoot(), "pkg", "govy", "test_data", name)
	data, err := os.ReadFile(filename)
	assert.Require(t, assert.NoError(t, err))
	return string(data)
}
