package main

import (
	"examplemodule/university"
	"fmt"

	"github.com/nobl9/govy/pkg/govyconfig"
)

type Teacher struct {
	Name       string
	University university.University `json:"university"`
}

func main() {
	john := Teacher{
		Name: "John",
		University: university.University{
			Name: "Oxford",
		},
	}
	err := teacherValidation.Validate(john)
	fmt.Println(err)
}

func init() {
	govyconfig.SetNameInferMode(govyconfig.NameInferModeRuntime)
}
