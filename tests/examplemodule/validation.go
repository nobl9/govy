package main

import (
	"examplemodule/university"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

var teacherValidation = govy.New(
	govy.For(func(t Teacher) string { return t.Name }).
		Rules(rules.EQ("Paul")),
	govy.For(func(t Teacher) string { return t.University.Name }).
		Rules(rules.EQ("Cambridge")),
	govy.For(func(t Teacher) university.University { return t.University }).
		Include(university.UniversityValidation),
)
