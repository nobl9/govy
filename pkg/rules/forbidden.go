package rules

import (
	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonschema"
)

// Forbidden ensures the property's value is its type's zero value, i.e. it's empty.
func Forbidden[T any]() govy.Rule[T] {
	tpl := messagetemplates.Get(messagetemplates.ForbiddenTemplate)

	return govy.NewRule(func(v T) error {
		if internal.IsEmpty(v) {
			return nil
		}
		return govy.NewRuleErrorTemplate(govy.TemplateVars{
			PropertyValue: v,
		})
	}).
		WithErrorCode(ErrorCodeForbidden).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{})).
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			value, err := jsonSchemaValue(*new(T))
			if err != nil {
				return nil, err
			}
			return &jsonschema.Schema{Const: ptr(value)}, nil
		})
}
