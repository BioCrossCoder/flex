package itertools

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccumulate(t *testing.T) {
	convey.Convey("accumulate all", t, func() {
		entry := []int{1, 2, 3}
		f := func(x, y int) int {
			return x + y
		}
		iterator := Accumulate(entry, f)
		result := iterator.Pour()
		expected := 0
		for _, v := range entry {
			expected += v
		}
		assert.Equal(t, expected, result)
		assert.False(t, iterator.Next())
		assert.Zero(t, iterator.Value())
	})
}

func ExampleAccumulate() {
	arr := []int{1, 2, 3}
	f := func(x, y int) int {
		return x + y
	}
	iter := Accumulate(arr, f)
	fmt.Println(iter.Value())
	for iter.Next() {
		fmt.Println(iter.Value())
	}
	// Output:
	// 1
	// 3
	// 6
}
