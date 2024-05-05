package linkedlist

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopy(t *testing.T) {
	convey.Convey("copy deque", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 5)
		d2 := d.Copy()
		assert.True(t, d.Equal(d2))
	})
}

func TestConcat(t *testing.T) {
	convey.Convey("link two deques", t, func() {
		d1 := NewLinkedList(1, 2, 3)
		d2 := NewLinkedList(4, 5, 6)
		assert.Equal(t, d1.Concat(*d2), *NewLinkedList(1, 2, 3, 4, 5, 6))
		assert.Equal(t, d2.Concat(*d1), *NewLinkedList(4, 5, 6, 1, 2, 3))
	})
}

func TestSlice(t *testing.T) {
	d := NewLinkedList(1, 3, 5, 7, 9)
	convey.Convey("slice without args", t, func() {
		assert.True(t, d.Equal(d.Slice()))
	})
	convey.Convey("slice with start index", t, func() {
		assert.True(t, NewLinkedList(5, 7, 9).Equal(d.Slice(2)))
	})
	convey.Convey("slice with start and end index", t, func() {
		assert.True(t, NewLinkedList(5, 7).Equal(d.Slice(2, 4)))
	})
	convey.Convey("slice with start index, end index and step", t, func() {
		assert.True(t, NewLinkedList(1, 5, 9).Equal(d.Slice(0, 6, 2)))
		assert.True(t, NewLinkedList(9, 5).Equal(d.Slice(-1, -5, -2)))
	})
}

func TestToSpliced(t *testing.T) {
	convey.Convey("to spliced deque", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 9)
		assert.True(t, NewLinkedList(1, 2, 4, 6, 9).Equal(d.ToSpliced(1, 3, 2, 4, 6)))
	})
}

func TestToReversed(t *testing.T) {
	convey.Convey("to reversed deque", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 5)
		assert.True(t, NewLinkedList(5, 4, 3, 2, 1).Equal(d.ToReversed()))
	})
}

func TestWith(t *testing.T) {
	convey.Convey("deque with a replaced element", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 5)
		assert.True(t, NewLinkedList(1, 2, 6, 4, 5).Equal(d.With(2, 6)))
	})
}
