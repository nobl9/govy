package govy

import "fmt"

// whenConfig defines optional configuration options for the When conditions.
type whenConfig struct {
	description string
}

// WhenOption applies selected option to [whenConfig].
type WhenOption func(whenConfig) whenConfig

// WhenDescription sets the description for the When condition.
func WhenDescription(description string) func(whenConfig) whenConfig {
	return func(opt whenConfig) whenConfig {
		opt.description = description
		return opt
	}
}

// WhenDescriptionf sets the description for the When condition.
func WhenDescriptionf(format string, a ...any) func(whenConfig) whenConfig {
	return func(opt whenConfig) whenConfig {
		opt.description = fmt.Sprintf(format, a...)
		return opt
	}
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

func (p predicateMatcher[T]) when(predicate Predicate[T], opts ...WhenOption) predicateMatcher[T] {
	container := predicateContainer[T]{predicate: predicate}
	options := whenConfig{}
	for _, opt := range opts {
		options = opt(options)
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
