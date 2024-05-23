package linkedlist

import (
	"fmt"
	"github.com/biocrosscoder/flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	convey.Convey("mapping deque", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		f := func(x int) int {
			return x * 3
		}
		assert.True(t, NewLinkedList(3, 6, 9, 12, 15, 18, 21, 24, 27, 30).Equal(d.Map(f)))
	})
}

func TestReduce(t *testing.T) {
	convey.Convey("reduce deque", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		f := func(x, y int) int {
			return x - y
		}
		convey.Convey("normal reduce", func() {
			result, err := d.Reduce(f)
			assert.Nil(t, err)
			assert.Equal(t, -53, result)
		})
		convey.Convey("reduce with initial value", func() {
			result, err := d.Reduce(f, 100)
			assert.Nil(t, err)
			assert.Equal(t, 45, result)
		})
		convey.Convey("reduce from right", func() {
			result, err := d.ReduceRight(f)
			assert.Nil(t, err)
			assert.Equal(t, -35, result)
		})
		convey.Convey("reduce from right with initial value", func() {
			result, err := d.ReduceRight(f, 100)
			assert.Nil(t, err)
			assert.Equal(t, 45, result)
		})
		convey.Convey("reduce on empty deque", func() {
			d := NewLinkedList[int]()
			result, err := d.Reduce(f)
			assert.Equal(t, common.ErrEmptyList, err)
			assert.Zero(t, result)
			result, err = d.ReduceRight(f)
			assert.Equal(t, common.ErrEmptyList, err)
			assert.Zero(t, result)
		})
		convey.Convey("too many arguments", func() {
			result, err := d.Reduce(f, 1, 2)
			assert.Equal(t, common.ErrTooManyArguments, err)
			assert.Zero(t, result)
			result, err = d.ReduceRight(f, 1, 2)
			assert.Equal(t, common.ErrTooManyArguments, err)
			assert.Zero(t, result)
		})
	})
}

func TestFilter(t *testing.T) {
	convey.Convey("filter deque", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		f := func(x int) bool {
			return x > 5
		}
		assert.True(t, NewLinkedList(6, 7, 8, 9, 10).Equal(d.Filter(f)))
	})
}

func TestSomeAndAny(t *testing.T) {
	convey.Convey("check condition on deque", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		f := func(x int) bool {
			return x > 5
		}
		convey.Convey("at least one element in deque satisfies the condition", func() {
			assert.True(t, d.Some(f))
			assert.False(t, d.Some(func(x int) bool {
				return x > 10
			}))
		})
		convey.Convey("all elements in deque satisfy the condition", func() {
			assert.False(t, d.Every(f))
			assert.True(t, d.Every(func(x int) bool {
				return x > 0
			}))
		})
	})
}

func ExampleLinkedList_Map() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	handler := func(val int) int {
		return val * 2
	}
	newList := list.Map(handler)
	fmt.Println(newList)
	fmt.Println(list)
	// Output:
	// [2 4 6 8 10]
	// [1 2 3 4 5]
}

func ExampleLinkedList_Reduce() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	// Example 1: Summing up the elements of the list
	sum, _ := list.Reduce(func(a, b int) int {
		return a + b
	})
	fmt.Println(sum)
	// Example 2: Finding the maximum element in the list
	max, _ := list.Reduce(func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}, 0)
	fmt.Println(max)
	// Example 3: Summing up the elements of the list with initial value
	sum, _ = list.Reduce(func(a, b int) int {
		return a + b
	}, 10)
	fmt.Println(sum)
	// Output:
	// 15
	// 5
	// 25
}

func ExampleLinkedList_ReduceRight() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	f := func(a, b int) int {
		return a - b
	}
	result1, _ := list.Reduce(f)
	fmt.Println(result1)
	result2, _ := list.ReduceRight(f)
	fmt.Println(result2)
	// Output:
	// -13
	// -5
}

func ExampleLinkedList_Filter() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	condition := func(item int) bool {
		return item%2 == 0
	}
	filteredList := list.Filter(condition)
	fmt.Println(filteredList)
	// Output: [2 4]
}

func ExampleLinkedList_Some() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	result := list.Some(func(i int) bool {
		return i > 3
	})
	fmt.Println(result)
	result = list.Some(func(i int) bool {
		return i > 5
	})
	fmt.Println(result)
	// Output:
	// true
	// false
}

func ExampleLinkedList_Every() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	allEven := list.Every(func(item int) bool {
		return item%2 == 0
	})
	fmt.Println(allEven)
	allPositive := list.Every(func(item int) bool {
		return item > 0
	})
	fmt.Println(allPositive)
	// Output:
	// false
	// true
}
