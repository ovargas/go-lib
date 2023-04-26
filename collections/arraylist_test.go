package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayList_Add(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	assert.Equal(t, 3, list.Size())
}

func TestArrayList_AddAt(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	list.AddAt(1, 4)

	assert.Equal(t, 4, list.Size())
	assert.Equal(t, 4, list.Get(1))
	assert.Equal(t, []int{1, 4, 2, 3}, list.ToArray())
}

func TestArrayList_AddAll(t *testing.T) {
	list := NewArrayList[int]()
	list.AddAll([]int{1, 2, 3})
	list.AddAll([]int{4, 5, 6})

	assert.Equal(t, 6, list.Size())
}

func TestArrayList_Clear(t *testing.T) {
	list := NewArrayList[int]()
	list.AddAll([]int{1, 2, 3})
	list.AddAll([]int{4, 5, 6})

	list.Clear()

	assert.Equal(t, 0, list.Size())
}

func TestArrayList_Contains(t *testing.T) {
	list := NewArrayList[int]()
	list.AddAll([]int{1, 2, 3})
	list.AddAll([]int{4, 5, 6})

	assert.True(t, list.Contains(3))
	assert.False(t, list.Contains(7))
}

func TestArrayList_Get(t *testing.T) {
	list := NewArrayList[int]()
	list.AddAll([]int{1, 2, 3})
	list.AddAll([]int{4, 5, 6})

	assert.Equal(t, 3, list.Get(2))
}

func TestArrayList_IndexOf(t *testing.T) {
	list := NewArrayList[int]()
	list.AddAll([]int{1, 2, 3})
	list.AddAll([]int{4, 5, 6})

	assert.Equal(t, 2, list.IndexOf(3))
}

func TestArrayList_IsEmpty(t *testing.T) {
	list := NewArrayList[int]()
	assert.True(t, list.IsEmpty())

	list.AddAll([]int{1, 2, 3})
	list.AddAll([]int{4, 5, 6})

	assert.False(t, list.IsEmpty())
}

func TestArrayList_Iterator(t *testing.T) {
	list := NewArrayList[int]()
	list.AddAll([]int{1, 2, 3})
	list.AddAll([]int{4, 5, 6})

	var actual int
	for i := list.Iterator(); i.HasNext(); {
		actual += i.Next()
	}
	assert.Equal(t, 21, actual)
}

func TestArrayList_Remove(t *testing.T) {
	list := NewArrayList[int]()
	list.AddAll([]int{1, 2, 3})
	list.AddAll([]int{4, 5, 6})

	assert.True(t, list.Remove(3))

	assert.Equal(t, 5, list.Size())
	assert.False(t, list.Contains(3))

	assert.False(t, list.Remove(10))
	assert.Equal(t, 5, list.Size())
}

func TestArrayList_RemoveAt(t *testing.T) {
	list := NewArrayList[int]()
	list.AddAll([]int{1, 2, 3})
	list.AddAll([]int{4, 5, 6})

	list.RemoveAt(2)

	assert.Equal(t, 5, list.Size())
	assert.False(t, list.Contains(3))
}

func TestArrayList_RemoveIf(t *testing.T) {
	list := NewArrayList[int]()
	list.AddAll([]int{1, 2, 3})
	list.AddAll([]int{4, 5, 6})

	list.RemoveIf(func(i int) bool {
		return i%2 == 0
	})

	assert.Equal(t, 3, list.Size())
	assert.False(t, list.Contains(4))
	assert.False(t, list.Contains(6))
}

func TestArrayList_Set(t *testing.T) {
	list := NewArrayList[int]()
	list.AddAll([]int{1, 2, 3})
	list.AddAll([]int{4, 5, 6})

	list.Set(2, 7)

	assert.Equal(t, 7, list.Get(2))
}

func TestArrayList_Size(t *testing.T) {
	list := NewArrayList[int]()
	list.AddAll([]int{1, 2, 3})
	list.AddAll([]int{4, 5, 6})

	assert.Equal(t, 6, list.Size())
}

func TestArrayList_ToArray(t *testing.T) {
	list := NewArrayList[int]()
	list.AddAll([]int{1, 2, 3})
	list.AddAll([]int{4, 5, 6})

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, list.ToArray())
}

func TestArrayList_WithElements(t *testing.T) {
	list := NewArrayListWithElements[int]([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, 6, list.Size())
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, list.ToArray())
}
