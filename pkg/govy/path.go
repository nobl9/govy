package govy

import "github.com/nobl9/govy/internal/jsonpath"

// Path is a builder for constructing valid JSONPath property paths.
// It ensures proper escaping and formatting of path segments.
//
// Usage:
//
//	govy.NewPath().Name("metadata").Name("name")         // metadata.name
//	govy.NewPath().Name("metadata").Name("labels").Index(0) // metadata.labels[0]
//	govy.NewPath().Name("complex.key")                   // ['complex.key']
type Path = jsonpath.Path

// NewPath creates a new empty [Path].
var NewPath = jsonpath.NewPath
