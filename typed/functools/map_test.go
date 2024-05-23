package functools

import (
	"fmt"
	"github.com/biocrosscoder/flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	convey.Convey("Call Map on array", t, func() {
		arr := []int{1, 2, 3, 4, 5}
		f := func(x int) float64 {
			return float64(x*2) / 10
		}
		expected := []float64{0.2, 0.4, 0.6, 0.8, 1}
		assert.Equal(t, expected, Map(f, arr))
	})
	convey.Convey("Call Maps on arrays", t, func() {
		arr1 := []int{1, 2, 3}
		arr2 := [][]int{
			{1, 2, 2, 1},
			{2, 5, 7, 3},
			{4, 5, 6, 7},
		}
		f := func(x int, y []int) bool {
			for _, v := range y {
				if x == v {
					return true
				}
			}
			return false
		}
		expected := []bool{true, true, false}
		actual, err := Maps(f, arr1, arr2)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	convey.Convey("Call Maps failed", t, func() {
		arr1 := []int{1, 2, 3}
		arr2 := []int{6}
		f := func(x int, y int) bool {
			return x == y
		}
		actual, err := Maps(f, arr1, arr2)
		assert.Equal(t, err, common.ErrListLengthMismatch)
		assert.Nil(t, actual)
	})
}

func ExampleMap() {
	str := "hello"
	arr := []int{1, 2, 3, 4, 5}
	f := func(x int) string {
		return string(str[x-1])
	}
	fmt.Println(Map(f, arr))
	// Output: [h e l l o]
}

func ExampleMaps() {
	arr1 := []int{1, 2, 3}
	arr2 := [][]any{{"a", -1}, {"b", -2}, {"c", -3}}
	f := func(x int, y []any) []any {
		return []any{x, y[0], y[1]}
	}
	m, _ := Maps(f, arr1, arr2)
	fmt.Println(m)
	// Output: [[1 a -1] [2 b -2] [3 c -3]]
}
