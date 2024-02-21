package generics

import "errors"

// Predicate is a type used by Filter, FindFirst functions.
type Predicate[T any] func(x T) bool

// ErrPredicateIsNotProvided is an error telling that predicate for Filter equals nil.
var ErrPredicateIsNotProvided = errors.New("predicate is not provided")

// Filter makes new slice with those values that passed by predicate.
//
// Filter iterates through x and calls p(el).
//
// If p(el) == true then resulting slice will contain el.
//
// If p(el) == false then resulting slice will not contain el.
//
// # Edge cases:
//
// If x == nil returns nil.
//
// If p == nil Filter panics with ErrPredicateIsNotProvided.
func Filter[S ~[]E, E any](x S, p Predicate[E]) S {
	if x == nil {
		return nil
	}

	if p == nil {
		panic(ErrPredicateIsNotProvided)
	}

	result := make(S, 0, len(x))
	for _, el := range x {
		if p(el) {
			result = append(result, el)
		}
	}
	return result
}

// FindFirst return first element which is acceptable by predicate.
//
// FindFirst iterates through x and calls p(el).
//
// If p(el) == true then returns el, true.
//
// If p(el) == false then returns zero value of type E, false.
//
// # Edge cases:
//
// If x == nil returns zero value, false.
//
// If p == nil FindFirst panics with ErrPredicateIsNotProvided.
func FindFirst[S ~[]E, E any](x S, p Predicate[E]) (E, bool) {
	if x == nil {
		return *new(E), false
	}

	if p == nil {
		panic(ErrPredicateIsNotProvided)
	}

	for _, el := range x {
		if p(el) {
			return el, true
		}
	}
	return *new(E), false
}
