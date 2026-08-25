package rules

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonschema"
)

// DurationPrecision ensures the duration is defined with the specified precision.
func DurationPrecision(precision time.Duration) govy.Rule[time.Duration] {
	if precision <= 0 {
		panic("precision must be greater than 0")
	}
	tpl := messagetemplates.Get(messagetemplates.DurationPrecisionTemplate)

	return govy.NewRule(func(v time.Duration) error {
		if v%precision != 0 {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   v,
				ComparisonValue: precision,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeDurationPrecision).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			PropertyValue: precision,
		})).
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			value, err := jsonSchemaValue(precision)
			if err != nil {
				return nil, err
			}
			number, ok := value.(json.Number)
			if !ok {
				return nil, fmt.Errorf(
					"value marshaled as %T instead of a JSON number",
					value,
				)
			}
			return &jsonschema.Schema{MultipleOf: number}, nil
		})
}
