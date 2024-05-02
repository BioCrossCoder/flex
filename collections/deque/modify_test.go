package deque

import (
	"flex/collections/list"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemove(t *testing.T) {
	d := NewDeque(1, 2, 3, 2, 4, 2, 3, 2, 1)
	entry, err := d.At(1)
	assert.Nil(t, err)
	count := d.Count(entry)
	length := d.Len()
	convey.Convey("remove the first specified element from deque", t, func() {
		d2 := d.Copy()
		assert.Equal(t, 1, d2.IndexOf(entry))
		_ = d2.Remove(entry)
		assert.Equal(t, count-1, d2.Count(entry))
		assert.Equal(t, length-1, d2.Len())
		assert.Equal(t, 2, d2.IndexOf(entry))
	})
	convey.Convey("remove two specified element from deque", t, func() {
		d2 := d.Copy()
		_ = d2.Remove(entry, 2)
		assert.Equal(t, count-2, d2.Count(entry))
		assert.Equal(t, length-2, d2.Len())
	})
	convey.Convey("remove all specified element from deque", t, func() {
		d2 := d.Copy()
		_ = d2.Remove(entry, -1)
		assert.Equal(t, 0, d2.Count(entry))
		assert.Equal(t, length-count, d2.Len())
	})
	convey.Convey("remove the last specified element from deque", t, func() {
		d2 := d.Copy()
		assert.Equal(t, 7, d2.LastIndexOf(entry))
		_ = d2.RemoveRight(entry, 2)
		assert.Equal(t, count-2, d2.Count(entry))
		assert.Equal(t, length-2, d2.Len())
		assert.Equal(t, 3, d2.LastIndexOf(entry))
	})
	convey.Convey("remove all elements from deque", t, func() {
		d2 := d.Copy()
		assert.False(t, d2.Empty())
		_ = d2.Clear()
		assert.True(t, d2.Empty())
	})
}

func TestAddOrCutElement(t *testing.T) {
	d := NewDeque(1, 2, 3, 4, 5)
	length := d.Len()
	convey.Convey("add one element to deque tail", t, func() {
		d2 := d.Copy()
		_ = d2.Append(6)
		assert.Equal(t, length+1, d2.Len())
		lastElement, err := d2.Tail()
		assert.Nil(t, err)
		assert.Equal(t, 6, lastElement)
	})
	convey.Convey("add one element to deque head", t, func() {
		d2 := d.Copy()
		_ = d2.AppendLeft(7)
		assert.Equal(t, length+1, d2.Len())
		firstElement, err := d2.Head()
		assert.Nil(t, err)
		assert.Equal(t, 7, firstElement)
	})
	convey.Convey("remove one element from deque tail", t, func() {
		d2 := d.Copy()
		lastElement, err := d2.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 5, lastElement)
		assert.Equal(t, length-1, d2.Len())
	})
	convey.Convey("remove one element from deque head", t, func() {
		d2 := d.Copy()
		firstElement, err := d2.PopLeft()
		assert.Nil(t, err)
		assert.Equal(t, 1, firstElement)
		assert.Equal(t, length-1, d2.Len())
	})
	convey.Convey("remove one element from deque by index", t, func() {
		d2 := d.Copy()
		value, err := d2.RemoveByIndex(1)
		assert.Nil(t, err)
		assert.Equal(t, 2, value)
		assert.Equal(t, length-1, d2.Len())
	})
	convey.Convey("add one element to deque by index", t, func() {
		d2 := d.Copy()
		_ = d2.Insert(2, 9)
		assert.Equal(t, length+1, d2.Len())
		value, err := d2.At(2)
		assert.Nil(t, err)
		assert.Equal(t, 9, value)
	})
}

func TestExtend(t *testing.T) {
	d := NewDeque(1, 2, 3, 4, 5)
	length := d.Len()
	convey.Convey("extend deque tail with another deque", t, func() {
		d2 := d.Copy()
		d3 := NewDeque(6, 7, 8)
		_ = d2.Extend(d3)
		assert.Equal(t, length+3, d2.Len())
		assert.True(t, d2.Equal(d.Concat(*d3)))
	})
	convey.Convey("extend deque head with another deque", t, func() {
		d2 := d.Copy()
		d3 := NewDeque(0, -1, -2)
		_ = d2.ExtendLeft(d3)
		assert.Equal(t, length+3, d2.Len())
		assert.True(t, d2.Equal(d3.Reverse().Concat(*d)))
	})
}

func TestRotate(t *testing.T) {
	d := NewDeque(1, 2, 3, 4, 5)
	length := d.Len()
	convey.Convey("rotate deque to the left", t, func() {
		steps := 2
		d2 := d.Copy()
		_ = d2.Rotate(steps)
		assert.Equal(t, length, d2.Len())
		assert.True(t, NewDeque(4, 5, 1, 2, 3).Equal(d2))
	})
	convey.Convey("rotate deque to the right", t, func() {
		steps := -2
		d2 := d.Copy()
		_ = d2.Rotate(steps)
		assert.Equal(t, length, d2.Len())
		assert.True(t, NewDeque(3, 4, 5, 1, 2).Equal(d2))
	})
}

func TestForEach(t *testing.T) {
	convey.Convey("convert deque elements", t, func() {
		d := NewDeque(1, 2, 3, 4, 5)
		f := func(x any) any {
			return x.(int) * 2
		}
		_ = d.ForEach(f)
		assert.True(t, NewDeque(2, 4, 6, 8, 10).Equal(*d))
	})
}

func TestReplace(t *testing.T) {
	d := NewDeque(list.Repeat(1, 10)...)
	convey.Convey("replace the first specified element with another element", t, func() {
		d2 := d.Copy()
		assert.Equal(t, 10, d2.Count(1))
		assert.Equal(t, 0, d2.IndexOf(1))
		_ = d2.Replace(1, 9)
		value, err := d2.Head()
		assert.Nil(t, err)
		assert.Equal(t, 9, value)
		assert.Equal(t, 9, d2.Count(1))
		assert.Equal(t, 1, d2.IndexOf(1))
	})
	convey.Convey("replace the last specified element with another element", t, func() {
		d2 := d.Copy()
		assert.Equal(t, 10, d2.Count(1))
		assert.Equal(t, 9, d2.LastIndexOf(1))
		_ = d2.ReplaceRight(1, 9)
		value, err := d2.Tail()
		assert.Nil(t, err)
		assert.Equal(t, 9, value)
		assert.Equal(t, 9, d2.Count(1))
		assert.Equal(t, 8, d2.LastIndexOf(1))
	})
	convey.Convey("replace two specified elements with another element", t, func() {
		d2 := d.Copy()
		assert.Equal(t, 10, d2.Count(1))
		assert.Equal(t, 0, d2.IndexOf(1))
		_ = d2.Replace(1, 9, 2)
		assert.Equal(t, 8, d2.Count(1))
		assert.Equal(t, 2, d2.IndexOf(1))
	})
	convey.Convey("replace all specified elements with another element", t, func() {
		d2 := d.Copy()
		assert.Equal(t, 10, d2.Count(1))
		assert.Equal(t, 0, d2.IndexOf(1))
		_ = d2.Replace(1, 9, -1)
		assert.Equal(t, 0, d2.Count(1))
		assert.Equal(t, -1, d2.IndexOf(1))
		assert.False(t, d2.Includes(1))
	})
}

func TestSplice(t *testing.T) {
	d := NewDeque(1, 2, 3, 4, 5)
	convey.Convey("remove elements from deque by area", t, func() {
		d2 := d.Copy()
		_ = d2.Splice(1, 4)
		assert.True(t, NewDeque(1).Equal(d2))
	})
	convey.Convey("remove elements from deque by area and insert new elements", t, func() {
		d2 := d.Copy()
		_ = d2.Splice(1, 2, 6, 7)
		assert.True(t, NewDeque(1, 6, 7, 4, 5).Equal(d2))
	})
}

func TestFill(t *testing.T) {
	d := NewDeque(list.Repeat(1, 5)...)
	convey.Convey("fill the deque with the specified element", t, func() {
		d2 := d.Copy()
		_ = d2.Fill(2)
		assert.True(t, NewDeque(2, 2, 2, 2, 2).Equal(d2))
	})
	convey.Convey("fill the deque with the specified element and start index", t, func() {
		d2 := d.Copy()
		_ = d2.Fill(2, 2)
		assert.True(t, NewDeque(1, 1, 2, 2, 2).Equal(d2))
	})
	convey.Convey("fill the deque with the specified element from start index to end index", t, func() {
		d2 := d.Copy()
		_ = d2.Fill(2, 1, 3)
		assert.True(t, NewDeque(1, 2, 2, 1, 1).Equal(d2))
	})
}

func TestReverse(t *testing.T) {
	convey.Convey("reverse the deque", t, func() {
		d := NewDeque(1, 2, 3, 4, 5)
		_ = d.Reverse()
		assert.True(t, NewDeque(5, 4, 3, 2, 1).Equal(*d))
	})
}

func TestSet(t *testing.T) {
	convey.Convey("set the element of deque by index", t, func() {
		d := NewDeque(1, 2, 3, 4, 5)
		assert.Nil(t, d.Set(-2, 9))
		result, err := d.At(-2)
		assert.Nil(t, err)
		assert.Equal(t, 9, result)
	})
}
