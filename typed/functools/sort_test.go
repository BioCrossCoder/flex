package functools

import (
	"fmt"
	"github.com/biocrosscoder/flex/typed/collections/arraylist"
	"github.com/biocrosscoder/flex/typed/collections/sortedcontainers/sortedlist"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSort(t *testing.T) {
	convey.Convey("simple sort", t, func() {
		a := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		b := Sorted(a, sortedlist.AscendOrder[int])
		Sort(a, sortedlist.AscendOrder[int])
		assert.Equal(t, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}, a)
		assert.Equal(t, a, b)
	})
	convey.Convey("complex sort", t, func() {
		arr := make(arraylist.ArrayList[[3]int], 10)
		for i := 0; i < 10; i++ {
			for j := 0; j < 3; j++ {
				arr[i][j] = rand.Intn(10)
			}
		}
		cmps := []func(a, b [3]int) int{
			func(a, b [3]int) int {
				return sortedlist.AscendOrder(a[0], b[0])
			}, func(a, b [3]int) int {
				return sortedlist.AscendOrder(a[1], b[1])
			}, func(a, b [3]int) int {
				return sortedlist.AscendOrder(a[2], b[2])
			},
		}
		Sort(arr, cmps...)
		assert.True(t, IsSorted(arr, cmps...))
	})
}

func ExampleSort() {
	arr := [][]int{{3, 1, 4}, {1, 5, 9}, {2, 6, 5}, {3, 5, 5}}
	f1 := func(a, b []int) int {
		return sortedlist.AscendOrder(a[0], b[0])
	}
	f2 := func(a, b []int) int {
		return sortedlist.AscendOrder(a[1], b[1])
	}
	fmt.Println(arr)
	Sort(arr, f1, f2)
	fmt.Println(arr)
	// Output:
	// [[3 1 4] [1 5 9] [2 6 5] [3 5 5]]
	// [[1 5 9] [2 6 5] [3 1 4] [3 5 5]]
}

func ExampleSorted() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	arr2 := Sorted(arr, sortedlist.AscendOrder[int])
	fmt.Println(arr)
	fmt.Println(arr2)
	// Output:
	// [3 1 4 1 5 9 2 6 5 3 5]
	// [1 1 2 3 3 4 5 5 5 6 9]
}

func ExampleIsSorted() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println(IsSorted(arr, sortedlist.AscendOrder[int]))
	Sort(arr, sortedlist.AscendOrder[int])
	fmt.Println(IsSorted(arr, sortedlist.AscendOrder[int]))
	// Output:
	// false
	// true
}
