package optional

import "reflect"

type (
	Optional[T any] struct {
		value T
	}
)

// Of returns an Optional with the given value.
func Of[T any](value T) *Optional[T] {
	return &Optional[T]{value: value}
}

// OfNullable returns an Optional with the given value or an empty Optional if the given value is nil.
func OfNullable[T any](value T) *Optional[T] {
	if reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil() {
		return Empty[T]()
	}
	return &Optional[T]{value: value}
}

// IsPresent returns true if the Optional contains a value.
func (o *Optional[T]) IsPresent() bool {

	if reflect.ValueOf(o.value).Kind() == reflect.Ptr {
		return !reflect.ValueOf(o.value).IsNil()
	}

	return true
}

// Get returns the value of the Optional.
func (o *Optional[T]) Get() T {
	return o.value
}

// IfPresent calls the given function if the Optional contains a value.
func (o *Optional[T]) IfPresent(f func(T)) {
	if o.IsPresent() {
		f(o.value)
	}
}

// IfPresentOrElse calls the given function if the Optional contains a value or an alternative function if the Optional is empty.
func (o *Optional[T]) IfPresentOrElse(f func(T), orElse func()) {
	if o.IsPresent() {
		f(o.value)
	} else {
		orElse()
	}
}

// OrElse returns the value of the Optional or the given value if the Optional is empty.
func (o *Optional[T]) OrElse(other T) T {
	if o.IsPresent() {
		return o.value
	}
	return other
}

// OrElseGet returns the value of the Optional or the value returned by the given function if the Optional is empty.
func (o *Optional[T]) OrElseGet(f func() T) T {
	if o.IsPresent() {
		return o.value
	}
	return f()
}

// Empty returns an empty Optional.
func Empty[T any]() *Optional[T] {
	return &Optional[T]{}
}

// Map returns an Optional with the result of the given function applied to the value of the Optional.
func Map[T any, R any](o *Optional[T], f func(T) R) *Optional[R] {
	if o.IsPresent() {
		return Of[R](f(o.value))
	}
	return Empty[R]()
}
