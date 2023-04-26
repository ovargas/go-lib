package collections

type (
	// Predicate is a function that returns true if the specified element satisfies the predicate.
	Predicate[T any] func(T) bool

	// Iterable is an interface that represents a collection of elements.
	Iterable[T any] interface {
		Iterator() Iterator[T]
	}

	// Collection is an interface that represents a collection of elements.
	Collection[T any] interface {
		Iterable[T]

		// Add adds the specified element to this collection.
		Add(T) bool

		// AddAll adds all the elements in the specified collection to this collection.
		AddAll([]T) bool

		// Clear removes all the elements from this collection.
		Clear()

		// Contains returns true if this collection contains the specified element.
		Contains(T) bool

		// IsEmpty returns true if this collection contains no elements.
		IsEmpty() bool

		// Remove removes the first occurrence of the specified element from this collection, if it is present.
		Remove(T) bool

		// RemoveIf removes all the elements that satisfy the given predicate.
		RemoveIf(Predicate[T]) bool

		// Size returns the number of elements in this collection.
		Size() int

		// ToArray returns an array containing all the elements in this collection.
		ToArray() []T
	}

	// List is an interface that represents a collection of elements that can be accessed by index.
	List[T any] interface {
		Collection[T]

		// Get returns the element at the specified position in this list.
		Get(int) T

		// AddAt adds the specified element at the specified position in this list.
		AddAt(int, T) bool

		// Set replaces the element at the specified position in this list with the specified element.
		Set(int, T) T

		// IndexOf returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
		IndexOf(T) int

		// RemoveAt removes the element at the specified position in this list.
		RemoveAt(int) T
	}

	// Set is an interface that represents a collection of elements that cannot contain duplicate elements.
	Set[T comparable] interface {
		Collection[T]
	}

	// Iterator is an interface that represents an iterator over a collection.
	Iterator[T any] interface {

		// HasNext returns true if there are more elements to iterate over.
		HasNext() bool

		// Next returns the next element.
		Next() T
	}

	collectionInitializer[T any] interface {
		withElements([]T)
	}
)

// WithElements returns an Option that initializes a collection with the specified elements.
func WithElements[T any](e []T) Option[T] {
	return func(c collectionInitializer[T]) {
		c.withElements(e)
	}
}
