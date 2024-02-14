package generics

import "errors"

// Transformation is a callback type that used by Map, MustMap.
// CallBack of this type transforms element of type `From` to `To`.
// If transformation is successful, function must return desired element and nil error.
// If not, then return *new(T) and corresponding error.
type Transformation[From, To any] func(From) (To, error)

var (
	// ErrTransformationIsNotProvided is an error telling that callback for TryMap equals nil.
	ErrTransformationIsNotProvided = errors.New("transformation is not provided")
)

// Map creates new slice of type []To using provided Transformation.
//
// Map iterates through x and calls cb(el) adding its value into resulting slice.
// If cb(el) returns error then this transformation will be skipped.
//
// If x == nil returns nil.
//
// If cb == nil Map panics with ErrTransformationIsNotProvided.
func Map[S ~[]From, From, To any](x S, cb Transformation[From, To]) []To {
	if x == nil {
		return nil
	}

	if cb == nil {
		panic(ErrTransformationIsNotProvided)
	}

	result := make([]To, 0, len(x))
	for _, el := range x {
		newEl, err := cb(el)
		if err != nil {
			continue
		}
		result = append(result, newEl)
	}
	return result
}

// TryMap is analog of Map returning error if it happens.
//
// TryMap iterates through x and calls cb(el) adding its value into resulting slice.
//
// If cb(el) returns error, TryMap stops
// and returns slice (with elements that have already been processed without error) and error as it is from cb(el).
//
// If x == nil returns nil, nil.
//
// If cb == nil TryMap returns nil, ErrTransformationIsNotProvided.
func TryMap[S ~[]From, From, To any](x S, cb Transformation[From, To]) ([]To, error) {
	if x == nil {
		return nil, nil
	}

	if cb == nil {
		return nil, ErrTransformationIsNotProvided
	}

	result := make([]To, 0, len(x))
	for _, el := range x {
		newEl, err := cb(el)
		if err != nil {
			return result, err
		}
		result = append(result, newEl)
	}
	return result, nil
}
