package itertools

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatten(t *testing.T) {
	convey.Convey("flatten list", t, func() {
		entry := []any{
			123,
			[3]int{6, 6, 6},
			"hello",
			[]any{
				[]string{"foo", "bar"},
				map[string]int{"one": 1, "two": 2},
			},
		}
		result, err := Flatten(entry)
		assert.Nil(t, err)
		assert.Equal(t, []any{123, 6, 6, 6, "hello", "foo", "bar", map[string]int{"one": 1, "two": 2}}, result)
	})
}

func ExampleFlatten() {
	entry := []any{1, []any{2, []int{3}}}
	f, _ := Flatten(entry)
	fmt.Println(f)
	// Output: [1 2 3]
}
