package arraylist

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemove(t *testing.T) {
	l := ArrayList{1, 2, 3, 2, 4, 2, 3, 2, 1}
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
		f := func(x any) bool {
			return x.(int)%2 == 0
		}
		removed := l2.RemoveIf(f, 3)
		assert.Equal(t, removed, ArrayList{2, 2, 4})
		assert.Equal(t, l2, ArrayList{1, 3, 2, 3, 2, 1, 6, 8})
		removed = l3.RemoveRightIf(f, 3)
		assert.Equal(t, removed, ArrayList{8, 6, 2})
		assert.Equal(t, l3, ArrayList{1, 2, 3, 2, 4, 2, 3, 1})
	})
}

func TestAddOrCutElement(t *testing.T) {
	l := ArrayList{1, 2, 3, 4, 5}
	length := l.Len()
	convey.Convey("add elements to list tail", t, func() {
		l2 := l.Copy()
		_ = l2.Push(1, 2, 3)
		assert.Equal(t, length+3, l2.Len())
		assert.Equal(t, l2[l2.Len()-3:], Of(1, 2, 3))
	})
	convey.Convey("add elements to list head", t, func() {
		l2 := l.Copy()
		s := ArrayList{1, 2, 3}
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
		l := ArrayList{1, 2, 3, 4, 5}
		f := func(x any) any {
			return -x.(int)
		}
		_ = l.ForEach(f)
		assert.Equal(t, ArrayList{-1, -2, -3, -4, -5}, l)
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
		f := func(x any) bool {
			return x.(int)%2 == 1
		}
		replaced := l2.ReplaceIf(f, 0, 3)
		assert.Equal(t, replaced, Repeat(1, 3))
		assert.Equal(t, l2, ArrayList{0, 0, 0, 1, 1, 3, 5, 7})
		replaced = l3.ReplaceRightIf(f, 0, 3)
		assert.Equal(t, replaced, ArrayList{7, 5, 3})
		assert.Equal(t, l3, ArrayList{1, 1, 1, 1, 1, 0, 0, 0})
	})
}

func TestSplice(t *testing.T) {
	l := ArrayList{1, 2, 3, 4, 5}
	convey.Convey("remove elements from list by area", t, func() {
		l2 := l.Copy()
		_ = l2.Splice(1, 3)
		assert.Equal(t, l2, ArrayList{1, 5})
	})
	convey.Convey("remove elements from list by area and insert new elements", t, func() {
		l2 := l.Copy()
		_ = l2.Splice(1, 2, 6, 0)
		assert.Equal(t, l2, ArrayList{1, 6, 0, 4, 5})
	})
}

func TestFill(t *testing.T) {
	l := make(ArrayList, 5)
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
		assert.Equal(t, l2, ArrayList{nil, 6, 6, nil, nil})
	})
}

func TestReverse(t *testing.T) {
	convey.Convey("reverse the list", t, func() {
		l := ArrayList{1, 2, 3, 4, 5}
		_ = l.Reverse()
		assert.Equal(t, l, ArrayList{5, 4, 3, 2, 1})
	})
}

func TestSet(t *testing.T) {
	convey.Convey("set the element of list by index", t, func() {
		l := ArrayList{1, 2, 3, 4, 5}
		assert.Nil(t, l.Set(-2, 6))
		result, err := l.At(-2)
		assert.Nil(t, err)
		assert.Equal(t, result, 6)
	})
}

func ExampleArrayList_Remove() {
	l := ArrayList{1, 2, 2, 3, 4, 4, 4, 5, 5, 5, 5, 5}
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

func ExampleArrayList_RemoveRight() {
	l := ArrayList{1, 2, 3, 4, 5, 4, 3, 2, 1}
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

func ExampleArrayList_Clear() {
	list := ArrayList{1, 2, 3, 4, 5}
	fmt.Println(list)
	list.Clear()
	fmt.Println(list)
	// Output:
	// [1 2 3 4 5]
	// []
}

func ExampleArrayList_Push() {
	list := ArrayList{1, 2, 3}
	fmt.Println(list)
	list.Push(4, 5, 6)
	fmt.Println(list)
	// Output:
	// [1 2 3]
	// [1 2 3 4 5 6]
}

func ExampleArrayList_Pop() {
	l := ArrayList{1, 2, 3, 4, 5}
	fmt.Println(l)
	// Remove the last element
	element, _ := l.Pop()
	fmt.Println(element)
	fmt.Println(l)
	// Remove the element at a specific index
	element, _ = l.Pop(2)
	fmt.Println(element)
	fmt.Println(l)
	// Output:
	// [1 2 3 4 5]
	// 5
	// [1 2 3 4]
	// 3
	// [1 2 4]
}

func ExampleArrayList_Unshift() {
	l := ArrayList{1, 2, 3}
	fmt.Println(l)
	l.Unshift(4, 5)
	fmt.Println(l)
	// Output:
	// [1 2 3]
	// [4 5 1 2 3]
}

func ExampleArrayList_Shift() {
	list := ArrayList{1, 2, 3, 4, 5}
	fmt.Println(list)
	element, _ := list.Shift()
	fmt.Println(element)
	fmt.Println(list)
	// Output:
	// [1 2 3 4 5]
	// 1
	// [2 3 4 5]
}

func ExampleArrayList_Insert() {
	list := ArrayList{1, 2, 3}
	fmt.Println(list)
	list.Insert(1, 5)
	fmt.Println(list)
	// Output:
	// [1 2 3]
	// [1 5 2 3]
}

func ExampleArrayList_ForEach() {
	list := ArrayList{1, 2, 3, 4, 5}
	fmt.Println(list)
	list.ForEach(func(item any) any {
		return item.(int) * 2
	})
	fmt.Println(list)
	// Output:
	// [1 2 3 4 5]
	// [2 4 6 8 10]
}

func ExampleArrayList_Replace() {
	l := ArrayList{1, 2, 2, 3, 4, 4, 4, 5, 5, 5, 5, 5}
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

func ExampleArrayList_ReplaceRight() {
	l := ArrayList{1, 2, 3, 4, 5, 4, 3, 2, 1}
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

func ExampleArrayList_Splice() {
	arr := ArrayList{1, 2, 3, 4, 5}
	fmt.Println(arr)
	arr.Splice(2, 2, 6, 7, 8)
	fmt.Println(arr)
	// Output:
	// [1 2 3 4 5]
	// [1 2 6 7 8 5]
}

func ExampleArrayList_Fill() {
	list := ArrayList{1, 2, 3, 4, 5}
	fmt.Println(list)
	list.Fill(0, 1, 3)
	fmt.Println(list)
	// Output:
	// [1 2 3 4 5]
	// [1 0 0 4 5]
}

func ExampleArrayList_Reverse() {
	l := ArrayList{1, 2, 3, 4, 5}
	fmt.Println(l)
	l.Reverse()
	fmt.Println(l)
	// Output:
	// [1 2 3 4 5]
	// [5 4 3 2 1]
}

func ExampleArrayList_Set() {
	list := ArrayList{1, 2, 3}
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

func ExampleArrayList_RemoveIf() {
	lst := ArrayList{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenCondition := func(val any) bool {
		num := val.(int)
		return num%2 == 0
	}
	fmt.Println(lst)
	removed := lst.RemoveIf(evenCondition, -1)
	fmt.Println(lst)
	fmt.Println(removed)
	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
	// [1 3 5 7 9]
	// [2 4 6 8 10]
}

func ExampleArrayList_RemoveRightIf() {
	list := ArrayList{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
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

func ExampleArrayList_ReplaceIf() {
	lst := ArrayList{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenCondition := func(val any) bool {
		num := val.(int)
		return num%2 == 0
	}
	fmt.Println(lst)
	replaced := lst.ReplaceIf(evenCondition, -1, -1)
	fmt.Println(lst)
	fmt.Println(replaced)
	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
	// [1 -1 3 -1 5 -1 7 -1 9 -1]
	// [2 4 6 8 10]
}

func ExampleArrayList_ReplaceRightIf() {
	list := ArrayList{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
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
