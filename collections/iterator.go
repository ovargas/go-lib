package collections

type (
	// ChannelIterator is an iterator for a slice.
	ChannelIterator[T any] struct {
		channel chan T
	}
)

// IteratorFromSlice instantiates a new iterator from a slice.
func IteratorFromSlice[T any](elements []T) *ChannelIterator[T] {
	channel := make(chan T, len(elements))
	defer close(channel)
	for _, e := range elements {
		channel <- e
	}

	return &ChannelIterator[T]{channel: channel}
}

func IteratorFromSet[T comparable](elements map[T]struct{}) *ChannelIterator[T] {
	channel := make(chan T, len(elements))
	defer close(channel)
	for e := range elements {
		channel <- e
	}

	return &ChannelIterator[T]{channel: channel}
}

func IteratorFromMap[K comparable, T any](elements map[K]T) *ChannelIterator[*Entry[K, T]] {
	channel := make(chan *Entry[K, T], len(elements))
	defer close(channel)
	for k, v := range elements {
		channel <- &Entry[K, T]{key: k, value: v}
	}

	return &ChannelIterator[*Entry[K, T]]{channel: channel}
}

// HasNext returns true if there are more elements to iterate over.
func (i *ChannelIterator[T]) HasNext() bool {
	return len(i.channel) > 0
}

// Next returns the next element.
func (i *ChannelIterator[T]) Next() T {
	return <-i.channel
}
