package linkedlist

import (
	"fmt"
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
		expected := NewLinkedList(5, 7, 9)
		assert.True(t, expected.Equal(d.Slice(2)))
		assert.True(t, expected.Equal(d.Slice(-3)))
	})
	convey.Convey("slice with start and end index", t, func() {
		expected := NewLinkedList(5, 7)
		assert.True(t, expected.Equal(d.Slice(2, 4)))
		assert.True(t, expected.Equal(d.Slice(2, -1)))
		assert.True(t, expected.Equal(d.Slice(-3, 4)))
		assert.True(t, expected.Equal(d.Slice(-3, -1)))
	})
	convey.Convey("slice with start index, end index and step", t, func() {
		expected := NewLinkedList(1, 5, 9)
		assert.True(t, expected.Equal(d.Slice(0, 10, 2)))
		assert.True(t, expected.Equal(d.Slice(-6, 5, 2)))
		assert.True(t, expected.ToReversed().Equal(d.Slice(-1, -10, -2)))
		assert.True(t, expected.ToReversed().Equal(d.Slice(4, -6, -2)))
		assert.True(t, NewLinkedList().Equal(d.Slice(-1, 2, 1)))
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

func ExampleLinkedList_Copy() {
	l := NewLinkedList(1, 2, 3, 4, 5)
	backup := l.Copy()
	fmt.Println(backup)
	fmt.Println(l.Equal(backup))
	// Output:
	// [1 2 3 4 5]
	// true
}

func ExampleLinkedList_Concat() {
	l1 := NewLinkedList(1, 2, 3)
	l2 := NewLinkedList(4, 5, 6)
	concatenated := l1.Concat(*l2)
	fmt.Println(concatenated)
	// Output: [1 2 3 4 5 6]
}

func ExampleLinkedList_Slice() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	slice1 := list.Slice(1, 4)
	fmt.Println(slice1)
	slice2 := list.Slice(0, 3, 2)
	fmt.Println(slice2)
	slice3 := list.Slice(3, 0, -1)
	fmt.Println(slice3)
	// Output:
	// [2 3 4]
	// [1 3]
	// [4 3 2]
}

func ExampleLinkedList_ToSpliced() {
	arr := NewLinkedList(1, 2, 3, 4, 5)
	newArr := arr.ToSpliced(2, 2, 6, 7, 8)
	fmt.Println(newArr)
	fmt.Println(arr)
	// Output:
	// [1 2 6 7 8 5]
	// [1 2 3 4 5]
}

func ExampleLinkedList_ToReversed() {
	l := NewLinkedList(1, 2, 3, 4, 5)
	reversed := l.ToReversed()
	fmt.Println(reversed)
	fmt.Println(l)
	// Output:
	// [5 4 3 2 1]
	// [1 2 3 4 5]
}

func ExampleLinkedList_With() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	newList := list.With(2, 10)
	fmt.Println(newList)
	fmt.Println(list)
	// Output:
	// [1 2 10 4 5]
	// [1 2 3 4 5]
}
