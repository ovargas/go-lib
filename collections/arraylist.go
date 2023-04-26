package collections

import (
	"reflect"
)

type (

	// ArrayList is a resizable-array implementation of the List interface.
	ArrayList[T any] struct {
		elements []T
	}

	EqualFn[T any] func(T, T) bool

	// Option is a functional option type for ArrayList.
	Option[T any] func(collectionInitializer[T])
)

var _ List[any] = (*ArrayList[any])(nil)

// NewArrayList returns a new ArrayList.
func NewArrayList[T any]() *ArrayList[T] {
	var a = &ArrayList[T]{
		elements: make([]T, 0),
	}
	return a
}

// NewArrayListWithElements returns a new ArrayList with the specified elements.
func NewArrayListWithElements[T any](elements []T) *ArrayList[T] {
	var a = &ArrayList[T]{
		elements: make([]T, len(elements)),
	}
	copy(a.elements, elements)
	return a
}

// Add adds the specified element to this list.
func (a *ArrayList[T]) Add(t T) bool {
	a.elements = append(a.elements, t)
	return true
}

// AddAt adds the specified element at the specified position in this list.
func (a *ArrayList[T]) AddAt(i int, t T) bool {
	if i < 0 || i > len(a.elements) {
		return false
	}
	a.elements = append(a.elements[:i+1], a.elements[i:]...)
	a.elements[i] = t
	return true
}

// AddAll adds all the elements in the specified collection to this list.
func (a *ArrayList[T]) AddAll(t []T) bool {
	a.elements = append(a.elements, t...)
	return true
}

// Remove removes the first occurrence of the specified element from this list, if it is present.
func (a *ArrayList[T]) Remove(t T) bool {
	for i, e := range a.elements {
		if Equal[T](e, t) {
			a.elements = append(a.elements[:i], a.elements[i+1:]...)
			return true
		}
	}
	return false
}

// RemoveAt removes the element at the specified position in this list.
func (a *ArrayList[T]) RemoveAt(i int) T {
	var e = a.elements[i]
	a.elements = append(a.elements[:i], a.elements[i+1:]...)
	return e
}

// RemoveIf removes all the elements that satisfy the given predicate.
func (a *ArrayList[T]) RemoveIf(f Predicate[T]) bool {
	removed := false
	for i := 0; i < len(a.elements); i++ {
		if f(a.elements[i]) {
			a.elements = append(a.elements[:i], a.elements[i+1:]...)
			i -= 1
			removed = true
		}
	}

	return removed
}

// Contains returns true if this list contains the specified element.
func (a *ArrayList[T]) Contains(t T) bool {
	return a.IndexOf(t) != -1
}

// IndexOf returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (a *ArrayList[T]) IndexOf(t T) int {
	for i, e := range a.elements {
		if Equal[T](e, t) {
			return i
		}
	}
	return -1
}

// IsEmpty returns true if this list contains no elements.
func (a *ArrayList[T]) IsEmpty() bool {
	return a.Size() == 0
}

// Clear removes all the elements from this list.
func (a *ArrayList[T]) Clear() {
	a.elements = nil
}

// Size returns the number of elements in this list.
func (a *ArrayList[T]) Size() int {
	return len(a.elements)
}

// Get returns the element at the specified position in this list.
func (a *ArrayList[T]) Get(i int) T {
	return a.elements[i]
}

// Set replaces the element at the specified position in this list with the specified element.
func (a *ArrayList[T]) Set(i int, t T) T {
	old := a.elements[i]
	a.elements[i] = t
	return old
}

// ToArray returns an array containing all the elements in this list in proper sequence.
func (a *ArrayList[T]) ToArray() []T {
	arrays := make([]T, a.Size())
	copy(arrays, a.elements)
	return arrays
}

// Iterator returns an iterator over the elements in this list in proper sequence.
func (a *ArrayList[T]) Iterator() Iterator[T] {
	return IteratorFromSlice[T](a.elements)
}

// Equal returns true if the two values are equal.
func Equal[T any](a, b T) bool {
	return reflect.DeepEqual(a, b)
}
