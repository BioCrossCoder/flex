package list

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemove(t *testing.T) {
	l := List{1, 2, 3, 2, 4, 2, 3, 2, 1}
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
}

func TestAddOrCutElement(t *testing.T) {
	l := List{1, 2, 3, 4, 5}
	length := l.Len()
	convey.Convey("add elements to list tail", t, func() {
		l2 := l.Copy()
		_ = l2.Push(1, 2, 3)
		assert.Equal(t, length+3, l2.Len())
		assert.Equal(t, l2[l2.Len()-3:], Of(1, 2, 3))
	})
	convey.Convey("add elements to list head", t, func() {
		l2 := l.Copy()
		s := List{1, 2, 3}
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
		l := List{1, 2, 3, 4, 5}
		f := func(x any) any {
			return -x.(int)
		}
		_ = l.ForEach(f)
		assert.Equal(t, List{-1, -2, -3, -4, -5}, l)
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
}

func TestSplice(t *testing.T) {
	l := List{1, 2, 3, 4, 5}
	convey.Convey("remove elements from list by area", t, func() {
		l2 := l.Copy()
		_ = l2.Splice(1, 3)
		assert.Equal(t, l2, List{1, 5})
	})
	convey.Convey("remove elements from list by area and insert new elements", t, func() {
		l2 := l.Copy()
		_ = l2.Splice(1, 2, 6, 0)
		assert.Equal(t, l2, List{1, 6, 0, 4, 5})
	})
}

func TestFill(t *testing.T) {
	l := make(List, 5)
	convey.Convey("fill the list with specified element", t, func() {
		l2 := l.Copy()
		l2.Fill(6)
		assert.Equal(t, l2, Repeat(6, 5))
	})
	convey.Convey("fill the list with specified element from start index", t, func() {
		l2 := l.Copy()
		l2.Fill(6, 1)
		assert.Equal(t, l2, Repeat(6, 5).With(0, nil))
	})
	convey.Convey("fill the list with specified element from start index to end index", t, func() {
		l2 := l.Copy()
		l2.Fill(6, 1, 3)
		assert.Equal(t, l2, List{nil, 6, 6, nil, nil})
	})
}

func TestReverse(t *testing.T) {
	convey.Convey("reverse the list", t, func() {
		l := List{1, 2, 3, 4, 5}
		_ = l.Reverse()
		assert.Equal(t, l, List{5, 4, 3, 2, 1})
	})
}
