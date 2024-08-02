// nolint: lll
package validation_test

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/nobl9/go-yaml"

	"github.com/nobl9/nobl9-go/internal/validation"
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

// In order to create a new [Validator] use [New] constructor.
// Let's define simple [PropertyRules] for [Teacher.Name].
// For now, it will be always failing.
func ExampleNew() {
	v := validation.New[Teacher](
		validation.For(func(t Teacher) string { return t.Name }).
			Rules(validation.NewSingleRule(func(name string) error { return fmt.Errorf("always fails") })),
	)

	err := v.Validate(Teacher{})
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - always fails
}

// To associate [Validator] with an entity name use [Validator.WithName] function.
// When any of the rules fails, the error will contain the entity name you've provided.
func ExampleValidator_WithName() {
	v := validation.New[Teacher](
		validation.For(func(t Teacher) string { return t.Name }).
			Rules(validation.NewSingleRule(func(name string) error { return fmt.Errorf("always fails") })),
	).WithName("Teacher")

	err := v.Validate(Teacher{})
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - always fails
}

// You can also add [Validator] name during runtime,
// by calling [ValidatorError.WithName] function on the returned error.
//
// NOTE: We left the previous "Teacher" name assignment, to demonstrate that
// the [ValidatorError.WithName] function call will shadow it.
//
// NOTE: This would also work:
//
//	err := v.WithName("Jake").Validate(Teacher{})
//
// Validation package, aside from errors handling,
// tries to follow immutability principle. Calling any function on [Validator]
// will not change its previous declaration (unless you assign it back to 'v').
func ExampleValidatorError_WithName() {
	v := validation.New[Teacher](
		validation.For(func(t Teacher) string { return t.Name }).
			Rules(validation.NewSingleRule(func(name string) error { return fmt.Errorf("always fails") })),
	).WithName("Teacher")

	err := v.Validate(Teacher{})
	if err != nil {
		fmt.Println(err.WithName("Jake"))
	}

	// Output:
	// Validation for Jake has failed for the following properties:
	//   - always fails
}

// [Validator] rules can be evaluated on condition, to specify the predicate use [Validator.When] function.
//
// In this example, validation for [Teacher] instance will only be evaluated
// if the [Age] property is less than 50 years.
func ExampleValidator_When() {
	v := validation.New[Teacher](
		validation.For(func(t Teacher) string { return t.Name }).
			Rules(validation.NewSingleRule(func(name string) error { return fmt.Errorf("always fails") })),
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

// So far we've been using a very simple [PropertyRules] instance:
//
//	validation.For(func(t Teacher) string { return t.Name }).
//		Rules(validation.NewSingleRule(func(name string) error { return fmt.Errorf("always fails") }))
//
// The error message returned by this property rule does not tell us
// which property is failing.
// Let's change that by adding property name using [PropertyRules.WithName].
//
// We can also change the [Rule] to be something more real.
// Validation package comes with a number of predefined [Rule], we'll use
// [EqualTo] which accepts a single argument, value to compare with.
func ExamplePropertyRules_WithName() {
	v := validation.New[Teacher](
		validation.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(validation.EqualTo("Tom")),
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

// [For] constructor creates new [PropertyRules] instance.
// It's only argument, [PropertyGetter] is used to extract the property value.
// It works fine for direct values, but falls short when working with pointers.
// Often times we use pointers to indicate that a property is optional,
// or we want to discern between nil and zero values.
// In either case we want our validation rules to work on direct values,
// not the pointer, otherwise we'd have to always check if pointer != nil.
//
// [ForPointer] constructor can be used to solve this problem and allow
// us to work with the underlying value in our rules.
// Under the hood it wraps [PropertyGetter] and safely extracts the underlying value.
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
	v := validation.New[Teacher](
		validation.ForPointer(func(t Teacher) *string { return t.MiddleName }).
			WithName("middleName").
			Rules(validation.StringMaxLength(5)),
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

// By default, when [PropertyRules] is constructed using [ForPointer]
// it will skip validation of the property if the pointer is nil.
// To enforce a value is set for pointer use [PropertyRules.Required].
//
// You may ask yourself why not just use [validation.Required] rule instead?
// If we were to do that, we'd be forced to operate on pointer in all of our rules.
// Other than checking if the pointer is nil, there aren't any rules which would
// benefit from working on the pointer instead of the underlying value.
//
// If you want to also make sure the underlying value is filled,
// i.e. it's not a zero value, you can also use [validation.Required] rule
// on top of [PropertyRules.Required].
//
// [PropertyRules.Required] when used with [For] constructor, will ensure
// the property does not contain a zero value.
//
// NOTE: [PropertyRules.Required] is introducing a short circuit.
// If the assertion fails, validation will stop and return [validation.ErrorCodeRequired].
// None of the rules you've defined would be evaluated.
//
// NOTE: Placement of [PropertyRules.Required] does not matter,
// it's not evaluated in a sequential loop, unlike standard [Rule].
// However, we recommend you always place it below [PropertyRules.WithName]
// to make your rules more readable.
func ExamplePropertyRules_Required() {
	alwaysFailingRule := validation.NewSingleRule(func(string) error {
		return fmt.Errorf("always fails")
	})

	v := validation.New[Teacher](
		validation.ForPointer(func(t Teacher) *string { return t.MiddleName }).
			WithName("middleName").
			Required().
			Rules(alwaysFailingRule),
		validation.For(func(t Teacher) string { return t.Name }).
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

// While [ForPointer] will by default omit validation for nil pointers,
// it might be useful to have a similar behavior for optional properties
// which are direct values.
// [PropertyRules.OmitEmpty] will do the trick.
//
// NOTE: [PropertyRules.OmitEmpty] will have no effect on pointers handled
// by [ForPointer], as they already behave in the same way.
func ExamplePropertyRules_OmitEmpty() {
	alwaysFailingRule := validation.NewSingleRule(func(string) error {
		return fmt.Errorf("always fails")
	})

	v := validation.New[Teacher](
		validation.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			OmitEmpty().
			Rules(alwaysFailingRule),
		validation.ForPointer(func(t Teacher) *string { return t.MiddleName }).
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

// If you want to access the value of the entity you're writing the [Validator] for,
// you can use [GetSelf] function which is a convenience [PropertyGetter] that returns self.
// Note that we don't call [PropertyRules.WithName] here,
// as we're comparing two properties in our top level, [Teacher] scope.
//
// You can provide your own rules using [NewSingleRule] constructor.
// It returns new [SingleRule] instance which wraps your validation function.
func ExampleGetSelf() {
	customRule := validation.NewSingleRule(func(v Teacher) error {
		return fmt.Errorf("now I have access to the whole teacher")
	})

	v := validation.New[Teacher](
		validation.For(validation.GetSelf[Teacher]()).
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

// You can use [SingleRule.WithDetails] to add additional details to the error message.
// This allows you to extend existing rules by adding your use case context.
// Let's give a regex validation some more clarity.
func ExampleSingleRule_WithDetails() {
	v := validation.New[Teacher](
		validation.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(validation.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
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
// Enter [ErrorCode], which is a simple string type alias used to ease testing,
// but also potentially allow third parties to integrate with your validation results.
// Use [SingleRule.WithErrorCode] to associate [ErrorCode] with a [SingleRule].
// Notice that our modified version of [StringMatchRegexp] will now return a new [ErrorCode].
// Predefined rules have [ErrorCode] already associated with them.
// To view the list of predefined [ErrorCode] checkout error_codes.go file.
func ExampleSingleRule_WithErrorCode() {
	v := validation.New[Teacher](
		validation.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(validation.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
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

// Sometimes it's useful to build a [Rule] using other rules.
// To do that we'll use [RuleSet] and [NewRuleSet] constructor.
// RuleSet is a simple container for multiple [Rule].
// It is later on unpacked and each [RuleError] is reported separately.
// When [RuleSet.WithErrorCode] or [RuleSet.WithDetails] are used,
// error code and details are added to each [RuleError].
// Note that validation package uses similar syntax to wrapped errors in Go;
// a ':' delimiter is used to chain error codes together.
func ExampleRuleSet() {
	teacherNameRule := validation.NewRuleSet[string](
		validation.StringLength(1, 5),
		validation.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
			WithDetails("Teacher can be either Tom or Jerry :)"),
	).
		WithErrorCode("teacher_name").
		WithDetails("I will add that to both rules!")

	v := validation.New[Teacher](
		validation.For(func(t Teacher) string { return t.Name }).
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

// To inspect if an error contains a given [validation.ErrorCode], use [HasErrorCode] function.
// This function will also return true if the expected [ErrorCode]
// is part of a chain of wrapped error codes.
// In this example we're dealing with two error code chains:
// - 'teacher_name:string_length'
// - 'teacher_name:string_match_regexp'
func ExampleHasErrorCode() {
	teacherNameRule := validation.NewRuleSet[string](
		validation.StringLength(1, 5),
		validation.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")),
	).
		WithErrorCode("teacher_name")

	v := validation.New[Teacher](
		validation.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(teacherNameRule),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jonathan",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		for _, code := range []validation.ErrorCode{
			"teacher_name",
			"string_length",
			"string_match_regexp",
		} {
			if validation.HasErrorCode(err, code) {
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
// One of the ways to do that is to use [NewPropertyError]
// and return [PropertyError] from your validation rule.
// Note that you can still use [ErrorCode] and pass [RuleError] to the constructor.
// You can pass any number of [RuleError].
func ExampleNewPropertyError() {
	v := validation.New[Teacher](
		validation.For(validation.GetSelf[Teacher]()).
			Rules(validation.NewSingleRule(func(t Teacher) error {
				if t.Name == "Jake" {
					return validation.NewPropertyError(
						"name",
						t.Name,
						validation.NewRuleError("name cannot be Jake", "error_code_jake"),
						validation.NewRuleError("you can pass me too!"))
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
// We can use [PropertyRules.Include] to include another [Validator] in our [PropertyRules].
//
// Let's extend our [Teacher] struct to include a nested [University] property.
// [University] in of itself is another struct with its own validation rules.
//
// Notice how the nested property path is automatically built for you,
// each segment separated by a dot.
func ExamplePropertyRules_Include() {
	universityValidation := validation.New[University](
		validation.For(func(u University) string { return u.Address }).
			WithName("address").
			Required(),
	)
	teacherValidation := validation.New[Teacher](
		validation.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(validation.EqualTo("Tom")),
		validation.For(func(t Teacher) University { return t.University }).
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
// You can use [ForSlice] function to do just that.
// It returns a new struct [PropertyRulesForSlice] which behaves exactly
// the same as [PropertyRules], but extends its API slightly.
//
// To define rules for each element use:
// - [PropertyRulesForSlice.RulesForEach]
// - [PropertyRulesForSlice.IncludeForEach]
// These work exactly the same way as [PropertyRules.Rules] and [PropertyRules.Include]
// verifying each slice element.
//
// [PropertyRulesForSlice.Rules] is in turn used to define rules for the whole slice.
//
// NOTE: [PropertyRulesForSlice] does not implement Include function for the whole slice.
//
// In the below example, we're defining that students slice must have at most 2 elements
// and that each element's index must be unique.
// For each element we're also including [Student] [Validator].
// Notice that property path for slices has the following format:
// <slice_name>[<index>].<slice_property_name>
func ExampleForSlice() {
	studentValidator := validation.New[Student](
		validation.For(func(s Student) string { return s.Index }).
			WithName("index").
			Rules(validation.StringLength(9, 9)),
	)
	teacherValidator := validation.New[Teacher](
		validation.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			Rules(
				validation.SliceMaxLength[[]Student](2),
				validation.SliceUnique(func(v Student) string { return v.Index })).
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
// You can use [ForMap] function to define rules for all the aforementioned iterators.
// It returns a new struct [PropertyRulesForMap] which behaves similar to
// [PropertyRulesForSlice]..
//
// To define rules for keys use:
// - [PropertyRulesForMap.RulesForKeys]
// - [PropertyRulesForMap.IncludeForKeys]
// - [PropertyRulesForMap.RulesForValues]
// - [PropertyRulesForMap.IncludeForValues]
// - [PropertyRulesForMap.RulesForItems]
// - [PropertyRulesForMap.IncludeForItems]
// These work exactly the same way as [PropertyRules.Rules] and [PropertyRules.Include]
// verifying each map's key, value or [MapItem].
//
// [PropertyRulesForMap.Rules] is in turn used to define rules for the whole map.
//
// NOTE: [PropertyRulesForMap] does not implement Include function for the whole map.
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
	teacherValidator := validation.New[Teacher](
		validation.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(validation.NotEqualTo("Eve")),
	)
	tutoringValidator := validation.New[Tutoring](
		validation.ForMap(func(t Tutoring) map[string]Teacher { return t.StudentIndexToTeacher }).
			WithName("students").
			Rules(
				validation.MapMaxLength[map[string]Teacher](2),
			).
			RulesForKeys(
				validation.StringLength(9, 9),
			).
			IncludeForValues(teacherValidator).
			RulesForItems(validation.NewSingleRule(func(v validation.MapItem[string, Teacher]) error {
				if v.Key == "918230013" && v.Value.Name == "Joan" {
					return validation.NewRuleError(
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

// To only run property validation on condition, use [PropertyRules.When].
// Predicates set through [PropertyRules.When] are evaluated in the order they are provided.
// If any predicate is not met, validation rules are not evaluated for the whole [PropertyRules].
//
// It's recommended to define [PropertyRules.When] before [PropertyRules.Rules] declaration.
func ExamplePropertyRules_When() {
	v := validation.New[Teacher](
		validation.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			When(func(t Teacher) bool { return t.Name == "Jerry" }).
			Rules(validation.NotEqualTo("Jerry")),
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

// To customize how [Rule] are evaluated use [PropertyRules.Cascade].
// Use [CascadeModeStop] to stop validation after the first error.
// If you wish to revert to the default behavior, use [CascadeModeContinue].
func ExamplePropertyRules_Cascade() {
	alwaysFailingRule := validation.NewSingleRule(func(string) error {
		return fmt.Errorf("always fails")
	})

	v := validation.New[Teacher](
		validation.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Cascade(validation.CascadeModeStop).
			Rules(validation.NotEqualTo("Jerry")).
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

// Bringing it all (mostly) together, let's create a fully fledged [Validator] for [Teacher].
func ExampleValidator() {
	universityValidation := validation.New[University](
		validation.For(func(u University) string { return u.Address }).
			WithName("address").
			Required(),
	)
	studentValidator := validation.New[Student](
		validation.For(func(s Student) string { return s.Index }).
			WithName("index").
			Rules(validation.StringLength(9, 9)),
	)
	teacherValidator := validation.New[Teacher](
		validation.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Required().
			Rules(
				validation.StringNotEmpty(),
				validation.OneOf("Jake", "George")),
		validation.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			Rules(
				validation.SliceMaxLength[[]Student](2),
				validation.SliceUnique(func(v Student) string { return v.Index })).
			IncludeForEach(studentValidator),
		validation.For(func(t Teacher) University { return t.University }).
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
// property has specific value, it's recommended to use [PropertyRules.When] and [PropertyRules.Include]
// to separate validation paths into non-overlapping branches.
//
// Notice how in the below example [File.Format] is the common,
// shared property between [CSV] and [JSON] files.
// We define separate [Validator] for [CSV] and [JSON] and use [PropertyRules.When] to only validate
// their included [Validator] if the correct [File.Format] is provided.
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

	csvValidation := validation.New[CSV](
		validation.For(func(c CSV) string { return c.Separator }).
			WithName("separator").
			Required().
			Rules(validation.OneOf(",", ";")),
	)

	jsonValidation := validation.New[JSON](
		validation.For(func(j JSON) string { return j.Indent }).
			WithName("indent").
			Required().
			Rules(validation.StringMatchRegexp(regexp.MustCompile(`^\s*$`))),
	)

	fileValidation := validation.New[File](
		validation.ForPointer(func(f File) *CSV { return f.CSV }).
			When(func(f File) bool { return f.Format == "csv" }).
			Include(csvValidation),
		validation.ForPointer(func(f File) *JSON { return f.JSON }).
			When(func(f File) bool { return f.Format == "json" }).
			Include(jsonValidation),
		validation.For(func(f File) string { return f.Format }).
			WithName("format").
			Required().
			Rules(validation.OneOf("csv", "json")),
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
// We can achieve that by using [Plan] function!
//
// There are multiple ways to improve the generated documentation:
//   - Use [PropertyRules.WithExamples] to provide a list of example values for the property.
//   - Use [SingleRule.WithDescription] to provide a plan-only description for your rule.
//     For builtin rules, the description is already provided.
//   - Use [WhenDescription] to provide a plan-only description for your when conditions.
func ExamplePlan() {
	v := validation.New[Teacher](
		validation.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			WithExamples("Jake", "John").
			When(
				func(t Teacher) bool { return t.Name == "Jerry" },
				validation.WhenDescription("name is Jerry"),
			).
			Rules(
				validation.NotEqualTo("Jerry").
					WithDetails("Jerry is just a name!"),
			),
	)

	properties := validation.Plan(v)
	_ = yaml.NewEncoder(os.Stdout, yaml.Indent(2)).Encode(properties)

	// Output:
	// - path: $.name
	//   type: string
	//   examples:
	//   - Jake
	//   - John
	//   rules:
	//   - description: should be not equal to 'Jerry'
	//     details: Jerry is just a name!
	//     errorCode: not_equal_to
	//     conditions:
	//     - name is Jerry
}
