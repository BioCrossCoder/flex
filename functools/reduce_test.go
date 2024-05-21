package functools

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func f(a, b string) string {
	return a + b
}

func TestReduce(t *testing.T) {
	convey.Convey("Call Reduce on [array]", t, func() {
		s := [3]string{"a", "b", "c"}
		expected := "abc"
		actual, err := Reduce(f, s)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	convey.Convey("Call Reduce on [slice]", t, func() {
		s := []string{"h", "e", "l", "l", "o"}
		expected := "hello"
		actual, err := Reduce(f, s)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	convey.Convey("Call Reduce on [string]", t, func() {
		s := "hello"
		actual, err := Reduce(f, s)
		assert.Nil(t, err)
		assert.Equal(t, s, actual)
	})
}

func ExampleReduce() {
	arr := []int{1, 2, 3, 4, 5}
	// sum
	r, _ := Reduce(func(a, b int) int {
		return a + b
	}, arr)
	fmt.Println(r)
	// product
	r, _ = Reduce(func(a, b int) int {
		return a * b
	}, arr)
	fmt.Println(r)
	// Output:
	// 15
	// 120
}
