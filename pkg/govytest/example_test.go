package govytest_test

import (
	"fmt"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/govytest"
	"github.com/nobl9/govy/pkg/rules"
)

type Teacher struct {
	Name       string     `json:"name"`
	University University `json:"university"`
}

type University struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// You can use [govytest.AssertNoError] to ensure no error was produced by [govy.Validator.Validate].
// If an error was produced, it will be printed to the stdout in JSON format.
//
// To demonstrate the erroneous output of [govytest.AssertNoError] we'll fail the assertion.
func ExampleAssertNoError() {
	teacherValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Required().
			Rules(
				rules.StringNotEmpty(),
				rules.OneOf("Jake", "George")),
		govy.For(func(t Teacher) University { return t.University }).
			WithName("university").
			Include(govy.New(
				govy.For(func(u University) string { return u.Address }).
					WithName("address").
					Required(),
			)),
	)

	teacher := Teacher{
		Name: "John",
		University: University{
			Name:    "Poznan University of Technology",
			Address: "",
		},
	}

	// We'll use a mock testing.T to capture the error produced by the assertion.
	mt := new(mockTestingT)

	err := teacherValidator.WithName("John").Validate(teacher)
	govytest.AssertNoError(mt, err)

	// This will print the error produced by the assertion.
	fmt.Println(mt.recordedError)

	// Output:
	// Received unexpected error:
	// {
	//   "errors": [
	//     {
	//       "propertyName": "name",
	//       "propertyValue": "John",
	//       "errors": [
	//         {
	//           "error": "must be one of [Jake, George]",
	//           "code": "one_of",
	//           "description": "must be one of: Jake, George"
	//         }
	//       ]
	//     },
	//     {
	//       "propertyName": "university.address",
	//       "errors": [
	//         {
	//           "error": "property is required but was empty",
	//           "code": "required"
	//         }
	//       ]
	//     }
	//   ],
	//   "name": "John"
	// }
}

// Verifying that expected errors were produced by [govy.Validator.Validate] can be a tedious task.
// Often times we might only care about [govy.ErrorCode] and not the message or description of the error.
// To help in that process, [govytest.AssertError] can be used to ensure that the expected errors were produced.
// It accepts multiple [govytest.ExpectedRuleError], each being a short and concise
// representation of the error we're expecting to occur.
// For more details on how to use [govytest.ExpectedRuleError], see its code documentation.
//
// To demonstrate the erroneous output of [govytest.AssertError] we'll fail the assertion.
func ExampleAssertError() {
	teacherValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Required().
			Rules(
				rules.StringNotEmpty(),
				rules.OneOf("Jake", "George")),
		govy.For(func(t Teacher) University { return t.University }).
			WithName("university").
			Include(govy.New(
				govy.For(func(u University) string { return u.Address }).
					WithName("address").
					Required(),
			)),
	)

	teacher := Teacher{
		Name: "John",
		University: University{
			Name:    "Poznan University of Technology",
			Address: "",
		},
	}

	// We'll use a mock testing.T to capture the error produced by the assertion.
	mt := new(mockTestingT)

	err := teacherValidator.WithName("John").Validate(teacher)
	govytest.AssertError(mt, err,
		govytest.ExpectedRuleError{
			PropertyName:    "name",
			ContainsMessage: "one of",
		},
		govytest.ExpectedRuleError{
			PropertyName: "university.address",
			Code:         "greater_than",
		},
	)

	// This will print the error produced by the assertion.
	fmt.Println(mt.recordedError)

	// Output:
	// Expected error was not found.
	// EXPECTED:
	// {
	//   "propertyName": "university.address",
	//   "code": "greater_than"
	// }
	// ACTUAL:
	// [
	//   {
	//     "propertyName": "name",
	//     "propertyValue": "John",
	//     "errors": [
	//       {
	//         "error": "must be one of [Jake, George]",
	//         "code": "one_of",
	//         "description": "must be one of: Jake, George"
	//       }
	//     ]
	//   },
	//   {
	//     "propertyName": "university.address",
	//     "errors": [
	//       {
	//         "error": "property is required but was empty",
	//         "code": "required"
	//       }
	//     ]
	//   }
	// ]
}