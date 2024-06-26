package functools

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountBy(t *testing.T) {
	convey.Convey("count element satisfying condition", t, func() {
		assert.Equal(t, CountBy([]int{1, 2, 3}, func(a any) bool {
			return fmt.Sprint(a) == "1"
		}), 1)
		assert.Equal(t, CountBy("hello", func(a any) bool {
			return a.(string) < "l"
		}), 2)
		assert.Equal(t, CountBy(123, func(any) bool { return true }), -1)
	})
}

func ExampleCountBy() {
	// string
	fmt.Println(CountBy("hello", func(a any) bool {
		return a.(string) != "l"
	}))
	// slice
	fmt.Println(CountBy([]int{1, 2, 3, 4, 5}, func(a any) bool {
		return a.(int)%2 == 0
	}))
	// Output:
	// 3
	// 2
}
