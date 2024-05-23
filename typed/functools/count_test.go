package functools

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCount(t *testing.T) {
	convey.Convey("count element satisfying condition", t, func() {
		assert.Equal(t, CountBy([]int{1, 2, 3}, func(a int) bool {
			return fmt.Sprint(a) == "1"
		}), 1)
		assert.Equal(t, Count([]int{1, 2, 3}, 1), 1)
	})
}

func ExampleCount() {
	fmt.Println(Count([]int{1, 2, 3, 1}, 1))
	// Output: 2
}

func ExampleCountBy() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	f := func(a int) bool {
		return a%2 == 0
	}
	fmt.Println(CountBy(arr, f))
	// Output: 4
}
