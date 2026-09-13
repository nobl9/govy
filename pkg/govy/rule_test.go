package govy_test

import (
	"errors"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"text/template"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

func TestRule(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		r := govy.NewRule(func(v int) error {
			if v < 0 {
				return errors.New("must be positive")
			}
			return nil
		})

		err := r.Validate(0)
		assert.NoError(t, err)
		err = r.Validate(-1)
		assert.EqualError(t, err, "must be positive")
	})
	t.Run("full", func(t *testing.T) {
		r := govy.NewRule(func(v string) error {
			return errors.New("error")
		}).
			WithMessagef("my message %s", "foo").
			WithDetailsf("some details %d", 1).
			WithExamples("foo", "bar")

		err := r.Validate("baz")
		assert.Require(t, assert.Error(t, err))
		assert.Equal(t, "my message foo (e.g. 'foo', 'bar'); some details 1", err.Error())
	})
}

func TestRule_WithErrorCode(t *testing.T) {
	r := govy.NewRule(func(v int) error {
		if v < 0 {
			return errors.New("must be positive")
		}
		return nil
	}).WithErrorCode("test")

	err := r.Validate(0)
	assert.NoError(t, err)
	err = r.Validate(-1)
	assert.EqualError(t, err, "must be positive")
	assert.Equal(t, govy.ErrorCode("test"), err.(*govy.RuleError).Code)
}

func TestRule_WithMessage(t *testing.T) {
	tests := []struct {
		Error         string
		Message       string
		ExpectedError string
	}{
		{
			Error:         "this is error",
			Message:       "",
			ExpectedError: "this is error",
		},
		{
			Error:         "this is error",
			Message:       "this is message",
			ExpectedError: "this is message",
		},
		{
			Error:         "",
			Message:       "message",
			ExpectedError: "message",
		},
	}

	for _, test := range tests {
		r := govy.NewRule(func(v int) error {
			if v < 0 {
				return errors.New(test.Error)
			}
			return nil
		}).
			WithErrorCode("test").
			WithMessage(test.Message)

		err := r.Validate(0)
		assert.NoError(t, err)
		err = r.Validate(-1)
		assert.EqualError(t, err, test.ExpectedError)
		assert.Equal(t, govy.ErrorCode("test"), err.(*govy.RuleError).Code)
	}
}

func TestRule_WithMessagef(t *testing.T) {
	tests := []struct {
		Error         string
		Message       string
		MessageArgs   []any
		ExpectedError string
	}{
		{
			Error:         "",
			Message:       "message %s %d",
			MessageArgs:   []any{"foo", 1},
			ExpectedError: "message foo 1",
		},
		{
			Error:         "this is error",
			Message:       "message %s",
			MessageArgs:   []any{"bar"},
			ExpectedError: "message bar",
		},
	}

	for _, test := range tests {
		r := govy.NewRule(func(v int) error {
			if v < 0 {
				return errors.New(test.Error)
			}
			return nil
		}).
			WithErrorCode("test").
			WithMessagef(test.Message, test.MessageArgs...)

		err := r.Validate(0)
		assert.NoError(t, err)
		err = r.Validate(-1)
		assert.EqualError(t, err, test.ExpectedError)
		assert.Equal(t, govy.ErrorCode("test"), err.(*govy.RuleError).Code)
	}
}

func TestRule_WithDetails(t *testing.T) {
	tests := []struct {
		Error         string
		Details       string
		ExpectedError string
	}{
		{
			Error:         "this is error",
			Details:       "details",
			ExpectedError: "this is error; details",
		},
		{
			Error:         "this is error",
			Details:       "",
			ExpectedError: "this is error",
		},
		{
			Error:         "",
			Details:       "details",
			ExpectedError: "details",
		},
	}

	for _, test := range tests {
		r := govy.NewRule(func(v int) error {
			if v < 0 {
				return errors.New(test.Error)
			}
			return nil
		}).
			WithErrorCode("test").
			WithDetails(test.Details)

		err := r.Validate(0)
		assert.NoError(t, err)
		err = r.Validate(-1)
		assert.EqualError(t, err, test.ExpectedError)
		assert.Equal(t, govy.ErrorCode("test"), err.(*govy.RuleError).Code)
	}
}

func TestRule_WithDetailsf(t *testing.T) {
	tests := []struct {
		Error         string
		Details       string
		DetailsArgs   []any
		ExpectedError string
	}{
		{
			Error:         "",
			Details:       "details %s %d",
			DetailsArgs:   []any{"foo", 1},
			ExpectedError: "details foo 1",
		},
		{
			Error:         "this is error",
			Details:       "details %s",
			DetailsArgs:   []any{"bar"},
			ExpectedError: "this is error; details bar",
		},
	}

	for _, test := range tests {
		r := govy.NewRule(func(v int) error {
			if v < 0 {
				return errors.New(test.Error)
			}
			return nil
		}).
			WithErrorCode("test").
			WithDetailsf(test.Details, test.DetailsArgs...)

		err := r.Validate(0)
		assert.NoError(t, err)
		err = r.Validate(-1)
		assert.EqualError(t, err, test.ExpectedError)
		assert.Equal(t, govy.ErrorCode("test"), err.(*govy.RuleError).Code)
	}
}

func TestRule_WithDescription(t *testing.T) {
	r := govy.NewRule(func(v int) error {
		if v < 0 {
			return errors.New("must be positive")
		}
		return nil
	}).
		WithErrorCode("test").
		WithDetails("some details").
		WithDescription("the integer must be positive")

	err := r.Validate(-1)
	assert.Require(t, assert.Error(t, err))
	assert.Equal(t, &govy.RuleError{
		Message:     "must be positive; some details",
		Code:        "test",
		Description: "the integer must be positive",
	}, err)
}

func TestRule_WithDescriptionTemplate(t *testing.T) {
	t.Run("nil template panics during configuration", func(t *testing.T) {
		defer func() {
			if recovered := recover(); recovered != "description template must not be nil" {
				t.Fatalf("unexpected panic: %v", recovered)
			}
		}()
		_ = govy.NewRule(func(int) error { return nil }).
			WithDescriptionTemplate(nil, govy.TemplateVars{})
	})

	t.Run("deferred and cached", func(t *testing.T) {
		var executions atomic.Int32
		tpl := template.Must(template.New("description").
			Funcs(template.FuncMap{
				"render": func(value string) string {
					executions.Add(1)
					return value
				},
			}).
			Parse("must be {{ render .Custom.Requirement }}"))
		requirements := map[string]string{"Requirement": "positive"}
		vars := govy.TemplateVars{Custom: requirements}
		rule := govy.NewRule(func(v int) error {
			if v < 0 {
				return errors.New("invalid")
			}
			return nil
		}).WithDescriptionTemplate(tpl, vars)

		assert.NoError(t, rule.Validate(1))
		assert.Equal(t, int32(0), executions.Load())

		err := rule.Validate(-1)
		assert.Require(t, assert.Error(t, err))
		assert.Equal(t, "must be positive", err.(*govy.RuleError).Description)
		assert.Equal(t, int32(1), executions.Load())

		err = rule.Validate(-1)
		assert.Require(t, assert.Error(t, err))
		assert.Equal(t, "must be positive", err.(*govy.RuleError).Description)

		validator := govy.New(
			govy.For(func(value int) int { return value }).
				WithName("value").
				Rules(rule),
		)
		plan, planErr := govy.Plan(validator)
		assert.Require(t, assert.NoError(t, planErr))
		if len(plan.Properties) != 1 || len(plan.Properties[0].Rules) != 1 {
			t.Fatalf("unexpected plan shape: %#v", plan)
		}
		assert.Equal(t, "must be positive", plan.Properties[0].Rules[0].Description)
		assert.Equal(t, int32(1), executions.Load())
	})

	t.Run("execution failure is cached and panics for every consumer", func(t *testing.T) {
		var executions atomic.Int32
		renderErr := errors.New("render failed")
		tpl := template.Must(template.New("description").
			Funcs(template.FuncMap{
				"fail": func() (string, error) {
					executions.Add(1)
					return "", renderErr
				},
			}).
			Parse("partial {{ fail }}"))
		rule := govy.NewRule(func(int) error { return errors.New("invalid") }).
			WithDescriptionTemplate(tpl, govy.TemplateVars{})
		validator := govy.New(
			govy.For(func(value int) int { return value }).
				WithName("value").
				Rules(rule),
		)
		assertExecutionPanic := func(call func()) {
			t.Helper()
			defer func() {
				recovered := recover()
				executionErr, ok := recovered.(error)
				if !ok {
					t.Fatalf("unexpected panic: %v", recovered)
				}
				if !errors.Is(executionErr, renderErr) {
					t.Fatalf("panic does not wrap the execution error: %v", executionErr)
				}
				if !strings.Contains(
					executionErr.Error(),
					`failed to execute description template "description"`,
				) {
					t.Fatalf("unexpected panic: %s", executionErr)
				}
			}()
			call()
		}

		assertExecutionPanic(func() {
			_ = rule.Validate(0)
		})
		assertExecutionPanic(func() {
			_, _ = govy.Plan(validator)
		})
		assert.Equal(t, int32(1), executions.Load())
	})

	t.Run("last description setter wins", func(t *testing.T) {
		var executions atomic.Int32
		tpl := template.Must(template.New("description").
			Funcs(template.FuncMap{
				"render": func() string {
					executions.Add(1)
					return "templated"
				},
			}).
			Parse("{{ render }}"))
		newRule := func() govy.Rule[int] {
			return govy.NewRule(func(int) error { return errors.New("invalid") })
		}

		err := newRule().
			WithDescriptionTemplate(tpl, govy.TemplateVars{}).
			WithDescription("eager").
			Validate(0)
		assert.Require(t, assert.Error(t, err))
		assert.Equal(t, "eager", err.(*govy.RuleError).Description)
		assert.Equal(t, int32(0), executions.Load())

		err = newRule().
			WithDescription("eager").
			WithDescriptionTemplate(tpl, govy.TemplateVars{}).
			Validate(0)
		assert.Require(t, assert.Error(t, err))
		assert.Equal(t, "templated", err.(*govy.RuleError).Description)
		assert.Equal(t, int32(1), executions.Load())
	})

	t.Run("plan can resolve first", func(t *testing.T) {
		var executions atomic.Int32
		tpl := template.Must(template.New("description").
			Funcs(template.FuncMap{
				"render": func() string {
					executions.Add(1)
					return "templated"
				},
			}).
			Parse("{{ render }}"))
		rule := govy.NewRule(func(int) error { return errors.New("invalid") }).
			WithDescriptionTemplate(tpl, govy.TemplateVars{})
		validator := govy.New(
			govy.For(func(value int) int { return value }).
				WithName("value").
				Rules(rule),
		)

		plan, planErr := govy.Plan(validator)
		assert.Require(t, assert.NoError(t, planErr))
		if len(plan.Properties) != 1 || len(plan.Properties[0].Rules) != 1 {
			t.Fatalf("unexpected plan shape: %#v", plan)
		}
		assert.Equal(t, "templated", plan.Properties[0].Rules[0].Description)

		err := rule.Validate(0)
		assert.Require(t, assert.Error(t, err))
		assert.Equal(t, "templated", err.(*govy.RuleError).Description)
		assert.Equal(t, int32(1), executions.Load())
	})

	t.Run("copies share concurrent resolution", func(t *testing.T) {
		var executions atomic.Int32
		tpl := template.Must(template.New("description").
			Funcs(template.FuncMap{
				"render": func() string {
					executions.Add(1)
					return "templated"
				},
			}).
			Parse("{{ render }}"))
		rule := govy.NewRule(func(v int) error {
			if v < 0 {
				return errors.New("invalid")
			}
			return nil
		}).WithDescriptionTemplate(tpl, govy.TemplateVars{})
		pointerRule := govy.RuleToPointer(rule)

		errs := make([]error, 64)
		var wg sync.WaitGroup
		for i := range errs {
			wg.Go(func() {
				if i%2 == 0 {
					errs[i] = rule.Validate(-1)
					return
				}
				value := -1
				errs[i] = pointerRule.Validate(&value)
			})
		}
		wg.Wait()

		for i, err := range errs {
			ruleErr, ok := err.(*govy.RuleError)
			if !ok {
				t.Fatalf("error %d has type %T, expected *govy.RuleError", i, err)
			}
			assert.Equal(t, "templated", ruleErr.Description)
		}
		assert.Equal(t, int32(1), executions.Load())
	})
}

func TestRule_WithDescriptionTemplateErrorBranches(t *testing.T) {
	for _, tc := range []struct {
		name                string
		validate            func(int) error
		withMessageTemplate bool
	}{
		{
			name:     "RuleError",
			validate: func(int) error { return &govy.RuleError{Message: "invalid"} },
		},
		{
			name: "RuleErrorTemplate",
			validate: func(int) error {
				return govy.NewRuleErrorTemplate(govy.TemplateVars{})
			},
			withMessageTemplate: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rule := govy.NewRule(tc.validate)
			if tc.withMessageTemplate {
				rule = rule.WithMessageTemplate(
					template.Must(template.New("message").Parse("invalid")),
				)
			}
			rule = rule.WithDescriptionTemplate(
				template.Must(template.New("description").Parse("templated")),
				govy.TemplateVars{},
			)

			err := rule.Validate(0)
			assert.Require(t, assert.Error(t, err))
			ruleErr, ok := err.(*govy.RuleError)
			if !ok {
				t.Fatalf("expected *govy.RuleError, got %T", err)
			}
			assert.Equal(t, "templated", ruleErr.Description)
		})
	}
}

func TestRule_WithExamples(t *testing.T) {
	r := govy.NewRule(func(v string) error {
		if v != "foo" && v != "bar" {
			return errors.New("must be foo or bar")
		}
		return nil
	}).
		WithErrorCode("test").
		WithDetails("some details").
		WithExamples("foo", "bar").
		WithDescription("string must be foo or bar")

	err := r.Validate("baz")
	assert.Require(t, assert.Error(t, err))
	assert.Equal(t, &govy.RuleError{
		Message:     "must be foo or bar (e.g. 'foo', 'bar'); some details",
		Code:        "test",
		Description: "string must be foo or bar",
	}, err)
}

func TestRule_WithMessageTemplate(t *testing.T) {
	tpl, err := template.New("").Parse("This is an {{ .Error }}")
	assert.Require(t, assert.NoError(t, err))

	rule := govy.NewRule(func(v string) error {
		return govy.NewRuleErrorTemplate(govy.TemplateVars{
			Error: "error",
		})
	}).
		WithErrorCode("my-code").
		WithExamples("This").
		WithMessageTemplate(tpl)

	err = rule.Validate("")
	assert.Require(t, assert.Error(t, err))
	assert.Equal(t, &govy.RuleError{
		Message: "This is an error",
		Code:    "my-code",
	}, err)
}

func TestRule_WithMessageTemplateString(t *testing.T) {
	rule := govy.NewRule(func(v string) error {
		return govy.NewRuleErrorTemplate(govy.TemplateVars{
			Error: "error",
		})
	}).
		WithErrorCode("my-code").
		WithMessageTemplateString("This is an {{ .Error }}")

	err := rule.Validate("")
	assert.Require(t, assert.Error(t, err))
	assert.Equal(t, &govy.RuleError{
		Message: "This is an error",
		Code:    "my-code",
	}, err)
}

func TestRuleToPointer(t *testing.T) {
	r := govy.NewRule(func(v int) error {
		if v < 0 {
			return errors.New("must be positive")
		}
		return nil
	}).
		WithErrorCode("test")
	rp := govy.RuleToPointer(r)
	t.Run("passes", func(t *testing.T) {
		err := rp.Validate(ptr(0))
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := rp.Validate(ptr(-1))
		assert.Require(t, assert.Error(t, err))
		t.Log(err)
		assert.Equal(t, govy.RuleError{
			Message: "must be positive",
			Code:    "test",
		}, *err.(*govy.RuleError))
	})
}
