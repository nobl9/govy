package rules

import (
	"time"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// DurationPrecision ensures the duration is defined with the specified precision.
func DurationPrecision(precision time.Duration) govy.Rule[time.Duration] {
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
		}))
}
