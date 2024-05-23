package arraylist

import (
	"fmt"
	"github.com/biocrosscoder/flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	convey.Convey("mapping list", t, func() {
		l := ArrayList[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		f := func(x int) int {
			return x * 2
		}
		assert.Equal(t, ArrayList[int]{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, l.Map(f))
	})
}

func TestReduce(t *testing.T) {
	convey.Convey("reduce list", t, func() {
		l := ArrayList[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		f := func(x, y int) int {
			return x - y
		}
		convey.Convey("normal reduce", func() {
			result, err := l.Reduce(f)
			assert.Nil(t, err)
			assert.Equal(t, -53, result)
		})
		convey.Convey("reduce with initial value", func() {
			result, err := l.Reduce(f, 10)
			assert.Nil(t, err)
			assert.Equal(t, -45, result)
		})
		convey.Convey("reduce from right", func() {
			result, err := l.ReduceRight(f)
			assert.Nil(t, err)
			assert.Equal(t, -35, result)
		})
		convey.Convey("reduce from right with initial value", func() {
			result, err := l.ReduceRight(f, 10)
			assert.Nil(t, err)
			assert.Equal(t, -45, result)
		})
		convey.Convey("reduce on empty list", func() {
			l := ArrayList[int]{}
			result, err := l.Reduce(f)
			assert.Equal(t, common.ErrEmptyList, err)
			assert.Zero(t, result)
			result, err = l.ReduceRight(f)
			assert.Equal(t, common.ErrEmptyList, err)
			assert.Zero(t, result)
		})
		convey.Convey("too many arguments", func() {
			result, err := l.Reduce(f, 1, 2)
			assert.Equal(t, common.ErrTooManyArguments, err)
			assert.Zero(t, result)
			result, err = l.ReduceRight(f, 1, 2)
			assert.Equal(t, common.ErrTooManyArguments, err)
			assert.Zero(t, result)
		})
	})
}

func TestFilter(t *testing.T) {
	convey.Convey("filter list", t, func() {
		l := ArrayList[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		f := func(x int) bool {
			return x%2 == 0
		}
		assert.Equal(t, ArrayList[int]{2, 4, 6, 8, 10}, l.Filter(f))
	})
}

func TestSomeAndEvery(t *testing.T) {
	convey.Convey("check condition on list", t, func() {
		l := ArrayList[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		f := func(x int) bool {
			return x > 9
		}
		convey.Convey("at least one element in list satisfies the condition", func() {
			assert.True(t, l.Some(f))
			assert.False(t, l.Some(func(x int) bool {
				return x > 50
			}))
		})
		convey.Convey("all elements in list satisfy the condition", func() {
			assert.False(t, l.Every(f))
			assert.True(t, l.Every(func(x int) bool {
				return x > 0
			}))
		})
	})
}

func ExampleArrayList_Map() {
	list := ArrayList[int]{1, 2, 3, 4, 5}
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

func ExampleArrayList_Reduce() {
	list := ArrayList[int]{1, 2, 3, 4, 5}
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

func ExampleArrayList_ReduceRight() {
	list := ArrayList[int]{1, 2, 3, 4, 5}
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

func ExampleArrayList_Filter() {
	list := ArrayList[int]{1, 2, 3, 4, 5}
	condition := func(item int) bool {
		return item%2 == 0
	}
	filteredList := list.Filter(condition)
	fmt.Println(filteredList)
	// Output: [2 4]
}

func ExampleArrayList_Some() {
	list := ArrayList[int]{1, 2, 3, 4, 5}
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

func ExampleArrayList_Every() {
	list := ArrayList[int]{1, 2, 3, 4, 5}
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
