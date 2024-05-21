package itertools

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhile(t *testing.T) {
	f := func(x any) bool {
		return x.(int) > 0
	}
	entry := []int{1, 2, 3, 0, 4, 5, 6}
	convey.Convey("drop while", t, func() {
		iter, err := DropWhile(f, entry)
		assert.Nil(t, err)
		assert.Equal(t, []any{0, 4, 5, 6}, iter.Pour())
	})
	convey.Convey("take while", t, func() {
		iter, err := TakeWhile(f, entry)
		assert.Nil(t, err)
		assert.Equal(t, []any{1, 2, 3}, iter.Pour())
	})
}

func ExampleDropWhile() {
	arr := []int{1, 2, 3, 0, 4, 5, 6}
	f := func(x any) bool {
		return x.(int) > 0
	}
	iter, _ := DropWhile(f, arr)
	for iter.Next() {
		fmt.Println(iter.Value())
	}
	// Output:
	// 0
	// 4
	// 5
	// 6
}

func ExampleTakeWhile() {
	arr := []int{1, 2, 3, 0, 4, 5, 6}
	f := func(x any) bool {
		return x.(int) > 0
	}
	iter, _ := TakeWhile(f, arr)
	for iter.Next() {
		fmt.Println(iter.Value())
	}
	// Output:
	// 1
	// 2
	// 3
}
