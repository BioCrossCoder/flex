package itertools

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestAccumulator(t *testing.T) {
	convey.Convey("accumulate all", t, func() {
		entry := []any{1, 2, 3}
		f := func(x, y int) int {
			return x + y
		}
		iterator := NewAccumulator(entry, func(p1, p2 any) any {
			params := []reflect.Value{reflect.ValueOf(p1), reflect.ValueOf(p2)}
			return reflect.ValueOf(f).Call(params)[0].Interface()
		})
		result := iterator.Pour()
		expected := 0
		for _, v := range entry {
			expected += v.(int)
		}
		assert.Equal(t, expected, result)
		assert.False(t, iterator.Next())
		assert.Nil(t, iterator.Value())
	})
}

func ExampleAccumulate() {
	arr := []int{1, 2, 3}
	f := func(x, y int) int {
		return x + y
	}
	iter, _ := Accumulate(f, arr)
	fmt.Println(iter.Value())
	for iter.Next() {
		fmt.Println(iter.Value())
	}
	// Output:
	// 1
	// 3
	// 6
}
