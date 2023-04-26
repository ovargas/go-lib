package slice

import "fmt"

type (
	Predicate[T any] func(T) bool

	CovertError struct {
		Index int
	}
)

func (e *CovertError) Error() string {
	return fmt.Sprintf("cannot convert element at index %d", e.Index)
}

// ToInterface converts a slice of any type to a slice of interface{}
func ToInterface[T any](slice []T) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

// FromInterface converts a slice of interface{} to a slice of any type
func FromInterface[T any](slice []interface{}) ([]T, error) {
	result := make([]T, len(slice))
	var ok bool
	for i, v := range slice {
		result[i], ok = v.(T)
		if !ok {
			return nil, &CovertError{Index: i}
		}
	}
	return result, nil
}

// Filter filters a slice based on a predicate
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
