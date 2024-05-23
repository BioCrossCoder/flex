package linkedlist

import (
	"fmt"
	"github.com/biocrosscoder/flex/collections/arraylist"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemove(t *testing.T) {
	d := NewLinkedList(1, 2, 3, 2, 4, 2, 3, 2, 1)
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
	convey.Convey("remove elements satisfy the condition from list", t, func() {
		d2 := d.Copy().Concat(*NewLinkedList(6, 8))
		d3 := d2.Copy()
		f := func(x any) bool {
			return x.(int)%2 == 0
		}
		removed := d2.RemoveIf(f, 3)
		assert.True(t, NewLinkedList(2, 2, 4).Equal(removed))
		assert.True(t, NewLinkedList(1, 3, 2, 3, 2, 1, 6, 8).Equal(d2))
		removed = d3.RemoveRightIf(f, 3)
		assert.True(t, NewLinkedList(8, 6, 2).Equal(removed))
		assert.True(t, NewLinkedList(1, 2, 3, 2, 4, 2, 3, 1).Equal(d3))
	})
}

func TestAddOrCutElement(t *testing.T) {
	d := NewLinkedList(1, 2, 3, 4, 5)
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
	d := NewLinkedList(1, 2, 3, 4, 5)
	length := d.Len()
	convey.Convey("extend deque tail with another deque", t, func() {
		d2 := d.Copy()
		d3 := NewLinkedList(6, 7, 8)
		_ = d2.Extend(d3)
		assert.Equal(t, length+3, d2.Len())
		assert.True(t, d2.Equal(d.Concat(*d3)))
	})
	convey.Convey("extend deque head with another deque", t, func() {
		d2 := d.Copy()
		d3 := NewLinkedList(0, -1, -2)
		_ = d2.ExtendLeft(d3)
		assert.Equal(t, length+3, d2.Len())
		assert.True(t, d2.Equal(d3.Reverse().Concat(*d)))
	})
}

func TestRotate(t *testing.T) {
	d := NewLinkedList(1, 2, 3, 4, 5)
	length := d.Len()
	convey.Convey("rotate deque to the left", t, func() {
		steps := 2
		d2 := d.Copy()
		_ = d2.Rotate(steps)
		assert.Equal(t, length, d2.Len())
		assert.True(t, NewLinkedList(4, 5, 1, 2, 3).Equal(d2))
	})
	convey.Convey("rotate deque to the right", t, func() {
		steps := -2
		d2 := d.Copy()
		_ = d2.Rotate(steps)
		assert.Equal(t, length, d2.Len())
		assert.True(t, NewLinkedList(3, 4, 5, 1, 2).Equal(d2))
	})
}

func TestForEach(t *testing.T) {
	convey.Convey("convert deque elements", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 5)
		f := func(x any) any {
			return x.(int) * 2
		}
		_ = d.ForEach(f)
		assert.True(t, NewLinkedList(2, 4, 6, 8, 10).Equal(*d))
	})
}

func TestReplace(t *testing.T) {
	d := NewLinkedList(arraylist.Repeat(1, 10)...)
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
	convey.Convey("replace elements satisfy the condition with another element", t, func() {
		d2 := d.Copy().Concat(*NewLinkedList(3, 5, 7))
		d3 := d2.Copy()
		f := func(x any) bool {
			return x.(int)%2 == 1
		}
		replaced := d2.ReplaceIf(f, 0, 3)
		assert.True(t, NewLinkedList(arraylist.Repeat(1, 3)...).Equal(replaced))
		assert.True(t, NewLinkedList(0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 3, 5, 7).Equal(d2))
		replaced = d3.ReplaceRightIf(f, 0, 3)
		assert.True(t, NewLinkedList(7, 5, 3).Equal(replaced))
		assert.True(t, NewLinkedList(1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0).Equal(d3))
	})
}

func TestSplice(t *testing.T) {
	d := NewLinkedList(1, 2, 3, 4, 5)
	convey.Convey("remove elements from deque by area", t, func() {
		d2 := d.Copy()
		_ = d2.Splice(1, 4)
		assert.True(t, NewLinkedList(1).Equal(d2))
	})
	convey.Convey("remove elements from deque by area and insert new elements", t, func() {
		d2 := d.Copy()
		_ = d2.Splice(1, 2, 6, 7)
		assert.True(t, NewLinkedList(1, 6, 7, 4, 5).Equal(d2))
	})
}

func TestFill(t *testing.T) {
	d := NewLinkedList(arraylist.Repeat(1, 5)...)
	convey.Convey("fill the deque with the specified element", t, func() {
		d2 := d.Copy()
		_ = d2.Fill(2)
		assert.True(t, NewLinkedList(2, 2, 2, 2, 2).Equal(d2))
	})
	convey.Convey("fill the deque with the specified element and start index", t, func() {
		d2 := d.Copy()
		_ = d2.Fill(2, 2)
		assert.True(t, NewLinkedList(1, 1, 2, 2, 2).Equal(d2))
	})
	convey.Convey("fill the deque with the specified element from start index to end index", t, func() {
		d2 := d.Copy()
		_ = d2.Fill(2, 1, 3)
		assert.True(t, NewLinkedList(1, 2, 2, 1, 1).Equal(d2))
	})
}

func TestReverse(t *testing.T) {
	convey.Convey("reverse the deque", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 5)
		_ = d.Reverse()
		assert.True(t, NewLinkedList(5, 4, 3, 2, 1).Equal(*d))
	})
}

func TestSet(t *testing.T) {
	convey.Convey("set the element of deque by index", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 5)
		assert.Nil(t, d.Set(-2, 9))
		result, err := d.At(-2)
		assert.Nil(t, err)
		assert.Equal(t, 9, result)
	})
}

func ExampleLinkedList_Remove() {
	l := NewLinkedList(1, 2, 2, 3, 4, 4, 4, 5, 5, 5, 5, 5)
	fmt.Println(l)
	l.Remove(2)
	fmt.Println(l)
	l.Remove(4, -1)
	fmt.Println(l)
	l.Remove(5, 2)
	fmt.Println(l)
	// Output:
	// [1 2 2 3 4 4 4 5 5 5 5 5]
	// [1 2 3 4 4 4 5 5 5 5 5]
	// [1 2 3 5 5 5 5 5]
	// [1 2 3 5 5 5]
}

func ExampleLinkedList_RemoveRight() {
	l := NewLinkedList(1, 2, 3, 4, 5, 4, 3, 2, 1)
	l2 := l.Copy()
	fmt.Println(l)
	l.Remove(4)
	l2.RemoveRight(4)
	fmt.Println(l)
	fmt.Println(l2)
	// Output:
	// [1 2 3 4 5 4 3 2 1]
	// [1 2 3 5 4 3 2 1]
	// [1 2 3 4 5 3 2 1]
}

func ExampleLinkedList_Clear() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	fmt.Println(list)
	list.Clear()
	fmt.Println(list)
	// Output:
	// [1 2 3 4 5]
	// []
}

func ExampleLinkedList_Append() {
	list := NewLinkedList(1, 2, 3)
	fmt.Println(list)
	list.Append(4)
	fmt.Println(list)
	// Output:
	// [1 2 3]
	// [1 2 3 4]
}

func ExampleLinkedList_AppendLeft() {
	list := NewLinkedList(1, 2, 3)
	fmt.Println(list)
	list.AppendLeft(4)
	fmt.Println(list)
	// Output:
	// [1 2 3]
	// [4 1 2 3]
}

func ExampleLinkedList_Pop() {
	l := NewLinkedList(1, 2, 3, 4, 5)
	fmt.Println(l)
	element, _ := l.Pop()
	fmt.Println(element)
	fmt.Println(l)
	// Output:
	// [1 2 3 4 5]
	// 5
	// [1 2 3 4]
}

func ExampleLinkedList_PopLeft() {
	l := NewLinkedList(1, 2, 3, 4, 5)
	fmt.Println(l)
	element, _ := l.PopLeft()
	fmt.Println(element)
	fmt.Println(l)
	// Output:
	// [1 2 3 4 5]
	// 1
	// [2 3 4 5]
}

func ExampleLinkedList_Extend() {
	l1 := NewLinkedList(1, 2, 3)
	fmt.Println(l1)
	l2 := NewLinkedList(4, 5, 6)
	l1.Extend(l2)
	fmt.Println(l1)
	// Output:
	// [1 2 3]
	// [1 2 3 4 5 6]
}

func ExampleLinkedList_ExtendLeft() {
	l1 := NewLinkedList(1, 2, 3)
	fmt.Println(l1)
	l2 := NewLinkedList(4, 5, 6)
	l1.ExtendLeft(l2)
	fmt.Println(l1)
	// Output:
	// [1 2 3]
	// [6 5 4 1 2 3]
}

func ExampleLinkedList_Insert() {
	list := NewLinkedList(1, 2, 3)
	fmt.Println(list)
	list.Insert(1, 5)
	fmt.Println(list)
	// Output:
	// [1 2 3]
	// [1 5 2 3]
}

func ExampleLinkedList_RemoveByIndex() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	fmt.Println(list)
	value, _ := list.RemoveByIndex(2)
	fmt.Println(value)
	fmt.Println(list)
	// Output:
	// [1 2 3 4 5]
	// 3
	// [1 2 4 5]
}

func ExampleLinkedList_Rotate() {
	list := NewLinkedList(1, 2, 3)
	fmt.Println(list)
	// Rotate the linked list
	list.Rotate()
	fmt.Println(list)
	// Rotate the linked list by -2 steps
	list.Rotate(-2)
	fmt.Println(list)
	// Output:
	// [1 2 3]
	// [3 1 2]
	// [2 3 1]
}

func ExampleLinkedList_Reverse() {
	l := NewLinkedList(1, 2, 3, 4, 5)
	fmt.Println(l)
	l.Reverse()
	fmt.Println(l)
	// Output:
	// [1 2 3 4 5]
	// [5 4 3 2 1]
}

func ExampleLinkedList_ForEach() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	fmt.Println(list)
	list.ForEach(func(item any) any {
		return item.(int) * 2
	})
	fmt.Println(list)
	// Output:
	// [1 2 3 4 5]
	// [2 4 6 8 10]
}

func ExampleLinkedList_Replace() {
	l := NewLinkedList(1, 2, 2, 3, 4, 4, 4, 5, 5, 5, 5, 5)
	fmt.Println(l)
	l.Replace(2, -2)
	fmt.Println(l)
	l.Replace(4, -4, -1)
	fmt.Println(l)
	l.Replace(5, -5, 2)
	fmt.Println(l)
	// Output:
	// [1 2 2 3 4 4 4 5 5 5 5 5]
	// [1 -2 2 3 4 4 4 5 5 5 5 5]
	// [1 -2 2 3 -4 -4 -4 5 5 5 5 5]
	// [1 -2 2 3 -4 -4 -4 -5 -5 5 5 5]
}

func ExampleLinkedList_ReplaceRight() {
	l := NewLinkedList(1, 2, 3, 4, 5, 4, 3, 2, 1)
	l2 := l.Copy()
	fmt.Println(l)
	l.Replace(4, 0)
	l2.ReplaceRight(4, 0)
	fmt.Println(l)
	fmt.Println(l2)
	// Output:
	// [1 2 3 4 5 4 3 2 1]
	// [1 2 3 0 5 4 3 2 1]
	// [1 2 3 4 5 0 3 2 1]
}

func ExampleLinkedList_Splice() {
	arr := NewLinkedList(1, 2, 3, 4, 5)
	fmt.Println(arr)
	arr.Splice(2, 2, 6, 7, 8)
	fmt.Println(arr)
	// Output:
	// [1 2 3 4 5]
	// [1 2 6 7 8 5]
}

func ExampleLinkedList_Fill() {
	list := NewLinkedList(1, 2, 3, 4, 5)
	fmt.Println(list)
	list.Fill(0, 1, 3)
	fmt.Println(list)
	// Output:
	// [1 2 3 4 5]
	// [1 0 0 4 5]
}

func ExampleLinkedList_Set() {
	list := NewLinkedList(1, 2, 3)
	fmt.Println(list)
	_ = list.Set(2, 6)
	fmt.Println(list)
	_ = list.Set(-2, 5)
	fmt.Println(list)
	err := list.Set(5, 6)
	fmt.Println(err)
	fmt.Println(list)
	// Output:
	// [1 2 3]
	// [1 2 6]
	// [1 5 6]
	// the index is out of range
	// [1 5 6]
}

func ExampleLinkedList_RemoveIf() {
	list := NewLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	evenCondition := func(val any) bool {
		num := val.(int)
		return num%2 == 0
	}
	fmt.Println(list)
	removed := list.RemoveIf(evenCondition, -1)
	fmt.Println(list)
	fmt.Println(removed)
	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
	// [1 3 5 7 9]
	// [2 4 6 8 10]
}

func ExampleLinkedList_RemoveRightIf() {
	list := NewLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	list2 := list.Copy()
	condition := func(val any) bool {
		return val.(int)%2 == 0
	}
	fmt.Println(list)
	removed1 := list.RemoveRightIf(condition, 3)
	removed2 := list2.RemoveIf(condition, 3)
	fmt.Println(list)
	fmt.Println(removed1)
	fmt.Println(list2)
	fmt.Println(removed2)
	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
	// [1 2 3 4 5 7 9]
	// [10 8 6]
	// [1 3 5 7 8 9 10]
	// [2 4 6]
}

func ExampleLinkedList_ReplaceIf() {
	list := NewLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	evenCondition := func(val any) bool {
		num := val.(int)
		return num%2 == 0
	}
	fmt.Println(list)
	replaced := list.ReplaceIf(evenCondition, -1, -1)
	fmt.Println(list)
	fmt.Println(replaced)
	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
	// [1 -1 3 -1 5 -1 7 -1 9 -1]
	// [2 4 6 8 10]
}

func ExampleLinkedList_ReplaceRightIf() {
	list := NewLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	list2 := list.Copy()
	condition := func(val any) bool {
		return val.(int)%2 == 0
	}
	fmt.Println(list)
	replaced1 := list.ReplaceRightIf(condition, 0, 3)
	replaced2 := list2.ReplaceIf(condition, 0, 3)
	fmt.Println(list)
	fmt.Println(replaced1)
	fmt.Println(list2)
	fmt.Println(replaced2)
	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
	// [1 2 3 4 5 0 7 0 9 0]
	// [10 8 6]
	// [1 0 3 0 5 0 7 8 9 10]
	// [2 4 6]
}
