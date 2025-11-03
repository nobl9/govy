package example_test

import (
	"fmt"
	"reflect"
	"regexp"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

type Service struct {
	APIVersion string          `json:"apiVersion" validate:"required,eq=n9/v1alpha"`
	Kind       string          `json:"kind"       validate:"required,eq=Service"`
	Metadata   ServiceMetadata `json:"metadata"`
	Spec       ServiceSpec     `json:"spec"`
}

type ServiceMetadata struct {
	Name        string `json:"name"                  validate:"required,objectName"`
	DisplayName string `json:"displayName,omitempty" validate:"omitempty,min=0,max=63"`
	Project     string `json:"project"               validate:"required,objectName"`
	Labels      Labels `json:"labels"                validate:"omitempty,labels"`
}

type ServiceSpec struct {
	Description string `json:"description" validate:"description"`
}

type Labels map[string][]string

var serviceInstance = Service{
	APIVersion: "n9/v2alpha",
	Kind:       "Project",
	Metadata: ServiceMetadata{
		Name:        "slo-status api",
		DisplayName: "SLO Status API",
		Project:     "default project",
		Labels: Labels{
			"ke y": {"value1", "value 2"},
		},
	},
	Spec: ServiceSpec{
		Description: "Status API allows users to retrieve the latest metrics for defined SLOs",
	},
}

func Example_goPlaygroundValidator() {
	const (
		minLabelKeyLength   = 1
		maxLabelKeyLength   = 63
		maxLabelValueLength = 200
	)

	var (
		labelKeyRegexp            = regexp.MustCompile(`^\p{Ll}([_\-0-9\p{Ll}]*[0-9\p{Ll}])?$`)
		hasUpperCaseLettersRegexp = regexp.MustCompile(`[A-Z]+`)
		dns1123LabelRegexp        = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`)
	)

	isValidObjectName := func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		if len(value) > 63 {
			return false
		}
		if !dns1123LabelRegexp.MatchString(value) {
			return false
		}
		return true
	}

	isValidDescription := func(fl validator.FieldLevel) bool {
		return utf8.RuneCountInString(fl.Field().String()) <= 1050
	}

	isLabelKeyValid := func(key string) bool {
		if len(key) > maxLabelKeyLength || len(key) < minLabelKeyLength {
			return false
		}
		if !labelKeyRegexp.MatchString(key) {
			return false
		}
		if hasUpperCaseLettersRegexp.MatchString(key) {
			return false
		}
		return true
	}

	isLabelValueValid := func(value string) bool {
		return len(value) == 0 || utf8.RuneCountInString(value) <= maxLabelValueLength
	}

	areLabelsUnique := func(labelValues []string) bool {
		uniqueValues := make(map[string]struct{})
		for _, value := range labelValues {
			if _, exists := uniqueValues[value]; exists {
				return false
			}
			uniqueValues[value] = struct{}{}
		}
		return true
	}

	areValidLabels := func(fl validator.FieldLevel) bool {
		labels := fl.Field().Interface().(Labels)
		for key, values := range labels {
			if !isLabelKeyValid(key) {
				return false
			}
			if !areLabelsUnique(values) {
				return false
			}
			for _, value := range values {
				if !isLabelValueValid(value) {
					return false
				}
			}
		}
		return true
	}

	goPlaygroundValidator := func() *validator.Validate {
		v := validator.New()

		_ = v.RegisterValidation("description", isValidDescription)
		_ = v.RegisterValidation("objectName", isValidObjectName)
		_ = v.RegisterValidation("labels", areValidLabels)

		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := field.Tag.Get("json")
			if name == "" {
				name = field.Tag.Get("validate")
			}
			return name
		})
		return v
	}()

	if err := goPlaygroundValidator.Struct(serviceInstance); err != nil {
		fmt.Println(err)
	}

	// Output:
	// Key: 'Service.apiVersion' Error:Field validation for 'apiVersion' failed on the 'eq' tag
	// Key: 'Service.kind' Error:Field validation for 'kind' failed on the 'eq' tag
	// Key: 'Service.metadata.name' Error:Field validation for 'name' failed on the 'objectName' tag
	// Key: 'Service.metadata.project' Error:Field validation for 'project' failed on the 'objectName' tag
	// Key: 'Service.metadata.labels' Error:Field validation for 'labels' failed on the 'labels' tag
}

func Example_govy() {
	const (
		minLabelKeyLength   = 1
		maxLabelKeyLength   = 63
		maxLabelValueLength = 200
	)

	labelKeyRegexp := regexp.MustCompile(`^\p{Ll}([_\-0-9\p{Ll}]*[0-9\p{Ll}])?$`)

	labelsValidator := govy.New(
		govy.ForMap(govy.GetSelf[Labels]()).
			RulesForKeys(
				rules.StringLength(minLabelKeyLength, maxLabelKeyLength),
				rules.StringMatchRegexp(labelKeyRegexp),
			).
			IncludeForValues(govy.New(
				govy.ForSlice(govy.GetSelf[[]string]()).
					Rules(rules.SliceUnique(rules.HashFuncSelf[string]())).
					RulesForEach(
						rules.StringMaxLength(maxLabelValueLength),
					),
			)),
	)

	govyValidator := govy.New(
		govy.For(func(s Service) string { return s.APIVersion }).
			WithName("apiVersion").
			Required().
			Rules(rules.EQ("n9/v1alpha")),
		govy.For(func(s Service) string { return s.Kind }).
			WithName("kind").
			Required().
			Rules(rules.EQ("Service")),
		govy.For(func(s Service) ServiceMetadata { return s.Metadata }).
			WithName("metadata").
			Include(govy.New(
				govy.For(func(m ServiceMetadata) string { return m.Name }).
					WithName("name").
					Required().
					Rules(rules.StringDNSLabel()),
				govy.For(func(m ServiceMetadata) string { return m.DisplayName }).
					WithName("displayName").
					OmitEmpty().
					Rules(rules.StringMaxLength(63)),
				govy.For(func(m ServiceMetadata) string { return m.Project }).
					WithName("project").
					Required().
					Rules(rules.StringDNSLabel()),
				govy.For(func(m ServiceMetadata) Labels { return m.Labels }).
					WithName("labels").
					OmitEmpty().
					Include(labelsValidator),
			)),
		govy.For(func(s Service) ServiceSpec { return s.Spec }).
			WithName("spec").
			Include(govy.New(
				govy.For(func(s ServiceSpec) string { return s.Description }).
					WithName("description").
					Required().
					Rules(rules.StringMaxLength(1050)),
			)),
	)

	if err := govyValidator.Validate(serviceInstance); err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'apiVersion' with value 'n9/v2alpha':
	//     - should be equal to 'n9/v1alpha'
	//   - 'kind' with value 'Project':
	//     - should be equal to 'Service'
	//   - 'metadata.name' with value 'slo-status api':
	//     - string must match regular expression: '^[a-z0-9]([-a-z0-9]*[a-z0-9])?$' (e.g. 'my-name', '123-abc'); an RFC-1123 compliant label name must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character
	//   - 'metadata.project' with value 'default project':
	//     - string must match regular expression: '^[a-z0-9]([-a-z0-9]*[a-z0-9])?$' (e.g. 'my-name', '123-abc'); an RFC-1123 compliant label name must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character
	//   - 'metadata.labels.['ke y']' with key 'ke y':
	//     - string must match regular expression: '^\p{Ll}([_\-0-9\p{Ll}]*[0-9\p{Ll}])?$'
}
