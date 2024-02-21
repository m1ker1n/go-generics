package generics

// MapKeys returns slice of keys contained in map m.
//
// Order of keys is random.
//
// # Edge cases:
//
// If m == nil returns nil.
func MapKeys[M map[K]V, K comparable, V any](m M) []K {
	if m == nil {
		return nil
	}

	result := make([]K, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// MapValues returns slice of values contained in map m.
//
// Order of values is random.
//
// # Edge cases:
//
// If m == nil returns nil.
func MapValues[M map[K]V, K comparable, V any](m M) []V {
	if m == nil {
		return nil
	}

	result := make([]V, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
