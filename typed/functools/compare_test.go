package functools

import (
	"fmt"
	"github.com/biocrosscoder/flex/typed/collections/sortedcontainers/sortedlist"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompare(t *testing.T) {
	convey.Convey("get max/min element", t, func() {
		arr := []int{1, 2, 3, 4, 5}
		assert.Equal(t, 1, Max(sortedlist.DescendOrder[int], arr...))
		assert.Equal(t, 1, Min(sortedlist.AscendOrder[int], arr...))
		assert.Equal(t, 5, Max(sortedlist.AscendOrder[int], arr...))
		assert.Equal(t, 5, Min(sortedlist.DescendOrder[int], arr...))
	})
	convey.Convey("compare elements", t, func() {
		cmp := func(a, b []int) int {
			return sortedlist.AscendOrder(a[0], b[0])
		}
		arr1 := []int{1, 2}
		arr2 := []int{1, 3}
		arr3 := []int{1, 2, 3}
		arr4 := []int{2, 3}
		assert.True(t, Equals(cmp, arr1, arr2, arr3))
		assert.False(t, Equals(cmp, arr1, arr4, arr3))
	})
	convey.Convey("check monotonicity of array", t, func() {
		arr := []int{1, 2, 3, 4, 5}
		assert.True(t, IsMonotonic(arr, sortedlist.AscendOrder[int], true))
		assert.True(t, IsMonotonic(arr, sortedlist.AscendOrder[int], false))
		arr2 := []int{1, 2, 3, 4, 4}
		assert.False(t, IsMonotonic(arr2, sortedlist.AscendOrder[int], true))
		assert.True(t, IsMonotonic(arr2, sortedlist.AscendOrder[int], false))
		arr3 := []int{5, 4, 3, 2, 1}
		assert.True(t, IsMonotonic(arr3, sortedlist.AscendOrder[int], true))
		assert.True(t, IsMonotonic(arr3, sortedlist.AscendOrder[int], false))
		arr4 := []int{5, 3, 3, 2, 1}
		assert.False(t, IsMonotonic(arr4, sortedlist.AscendOrder[int], true))
		assert.True(t, IsMonotonic(arr4, sortedlist.AscendOrder[int], false))
	})
}

func ExampleMax() {
	max := Max(sortedlist.DescendOrder[int], 5, 3, 9, 2, 7)
	fmt.Println(max)
	// Output: 2
}

func ExampleMin() {
	min := Min(sortedlist.DescendOrder[int], 5, 3, 9, 2, 7)
	fmt.Println(min)
	// Output: 9
}

func ExampleEquals() {
	eq := func(a, b int) int {
		return sortedlist.AscendOrder(a%3, b%3)
	}
	equal := Equals(eq, 3, 6, 9)
	fmt.Println(equal)
	// Output: true
}

func ExampleEqual() {
	fmt.Println(Equal(3, 9, func(a, b int) int {
		return sortedlist.AscendOrder(a%3, b%3)
	}))
	// Output: true
}

func ExampleLess() {
	fmt.Println(Less(2, 3, func(a, b int) int {
		return sortedlist.AscendOrder(a%3, b%3)
	}))
	// Output: false
}

func ExampleGreater() {
	fmt.Println(Greater(5, 3, func(a, b int) int {
		return sortedlist.AscendOrder(a%2, b%2)
	}))
	// Output: false
}

func ExampleIsIncreasing() {
	arr := []int{1, 2, 3, 3, 4}
	fmt.Println(IsIncreasing(arr, sortedlist.AscendOrder[int], false))
	fmt.Println(IsIncreasing(arr, sortedlist.AscendOrder[int], true))
	// Output:
	// true
	// false
}

func ExampleIsDecreasing() {
	arr := []int{5, 4, 3, 3, 2}
	fmt.Println(IsDecreasing(arr, sortedlist.AscendOrder[int], false))
	fmt.Println(IsDecreasing(arr, sortedlist.AscendOrder[int], true))
	// Output:
	// true
	// false
}

func ExampleIsMonotonic() {
	arr := []int{1, 2, 2, 4, 5}
	fmt.Println(IsMonotonic(arr, sortedlist.AscendOrder[int], false))
	fmt.Println(IsMonotonic(arr, sortedlist.AscendOrder[int], true))
	fmt.Println(IsMonotonic(arr, sortedlist.DescendOrder[int], false))
	// Output:
	// true
	// false
	// true
}
