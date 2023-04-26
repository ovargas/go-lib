package collections

import "errors"

type (

	// Dictionary is an extension of a map to reduce the common boilerplate coding.
	Dictionary[K comparable, T any] map[K]T

	Entry[K comparable, T any] struct {
		key   K
		value T
	}
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

// Key returns the key of the entry
func (e *Entry[K, T]) Key() K {
	return e.key
}

// Value returns the value of the entry
func (e *Entry[K, T]) Value() T {
	return e.value
}

// Has returns true if the key exists in the dictionary
func (d Dictionary[K, T]) Has(key K) bool {
	_, ok := d[key]
	return ok
}

// Get returns the value of the key in the dictionary
func (d Dictionary[K, T]) Get(key K) (T, error) {
	v, ok := d[key]
	if !ok {
		return v, ErrKeyNotFound
	}
	return v, nil
}

// GetOrDefault returns the value of the key in the dictionary or the default value if the key does not exist
func (d Dictionary[K, T]) GetOrDefault(key K, value T) T {
	v, ok := d[key]
	if !ok {
		return value
	}

	return v
}

// Set sets the value of the key in the dictionary
func (d Dictionary[K, T]) Set(key K, value T) {
	d[key] = value
}

// Remove removes the key from the dictionary
func (d Dictionary[K, T]) Remove(key K) {
	delete(d, key)
}

// Keys returns the keys of the dictionary
func (d Dictionary[K, T]) Keys() []K {
	keys := make([]K, 0, len(d))
	for k := range d {
		keys = append(keys, k)
	}
	return keys
}

// Values returns the values of the dictionary
func (d Dictionary[K, T]) Values() []T {
	values := make([]T, 0, len(d))
	for _, v := range d {
		values = append(values, v)
	}
	return values
}

// Merge merges the other dictionary into the dictionary
func (d Dictionary[K, T]) Merge(other Dictionary[K, T]) {
	for k, v := range other {
		d[k] = v
	}
}

// Iterator returns an iterator for the dictionary
func (d Dictionary[K, T]) Iterator() *ChannelIterator[*Entry[K, T]] {
	return IteratorFromMap[K, T](d)
}
