package govy

import "fmt"

// WhenOptions defines optional parameters for the When conditions.
type WhenOptions struct {
	description string
}

// WhenDescription sets the description for the When condition.
func WhenDescription(format string, a ...any) WhenOptions {
	return WhenOptions{description: fmt.Sprintf(format, a...)}
}

// Predicate defines a function that returns a boolean value.
type Predicate[T any] func(T) bool

type predicateContainer[T any] struct {
	predicate   Predicate[T]
	description string
}

type predicateMatcher[T any] struct {
	predicates []predicateContainer[T]
}

func (p predicateMatcher[T]) when(predicate Predicate[T], opts ...WhenOptions) predicateMatcher[T] {
	container := predicateContainer[T]{predicate: predicate}
	for _, opt := range opts {
		if opt.description != "" {
			container.description = opt.description
		}
	}
	p.predicates = append(p.predicates, container)
	return p
}

func (p predicateMatcher[T]) matchPredicates(st T) bool {
	for _, predicate := range p.predicates {
		if !predicate.predicate(st) {
			return false
		}
	}
	return true
}
