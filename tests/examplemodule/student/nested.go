package student

import (
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

type Student struct {
	Name string `json:"name"`
}

var StudentValidation = govy.New(
	govy.For(func(s Student) string { return s.Name }).
		Rules(rules.EQ("Jake")),
)
