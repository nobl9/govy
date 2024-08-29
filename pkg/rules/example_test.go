// nolint: lll
package rules_test

import (
	"fmt"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

type Teacher struct {
	Students []Student `json:"students"`
}

type Student struct {
	Index string `json:"index"`
}

func ExampleSliceUnique() {
	v := govy.New(
		govy.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			Rules(rules.SliceUnique(func(v Student) string { return v.Index },
				"each student must have unique index")),
	)
	teacher := Teacher{
		Students: []Student{
			{Index: "foo"},
			{Index: "bar"}, // 2nd element
			{Index: "baz"},
			{Index: "bar"}, // 4th element
		},
	}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'students' with value '[{"index":"foo"},{"index":"bar"},{"index":"baz"},{"index":"bar"}]':
	//     - elements are not unique, 2nd and 4th elements collide based on constraints: each student must have unique index
}
