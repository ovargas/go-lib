package collections

type (

	// ValueSet is a set of values.
	ValueSet[T comparable] struct {
		elements map[T]struct{}
	}
)

var _ Set[string] = (*ValueSet[string])(nil)

// NewValueSet returns a new ValueSet.
func NewValueSet[T comparable]() *ValueSet[T] {
	var a = &ValueSet[T]{
		elements: make(map[T]struct{}),
	}
	return a
}

// NewValueSetWithElements returns a new ValueSet with the specified elements.
func NewValueSetWithElements[T comparable](elements []T) *ValueSet[T] {
	var a = &ValueSet[T]{
		elements: make(map[T]struct{}),
	}
	a.AddAll(elements)
	return a
}

// Add adds the specified element to this set.
func (h *ValueSet[T]) Add(t T) bool {
	h.elements[t] = struct{}{}
	return true
}

// AddAll adds all the elements in the specified collection to this set.
func (h *ValueSet[T]) AddAll(ts []T) bool {
	for _, t := range ts {
		h.elements[t] = struct{}{}
	}
	return true
}

// Remove removes the first occurrence of the specified element from this set, if it is present.
func (h *ValueSet[T]) Remove(t T) bool {
	if _, ok := h.elements[t]; !ok {
		return false
	}
	delete(h.elements, t)
	return true
}

// RemoveIf removes all the elements that satisfy the given predicate.
func (h *ValueSet[T]) RemoveIf(f Predicate[T]) bool {
	removed := false
	for k := range h.elements {
		if f(k) {
			delete(h.elements, k)
			removed = true
		}
	}
	return removed
}

// Clear removes all the elements from this set.
func (h *ValueSet[T]) Clear() {
	for k := range h.elements {
		delete(h.elements, k)
	}
}

// Size returns the number of elements in this set.
func (h *ValueSet[T]) Size() int {
	return len(h.elements)
}

// IsEmpty returns true if this set contains no elements.
func (h *ValueSet[T]) IsEmpty() bool {
	return len(h.elements) == 0
}

// Contains returns true if this set contains the specified element.
func (h *ValueSet[T]) Contains(t T) bool {
	_, ok := h.elements[t]
	return ok
}

// Iterator returns an iterator over the elements in this set.
func (h *ValueSet[T]) Iterator() Iterator[T] {
	return IteratorFromSet(h.elements)
}

// ToArray returns an array containing all the elements in this set.
func (h *ValueSet[T]) ToArray() []T {
	array := make([]T, len(h.elements))
	i := 0
	for k := range h.elements {
		array[i] = k
		i++
	}
	return array
}
