package rules

import (
	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonschema"
)

// Required ensures the property's value is not empty (i.e. it's not its type's zero value).
func Required[T any]() govy.Rule[T] {
	tpl := messagetemplates.Get(messagetemplates.RequiredTemplate)

	return govy.NewRule(func(v T) error {
		if internal.IsEmpty(v) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: v,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeRequired).
		WithMessageTemplate(tpl).
		WithDescription(internal.RequiredDescription).
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return nil, nil
		})
}
