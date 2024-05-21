package functools

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAny(t *testing.T) {
	f := func(a any) bool {
		return a == 0
	}
	convey.Convey("true", t, func() {
		entry := []int{1, 0, 1}
		result, err := Any(entry, f)
		assert.Nil(t, err)
		assert.True(t, result)
	})
	convey.Convey("false", t, func() {
		entry := []int{1, 3, 2}
		result, err := Any(entry, f)
		assert.Nil(t, err)
		assert.False(t, result)
	})
}

func ExampleAny() {
	// string
	r, _ := Any("hello", func(a any) bool {
		return a.(string) > "n"
	})
	fmt.Println(r)
	r, _ = Any("hello", func(a any) bool {
		return a.(string) > "z"
	})
	fmt.Println(r)
	// slice
	r, _ = All([]int{1, 2, 3}, func(a any) bool {
		return a.(int) > 0
	})
	fmt.Println(r)
	r, _ = All([]int{1, 2, 3}, func(a any) bool {
		return a.(int) > 3
	})
	fmt.Println(r)
	// Output:
	// true
	// false
	// true
	// false
}
