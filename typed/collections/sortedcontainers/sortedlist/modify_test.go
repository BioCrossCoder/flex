package sortedlist

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemove(t *testing.T) {
	l := NewSortedList(AscendOrder, 1, 2, 3, 2, 4, 2, 3, 2, 1)
	entry := 2
	count := l.Count(entry)
	length := l.Len()
	finalIndex := l.LastIndexOf(entry)
	convey.Convey("remove the first specified element from list", t, func() {
		l2 := l.Copy()
		assert.Equal(t, finalIndex, l2.LastIndexOf(entry))
		_ = l2.Remove(entry)
		assert.Equal(t, count-1, l2.Count(entry))
		assert.Equal(t, length-1, l2.Len())
		assert.Equal(t, finalIndex-1, l2.LastIndexOf(entry))
	})
	convey.Convey("remove two specified element from list", t, func() {
		l2 := l.Copy()
		assert.Equal(t, finalIndex, l2.LastIndexOf(entry))
		_ = l2.Remove(entry, 2)
		assert.Equal(t, length-2, l2.Len())
		assert.Equal(t, count-2, l2.Count(entry))
		assert.Equal(t, finalIndex-2, l2.LastIndexOf(entry))
	})
	convey.Convey("remove all specified elements from list", t, func() {
		l2 := l.Copy()
		assert.True(t, l2.Includes(entry))
		_ = l2.Remove(entry, -1)
		assert.Equal(t, 0, l2.Count(entry))
		assert.Equal(t, length-count, l2.Len())
		assert.False(t, l2.Includes(entry))
	})
	convey.Convey("remove specified area of list", t, func() {
		l2 := l.Copy()
		removed := l2.RemoveRange(2, 4)
		assert.True(t, removed.Equal(l.Slice(2, 4)))
		left := l.Slice(0, 2).ToList()
		right := l.Slice(4).ToList()
		expected := left.Concat(right)
		assert.Equal(t, expected, l2.ToList())
	})
	convey.Convey("remove all elements from list", t, func() {
		l2 := l.Copy()
		assert.True(t, l2.Len() > 0)
		_ = l2.Clear()
		assert.True(t, l2.Empty())
	})
	convey.Convey("remove elements satisfy the condition from list", t, func() {
		l2 := l.Copy()
		l3 := l2.Copy()
		f := func(x int) bool {
			return x%2 == 0
		}
		removed := l2.RemoveIf(f, 3)
		assert.Equal(t, removed.ToArray(), []int{2, 2, 2})
		assert.Equal(t, l2.ToArray(), []int{1, 1, 2, 3, 3, 4})
		removed = l3.RemoveRightIf(f, 3)
		assert.Equal(t, removed.ToArray(), []int{4, 2, 2})
		assert.Equal(t, l3.ToArray(), []int{1, 1, 2, 2, 3, 3})
	})
}

func TestAddOrCutElement(t *testing.T) {
	l := NewSortedList(AscendOrder, 1, 2, 3, 4, 5)
	length := l.Len()
	convey.Convey("remove one element from list tail", t, func() {
		l2 := l.Copy()
		assert.False(t, l2.Empty())
		expected, err := l2.At(-1)
		assert.Nil(t, err)
		actual, err := l2.Pop()
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
		assert.Equal(t, length-1, l2.Len())
		assert.True(t, l.Slice(0, length-1).Equal(l2))
	})
	convey.Convey("remove one element from list by index", t, func() {
		l2 := l.Copy()
		assert.False(t, l2.Empty())
		expected, err := l2.At(-2)
		assert.Nil(t, err)
		actual, err := l2.Pop(-2)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
		assert.Equal(t, length-1, l2.Len())
		index := l.Len() - 2
		left := l.Slice(0, index).ToList()
		right := l.Slice(index + 1).ToList()
		assert.Equal(t, left.Concat(right), l2.ToList())
	})
	convey.Convey("insert one element to list", t, func() {
		l2 := l.Copy()
		entry := 2
		assert.True(t, l2.Includes(entry))
		count := l2.Count(entry)
		finalIndex := l2.LastIndexOf(2)
		_ = l2.Insert(2)
		assert.Equal(t, finalIndex+1, l2.LastIndexOf(entry))
		assert.Equal(t, count+1, l2.Count(entry))
		assert.Equal(t, length+1, l2.Len())
	})
}

func TestReverse(t *testing.T) {
	convey.Convey("reverse the list", t, func() {
		l := NewSortedList(AscendOrder, 1, 2, 3, 4, 5)
		_ = l.Reverse()
		assert.True(t, l.Equal(*NewSortedList(DescendOrder, 1, 2, 3, 4, 5)))
	})
}
