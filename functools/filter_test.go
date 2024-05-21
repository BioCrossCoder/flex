package functools

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	f := func(a string) bool {
		return a > "0"
	}
	convey.Convey("Call Filter on [array]", t, func() {
		s := [3]string{"0", "1", "2"}
		expected := []any{"1", "2"}
		actual, err := Filter(f, s)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	convey.Convey("Call Filter on [slice]", t, func() {
		s := []string{"0", "1", "2", "3"}
		expected := []any{"1", "2", "3"}
		actual, err := Filter(f, s)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	convey.Convey("Call Filter on [string]", t, func() {
		s := "012345"
		expected := []any{"1", "2", "3", "4", "5"}
		actual, err := Filter(f, s)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}

func ExampleFilter() {
	// string
	f, _ := Filter(func(a string) bool {
		return a != "l"
	}, "hello")
	fmt.Println(f)
	// slice
	f, _ = Filter(func(a int) bool {
		return a%2 == 0
	}, []int{1, 2, 3, 4, 5})
	fmt.Println(f)
	// Output:
	// [h e o]
	// [2 4]
}
