package govy_test

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/govyconfig"
	"github.com/nobl9/govy/pkg/rules"
)

type Teacher struct {
	Name       string        `json:"name"`
	Age        time.Duration `json:"age"`
	Students   []Student     `json:"students"`
	MiddleName *string       `json:"middleName,omitempty"`
	University University    `json:"university"`
}

type University struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Student struct {
	Index string `json:"index"`
}

type Tutoring struct {
	StudentIndexToTeacher map[string]Teacher `json:"studentIndexToTeacher"`
}

const year = 24 * 365 * time.Hour

// In order to create a new [govy.Validator] use [govy.New] constructor.
// Let's define simple [govy.PropertyRules] for [Teacher.Name].
// For now, it will be always failing.
func ExampleNew() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })),
	)

	err := v.Validate(Teacher{})
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed:
	//   - always fails
}

// To associate [govy.Validator] with an entity name use [govy.Validator.WithName] function.
// When any of the rules fails, the error will contain the entity name you've provided.
func ExampleValidator_WithName() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })),
	).WithName("Teacher")

	err := v.Validate(Teacher{})
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed:
	//   - always fails
}

// If statically defined name through [govy.Validator.WithName] is not enough,
// you can use [govy.Validator.WithNameFunc].
// The function receives the entity's instance you're validating and returns a string name.
func ExampleValidator_WithNameFunc() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })),
	).WithNameFunc(func(t Teacher) string { return "Teacher " + t.Name })

	err := v.Validate(Teacher{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher John has failed:
	//   - always fails
}

// You can also add [govy.Validator] name during runtime,
// by calling [govy.ValidatorError.WithName] function on the returned error.
//
// Note: We left the previous "Teacher" name assignment, to demonstrate that
// the [govy.ValidatorError.WithName] function call will overwrite it.
//
// Note: This would also work:
//
//	err := v.WithName("Jake").Validate(Teacher{})
//
// govy, excluding error handling, tries to follow immutability principle.
// Calling any method on [govy.Validator] will not change its declared instance,
// but rather create a copy of it.
func ExampleValidatorError_WithName() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })),
	).WithName("Teacher")

	err := v.Validate(Teacher{})
	if err != nil {
		fmt.Println(err.(*govy.ValidatorError).WithName("Jake"))
	}

	// Output:
	// Validation for Jake has failed:
	//   - always fails
}

// [govy.Validator] rules can be evaluated on condition, to specify the predicate use [govy.Validator.When] function.
//
// In this example, validation for [Teacher] instance will only be evaluated
// if the [Teacher.Age] property is less than 50 years.
func ExampleValidator_When() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })),
	).
		When(func(t Teacher) bool { return t.Age < (50 * year) })

	// Prepare teachers.
	teacherTom := Teacher{
		Name: "Tom",
		Age:  51 * year,
	}
	teacherJerry := Teacher{
		Name: "Jerry",
		Age:  30 * year,
	}

	// Run validation.
	err := v.Validate(teacherTom)
	if err != nil {
		fmt.Println(err.(*govy.ValidatorError).WithName("Tom"))
	}
	err = v.Validate(teacherJerry)
	if err != nil {
		fmt.Println(err.(*govy.ValidatorError).WithName("Jerry"))
	}

	// Output:
	// Validation for Jerry has failed:
	//   - always fails
}

// All errors returned by [govy.Validator] are of type [govy.ValidatorError].
// Type casting directly to [govy.ValidatorError] should be safe once an error
// was asserted to be non-nil.
// However, you shouldn't trust any API with such promises, and always type check in your
// type assignments.
//
// All error types return by govy are JSON serializable.
func ExampleValidatorError() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })).
			WithName("name"),
	).WithName("Teacher")

	err := v.Validate(Teacher{Name: "John"})
	if err != nil {
		if validatorErr, ok := err.(*govy.ValidatorError); ok {
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			if err = enc.Encode(validatorErr); err != nil {
				fmt.Printf("error encoding: %v\n", err)
			}
		}
	}

	// Output:
	// {
	//   "errors": [
	//     {
	//       "propertyName": "name",
	//       "propertyValue": "John",
	//       "errors": [
	//         {
	//           "error": "always fails"
	//         }
	//       ]
	//     }
	//   ],
	//   "name": "Teacher"
	// }
}

// If you want to validate a slice of entities, you can combine [govy.New] with [govy.ForSlice].
// The produced errors will contain information about the failing entity's index
// in their [govy.PropertyError.PropertyName].
func ExampleValidator_Validate_slice() {
	teacherValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })),
	)
	v := govy.New(
		govy.ForSlice(govy.GetSelf[[]Teacher]()).
			IncludeForEach(teacherValidator),
	)

	err := v.Validate([]Teacher{
		{Name: "John"},
		{Name: "Jake"},
	})
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - '[0].name' with value 'John':
	//     - always fails
	//   - '[1].name' with value 'Jake':
	//     - always fails
}

// So far we've been using a very simple [govy.PropertyRules] instance:
//
//	validation.For(func(t Teacher) string { return t.Name }).
//		Rules(validation.NewRule(func(name string) error { return fmt.Errorf("always fails") }))
//
// The error message returned by this property rule does not tell us
// which property is failing.
// Let's change that by adding property name using [govy.PropertyRules.WithName].
//
// We can also change the [govy.Rule] to be something more real.
// govy comes with a number of predefined [govy.Rule], we'll use
// [rules.EQ] which accepts a single argument, value to compare with.
func ExamplePropertyRules_WithName() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.EQ("Tom")),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jake':
	//     - should be equal to 'Tom'
}

// [govy.For] constructor creates new [govy.PropertyRules] instance.
// It's only argument, [govy.PropertyGetter] is used to extract the property value.
// It works fine for direct values, but falls short when working with pointers.
// Often times we use pointers to indicate that a property is optional,
// or we want to discern between nil and zero values.
// In either case we want our validation rules to work on direct values,
// not the pointer, otherwise we'd have to always check if pointer != nil.
//
// [govy.ForPointer] constructor can be used to solve this problem and allow
// us to work with the underlying value in our rules.
// Under the hood it wraps [govy.PropertyGetter] and safely extracts the underlying value.
// If the value was nil, it will not attempt to evaluate any rules for this property.
// The rationale for that is it doesn't make sense to evaluate any rules for properties
// which are essentially empty. The only rule that makes sense in this context is to
// ensure the property is required.
// We'll learn about a way to achieve that in the next example: [ExamplePropertyRules_Required].
//
// Let's define a rule for [Teacher.MiddleName] property.
// Not everyone has to have a middle name, that's why we've defined this field
// as a pointer to string, rather than a string itself.
func ExampleForPointer() {
	v := govy.New(
		govy.ForPointer(func(t Teacher) *string { return t.MiddleName }).
			WithName("middleName").
			Rules(rules.StringMaxLength(5)),
	).WithName("Teacher")

	middleName := "Thaddeus"
	teacher := Teacher{
		Name:       "Jake",
		Age:        51 * year,
		MiddleName: &middleName,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'middleName' with value 'Thaddeus':
	//     - length must be less than or equal to 5
}

// [govy.Transform] constructor can be used to transform the property's value
// before it's passed to the rules' evaluation.
// It's useful when you want to use rules that operate on a different type than the property's.
//
// Along with the standard [govy.PropertyGetter] it accepts a [govy.Transformer] function
// which takes the property value and returns the transformed value along with an error.
// If the error is not nil, the validation will fail with the error message returned by [govy.Transformer] error.
//
// In this example we'll use [time.ParseDuration] to transform the string value of [Clock.Duration] to [time.Duration].
// The first value we'll validate will force [govy.Transformer] to return an error,
// the second will succeed transformation, but it will fail the validation for [rules.DurationPrecision].
//
// Notice how the [govy.Transformer] shape adheres to a lot of standard library conversion/parsing functions.
func ExampleTransform() {
	type Clock struct {
		Duration string `json:"duration"`
	}
	v := govy.New(
		govy.Transform(func(c Clock) string { return c.Duration }, time.ParseDuration).
			WithName("duration").
			Rules(rules.DurationPrecision(time.Minute)),
	).WithName("MyClock")

	err := v.Validate(Clock{Duration: "bad duration!"})
	if err != nil {
		fmt.Println(err)
	}

	err = v.Validate(Clock{Duration: (256 * time.Second).String()})
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for MyClock has failed for the following properties:
	//   - 'duration' with value 'bad duration!':
	//     - time: invalid duration "bad duration!"
	// Validation for MyClock has failed for the following properties:
	//   - 'duration' with value '4m16s':
	//     - duration must be defined with 1m0s precision
}

// By default, when [govy.PropertyRules] is constructed using [govy.ForPointer]
// it will skip validation of the property if the pointer is nil.
// To enforce a value is set for pointer use [govy.PropertyRules.Required].
//
// You may ask yourself why not just use [rules.Required] rule instead?
// If we were to do that, we'd be forced to operate on pointer in all of our rules.
// Other than checking if the pointer is nil, there aren't any rules which would
// benefit from working on the pointer instead of the underlying value.
//
// If you want to also make sure the underlying value is filled,
// i.e. it's not a zero value, you can also use [rules.Required] rule
// on top of [govy.PropertyRules.Required].
//
// [govy.PropertyRules.Required] when used with [govy.For] constructor, will ensure
// the property does not contain a zero value.
//
// Note: [govy.PropertyRules.Required] is introducing a short circuit.
// If the assertion fails, validation will stop and return [govy.govy.ErrorCodeRequired].
// None of the rules you've defined would be evaluated.
//
// Note: Placement of [govy.PropertyRules.Required] does not matter,
// it's not evaluated in a sequential loop, unlike standard [govy.Rule].
// However, we recommend you always place it below [govy.PropertyRules.WithName]
// to make your rules more readable.
func ExamplePropertyRules_Required() {
	alwaysFailingRule := govy.NewRule(func(string) error {
		return fmt.Errorf("always fails")
	})

	v := govy.New(
		govy.ForPointer(func(t Teacher) *string { return t.MiddleName }).
			WithName("middleName").
			Required().
			Rules(alwaysFailingRule),
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Required().
			Rules(alwaysFailingRule),
	).WithName("Teacher")

	teacher := Teacher{
		Name:       "",
		Age:        51 * year,
		MiddleName: nil,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'middleName':
	//     - property is required but was empty
	//   - 'name':
	//     - property is required but was empty
}

// While [govy.ForPointer] will by default omit validation for nil pointers,
// it might be useful to have a similar behavior for optional properties
// which are direct values.
// [govy.PropertyRules.OmitEmpty] will do the trick.
//
// Note: [govy.PropertyRules.OmitEmpty] will have no effect on pointers handled
// by [govy.ForPointer], as they already behave in the same way.
func ExamplePropertyRules_OmitEmpty() {
	alwaysFailingRule := govy.NewRule(func(string) error {
		return fmt.Errorf("always fails")
	})

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			OmitEmpty().
			Rules(alwaysFailingRule),
		govy.ForPointer(func(t Teacher) *string { return t.MiddleName }).
			WithName("middleName").
			Rules(alwaysFailingRule),
	).WithName("Teacher")

	teacher := Teacher{
		Name:       "",
		Age:        51 * year,
		MiddleName: nil,
	}

	err := v.Validate(teacher)
	if err == nil {
		fmt.Println("no error! we skipped 'name' validation and 'middleName' is implicitly skipped")
	}

	// Output:
	// no error! we skipped 'name' validation and 'middleName' is implicitly skipped
}

// Sometimes you want to hide the value of the property in the error message.
// It can contain sensitive information, like a secret access key.
// You can use [govy.PropertyRules.HideValue] to achieve that.
//
// You can see that the error message now contains "[hidden]" instead of the actual value,
// and the property value is not included in the property bullet point (- 'name').
func ExamplePropertyRules_HideValue() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			HideValue().
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("that Jake is secret") })),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name':
	//     - that [hidden] is secret
}

// If you want to access the value of the entity you're writing the [govy.Validator] for,
// you can use [govy.GetSelf] function which is a convenience [govy.PropertyGetter] that returns self.
// Note that we don't call [govy.PropertyRules.WithName] here,
// as we're comparing two properties in our top level, [Teacher] scope.
//
// You can provide your own rules using [govy.NewRule] constructor.
// It returns new [govy.Rule] instance which wraps your validation function.
func ExampleGetSelf() {
	customRule := govy.NewRule(func(v Teacher) error {
		return fmt.Errorf("now I have access to the whole teacher")
	})

	v := govy.New(
		govy.For(govy.GetSelf[Teacher]()).
			Rules(customRule),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed:
	//   - now I have access to the whole teacher
}

// Govy comes with a set of predefined rules,
// which you can use out of the box by importing [rules] package.
//
// However, you can also create your own rules by using [govy.NewRule] constructor.
// It accepts a simple validation function which takes in a value
// and returns an error if the validation failed.
//
// Note: the [govy.Rule] struct has all its fields private,
// so you can only create and modify them using exported constructor and methods.
func ExampleRule() {
	myRule := govy.NewRule(func(name string) error {
		if name != "Tom" {
			return fmt.Errorf("Teacher can be only Tom")
		}
		return nil
	})
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(myRule),
	)

	teacher := Teacher{Name: "Jake"}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'name' with value 'Jake':
	//     - Teacher can be only Tom
}

// You can use [govy.Rule.WithDetails] to add additional details to the error message.
// This allows you to extend existing rules by adding your use case context.
// Let's give a regex validation some more clarity.
func ExampleRule_WithDetails() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
				WithDetails("Teacher can be either Tom or Jerry :)")),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jake':
	//     - string must match regular expression: '^(Tom|Jerry)$'; Teacher can be either Tom or Jerry :)
}

// You can use [govy.Rule.WithDetailsf] to add formatted details to the returned [govy.RuleError] error message.
func ExampleRule_WithDetailsf() {
	minLen := 3
	maxLen := 10
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringLength(minLen, maxLen).
				WithDetailsf("Teacher name must be between %d and %d characters", minLen, maxLen)),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jo",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jo':
	//     - length must be between 3 and 10; Teacher name must be between 3 and 10 characters
}

// You can use [govy.Rule.WithExamples] to add examples of valid inputs
// which pass the [govy.Rule].
// This can be useful for more complex rules, especially regex based, where
// it might not be immediately obvious how a valid value should look like.
//
// Note: examples are added between the error message and details
// (configured with [govy.Rule.WithDetails]).
func ExampleRule_WithExamples() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
				WithDetails("Teacher can be either Tom or Jerry :)").
				WithExamples("Tom", "Jerry")),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jake':
	//     - string must match regular expression: '^(Tom|Jerry)$' (e.g. 'Tom', 'Jerry'); Teacher can be either Tom or Jerry :)
}

// When testing, it can be tedious to always rely on error messages as these can change over time.
// Enter [govy.ErrorCode], which is a simple string type alias used to ease testing,
// but also potentially allow third parties to integrate with your validation results.
// Use [govy.Rule.WithErrorCode] to associate [govy.ErrorCode] with a [govy.Rule].
// Notice that our modified version of [rules.StringMatchRegexp] will now return a new [govy.ErrorCode].
// Predefined rules have [govy.ErrorCode] already associated with them.
// To view the list of predefined [govy.ErrorCode] checkout error_codes.go file.
func ExampleRule_WithErrorCode() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
				WithDetails("Teacher can be either Tom or Jerry :)").
				WithErrorCode("custom_code")),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		propertyErrors := err.(*govy.ValidatorError).Errors
		ruleErrors := propertyErrors[0].Errors
		fmt.Println(ruleErrors[0].Code)
	}

	// Output:
	// custom_code
}

// If you want to override the default error message, you can use [govy.Rule.WithMessage].
func ExampleRule_WithMessage() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
				WithDetails("Teacher can be either Tom or Jerry :)").
				WithMessage("unsupported name")),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jake':
	//     - unsupported name; Teacher can be either Tom or Jerry :)
}

// You can use [govy.Rule.WithMessagef] to override the default error message using printf-like formatting.
func ExampleRule_WithMessagef() {
	allowedNames := []string{"Tom", "Jerry"}
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
				WithMessagef("name must be one of: %v", allowedNames)),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jake':
	//     - name must be one of: [Tom Jerry]
}

// If you want to have more control over the resulting error message, but [govy.Rule.WithMessage]
// is not enough, you can utilize a template string which is parsed by [govy.Rule] into
// [template.Template] to construct a custom error message.
//
// Each builtin rule supports different variables.
// For instance, [rules.StringLength] supports 'MinLength' and 'MaxLength' variables.
// Refer to the rule's documentation to see which variables are supported.
//
// Note: Builtin functions provided by [govy.AddTemplateFunctions], like 'formatExamples',
// are automatically added to the parsed [template.Template].
func ExampleRule_WithMessageTemplateString() {
	tplString := `Teacher's name must be between {{ .MinLength }} and {{ .MaxLength }} characters {{ formatExamples .Examples }}.`

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringLength(5, 10).
				WithExamples("Joanna", "Angeline").
				WithMessageTemplateString(tplString)),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Eve",
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Eve':
	//     - Teacher's name must be between 5 and 10 characters (e.g. 'Joanna', 'Angeline').
}

// If you want to have more control over the [template.Template] used for error message creation,
// for instance, add custom functions, use [govy.Rule.WithMessageTemplate].
//
// In the example below, we're defining a custom template function 'join' which calls [strings.Join]
// under the hood to join a slice of strings with a comma.
//
// Note: 'Examples' field is a plain slice of strings, If you wish to format it the same way
// as the default message does, use 'formatExamples' function provided by [govy.AddTemplateFunctions].
func ExampleRule_WithMessageTemplate() {
	tplString := `Teacher's name '{{ .PropertyValue }}' is not supported. {{ .Details }} (e.g. {{ join .Examples ", " }}).`
	tpl := template.New("").Funcs(template.FuncMap{"join": strings.Join})
	tpl = template.Must(tpl.Parse(tplString))

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringLength(5, 10).
				WithDetails("Teacher's name must be between 5 and 10 characters").
				WithExamples("Joanna", "Angeline").
				WithMessageTemplate(tpl)),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Eve",
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Eve':
	//     - Teacher's name 'Eve' is not supported. Teacher's name must be between 5 and 10 characters (e.g. Joanna, Angeline).
}

// Under the hood builtin rules' message templates utilize a set of custom template functions.
// If you want to use them in your custom templates, you can add them to your [template.Template]
// instance by calling [govy.AddTemplateFunctions].
//
// An example of such function is 'formatExamples' which takes in a slice of strings
// and returns a formatted string.
//
// Note: Builtin functions are automatically added to the parsed [template.Template] if you're using
// [govy.Rule.WithMessageTemplateString].
//
// Note: [govy.AddTemplateFunctions] calls [template.Template.Funcs], which will not add the functions
// to your template If it was already parsed.
func ExampleAddTemplateFunctions() {
	tplString := `Teacher's name '{{ .PropertyValue }}' is not supported {{ formatExamples .Examples }}.`
	tpl := template.New("")
	tpl = govy.AddTemplateFunctions(tpl)
	tpl = template.Must(tpl.Parse(tplString))

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringLength(5, 10).
				WithExamples("Joanna", "Angeline").
				WithMessageTemplate(tpl)),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Eve",
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Eve':
	//     - Teacher's name 'Eve' is not supported (e.g. 'Joanna', 'Angeline').
}

func ExampleAddTemplateFunctions_formatExamples() {
	tplString := "{{ formatExamples .Examples }}"
	tpl := template.New("")
	tpl = govy.AddTemplateFunctions(tpl)
	tpl = template.Must(tpl.Parse(tplString))

	err := tpl.Execute(
		os.Stdout,
		map[string]any{"Examples": []string{"Joanna", "Angeline"}},
	)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// (e.g. 'Joanna', 'Angeline')
}

func ExampleAddTemplateFunctions_joinSlice() {
	tplString := `{{ joinSlice .Slice "'" }}`
	tpl := template.New("")
	tpl = govy.AddTemplateFunctions(tpl)
	tpl = template.Must(tpl.Parse(tplString))

	err := tpl.Execute(
		os.Stdout,
		map[string]any{"Slice": []string{"Joanna", "Angeline"}},
	)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// 'Joanna', 'Angeline'
}

func ExampleAddTemplateFunctions_indent() {
	tplString := "{{ indent 2 .Details }}"
	tpl := template.New("")
	tpl = govy.AddTemplateFunctions(tpl)
	tpl = template.Must(tpl.Parse(tplString))

	err := tpl.Execute(
		os.Stdout,
		map[string]any{"Details": "foo\nbar"},
	)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	//   foo
	//   bar
}

// [govy.Rule] error might be static, i.e. a single [govy.Rule] always returns
// the same exact error message, but they don't have to.
// For instance, consider a rule which parses a URL using [net/url] package.
//
// This makes it very hard to infer error message for [govy.RulePlan], if not
// impossible, since the exact error might only be known during runtime.
//
// To solve this problem, you can use [govy.Rule.WithDescription] to provide a
// verbose and informative rule description.
// It will be only included in the [govy.RulePlan] and otherwise not displayed in
// the default [govy.RuleError.Error].
// However, it is available in the structured [govy.RuleError].
func ExampleRule_WithDescription() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
				WithDetails("Teacher can be either Tom or Jerry :)").
				WithMessage("unsupported name")),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jake':
	//     - unsupported name; Teacher can be either Tom or Jerry :)
}

// The builtin rules, and most likely your custom rules as well, all operate on non-pointer values.
// This means you cannot use them on pointers to the same type.
//
// If for whatever reason you don't want to use [govy.ForPointer] constructor,
// you can use [govy.RuleToPointer] constructor and convert any [govy.Rule] to work on pointers.
//
// Note: [govy.RuleToPointer] will skip validation for nil pointers.
// If you want to enforce the value to be non-nil, you can use [rules.Required].
// This behavior is consistent with [govy.ForPointer] constructor, which will skip the validation
// unless you add [govy.PropertyRules.Required] to enforce the value to be a non-nil pointer.
func ExampleRuleToPointer() {
	type Pointer struct {
		Pointer *string `json:"pointer"`
	}
	validator := govy.New(
		govy.For(func(p Pointer) *string { return p.Pointer }).
			WithName("pointer").
			Rules(govy.RuleToPointer(rules.EQ("foo"))),
	)

	pointer := Pointer{Pointer: ptr("bar")}

	err := validator.Validate(pointer)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'pointer' with value 'bar':
	//     - should be equal to 'foo'
}

// Sometimes it's useful to aggregate multiple [govy.Rule] into a single, composite rule.
// To do that we'll use [govy.RuleSet] and [govy.NewRuleSet] constructor.
// RuleSet is a simple container for multiple [govy.Rule].
// During validation it is unpacked and each [govy.RuleError] is reported separately.
//
// Note that govy uses similar syntax to wrapped errors in Go;
// a ':' delimiter is used to chain error codes together.
func ExampleRuleSet() {
	teacherNameRule := govy.NewRuleSet(
		rules.StringLength(1, 5),
		rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
			WithDetails("Teacher can be either Tom or Jerry :)"),
	).
		WithErrorCode("teacher_name")

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(teacherNameRule),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jonathan",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		propertyErrors := err.(*govy.ValidatorError).Errors
		ruleErrors := propertyErrors[0].Errors
		fmt.Printf("Error codes: %s, %s\n\n", ruleErrors[0].Code, ruleErrors[1].Code)
		fmt.Println(err)
	}

	// Output:
	// Error codes: teacher_name:string_length, teacher_name:string_match_regexp
	//
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jonathan':
	//     - length must be between 1 and 5
	//     - string must match regular expression: '^(Tom|Jerry)$'; Teacher can be either Tom or Jerry :)
}

// Similar to [govy.RuleToPointer], you can use [govy.RuleSetToPointer] to convert
// [govy.RuleSet] to work with pointers.
//
// See [ExampleRuleToPointer] for more details.
func ExampleRuleSetToPointer() {
	type Pointer struct {
		Pointer *string `json:"pointer"`
	}
	ruleSet := govy.NewRuleSet(
		rules.StringStartsWith("f"),
		rules.StringEndsWith("o"),
	)
	validator := govy.New(
		govy.For(func(p Pointer) *string { return p.Pointer }).
			WithName("pointer").
			Rules(govy.RuleSetToPointer(ruleSet)),
	)

	pointer := Pointer{Pointer: ptr("bar")}

	err := validator.Validate(pointer)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'pointer' with value 'bar':
	//     - string must start with 'f' prefix
	//     - string must end with 'o' suffix
}

// If you wish to control how rules aggregated by [govy.RuleSet] evaluate
// you can use [govy.RuleSet.Cascade] to set a [govy.CascadeMode].
//
// Similar to how the cascade mode works when evaluating [govy.PropertyRules],
// the [govy.CascadeModeStop] will stop validation after the first encountered error.
//
// In the example below we can see that although both rules should fail,
// only the first one (order of definitions matters here!) returns an error.
func ExampleRuleSet_Cascade() {
	teacherNameRule := govy.NewRuleSet(
		rules.StringLength(1, 5),
		rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")),
	).
		Cascade(govy.CascadeModeStop)

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(teacherNameRule),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jonathan",
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jonathan':
	//     - length must be between 1 and 5
}

// To inspect if an error contains a given [govy.ErrorCode], use [govy.HasErrorCode] function.
// This function will also return true if the expected [govy.ErrorCode]
// is part of a chain of wrapped error codes.
// In this example we're dealing with two error code chains:
//   - 'teacher_name:string_length'
//   - 'teacher_name:string_match_regexp'
func ExampleHasErrorCode() {
	teacherNameRule := govy.NewRuleSet(
		rules.StringLength(1, 5),
		rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")),
	).
		WithErrorCode("teacher_name")

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(teacherNameRule),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jonathan",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		for _, code := range []govy.ErrorCode{
			"teacher_name",
			"string_length",
			"string_match_regexp",
		} {
			if govy.HasErrorCode(err, code) {
				fmt.Println("Has error code:", code)
			}
		}
	}

	// Output:
	// Has error code: teacher_name
	// Has error code: string_length
	// Has error code: string_match_regexp
}

// Sometimes you need top level context,
// but you want to scope the error to a specific, nested property.
// One of the ways to do that is to use [govy.NewPropertyError]
// and return [govy.PropertyError] from your validation rule.
// Note that you can still use [govy.ErrorCode] and pass [govy.RuleError] to the constructor.
// You can pass any number of [govy.RuleError].
func ExampleNewPropertyError() {
	v := govy.New(
		govy.For(govy.GetSelf[Teacher]()).
			Rules(govy.NewRule(func(t Teacher) error {
				if t.Name == "Jake" {
					return govy.NewPropertyError(
						"name",
						t.Name,
						govy.NewRuleError("name cannot be Jake", "error_code_jake"),
						govy.NewRuleError("you can pass me too!"))
				}
				return nil
			})),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		propertyErrors := err.(*govy.ValidatorError).Errors
		ruleErrors := propertyErrors[0].Errors
		fmt.Printf("Error code: %s\n\n", ruleErrors[0].Code)
		fmt.Println(err)
	}

	// Output:
	// Error code: error_code_jake
	//
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jake':
	//     - name cannot be Jake
	//     - you can pass me too!
}

// So far we've defined validation rules for simple, top-level properties.
// What If we want to define validation rules for nested properties?
// We can use [govy.PropertyRules.Include] to include another [govy.Validator] in our [govy.PropertyRules].
//
// Let's extend our [Teacher] struct to include a nested [University] property.
// [University] in of itself is another struct with its own validation rules.
//
// Notice how the nested property path is automatically built for you,
// each segment separated by a dot.
func ExamplePropertyRules_Include() {
	universityValidation := govy.New(
		govy.For(func(u University) string { return u.Address }).
			WithName("address").
			Required(),
	)
	teacherValidation := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.EQ("Tom")),
		govy.For(func(t Teacher) University { return t.University }).
			WithName("university").
			Include(universityValidation),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jerry",
		Age:  51 * year,
		University: University{
			Name:    "Poznan University of Technology",
			Address: "",
		},
	}

	err := teacherValidation.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jerry':
	//     - should be equal to 'Tom'
	//   - 'university.address':
	//     - property is required but was empty
}

// When dealing with slices we often want to both validate the whole slice
// and each of its elements.
// You can use [govy.ForSlice] function to do just that.
// It returns a new struct [govy.PropertyRulesForSlice] which behaves exactly
// the same as [govy.PropertyRules], but extends its API slightly.
//
// To define rules for each element use:
//   - [govy.PropertyRulesForSlice.RulesForEach]
//   - [govy.PropertyRulesForSlice.IncludeForEach]
//
// These work exactly the same way as [govy.PropertyRules.Rules] and [govy.PropertyRules.Include]
// verifying each slice element.
//
// [govy.PropertyRulesForSlice.Rules] is in turn used to define rules for the whole slice.
//
// Note: [govy.PropertyRulesForSlice] does not implement Include function for the whole slice.
//
// In the below example, we're defining that students slice must have at most 2 elements
// and that each element's index must be unique.
// For each element we're also including [Student] [govy.Validator].
// Notice that property path for slices has the following format:
// <slice_name>[<index>].<slice_property_name>
func ExampleForSlice() {
	studentValidator := govy.New(
		govy.For(func(s Student) string { return s.Index }).
			WithName("index").
			Rules(rules.StringLength(9, 9)),
	)
	teacherValidator := govy.New(
		govy.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			Rules(
				rules.SliceMaxLength[[]Student](2),
				rules.SliceUnique(func(v Student) string { return v.Index })).
			IncludeForEach(studentValidator),
	).When(func(t Teacher) bool { return t.Age < 50 })

	teacher := Teacher{
		Name: "John",
		Students: []Student{
			{Index: "918230014"},
			{Index: "9182300123"},
			{Index: "918230014"},
		},
	}

	err := teacherValidator.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'students' with value '[{"index":"918230014"},{"index":"9182300123"},{"index":"918230014"}]':
	//     - length must be less than or equal to 2
	//     - elements are not unique, 1st and 3rd elements collide
	//   - 'students[1].index' with value '9182300123':
	//     - length must be between 9 and 9
}

// When dealing with slices of pointers you may find it problematic to add [govy.Rule]
// with [govy.PropertyRulesForSlice.RulesForEach].
// The builtin rules, and most likely your custom rules as well, all operate on non-pointer values.
// This means you cannot use them on your slice's pointer elements.
//
// To solve this problem you can use [govy.ForPointer] constructor and convert any [govy.Rule]
// to work on pointers.
//
// In the below example we're defining two [govy.Validator] instances:
//   - 'faultyValidator' which will not fail for 'nil' value
//   - 'goodValidator' which will fail for 'nil' value by using [rules.Required] rule
//
// This behavior is consistent with [govy.ForPointer] constructor, which will skip the validation
// unless you add [govy.PropertyRules.Required] to enforce the value to be a non-nil pointer.
func ExampleForSlice_sliceOfPointers() {
	type Pointers struct {
		Pointers []*string `json:"pointers"`
	}
	pointersRules := govy.ForSlice(func(p Pointers) []*string { return p.Pointers }).
		WithName("pointers").
		Rules(rules.SliceMaxLength[[]*string](2)).
		RulesForEach(
			govy.RuleToPointer(rules.StringLength(9, 9)),
		)
	faultyValidator := govy.New(
		pointersRules,
	)
	goodValidator := govy.New(
		pointersRules.RulesForEach(rules.Required[*string]()),
	)

	pointers := Pointers{
		Pointers: []*string{ptr("918230014"), ptr("9182300123"), ptr("918230014"), nil},
	}

	err := faultyValidator.Validate(pointers)
	if err != nil {
		fmt.Println(err)
	}
	err = goodValidator.Validate(pointers)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'pointers' with value '["918230014","9182300123","918230014",null]':
	//     - length must be less than or equal to 2
	//   - 'pointers[1]' with value '9182300123':
	//     - length must be between 9 and 9
	// Validation has failed for the following properties:
	//   - 'pointers' with value '["918230014","9182300123","918230014",null]':
	//     - length must be less than or equal to 2
	//   - 'pointers[1]' with value '9182300123':
	//     - length must be between 9 and 9
	//   - 'pointers[3]':
	//     - property is required but was empty
}

// When dealing with maps there are three forms of iteration:
//   - keys
//   - values
//   - key-value pairs (items)
//
// You can use [govy.ForMap] function to define rules for all the aforementioned iterators.
// It returns a new struct [govy.PropertyRulesForMap] which behaves similar to
// [govy.PropertyRulesForSlice]..
//
// To define rules for keys use:
//   - [govy.PropertyRulesForMap.RulesForKeys]
//   - [govy.PropertyRulesForMap.IncludeForKeys]
//   - [govy.PropertyRulesForMap.RulesForValues]
//   - [govy.PropertyRulesForMap.IncludeForValues]
//   - [govy.PropertyRulesForMap.RulesForItems]
//   - [govy.PropertyRulesForMap.IncludeForItems]
//
// These work exactly the same way as [govy.PropertyRules.Rules] and [govy.PropertyRules.Include]
// verifying each map's key, value or [govy.MapItem].
//
// [govy.PropertyRulesForMap.Rules] is in turn used to define rules for the whole map.
//
// Note: [govy.PropertyRulesForMap] does not implement Include function for the whole map.
//
// In the below example, we're defining that student index to [Teacher] map:
//   - Must have at most 2 elements (map).
//   - Keys must have a length of 9 (keys).
//   - Eve cannot be a teacher for any student (values).
//   - Joan cannot be a teacher for student with index 918230013 (items).
//
// Notice that property path for maps has the following format:
// <map_name>.<key>.<map_property_name>
func ExampleForMap() {
	teacherValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.NEQ("Eve")),
	)
	tutoringValidator := govy.New(
		govy.ForMap(func(t Tutoring) map[string]Teacher { return t.StudentIndexToTeacher }).
			WithName("students").
			Rules(
				rules.MapMaxLength[map[string]Teacher](2),
			).
			RulesForKeys(
				rules.StringLength(9, 9),
			).
			IncludeForValues(teacherValidator).
			RulesForItems(govy.NewRule(func(v govy.MapItem[string, Teacher]) error {
				if v.Key == "918230013" && v.Value.Name == "Joan" {
					return govy.NewRuleError(
						"Joan cannot be a teacher for student with index 918230013",
						"joan_teacher",
					)
				}
				return nil
			})),
	)

	tutoring := Tutoring{
		StudentIndexToTeacher: map[string]Teacher{
			"918230013":  {Name: "Joan"},
			"9182300123": {Name: "Eve"},
			"918230014":  {Name: "Joan"},
		},
	}

	err := tutoringValidator.Validate(tutoring)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'students' with value '{"9182300123":{"name":"Eve","age":0,"students":null,"university":{"name":"","address":""}},"91823001...':
	//     - length must be less than or equal to 2
	//   - 'students.9182300123' with key '9182300123':
	//     - length must be between 9 and 9
	//   - 'students.9182300123.name' with value 'Eve':
	//     - should be not equal to 'Eve'
	//   - 'students.918230013' with value '{"name":"Joan","age":0,"students":null,"university":{"name":"","address":""}}':
	//     - Joan cannot be a teacher for student with index 918230013
}

// To only run property validation on condition, use [govy.PropertyRules.When].
// Predicates set through [govy.PropertyRules.When] are evaluated in the order they are provided.
// If any predicate is not met, validation rules are not evaluated for the whole [govy.PropertyRules].
//
// It's recommended to define [govy.PropertyRules.When] before [govy.PropertyRules.Rules] declaration.
func ExamplePropertyRules_When() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			When(func(t Teacher) bool { return t.Name == "Jerry" }).
			Rules(rules.NEQ("Jerry")),
	).WithName("Teacher")

	for _, name := range []string{"Tom", "Jerry", "Mickey"} {
		teacher := Teacher{Name: name}
		err := v.Validate(teacher)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jerry':
	//     - should be not equal to 'Jerry'
}

// To customize how [govy.Rule] are evaluated use [govy.PropertyRules.Cascade].
// Use [govy.CascadeModeStop] to stop validation after the first error.
// If you wish to revert to the default behavior, use [govy.CascadeModeContinue].
//
// Note: the cascade mode change only applies to the given [govy.PropertyRules] instance
// and not the parent [govy.Validator] or neighboring [govy.PropertyRules].
// It does however override the [govy.CascadeMode] set for [govy.Validator].
func ExamplePropertyRules_Cascade() {
	alwaysFailingRule := govy.NewRule(func(string) error {
		return fmt.Errorf("always fails")
	})

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Cascade(govy.CascadeModeStop).
			Rules(rules.NEQ("Jerry")).
			Rules(alwaysFailingRule),
	).WithName("Teacher")

	for _, name := range []string{"Tom", "Jerry"} {
		teacher := Teacher{Name: name}
		err := v.Validate(teacher)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Tom':
	//     - always fails
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jerry':
	//     - should be not equal to 'Jerry'
}

// If combining [govy.New] with [govy.ForSlice] is not verbose enough for you,
// you can use [govy.Validator.ValidateSlice] function.
// It will validate each element according to the rules defined by [govy.Validator].
// It returns [govy.ValidatorErrors].
//
// Note: If you need to perform additional validation on the whole slice,
// you should rather use [govy.New] with [govy.ForSlice] and [govy.GetSelf].
// [govy.Validator.ValidateSlice] is designed to be used for processing independent values.
//
// Note: Since each element is validated in isolation,
// the property names will not start with the slice index,
// they will instead start at the element's root.
func ExampleValidator_ValidateSlice() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })),
	).WithName("Teacher")

	err := v.ValidateSlice([]Teacher{
		{Name: "John"},
		{Name: "Jake"},
	})
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher at index 0 has failed for the following properties:
	//   - 'name' with value 'John':
	//     - always fails
	// Validation for Teacher at index 1 has failed for the following properties:
	//   - 'name' with value 'Jake':
	//     - always fails
}

// Unlike [govy.PropertyRules.Cascade] which works on [govy.PropertyRules] level,
// [govy.Validator.Cascade] propagates to all the properties of [govy.Validator] and
// furthermore, will stop evaluating the next property if any preceding property fails.
//
// If [govy.PropertyRules.Cascade] is set, the setting will take precedence over
// [govy.Validator] cascade mode.
//
// See [ExamplePropertyRules_Cascade] for more details on [govy.PropertyRules.Cascade].
func ExampleValidator_Cascade() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Cascade(govy.CascadeModeContinue).
			Rules(rules.NEQ("Jerry")).
			Rules(rules.EQ("Tom")),
		govy.For(func(t Teacher) time.Duration { return t.Age }).
			WithName("age").
			Rules(
				rules.GT(18*year),
				govy.NewRule(func(time.Duration) error {
					return fmt.Errorf("always fails")
				}),
			),
	).
		Cascade(govy.CascadeModeStop)

	for _, name := range []string{"Tom", "Jerry"} {
		teacher := Teacher{
			Name: name,
			Age:  17 * year,
		}
		err := v.WithName(name).Validate(teacher)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Output:
	// Validation for Tom has failed for the following properties:
	//   - 'age' with value '148920h0m0s':
	//     - should be greater than '157680h0m0s'
	// Validation for Jerry has failed for the following properties:
	//   - 'name' with value 'Jerry':
	//     - should be not equal to 'Jerry'
	//     - should be equal to 'Tom'
}

// [govy.Validator.ValidateSlice] outputs [govy.ValidatorErrors] which is a slice of [govy.ValidatorError].
// Each [govy.ValidatorError] has an additional property set: SliceIndex, which is a 0-based slice element index.
func ExampleValidatorErrors() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(govy.NewRule(func(name string) error {
				if name == "John" || name == "Jake" {
					return fmt.Errorf("fails for John and Jake")
				}
				return nil
			})),
	).WithName("Teacher")

	err := v.ValidateSlice([]Teacher{
		{Name: "John"},
		{Name: "George"},
		{Name: "Jake"},
	})
	if err != nil {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err = enc.Encode(err); err != nil {
			fmt.Printf("error encoding: %v\n", err)
		}
	}

	// Output:
	// [
	//   {
	//     "errors": [
	//       {
	//         "propertyName": "name",
	//         "propertyValue": "John",
	//         "errors": [
	//           {
	//             "error": "fails for John and Jake"
	//           }
	//         ]
	//       }
	//     ],
	//     "name": "Teacher",
	//     "sliceIndex": 0
	//   },
	//   {
	//     "errors": [
	//       {
	//         "propertyName": "name",
	//         "propertyValue": "Jake",
	//         "errors": [
	//           {
	//             "error": "fails for John and Jake"
	//           }
	//         ]
	//       }
	//     ],
	//     "name": "Teacher",
	//     "sliceIndex": 2
	//   }
	// ]
}

// Bringing it all (mostly) together, let's create a fully fledged [govy.Validator] for [Teacher].
func ExampleValidator() {
	universityValidation := govy.New(
		govy.For(func(u University) string { return u.Address }).
			WithName("address").
			Required(),
	)
	studentValidator := govy.New(
		govy.For(func(s Student) string { return s.Index }).
			WithName("index").
			Rules(rules.StringLength(9, 9)),
	)
	teacherValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Required().
			Rules(
				rules.StringNotEmpty(),
				rules.OneOf("Jake", "George")),
		govy.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			Rules(
				rules.SliceMaxLength[[]Student](2),
				rules.SliceUnique(func(v Student) string { return v.Index })).
			IncludeForEach(studentValidator),
		govy.For(func(t Teacher) University { return t.University }).
			WithName("university").
			Include(universityValidation),
	).When(func(t Teacher) bool { return t.Age < 50 })

	teacher := Teacher{
		Name: "John",
		Students: []Student{
			{Index: "918230014"},
			{Index: "9182300123"},
			{Index: "918230014"},
		},
		University: University{
			Name:    "Poznan University of Technology",
			Address: "",
		},
	}

	err := teacherValidator.WithName("John").Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for John has failed for the following properties:
	//   - 'name' with value 'John':
	//     - must be one of: Jake, George
	//   - 'students' with value '[{"index":"918230014"},{"index":"9182300123"},{"index":"918230014"}]':
	//     - length must be less than or equal to 2
	//     - elements are not unique, 1st and 3rd elements collide
	//   - 'students[1].index' with value '9182300123':
	//     - length must be between 9 and 9
	//   - 'university.address':
	//     - property is required but was empty
}

// What follows below is a collection of more complex examples and useful patterns.

// When dealing with properties that should only be validated if a certain other
// property has specific value, it's recommended to use [govy.PropertyRules.When] and [govy.PropertyRules.Include]
// to separate validation paths into non-overlapping branches.
//
// Notice how in the below example [File.Format] is the common,
// shared property between [CSV] and [JSON] files.
// We define separate [govy.Validator] for [CSV] and [JSON] and use [govy.PropertyRules.When] to only validate
// their included [govy.Validator] if the correct [File.Format] is provided.
func ExampleValidator_branchingPattern() {
	type (
		CSV struct {
			Separator string `json:"separator"`
		}
		JSON struct {
			Indent string `json:"indent"`
		}
		File struct {
			Format string `json:"format"`
			CSV    *CSV   `json:"csv,omitempty"`
			JSON   *JSON  `json:"json,omitempty"`
		}
	)

	csvValidation := govy.New(
		govy.For(func(c CSV) string { return c.Separator }).
			WithName("separator").
			Required().
			Rules(rules.OneOf(",", ";")),
	)

	jsonValidation := govy.New(
		govy.For(func(j JSON) string { return j.Indent }).
			WithName("indent").
			Required().
			Rules(rules.StringMatchRegexp(regexp.MustCompile(`^\s*$`))),
	)

	fileValidation := govy.New(
		govy.ForPointer(func(f File) *CSV { return f.CSV }).
			When(func(f File) bool { return f.Format == "csv" }).
			Include(csvValidation),
		govy.ForPointer(func(f File) *JSON { return f.JSON }).
			When(func(f File) bool { return f.Format == "json" }).
			Include(jsonValidation),
		govy.For(func(f File) string { return f.Format }).
			WithName("format").
			Required().
			Rules(rules.OneOf("csv", "json")),
	).WithName("File")

	file := File{
		Format: "json",
		CSV:    nil,
		JSON: &JSON{
			Indent: "invalid",
		},
	}

	err := fileValidation.Validate(file)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for File has failed for the following properties:
	//   - 'indent' with value 'invalid':
	//     - string must match regular expression: '^\s*$'
}

// When documenting an API it's often a struggle to keep consistency
// between the code and documentation we write for it.
// What If your code could be self-descriptive?
// Specifically, what If we could generate documentation out of our validation rules?
// We can achieve that by using [govy.Plan] function!
//
// There are multiple ways to improve the generated documentation:
//   - Use [govy.PropertyRules.WithExamples] to provide a list of example values for the property.
//   - Use [govy.Rule.WithDescription] to provide a plan-only description for your rule.
//     For builtin rules, the description is already provided.
//   - Use [govy.WhenDescription] to provide a plan-only description for your when conditions.
func ExamplePlan() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			WithExamples("Jake", "John").
			When(
				func(t Teacher) bool { return t.Name == "Jerry" },
				govy.WhenDescription("name is Jerry"),
			).
			Rules(
				rules.NEQ("Jerry").
					WithDetails("Jerry is just a name!"),
				govy.NewRule(func(v string) error {
					return fmt.Errorf("some custom error")
				}).
					WithDescription("this is a custom error!"),
			),
	).WithName("Teacher")

	properties, err := govy.Plan(v)
	if err != nil {
		panic(err)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	_ = enc.Encode(properties)

	// Output:
	// {
	//   "name": "Teacher",
	//   "properties": [
	//     {
	//       "path": "$.name",
	//       "typeInfo": {
	//         "name": "string",
	//         "kind": "string"
	//       },
	//       "examples": [
	//         "Jake",
	//         "John"
	//       ],
	//       "rules": [
	//         {
	//           "description": "should be not equal to 'Jerry'",
	//           "details": "Jerry is just a name!",
	//           "errorCode": "not_equal_to",
	//           "conditions": [
	//             "name is Jerry"
	//           ]
	//         },
	//         {
	//           "description": "this is a custom error!",
	//           "conditions": [
	//             "name is Jerry"
	//           ]
	//         }
	//       ]
	//     }
	//   ]
	// }
}

// You can enforce certain rules upon [govy.Plan].
// For instance, If you'd want to make sure every [govy.Predicate]
// has a description attached to it, provide [govy.Plan] with [govy.PlanRequirePredicateDescription] option.
//
// If you want to follow our best recommendations, use [govy.PlanStrictMode].
func ExamplePlan_validation() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			WithExamples("Jake", "John").
			When(func(t Teacher) bool { return t.Name == "Jerry" }).
			Rules(
				rules.NEQ("Jerry").
					WithDetails("Jerry is just a name!"),
				govy.NewRule(func(v string) error {
					return fmt.Errorf("some custom error")
				}).
					WithDescription("this is a custom error!"),
			),
	).
		When(func(t Teacher) bool { return t.Age > 18 }).
		WithName("Teacher")

	_, err := govy.Plan(v, govy.PlanStrictMode())
	fmt.Println(err)

	// Output:
	// predicates without description found at: validator level, $.name
}

// In the interactive tutorial for govy, we've been using
// [govy.PropertyRules.WithName] to provide the name for our properties.
//
// Ideally, we'd want to make sure the names govy assigns to each property
// match the name of the struct fields' names that the user defines the validation for.
// Go uses struct tags to achieve that,
// and libraries like [encoding/json] use these tags to encode/decode structs.
// Unfortunately, there's no easy way to tell what exact property we're returning from [govy.PropertyGetter].
//
// To solve this problem, govy provides a way to infer the name of the property (with a catch).
// The catch being that the name inference mechanism needs to parse the whole modules' AST.
// This can be a performance hit, especially for large projects if not done properly.
//
// By default govy **WILL NOT** attempt to infer **ANY** property names.
//
// So, how do we do that properly?
// Both [govy.Validator] and [govy.PropertyRules] (including variants) have a dedicated function -- TODO
//
// It depends on the [govy.InferNameMode] used:
//   - [govy.InferNameModeDisable], name inference is disabled (default), nothing to do here
//   - [govy.InferNameModeRuntime], the name is inferred during runtime, whenever [govy.For] is called.
//     This is the most flexible option, but also the slowest, although the slowdown
//     is incurred only once, whenever [govy.PropertyRules.Validate] is first called.
//     If you make sure that [govy.PropertyRules] is created only once and don't mind
//     the one-time performance hit, this should be enough for you.
//   - [govy.InferNameModeGenerate], the name is inferred during separate code generation phase.
//     This mode requires you to run 'cmd/govy nameinfer' BEFORE you run your code.
//     It will generate a file with inferred names for your structs which automatically
//     registers these names using [govy.SetInferredName].
//
// Since this tutorial is run as a test,
// we need to explicitly instruct govy to infer names from test files.
// By default test files are not parsed to improve performance.
// In order to do that, we use [govyconfig.SetInferNameIncludeTestFiles].
func ExampleInferNameMode() {
	govyconfig.SetInferNameIncludeTestFiles(true)
	defer govyconfig.SetInferNameIncludeTestFiles(false)

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Jerry")),
	).
		InferName(govy.InferNameModeRuntime).
		WithName("Teacher")

	teacher := Teacher{Name: "Tom"}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Tom':
	//     - should be equal to 'Jerry'
}

// In the previous example we've seen [govy.InferNameModeRuntime] in action.
// An alternative for the aforementioned mode which offers better runtime performance
// is [govy.InferNameModeGenerate].
//
// It comes at a cost of having to run the code generation utility before running your code.
// The utility generates code which uses [govyconfig.SetInferredName].
// We'll use this very function in this example to simulate the code generation step.
// The first validator, 'v1', is created with [govy.InferNameModeDisable],
// the second validator, 'v2' is created with [govy.InferNameModeGenerate].
// As you can see in the output, only the second validator, 'v2' has the inferred name.
func ExampleInferNameModeGenerate() {
	govyconfig.SetInferNameIncludeTestFiles(true)
	defer govyconfig.SetInferNameIncludeTestFiles(false)

	v1 := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Jerry")),
	).
		InferName(govy.InferNameModeGenerate).
		WithName("Teacher")

	govyconfig.SetInferredName(govyconfig.InferredName{
		Name: "name",
		File: "pkg/govyconfig/example_test.go",
		Line: 96,
	})

	v2 := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Thomas")),
	).WithName("NotTeacher")

	teacher := Teacher{Name: "Tom"}
	if err := v1.Validate(teacher); err != nil {
		fmt.Println(err)
	}
	if err := v2.Validate(teacher); err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed:
	//   - should be equal to 'Jerry'
	// Validation for NotTeacher has failed for the following properties:
	//   - 'name' with value 'Tom':
	//     - should be equal to 'Thomas'
}

// Knowing when to call [govy.Validator.InferName] is important.
// Beware that once a [govy.Validator.Validate] is called, it will cache the inferred name.
//
// To demonstrate this we'll set the [govy.InferNameModeDisable] and you will observe
// that the name is still inferred, although to be precise, it's not inferred anymore,
// it was inferred the first time [govy.Validator.Validate] was called and now it's cached.
func ExampleValidator_InferName_changeModeInRuntime() {
	govyconfig.SetInferNameIncludeTestFiles(true)
	defer govyconfig.SetInferNameIncludeTestFiles(false)

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Jerry")),
	).
		InferName(govy.InferNameModeDisable).
		WithName("Teacher")

	teacher := Teacher{Name: "Tom"}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("---\nAfter setting Runtime infer mode.\n---")
	err = v.InferName(govy.InferNameModeRuntime).Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed:
	//   - should be equal to 'Jerry'
	// ---
	// After setting Runtime infer mode.
	// ---
	// Validation for Teacher has failed:
	//   - should be equal to 'Jerry'
}
