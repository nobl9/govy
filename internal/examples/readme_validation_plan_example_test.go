package examples

import (
	"encoding/json"
	"os"
	"regexp"
	"time"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

func Example_validationPlan() {
	type University struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}
	type Student struct {
		Index string `json:"index"`
	}
	type Teacher struct {
		Name       string        `json:"name"`
		Age        time.Duration `json:"age"`
		Students   []Student     `json:"students"`
		MiddleName *string       `json:"middleName,omitempty"`
		University University    `json:"university"`
	}

	universityValidation := govy.New(
		govy.For(func(u University) string { return u.Name }).
			WithName("name").
			Required(),
		govy.For(func(u University) string { return u.Address }).
			WithName("address").
			Rules(rules.StringMatchRegexp(
				regexp.MustCompile(`[\w\s.]+, [0-9]{2}-[0-9]{3} \w+`),
			).
				WithDetails("Polish address format must consist of the main address and zip code").
				WithExamples("5 M. Skłodowska-Curie Square, 60-965 Poznan")).
			When(func(u University) bool { return u.Name == "PUT" },
				govy.WhenDescription("University name is PUT University")),
	)
	studentValidator := govy.New(
		govy.For(func(s Student) string { return s.Index }).
			WithName("index").
			Rules(rules.StringLength(9, 9)),
	)
	teacherValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(
				rules.StringNotEmpty(),
				rules.OneOf("Jake", "George")),
		govy.ForPointer(func(t Teacher) *string { return t.MiddleName }).
			WithName("middleName").
			Rules(rules.StringTitle()),
		govy.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			Rules(
				rules.SliceMaxLength[[]Student](2),
				rules.SliceUnique(func(v Student) string { return v.Index })).
			IncludeForEach(studentValidator),
		govy.For(func(t Teacher) University { return t.University }).
			WithName("university").
			Include(universityValidation).
			When(func(t Teacher) bool { return t.Name == "John" },
				govy.WhenDescription("Teacher name is John")),
	).
		WithName("Teacher")

	plan := govy.Plan(teacherValidator)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	_ = enc.Encode(plan)

	// Output:
	// {
	//   "name": "Teacher",
	//   "properties": [
	//     {
	//       "path": "$.middleName",
	//       "typeInfo": {
	//         "name": "string",
	//         "kind": "string"
	//       },
	//       "rules": [
	//         {
	//           "description": "property is optional",
	//           "errorCode": "optional"
	//         },
	//         {
	//           "description": "each word in a string must start with a capital letter",
	//           "errorCode": "string_title"
	//         }
	//       ]
	//     },
	//     {
	//       "path": "$.name",
	//       "typeInfo": {
	//         "name": "string",
	//         "kind": "string"
	//       },
	//       "values": [
	//         "Jake",
	//         "George"
	//       ],
	//       "rules": [
	//         {
	//           "description": "string cannot be empty",
	//           "errorCode": "string_not_empty"
	//         },
	//         {
	//           "description": "must be one of: Jake, George",
	//           "errorCode": "one_of"
	//         }
	//       ]
	//     },
	//     {
	//       "path": "$.students",
	//       "typeInfo": {
	//         "name": "[]Student",
	//         "kind": "[]struct",
	//         "package": "github.com/nobl9/govy/internal/examples"
	//       },
	//       "rules": [
	//         {
	//           "description": "length must be less than or equal to 2",
	//           "errorCode": "slice_max_length"
	//         },
	//         {
	//           "description": "elements must be unique",
	//           "errorCode": "slice_unique"
	//         }
	//       ]
	//     },
	//     {
	//       "path": "$.students[*].index",
	//       "typeInfo": {
	//         "name": "string",
	//         "kind": "string"
	//       },
	//       "rules": [
	//         {
	//           "description": "length must be between 9 and 9",
	//           "errorCode": "string_length"
	//         }
	//       ]
	//     },
	//     {
	//       "path": "$.university.address",
	//       "typeInfo": {
	//         "name": "string",
	//         "kind": "string"
	//       },
	//       "rules": [
	//         {
	//           "description": "string must match regular expression: '[\\w\\s.]+, [0-9]{2}-[0-9]{3} \\w+'",
	//           "details": "Polish address format must consist of the main address and zip code",
	//           "errorCode": "string_match_regexp",
	//           "conditions": [
	//             "Teacher name is John",
	//             "University name is PUT University"
	//           ],
	//           "examples": [
	//             "5 M. Skłodowska-Curie Square, 60-965 Poznan"
	//           ]
	//         }
	//       ]
	//     },
	//     {
	//       "path": "$.university.name",
	//       "typeInfo": {
	//         "name": "string",
	//         "kind": "string"
	//       },
	//       "rules": [
	//         {
	//           "description": "property is required",
	//           "errorCode": "required",
	//           "conditions": [
	//             "Teacher name is John"
	//           ]
	//         }
	//       ]
	//     }
	//   ]
	// }
}
