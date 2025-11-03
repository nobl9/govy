package examples

import (
	"fmt"
	"regexp"
	"time"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

func Example_basicUsage() {
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
			Required().
			Rules(rules.StringMatchRegexp(
				regexp.MustCompile(`[\w\s.]+, \d{2}-\d{3} \w+`),
			).
				WithDetails("Polish address format must consist of the main address and zip code").
				WithExamples("5 M. Skłodowska-Curie Square, 60-965 Poznan")),
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
			Include(universityValidation),
	).
		When(func(t Teacher) bool { return t.Age < 50 })

	teacher := Teacher{
		Name:       "John",
		MiddleName: nil, // Validation for nil pointers by default is skipped.
		Age:        48,
		Students: []Student{
			{Index: "918230014"},
			{Index: "9182300123"},
			{Index: "918230014"},
		},
		University: University{
			Name:    "",
			Address: "10th University St.",
		},
	}

	if err := teacherValidator.WithName("John").Validate(teacher); err != nil {
		fmt.Println(err)
	}
	// When condition is not met, no validation errors.
	johnFromTheFuture := teacher
	johnFromTheFuture.Age = 51
	if err := teacherValidator.WithName("John From The Future").Validate(johnFromTheFuture); err != nil {
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
	//   - 'university.name':
	//     - property is required but was empty
	//   - 'university.address' with value '10th University St.':
	//     - string must match regular expression: '[\w\s.]+, \d{2}-\d{3} \w+' (e.g. '5 M. Skłodowska-Curie Square, 60-965 Poznan'); Polish address format must consist of the main address and zip code
}
