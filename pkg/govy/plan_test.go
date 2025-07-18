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

	plan := govy.Plan(validator)

	actual := requireJSON(t, plan)
	assert.Equal(t, readExpectedPlan(t, "expected_pod_plan.json"), actual)
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

	plan := govy.Plan(validator)

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

	plan := govy.Plan(validator)

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

	plan := govy.Plan(validator)

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

	plan := govy.Plan(validator)

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

	plan := govy.Plan(validator)

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

func readExpectedPlan(t *testing.T, name string) string {
	filename := filepath.Join(internal.FindModuleRoot(), "pkg", "govy", "test_data", name)
	data, err := os.ReadFile(filename)
	assert.Require(t, assert.NoError(t, err))
	return string(data)
}
