package arraylist

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopy(t *testing.T) {
	convey.Convey("copy list", t, func() {
		l := ArrayList[int]{1, 3, 5, 7, 9}
		l2 := l.Copy()
		assert.Equal(t, l[:], l2)
	})
}

func TestConcat(t *testing.T) {
	convey.Convey("link two lists", t, func() {
		l1 := ArrayList[int]{1, 3, 5, 7, 9}
		l2 := ArrayList[int]{2, 4, 6, 8, 10}
		assert.Equal(t, ArrayList[int]{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}, l1.Concat(l2))
		assert.Equal(t, ArrayList[int]{2, 4, 6, 8, 10, 1, 3, 5, 7, 9}, l2.Concat(l1))
	})
}

func TestSlice(t *testing.T) {
	l := ArrayList[int]{1, 3, 5, 7, 9}
	convey.Convey("slice without args", t, func() {
		assert.Equal(t, l, l.Slice())
	})
	convey.Convey("slice with start index", t, func() {
		expected := ArrayList[int]{5, 7, 9}
		assert.Equal(t, expected, l.Slice(2))
		assert.Equal(t, expected, l.Slice(-3))
	})
	convey.Convey("slice with start and end index", t, func() {
		expected := ArrayList[int]{5, 7}
		assert.Equal(t, expected, l.Slice(2, 4))
		assert.Equal(t, expected, l.Slice(2, -1))
		assert.Equal(t, expected, l.Slice(-3, 4))
		assert.Equal(t, expected, l.Slice(-3, -1))
	})
	convey.Convey("slice with start index, end index and step", t, func() {
		expected := ArrayList[int]{1, 5, 9}
		assert.Equal(t, expected, l.Slice(0, 10, 2))
		assert.Equal(t, expected, l.Slice(-6, 5, 2))
		assert.Equal(t, expected.ToReversed(), l.Slice(-1, -10, -2))
		assert.Equal(t, expected.ToReversed(), l.Slice(4, -6, -2))
		assert.Equal(t, ArrayList[int]{}, l.Slice(-1, 2, 1))
	})
}

func TestToSpliced(t *testing.T) {
	convey.Convey("to spliced list", t, func() {
		l := ArrayList[int]{1, 3, 5, 7, 9}
		assert.Equal(t, ArrayList[int]{1, 2, 4, 6, 9}, l.ToSpliced(1, 3, 2, 4, 6))
	})
}

func TestToReversed(t *testing.T) {
	convey.Convey("to reversed list", t, func() {
		l := ArrayList[int]{1, 3, 5, 7, 9}
		assert.Equal(t, ArrayList[int]{9, 7, 5, 3, 1}, l.ToReversed())
	})
}

func TestWith(t *testing.T) {
	convey.Convey("list with a replaced element", t, func() {
		l := ArrayList[int]{1, 3, 5, 7, 9}
		assert.Equal(t, ArrayList[int]{1, 3, 5, 7, 10}, l.With(l.Len()-1, 10))
	})
}

func TestOf(t *testing.T) {
	convey.Convey("pack input arguments to a list", t, func() {
		assert.Equal(t, ArrayList[int]{1, 2, 3, 4, 5}, Of(1, 2, 3, 4, 5))
	})
}

func TestRepeat(t *testing.T) {
	convey.Convey("create a list with repeated elements", t, func() {
		assert.Equal(t, ArrayList[int]{1, 1, 1, 1, 1}, Repeat(1, 5))
	})
}
