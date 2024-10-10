package govy_test

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

//go:embed test_data/expected_pod_plan.json
var expectedPlanJSON string

type Pod struct {
	APIVersion string      `json:"apiVersion"`
	Kind       string      `json:"kind"`
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
			Required().
			Rules(rules.StringNotEmpty()),
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
			Required().
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
		govy.For(func(p Pod) string { return p.Kind }).
			WithName("kind").
			Required().
			Rules(rules.EQ("Pod")),
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

	t.Run("plan for Validator", func(t *testing.T) {
		plan := govy.Plan[Pod](validator)

		buf := bytes.Buffer{}
		enc := json.NewEncoder(&buf)
		enc.SetIndent("", "  ")
		err := enc.Encode(plan)
		assert.Require(t, assert.NoError(t, err))

		assert.Equal(t, expectedPlanJSON, buf.String())
	})

	t.Run("plan for ValidatorForSlice", func(t *testing.T) {
		validatorForSlice := govy.NewForSlice(validator)

		plan := govy.Plan[Pod](validatorForSlice)

		buf := bytes.Buffer{}
		enc := json.NewEncoder(&buf)
		enc.SetIndent("", "  ")
		err := enc.Encode(plan)
		assert.Require(t, assert.NoError(t, err))

		expectedSlicePlanJSON := expectedPlanJSON
		i := strings.Index(expectedSlicePlanJSON, `  "properties": [`)
		expectedSlicePlanJSON = expectedSlicePlanJSON[:i] + `  "isSlice": true,` + "\n" + expectedSlicePlanJSON[i:]
		expectedSlicePlanJSON = strings.ReplaceAll(expectedSlicePlanJSON, `"path": "$`, `"path": "$[*]`)

		assert.Equal(t, expectedSlicePlanJSON, buf.String())
	})
}
