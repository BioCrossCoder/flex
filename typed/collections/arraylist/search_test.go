package arraylist

import (
	"fmt"
	"github.com/biocrosscoder/flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchElement(t *testing.T) {
	convey.Convey("search element in list", t, func() {
		l := ArrayList[int]{1, 2, 3, 2, 4}
		convey.Convey("search by index", func() {
			assert.Equal(t, 1, l.IndexOf(2))
			assert.Equal(t, 3, l.LastIndexOf(2))
		})
		convey.Convey("get element by index", func() {
			num, err := l.At(-3)
			assert.Nil(t, err)
			assert.Equal(t, l[2], num)
		})
		convey.Convey("search by condition", func() {
			f := func(num int) bool {
				return num%2 == 0
			}
			v, found := l.Find(f)
			assert.True(t, found)
			assert.Equal(t, 2, v)
			v, found = l.FindLast(f)
			assert.True(t, found)
			assert.Equal(t, 4, v)
			assert.Equal(t, 1, l.FindIndex(f))
			assert.Equal(t, 4, l.FindLastIndex(f))
			assert.Equal(t, []int{1, 3, 4}, l.FindIndexes(f))
			assert.Equal(t, []int{4, 3, 1}, l.FindLastIndexes(f))
			assert.Equal(t, []int{2, 2, 4}, l.Finds(f))
			assert.Equal(t, []int{4, 2, 2}, l.FindLasts(f))
			assert.Equal(t, []int{1, 3}, l.FindIndexes(f, 2))

		})
		convey.Convey("get first/last element", func() {
			head, err := l.Head()
			assert.Nil(t, err)
			assert.Equal(t, l[0], head)
			tail, err := l.Tail()
			assert.Nil(t, err)
			assert.Equal(t, l[l.Len()-1], tail)
		})
		convey.Convey("element not found", func() {
			assert.Equal(t, -1, l.IndexOf(0))
			assert.Equal(t, -1, l.LastIndexOf(0))
			f := func(num int) bool {
				return num < 0
			}
			v, found := l.Find(f)
			assert.False(t, found)
			assert.Zero(t, v)
			v, found = l.FindLast(f)
			assert.False(t, found)
			assert.Zero(t, v)
			assert.Equal(t, -1, l.FindIndex(f))
			assert.Equal(t, -1, l.FindLastIndex(f))
		})
		convey.Convey("empty list", func() {
			l := ArrayList[any]{}
			head, err := l.Head()
			assert.Equal(t, err, common.ErrEmptyList)
			assert.Nil(t, head)
			tail, err := l.Tail()
			assert.Equal(t, err, common.ErrEmptyList)
			assert.Nil(t, tail)
		})
	})
}

func ExampleArrayList_IndexOf() {
	list := ArrayList[int]{1, 2, 3, 4, 5}
	fmt.Println(list.IndexOf(3))
	fmt.Println(list.IndexOf(6))
	// Output:
	// 2
	// -1
}

func ExampleArrayList_LastIndexOf() {
	list := ArrayList[int]{1, 2, 3, 4, 3}
	fmt.Println(list.IndexOf(3))
	fmt.Println(list.LastIndexOf(3))
	// Output:
	// 2
	// 4
}

func ExampleArrayList_At() {
	list := ArrayList[int]{1, 2, 3, 4, 5}
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

func ExampleArrayList_Find() {
	list := ArrayList[int]{1, 2, 3, 4, 5}
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

func ExampleArrayList_FindIndex() {
	list := ArrayList[int]{1, 2, 3, 4, 5}
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

func ExampleArrayList_FindLast() {
	list := ArrayList[int]{1, 2, 3, 4, 5}
	even := func(num int) bool {
		return num%2 == 0
	}
	fmt.Println(list.Find(even))
	fmt.Println(list.FindLast(even))
	// Output:
	// 2 true
	// 4 true
}

func ExampleArrayList_FindLastIndex() {
	list := ArrayList[int]{1, 2, 3, 4, 5}
	even := func(num int) bool {
		return num%2 == 0
	}
	fmt.Println(list.FindIndex(even))
	fmt.Println(list.FindLastIndex(even))
	// Output:
	// 1
	// 3
}

func ExampleArrayList_Head() {
	list := ArrayList[int]{1, 2, 3, 4, 5}
	fmt.Println(list.Head())
	fmt.Println(ArrayList[int]{}.Head())
	// Output:
	// 1 <nil>
	// 0 the input list is empty
}

func ExampleArrayList_Tail() {
	list := ArrayList[int]{1, 2, 3, 4, 5}
	fmt.Println(list.Tail())
	fmt.Println(ArrayList[int]{}.Tail())
	// Output:
	// 5 <nil>
	// 0 the input list is empty
}

func ExampleArrayList_FindIndexes() {
	l := ArrayList[int]{1, 2, 3, 4, 5, 2}
	condition := func(val int) bool {
		return val > 2
	}
	fmt.Println(l.FindIndexes(condition))
	fmt.Println(l.FindIndexes(condition, 2))
	// Output:
	// [2 3 4]
	// [2 3]
}

func ExampleArrayList_FindLastIndexes() {
	l := ArrayList[int]{1, 2, 3, 4, 5, 2}
	condition := func(val int) bool {
		return val > 2
	}
	fmt.Println(l.FindIndexes(condition))
	fmt.Println(l.FindLastIndexes(condition))
	// Output:
	// [2 3 4]
	// [4 3 2]
}

func ExampleArrayList_Finds() {
	l := ArrayList[int]{1, 2, 3, 4, 5, 2}
	condition := func(val int) bool {
		return val > 2
	}
	fmt.Println(l.Finds(condition))
	fmt.Println(l.Finds(condition, 2))
	// Output:
	// [3 4 5]
	// [3 4]
}

func ExampleArrayList_FindLasts() {
	l := ArrayList[int]{1, 2, 3, 4, 5, 2}
	condition := func(val int) bool {
		return val > 2
	}
	fmt.Println(l.Finds(condition))
	fmt.Println(l.FindLasts(condition))
	// Output:
	// [3 4 5]
	// [5 4 3]
}
