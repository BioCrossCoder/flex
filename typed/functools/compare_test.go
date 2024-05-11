package functools

import (
	"github.com/biocrosscoder/flex/typed/collections/sortedcontainers/sortedlist"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompare(t *testing.T) {
	convey.Convey("get max/min element", t, func() {
		arr := []int{1, 2, 3, 4, 5}
		assert.Equal(t, 1, Max(sortedlist.DescendOrder, arr...))
		assert.Equal(t, 1, Min(sortedlist.AscendOrder, arr...))
		assert.Equal(t, 5, Max(sortedlist.AscendOrder, arr...))
		assert.Equal(t, 5, Min(sortedlist.DescendOrder, arr...))
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
		assert.True(t, IsMonotonic(arr, sortedlist.AscendOrder, true))
		assert.True(t, IsMonotonic(arr, sortedlist.AscendOrder, false))
		arr2 := []int{1, 2, 3, 4, 4}
		assert.False(t, IsMonotonic(arr2, sortedlist.AscendOrder, true))
		assert.True(t, IsMonotonic(arr2, sortedlist.AscendOrder, false))
		arr3 := []int{5, 4, 3, 2, 1}
		assert.True(t, IsMonotonic(arr3, sortedlist.AscendOrder, true))
		assert.True(t, IsMonotonic(arr3, sortedlist.AscendOrder, false))
		arr4 := []int{5, 3, 3, 2, 1}
		assert.False(t, IsMonotonic(arr4, sortedlist.AscendOrder, true))
		assert.True(t, IsMonotonic(arr4, sortedlist.AscendOrder, false))
	})
}
