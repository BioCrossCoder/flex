package linkedlist

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLen(t *testing.T) {
	convey.Convey("check length", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 5)
		assert.Equal(t, 5, d.Len())
	})
}

func TestCount(t *testing.T) {
	convey.Convey("count specific element", t, func() {
		d := NewLinkedList(1, 2, 3, 2, 1, 4, 5, 4)
		assert.Equal(t, 2, d.Count(2))
		assert.Equal(t, 2, d.Count(4))
		assert.Equal(t, 0, d.Count(6))
	})
}

func TestIncludes(t *testing.T) {
	convey.Convey("check if deque includes element", t, func() {
		d := NewLinkedList(1, 2, 3, 2, 1, 4, 5, 4)
		assert.True(t, d.Includes(2))
		assert.False(t, d.Includes(6))
	})
}

func TestEmpty(t *testing.T) {
	convey.Convey("check if deque is empty", t, func() {
		d1 := NewLinkedList[any]()
		d2 := NewLinkedList(1, 2, 3, 4, 5)
		assert.True(t, d1.Empty())
		assert.False(t, d2.Empty())
	})
}

func TestToList(t *testing.T) {
	convey.Convey("convert deque to list", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 5)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, d.ToArray())
	})
}

func ExampleLinkedList_Len() {
	ll := NewLinkedList(1, 2, 3, 4, 5)
	fmt.Println(ll.Len())
	// Output: 5
}

func ExampleLinkedList_Count() {
	ll := NewLinkedList(1, 2, 3, 4, 5, 5, 5)
	fmt.Println(ll.Count(5))
	fmt.Println(ll.Count(6))
	// Output:
	// 3
	// 0
}

func ExampleLinkedList_Includes() {
	ll := NewLinkedList(1, 2, 3, 4, 5)
	fmt.Println(ll.Includes(3))
	fmt.Println(ll.Includes(6))
	// Output:
	// true
	// false
}

func ExampleLinkedList_Empty() {
	ll := NewLinkedList[int]()
	fmt.Println(ll.Empty())
	ll2 := NewLinkedList(1, 2, 3, 4, 5)
	fmt.Println(ll2.Empty())
	// Output:
	// true
	// false
}

func ExampleLinkedList_Equal() {
	ll1 := NewLinkedList(1, 2, 3)
	ll2 := NewLinkedList(1, 2, 3)
	ll3 := NewLinkedList(1, 2, 3, 4)
	ll4 := NewLinkedList(1, 2, 4)
	fmt.Println(ll1.Equal(*ll2))
	fmt.Println(ll1.Equal(*ll3))
	fmt.Println(ll1.Equal(*ll4))
	// Output:
	// true
	// false
	// false
}
