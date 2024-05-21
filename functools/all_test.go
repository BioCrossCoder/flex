package functools

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAll(t *testing.T) {
	f := func(a any) bool {
		return a.(int) > 0
	}
	convey.Convey("true", t, func() {
		entry := []int{1, 2, 3}
		result, err := All(entry, f)
		assert.Nil(t, err)
		assert.True(t, result)
	})
	convey.Convey("false", t, func() {
		entry := []int{1, 2, 0}
		result, err := All(entry, f)
		assert.Nil(t, err)
		assert.False(t, result)
	})
}

func ExampleAll() {
	// string
	r, _ := All("hello", func(a any) bool {
		return a.(string) > "a"
	})
	fmt.Println(r)
	r, _ = All("hello", func(a any) bool {
		return a.(string) > "g"
	})
	fmt.Println(r)
	// slice
	r, _ = All([]int{1, 2, 3}, func(a any) bool {
		return a.(int) > 0
	})
	fmt.Println(r)
	r, _ = All([]int{1, 2, 3}, func(a any) bool {
		return a.(int) > 2
	})
	fmt.Println(r)
	// Output:
	// true
	// false
	// true
	// false
}
