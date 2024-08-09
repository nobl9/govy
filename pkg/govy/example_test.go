// nolint: lll
package govy_test

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/nobl9/govy/pkg/govy"
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
			Rules(govy.NewSingleRule(func(name string) error { return fmt.Errorf("always fails") })),
	)

	err := v.Validate(Teacher{})
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - always fails
}

// To associate [govy.Validator] with an entity name use [govy.Validator.WithName] function.
// When any of the rules fails, the error will contain the entity name you've provided.
func ExampleValidator_WithName() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewSingleRule(func(name string) error { return fmt.Errorf("always fails") })),
	).WithName("Teacher")

	err := v.Validate(Teacher{})
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - always fails
}

// You can also add [govy.Validator] name during runtime,
// by calling [govy.ValidatorError.WithName] function on the returned error.
//
// NOTE: We left the previous "Teacher" name assignment, to demonstrate that
// the [govy.ValidatorError.WithName] function call will shadow it.
//
// NOTE: This would also work:
//
//	err := v.WithName("Jake").Validate(Teacher{})
//
// Validation package, aside from errors handling,
// tries to follow immutability principle. Calling any function on [govy.Validator]
// will not change its previous declaration (unless you assign it back to 'v').
func ExampleValidatorError_WithName() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewSingleRule(func(name string) error { return fmt.Errorf("always fails") })),
	).WithName("Teacher")

	err := v.Validate(Teacher{})
	if err != nil {
		fmt.Println(err.WithName("Jake"))
	}

	// Output:
	// Validation for Jake has failed for the following properties:
	//   - always fails
}

// [govy.Validator] rules can be evaluated on condition, to specify the predicate use [govy.Validator.When] function.
//
// In this example, validation for [Teacher] instance will only be evaluated
// if the [Teacher.Age] property is less than 50 years.
func ExampleValidator_When() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewSingleRule(func(name string) error { return fmt.Errorf("always fails") })),
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
		fmt.Println(err.WithName("Tom"))
	}
	err = v.Validate(teacherJerry)
	if err != nil {
		fmt.Println(err.WithName("Jerry"))
	}

	// Output:
	// Validation for Jerry has failed for the following properties:
	//   - always fails
}

// So far we've been using a very simple [govy.PropertyRules] instance:
//
//	validation.For(func(t Teacher) string { return t.Name }).
//		Rules(validation.NewSingleRule(func(name string) error { return fmt.Errorf("always fails") }))
//
// The error message returned by this property rule does not tell us
// which property is failing.
// Let's change that by adding property name using [govy.PropertyRules.WithName].
//
// We can also change the [govy.Rule] to be something more real.
// Validation package comes with a number of predefined [govy.Rule], we'll use
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
// NOTE: [govy.PropertyRules.Required] is introducing a short circuit.
// If the assertion fails, validation will stop and return [govy.govy.ErrorCodeRequired].
// None of the rules you've defined would be evaluated.
//
// NOTE: Placement of [govy.PropertyRules.Required] does not matter,
// it's not evaluated in a sequential loop, unlike standard [govy.Rule].
// However, we recommend you always place it below [govy.PropertyRules.WithName]
// to make your rules more readable.
func ExamplePropertyRules_Required() {
	alwaysFailingRule := govy.NewSingleRule(func(string) error {
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
// NOTE: [govy.PropertyRules.OmitEmpty] will have no effect on pointers handled
// by [govy.ForPointer], as they already behave in the same way.
func ExamplePropertyRules_OmitEmpty() {
	alwaysFailingRule := govy.NewSingleRule(func(string) error {
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

// If you want to access the value of the entity you're writing the [govy.Validator] for,
// you can use [govy.GetSelf] function which is a convenience [govy.PropertyGetter] that returns self.
// Note that we don't call [govy.PropertyRules.WithName] here,
// as we're comparing two properties in our top level, [Teacher] scope.
//
// You can provide your own rules using [govy.NewSingleRule] constructor.
// It returns new [govy.SingleRule] instance which wraps your validation function.
func ExampleGetSelf() {
	customRule := govy.NewSingleRule(func(v Teacher) error {
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
	// Validation for Teacher has failed for the following properties:
	//   - now I have access to the whole teacher
}

// You can use [govy.SingleRule.WithDetails] to add additional details to the error message.
// This allows you to extend existing rules by adding your use case context.
// Let's give a regex validation some more clarity.
func ExampleSingleRule_WithDetails() {
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

// When testing, it can be tedious to always rely on error messages as these can change over time.
// Enter [govy.ErrorCode], which is a simple string type alias used to ease testing,
// but also potentially allow third parties to integrate with your validation results.
// Use [govy.SingleRule.WithErrorCode] to associate [govy.ErrorCode] with a [govy.SingleRule].
// Notice that our modified version of [rules.StringMatchRegexp] will now return a new [govy.ErrorCode].
// Predefined rules have [govy.ErrorCode] already associated with them.
// To view the list of predefined [govy.ErrorCode] checkout error_codes.go file.
func ExampleSingleRule_WithErrorCode() {
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
		propertyErrors := err.Errors
		ruleErrors := propertyErrors[0].Errors
		fmt.Println(ruleErrors[0].Code)
	}

	// Output:
	// custom_code
}

// Sometimes it's useful to build a [govy.Rule] using other rules.
// To do that we'll use [govy.RuleSet] and [govy.NewRuleSet] constructor.
// RuleSet is a simple container for multiple [govy.Rule].
// It is later on unpacked and each [govy.RuleError] is reported separately.
// When [govy.RuleSet.WithErrorCode] or [govy.RuleSet.WithDetails] are used,
// error code and details are added to each [govy.RuleError].
// Note that validation package uses similar syntax to wrapped errors in Go;
// a ':' delimiter is used to chain error codes together.
func ExampleRuleSet() {
	teacherNameRule := govy.NewRuleSet(
		rules.StringLength(1, 5),
		rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
			WithDetails("Teacher can be either Tom or Jerry :)"),
	).
		WithErrorCode("teacher_name").
		WithDetails("I will add that to both rules!")

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
		propertyErrors := err.Errors
		ruleErrors := propertyErrors[0].Errors
		fmt.Printf("Error codes: %s, %s\n\n", ruleErrors[0].Code, ruleErrors[1].Code)
		fmt.Println(err)
	}

	// nolint: lll
	// Output:
	// Error codes: teacher_name:string_length, teacher_name:string_match_regexp
	//
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jonathan':
	//     - length must be between 1 and 5; I will add that to both rules!
	//     - string does not match regular expression: '^(Tom|Jerry)$'; Teacher can be either Tom or Jerry :); I will add that to both rules!
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
			Rules(govy.NewSingleRule(func(t Teacher) error {
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
		propertyErrors := err.Errors
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
// - [govy.PropertyRulesForSlice.RulesForEach]
// - [govy.PropertyRulesForSlice.IncludeForEach]
// These work exactly the same way as [govy.PropertyRules.Rules] and [govy.PropertyRules.Include]
// verifying each slice element.
//
// [govy.PropertyRulesForSlice.Rules] is in turn used to define rules for the whole slice.
//
// NOTE: [govy.PropertyRulesForSlice] does not implement Include function for the whole slice.
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
	//     - elements are not unique, index 0 collides with index 2
	//   - 'students[1].index' with value '9182300123':
	//     - length must be between 9 and 9
}

// When dealing with maps there are three forms of iteration:
// - keys
// - values
// - key-value pairs (items)
//
// You can use [govy.ForMap] function to define rules for all the aforementioned iterators.
// It returns a new struct [govy.PropertyRulesForMap] which behaves similar to
// [govy.PropertyRulesForSlice]..
//
// To define rules for keys use:
// - [govy.PropertyRulesForMap.RulesForKeys]
// - [govy.PropertyRulesForMap.IncludeForKeys]
// - [govy.PropertyRulesForMap.RulesForValues]
// - [govy.PropertyRulesForMap.IncludeForValues]
// - [govy.PropertyRulesForMap.RulesForItems]
// - [govy.PropertyRulesForMap.IncludeForItems]
// These work exactly the same way as [govy.PropertyRules.Rules] and [govy.PropertyRules.Include]
// verifying each map's key, value or [govy.MapItem].
//
// [govy.PropertyRulesForMap.Rules] is in turn used to define rules for the whole map.
//
// NOTE: [govy.PropertyRulesForMap] does not implement Include function for the whole map.
//
// In the below example, we're defining that student index to [Teacher] map:
// - Must have at most 2 elements (map).
// - Keys must have a length of 9 (keys).
// - Eve cannot be a teacher for any student (values).
// - Joan cannot be a teacher for student with index 918230013 (items).
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
			RulesForItems(govy.NewSingleRule(func(v govy.MapItem[string, Teacher]) error {
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
func ExamplePropertyRules_Cascade() {
	alwaysFailingRule := govy.NewSingleRule(func(string) error {
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
	//     - must be one of [Jake, George]
	//   - 'students' with value '[{"index":"918230014"},{"index":"9182300123"},{"index":"918230014"}]':
	//     - length must be less than or equal to 2
	//     - elements are not unique, index 0 collides with index 2
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

func ExampleValidator_branchingPattern() {
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
//   - Use [govy.SingleRule.WithDescription] to provide a plan-only description for your rule.
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
			),
	)

	properties := govy.Plan(v)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	_ = enc.Encode(properties)

	//[
	//  {
	//    "path": "$.name",
	//    "type": "string",
	//    "examples": [
	//      "Jake",
	//      "John"
	//    ],
	//    "rules": [
	//      {
	//        "description": "should be not equal to 'Jerry'",
	//        "details": "Jerry is just a name!",
	//        "errorCode": "not_equal_to",
	//        "conditions": [
	//          "name is Jerry"
	//        ]
	//      }
	//    ]
	//  }
	//]
}
