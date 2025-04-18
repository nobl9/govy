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

	// We're using a mock testing.T to capture the error produced by the assertion.
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
	//           "error": "must be one of: Jake, George",
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

	// We're using a mock testing.T to capture the error produced by the assertion.
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
	//         "error": "must be one of: Jake, George",
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

// [govytest.AssertError] can handle not only [govy.ValidatorError] but also a slice of these
// wrapped into [govy.ValidatorErrors].
//
// For instance, [govy.ValidatorErrors] is returned from [govy.Validator.ValidateSlice],
// you can also choose to construct this slice type yourself.
// In any case, in order to match [govytest.ExpectedRuleError] the assertion function needs
// to first match it with the right [govy.ValidatorError].
//
// In order to achieve that you need to set the following fields:
//   - [govytest.ExpectedRuleError.ValidatorName] which matches [govy.ValidatorError.Name]
//   - [govytest.ExpectedRuleError.ValidatorIndex] which matches [govy.ValidatorError.SliceIndex]
//
// You can set them both, but you need to provide at least one of them.
//
// Every [govytest.ExpectedRuleError] is aggregated per matched [govy.ValidatorError]
// and the function runs recursively for every [govy.ValidatorError] and epxected errors pair.
func ExampleAssertError_validatorErrors() {
	teacherValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Required().
			Rules(
				rules.StringNotEmpty(),
				rules.OneOf("Eve", "George")),
		govy.For(func(t Teacher) University { return t.University }).
			WithName("university").
			Include(govy.New(
				govy.For(func(u University) string { return u.Address }).
					WithName("address").
					Required(),
			)),
	)

	teacherEve := Teacher{
		Name: "Eve",
		University: University{
			Name:    "Poznan University of Technology",
			Address: "",
		},
	}
	teacherJohn := Teacher{
		Name: "John",
		University: University{
			Name:    "Poznan University of Technology",
			Address: "",
		},
	}
	teachers := []Teacher{teacherEve, teacherJohn}

	// We're using a mock testing.T to capture the error produced by the assertion.
	mt := new(mockTestingT)

	err := teacherValidator.WithNameFunc(func(s Teacher) string { return s.Name }).ValidateSlice(teachers)
	govytest.AssertError(mt, err,
		govytest.ExpectedRuleError{
			PropertyName:   "university.address",
			Code:           "greater_than",
			ValidatorName:  "Eve",
			ValidatorIndex: ptr(0),
		},
		govytest.ExpectedRuleError{
			PropertyName:    "name",
			ContainsMessage: "one of",
			ValidatorName:   "John",
			ValidatorIndex:  ptr(1),
		},
	)

	// This will print the error produced by the assertion.
	fmt.Println(mt.recordedError)

	// Output:
	// Expected error was not found.
  // EXPECTED:
  // {
  //   "propertyName": "university.address",
  //   "code": "greater_than",
  //   "validatorName": "Eve",
  //   "validatorIndex": 0
  // }
  // ACTUAL:
  // [
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

// If you don't want to verify all the errors returned by [govy.Validator],
// but ensure a single, expected error is produced use [govytest.AssertErrorContains]
// instead of [govytest.AssertError].
//
// To demonstrate the erroneous output of [govytest.AssertErrorContains]
// we'll first match the error and then fail the assertion.
func ExampleAssertErrorContains() {
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

	// We're using a mock testing.T to capture the error produced by the assertion.
	mt := new(mockTestingT)

	// Match the error.
	err := teacherValidator.WithName("John").Validate(teacher)
	govytest.AssertErrorContains(mt, err, govytest.ExpectedRuleError{
		PropertyName: "name",
		Code:         "one_of",
	})

	// Fail to match the error.
	err = teacherValidator.WithName("John").Validate(teacher)
	govytest.AssertErrorContains(mt, err, govytest.ExpectedRuleError{
		PropertyName: "university.address",
		Code:         "greater_than",
	})

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
	//         "error": "must be one of: Jake, George",
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
