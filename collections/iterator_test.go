package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SliceIterator(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		hasNext  bool
		result   int
	}{
		{
			name:     "Has elements",
			elements: []int{1, 2, 3},
			hasNext:  true,
			result:   6,
		},
		{
			name:     "Has no elements",
			elements: nil,
			hasNext:  false,
			result:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var iterator = IteratorFromSlice(tt.elements)
			assert.Equal(t, tt.hasNext, iterator.HasNext())

			var actual int
			for iterator.HasNext() {
				actual += iterator.Next()
			}
			assert.Equal(t, tt.result, actual)
		})
	}
}
