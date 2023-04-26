package collections

// Filter returns a new slice containing the elements of the slice that satisfy the given predicate.
func Filter[T any](it Iterator[T], predicate Predicate[T]) []T {
	var result []T
	for it.HasNext() {
		v := it.Next()
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
