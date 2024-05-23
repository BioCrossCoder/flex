package linkedlist

import (
	"fmt"
	"github.com/biocrosscoder/flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchElement(t *testing.T) {
	convey.Convey("search element in deque", t, func() {
		d := NewLinkedList(1, 2, 3, 2, 4)
		convey.Convey("search by index", func() {
			assert.Equal(t, 1, d.IndexOf(2))
			assert.Equal(t, 3, d.LastIndexOf(2))
		})
		convey.Convey("get element by index", func() {
			num, err := d.At(-3)
			assert.Nil(t, err)
			assert.Equal(t, 3, num)
		})
		convey.Convey("search by condition", func() {
			f := func(num int) bool {
				return num%2 == 0
			}
			v, found := d.Find(f)
			assert.True(t, found)
			assert.Equal(t, 2, v)
			v, found = d.FindLast(f)
			assert.True(t, found)
			assert.Equal(t, 4, v)
			assert.Equal(t, 1, d.FindIndex(f))
			assert.Equal(t, 4, d.FindLastIndex(f))
			assert.Equal(t, []int{1, 3, 4}, d.FindIndexes(f))
			assert.Equal(t, []int{4, 3, 1}, d.FindLastIndexes(f))
			assert.Equal(t, []int{2, 2, 4}, d.Finds(f))
			assert.Equal(t, []int{4, 2, 2}, d.FindLasts(f))
			assert.Equal(t, []int{1, 3}, d.FindIndexes(f, 2))

		})
		convey.Convey("get first/last element", func() {
			head, err := d.Head()
			assert.Nil(t, err)
			assert.Equal(t, 1, head)
			tail, err := d.Tail()
			assert.Nil(t, err)
			assert.Equal(t, 4, tail)
		})
		convey.Convey("element not found", func() {
			assert.Equal(t, -1, d.IndexOf(5))
			assert.Equal(t, -1, d.LastIndexOf(5))
			f := func(num int) bool {
				return num > 5
			}
			v, found := d.Find(f)
			assert.False(t, found)
			assert.Zero(t, v)
			v, found = d.FindLast(f)
			assert.False(t, found)
			assert.Zero(t, v)
			assert.Equal(t, -1, d.FindIndex(f))
			assert.Equal(t, -1, d.FindLastIndex(f))
		})
		convey.Convey("empty deque", func() {
			d := NewLinkedList[any]()
			head, err := d.Head()
			assert.Equal(t, err, common.ErrEmptyList)
			assert.Nil(t, head)
			tail, err := d.Tail()
			assert.Equal(t, err, common.ErrEmptyList)
			assert.Nil(t, tail)
		})
	})
}

func ExampleLinkedList_IndexOf() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	fmt.Println(list.IndexOf(3))
	fmt.Println(list.IndexOf(6))
	// Output:
	// 2
	// -1
}

func ExampleLinkedList_LastIndexOf() {
	list := NewLinkedList(1, 2, 3, 4, 3)
	fmt.Println(list.IndexOf(3))
	fmt.Println(list.LastIndexOf(3))
	// Output:
	// 2
	// 4
}

func ExampleLinkedList_At() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	element, _ := list.At(2)
	fmt.Println(element)
	element, _ = list.At(-1)
	fmt.Println(element)
	_, err := list.At(10)
	fmt.Println(err)
	// Output:
	// 3
	// 5
	// the index is out of range
}

func ExampleLinkedList_Find() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	even := func(num int) bool {
		return num%2 == 0
	}
	fmt.Println(list.Find(even))
	negative := func(num int) bool {
		return num < 0
	}
	fmt.Println(list.Find(negative))
	// Output:
	// 2 true
	// 0 false
}

func ExampleLinkedList_FindIndex() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	even := func(num int) bool {
		return num%2 == 0
	}
	fmt.Println(list.FindIndex(even))
	negative := func(num int) bool {
		return num < 0
	}
	fmt.Println(list.FindIndex(negative))
	// Output:
	// 1
	// -1
}

func ExampleLinkedList_FindLast() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	even := func(num int) bool {
		return num%2 == 0
	}
	fmt.Println(list.Find(even))
	fmt.Println(list.FindLast(even))
	// Output:
	// 2 true
	// 4 true
}

func ExampleLinkedList_FindLastIndex() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	even := func(num int) bool {
		return num%2 == 0
	}
	fmt.Println(list.FindIndex(even))
	fmt.Println(list.FindLastIndex(even))
	// Output:
	// 1
	// 3
}

func ExampleLinkedList_Head() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	fmt.Println(list.Head())
	fmt.Println(NewLinkedList[int]().Head())
	// Output:
	// 1 <nil>
	// 0 the input list is empty
}

func ExampleLinkedList_Tail() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	fmt.Println(list.Tail())
	fmt.Println(NewLinkedList[int]().Tail())
	// Output:
	// 5 <nil>
	// 0 the input list is empty
}

func ExampleLinkedList_FindIndexes() {
	l := NewLinkedList(1, 2, 3, 4, 5, 2)
	condition := func(val int) bool {
		return val > 2
	}
	fmt.Println(l.FindIndexes(condition))
	fmt.Println(l.FindIndexes(condition, 2))
	// Output:
	// [2 3 4]
	// [2 3]
}

func ExampleLinkedList_FindLastIndexes() {
	l := NewLinkedList(1, 2, 3, 4, 5, 2)
	condition := func(val int) bool {
		return val > 2
	}
	fmt.Println(l.FindIndexes(condition))
	fmt.Println(l.FindLastIndexes(condition))
	// Output:
	// [2 3 4]
	// [4 3 2]
}

func ExampleLinkedList_Finds() {
	l := NewLinkedList(1, 2, 3, 4, 5, 2)
	condition := func(val int) bool {
		return val > 2
	}
	fmt.Println(l.Finds(condition))
	fmt.Println(l.Finds(condition, 2))
	// Output:
	// [3 4 5]
	// [3 4]
}

func ExampleLinkedList_FindLasts() {
	l := NewLinkedList(1, 2, 3, 4, 5, 2)
	condition := func(val int) bool {
		return val > 2
	}
	fmt.Println(l.Finds(condition))
	fmt.Println(l.FindLasts(condition))
	// Output:
	// [3 4 5]
	// [5 4 3]
}
