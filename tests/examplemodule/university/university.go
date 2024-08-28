package university

import (
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

type University struct {
	Name string `json:"name"`
}

var UniversityValidation = govy.New(
	govy.For(func(u University) string { return u.Name }).
		Rules(rules.Required[string]()),
)
