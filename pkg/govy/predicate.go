package govy

import "fmt"

// whenOptions defines optional configuration options for the When conditions.
type whenOptions struct {
	description string
}

// WhenOption applies selected option to [whenConfig].
type WhenOption func(options whenOptions) whenOptions

// WhenDescription sets the description for the When condition.
func WhenDescription(description string) WhenOption {
	return func(options whenOptions) whenOptions {
		options.description = description
		return options
	}
}

// WhenDescriptionf sets the description for the When condition using printf-like formatting.
func WhenDescriptionf(format string, a ...any) WhenOption {
	return func(options whenOptions) whenOptions {
		options.description = fmt.Sprintf(format, a...)
		return options
	}
}

// Predicate defines a function that returns a boolean value.
type Predicate[T any] func(T) bool

type predicateContainer[T any] struct {
	whenOptions
	predicate Predicate[T]
}

type predicateMatcher[T any] struct {
	predicates []predicateContainer[T]
}

// when adds a [Predicate] to the [predicateMatcher] and applies all provided [WhenOption].
func (p predicateMatcher[T]) when(predicate Predicate[T], opts ...WhenOption) predicateMatcher[T] {
	container := predicateContainer[T]{predicate: predicate}
	for _, opt := range opts {
		container.whenOptions = opt(container.whenOptions)
	}
	p.predicates = append(p.predicates, container)
	return p
}

// matchPredicates evaluates each [Predicate] and returns false if ANY of them fails.
func (p predicateMatcher[T]) matchPredicates(st T) bool {
	for _, predicate := range p.predicates {
		if !predicate.predicate(st) {
			return false
		}
	}
	return true
}
