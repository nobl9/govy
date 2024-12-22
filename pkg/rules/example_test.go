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
	Index     string `json:"index,omitempty"`
	Name      string `json:"name,omitempty"`
	IndexCopy string `json:"indexCopy,omitempty"`
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

func ExampleMutuallyExclusive() {
	v := govy.New(
		govy.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			RulesForEach(rules.MutuallyExclusive(true, map[string]func(Student) any{
				"index": func(s Student) any { return s.Index },
				"name":  func(s Student) any { return s.Name },
			})),
	)
	teacher := Teacher{
		Students: []Student{
			{Index: "foo"},
			{Index: "bar", Name: "John"},
			{Name: "Eve"},
			{},
		},
	}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'students[1]' with value '{"index":"bar","name":"John"}':
	//     - [index, name] properties are mutually exclusive, provide only one of them
	//   - 'students[3]':
	//     - one of [index, name] properties must be set, none was provided
}

func ExampleOneOfProperties() {
	v := govy.New(
		govy.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			RulesForEach(rules.OneOfProperties(map[string]func(Student) any{
				"index": func(s Student) any { return s.Index },
				"name":  func(s Student) any { return s.Name },
			})),
	)
	teacher := Teacher{
		Students: []Student{
			{Index: "foo"},
			{},
			{Name: "John"},
			{Index: "bar", Name: "Eve"},
		},
	}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'students[1]':
	//     - one of [index, name] properties must be set, none was provided
}

func ExampleEqualProperties() {
	v := govy.New(
		govy.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			RulesForEach(rules.EqualProperties(rules.CompareFunc, map[string]func(Student) any{
				"index":     func(s Student) any { return s.Index },
				"indexCopy": func(s Student) any { return s.IndexCopy },
			})),
	)
	teacher := Teacher{
		Students: []Student{
			{Index: "foo", IndexCopy: "foo"},
			{Index: "bar"},
			{IndexCopy: "foo"},
			{}, // Both index and indexCopy are empty strings, and thus equal.
		},
	}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'students[1]' with value '{"index":"bar"}':
	//     - all of [index, indexCopy] properties must be equal, but 'index' is not equal to 'indexCopy'
	//   - 'students[2]' with value '{"indexCopy":"foo"}':
	//     - all of [index, indexCopy] properties must be equal, but 'index' is not equal to 'indexCopy'
}
