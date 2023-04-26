package collections

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestValueSet_Add(t *testing.T) {
	set := NewValueSet[int]()
	set.Add(1)
	set.Add(2)

	assert.Equal(t, 2, set.Size())
}

func TestValueSet_AddAll(t *testing.T) {
	set := NewValueSet[int]()
	set.AddAll([]int{1, 2, 3})
	set.AddAll([]int{4, 5, 6})

	assert.Equal(t, 6, set.Size())

}

func TestValueSet_Clear(t *testing.T) {
	set := NewValueSet[int]()
	set.AddAll([]int{1, 2, 3})
	set.AddAll([]int{4, 5, 6})

	set.Clear()

	assert.Equal(t, 0, set.Size())
}

func TestValueSet_Contains(t *testing.T) {
	set := NewValueSet[int]()
	set.AddAll([]int{1, 2, 3})
	set.AddAll([]int{4, 5, 6})

	assert.True(t, set.Contains(3))
}

func TestValueSet_IsEmpty(t *testing.T) {
	set := NewValueSet[int]()
	assert.True(t, set.IsEmpty())

	set.AddAll([]int{1, 2, 3})
	set.AddAll([]int{4, 5, 6})

	assert.False(t, set.IsEmpty())
}

func TestValueSet_Iterator(t *testing.T) {
	set := NewValueSet[int]()
	set.AddAll([]int{1, 2, 3})

	it := set.Iterator()

	sum := 0
	for it.HasNext() {
		sum += it.Next()
	}

	assert.Equal(t, 6, sum)
}

func TestValueSet_Remove(t *testing.T) {
	set := NewValueSet[int]()
	set.AddAll([]int{1, 2, 3})

	assert.True(t, set.Remove(2))

	assert.Equal(t, 2, set.Size())
	assert.False(t, set.Contains(2))

	assert.False(t, set.Remove(10))
	assert.Equal(t, 2, set.Size())
}

func TestValueSet_RemoveIf(t *testing.T) {
	set := NewValueSet[int]()
	set.AddAll([]int{1, 2, 3})
	set.AddAll([]int{4, 5, 6})

	set.RemoveIf(func(item int) bool {
		return item%2 == 0
	})

	assert.Equal(t, 3, set.Size())
	assert.False(t, set.Contains(2))
	assert.False(t, set.Contains(4))
	assert.False(t, set.Contains(6))

}

func TestValueSet_Size(t *testing.T) {
	set := NewValueSet[int]()
	set.AddAll([]int{1, 2, 3})
	set.AddAll([]int{4, 5, 6})

	assert.Equal(t, 6, set.Size())
}

func TestValueSet_ToArray(t *testing.T) {
	set := NewValueSet[int]()
	set.AddAll([]int{1, 2, 3})
	set.AddAll([]int{4, 5, 6})

	arr := set.ToArray()

	assert.Equal(t, 6, len(arr))
}

func TestValueSet_WithElements(t *testing.T) {
	set := NewValueSetWithElements[int]([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, 6, set.Size())

	arrays := set.ToArray()
	sort.Ints(arrays)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, arrays)
}
