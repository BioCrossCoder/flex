package arraylist

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemove(t *testing.T) {
	l := ArrayList[int]{1, 2, 3, 2, 4, 2, 3, 2, 1}
	entry := l[1]
	count := l.Count(entry)
	length := l.Len()
	convey.Convey("remove the first specified element from list", t, func() {
		l2 := l.Copy()
		assert.Equal(t, 1, l2.IndexOf(entry))
		_ = l2.Remove(entry)
		assert.Equal(t, count-1, l2.Count(entry))
		assert.Equal(t, length-1, l2.Len())
		assert.Equal(t, 2, l2.IndexOf(entry))
	})
	convey.Convey("remove two specified element from list", t, func() {
		l2 := l.Copy()
		_ = l2.Remove(entry, 2)
		assert.Equal(t, length-2, l2.Len())
		assert.Equal(t, count-2, l2.Count(entry))
	})
	convey.Convey("remove all specified elements from list", t, func() {
		l2 := l.Copy()
		_ = l2.Remove(entry, -1)
		assert.Equal(t, 0, l2.Count(entry))
		assert.Equal(t, length-count, l2.Len())
	})
	convey.Convey("remove the last specified element from list", t, func() {
		l2 := l.Copy()
		assert.Equal(t, 7, l2.LastIndexOf(entry))
		_ = l2.RemoveRight(entry, 2)
		assert.Equal(t, count-2, l2.Count(entry))
		assert.Equal(t, length-2, l2.Len())
		assert.Equal(t, 3, l2.LastIndexOf(entry))
	})
	convey.Convey("remove all elements from list", t, func() {
		l2 := l.Copy()
		assert.True(t, l2.Len() > 0)
		_ = l2.Clear()
		assert.True(t, l2.Empty())
	})
	convey.Convey("remove elements satisfy the condition from list", t, func() {
		l2 := l.Copy().Concat(Of(6, 8))
		l3 := l2.Copy()
		f := func(x int) bool {
			return x%2 == 0
		}
		removed := l2.RemoveIf(f, 3)
		assert.Equal(t, removed, ArrayList[int]{2, 2, 4})
		assert.Equal(t, l2, ArrayList[int]{1, 3, 2, 3, 2, 1, 6, 8})
		removed = l3.RemoveRightIf(f, 3)
		assert.Equal(t, removed, ArrayList[int]{8, 6, 2})
		assert.Equal(t, l3, ArrayList[int]{1, 2, 3, 2, 4, 2, 3, 1})
	})

}

func TestAddOrCutElement(t *testing.T) {
	l := ArrayList[int]{1, 2, 3, 4, 5}
	length := l.Len()
	convey.Convey("add elements to list tail", t, func() {
		l2 := l.Copy()
		_ = l2.Push(1, 2, 3)
		assert.Equal(t, length+3, l2.Len())
		assert.Equal(t, l2[l2.Len()-3:], Of(1, 2, 3))
	})
	convey.Convey("add elements to list head", t, func() {
		l2 := l.Copy()
		s := ArrayList[int]{1, 2, 3}
		_ = l2.Unshift(s...)
		assert.Equal(t, length+s.Len(), l2.Len())
		assert.Equal(t, l2[:3], s)
	})
	convey.Convey("remove one element from list tail", t, func() {
		l2 := l.Copy()
		assert.False(t, l2.Empty())
		expected, err := l2.Tail()
		assert.Nil(t, err)
		actual, err := l2.Pop()
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
		assert.Equal(t, length-1, l2.Len())
		assert.Equal(t, l[:length-1], l2)
	})
	convey.Convey("remove one element from list head", t, func() {
		l2 := l.Copy()
		assert.False(t, l2.Empty())
		expected, err := l2.Head()
		assert.Nil(t, err)
		actual, err := l2.Shift()
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
		assert.Equal(t, length-1, l2.Len())
		assert.Equal(t, l[1:], l2)
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
		assert.Equal(t, l[:index].Concat(l[index+1:]), l2)
	})
	convey.Convey("add one element to list by index", t, func() {
		l2 := l.Copy()
		_ = l2.Insert(2, 6)
		assert.Equal(t, l2[2], 6)
		assert.Equal(t, length+1, l2.Len())
		assert.Equal(t, l[:2].Concat(Of(6)).Concat(l[2:]), l2)
	})
}

func TestForEach(t *testing.T) {
	convey.Convey("convert list elements", t, func() {
		l := ArrayList[int]{1, 2, 3, 4, 5}
		f := func(x int) int {
			return -x
		}
		_ = l.ForEach(f)
		assert.Equal(t, ArrayList[int]{-1, -2, -3, -4, -5}, l)
	})
}

func TestReplace(t *testing.T) {
	l := Repeat(1, 5)
	convey.Convey("replace the first specified element with another element", t, func() {
		l2 := l.Copy()
		assert.Equal(t, l2.Count(1), 5)
		assert.Equal(t, l2.IndexOf(1), 0)
		_ = l2.Replace(1, 2)
		assert.Equal(t, l2.Count(1), 4)
		assert.Equal(t, l2.IndexOf(1), 1)
	})
	convey.Convey("replace two specified elements with another element", t, func() {
		l2 := l.Copy()
		assert.Equal(t, l2.Count(1), 5)
		assert.Equal(t, l2.IndexOf(1), 0)
		_ = l2.Replace(1, 2, 2)
		assert.Equal(t, l2.Count(1), 3)
		assert.Equal(t, l2.IndexOf(1), 2)
	})
	convey.Convey("replace all specified elements with another element", t, func() {
		l2 := l.Copy()
		assert.Equal(t, l2.Count(1), 5)
		assert.Equal(t, l2.IndexOf(1), 0)
		_ = l2.Replace(1, 2, -1)
		assert.Equal(t, l2.Count(1), 0)
		assert.False(t, l2.Includes(1))
	})
	convey.Convey("replace the last specified element with another element", t, func() {
		l2 := l.Copy()
		assert.Equal(t, l2.Count(1), 5)
		assert.Equal(t, l2.LastIndexOf(1), 4)
		_ = l2.ReplaceRight(1, 2)
		assert.Equal(t, l2.Count(1), 4)
		assert.Equal(t, l2.LastIndexOf(1), 3)
	})
	convey.Convey("replece elements satisfy the condition with another element", t, func() {
		l2 := l.Copy().Concat(Of(3, 5, 7))
		l3 := l2.Copy()
		f := func(x int) bool {
			return x%2 == 1
		}
		replaced := l2.ReplaceIf(f, 0, 3)
		assert.Equal(t, replaced, Repeat(1, 3))
		assert.Equal(t, l2, ArrayList[int]{0, 0, 0, 1, 1, 3, 5, 7})
		replaced = l3.ReplaceRightIf(f, 0, 3)
		assert.Equal(t, replaced, ArrayList[int]{7, 5, 3})
		assert.Equal(t, l3, ArrayList[int]{1, 1, 1, 1, 1, 0, 0, 0})
	})
}

func TestSplice(t *testing.T) {
	l := ArrayList[int]{1, 2, 3, 4, 5}
	convey.Convey("remove elements from list by area", t, func() {
		l2 := l.Copy()
		_ = l2.Splice(1, 3)
		assert.Equal(t, l2, ArrayList[int]{1, 5})
	})
	convey.Convey("remove elements from list by area and insert new elements", t, func() {
		l2 := l.Copy()
		_ = l2.Splice(1, 2, 6, 0)
		assert.Equal(t, l2, ArrayList[int]{1, 6, 0, 4, 5})
	})
}

func TestFill(t *testing.T) {
	l := make(ArrayList[int], 5)
	convey.Convey("fill the list with specified element", t, func() {
		l2 := l.Copy()
		l2.Fill(6)
		assert.Equal(t, l2, Repeat(6, 5))
	})
	convey.Convey("fill the list with specified element from start index", t, func() {
		l2 := l.Copy()
		l2.Fill(6, 1)
		assert.Equal(t, l2, Repeat(6, 5).With(0, 0))
	})
	convey.Convey("fill the list with specified element from start index to end index", t, func() {
		l2 := l.Copy()
		l2.Fill(6, 1, 3)
		assert.Equal(t, l2, ArrayList[int]{0, 6, 6, 0, 0})
	})
}

func TestReverse(t *testing.T) {
	convey.Convey("reverse the list", t, func() {
		l := ArrayList[int]{1, 2, 3, 4, 5}
		_ = l.Reverse()
		assert.Equal(t, l, ArrayList[int]{5, 4, 3, 2, 1})
	})
}

func TestSet(t *testing.T) {
	convey.Convey("set the element of list by index", t, func() {
		l := ArrayList[int]{1, 2, 3, 4, 5}
		assert.Nil(t, l.Set(-2, 6))
		result, err := l.At(-2)
		assert.Nil(t, err)
		assert.Equal(t, result, 6)
	})
}
