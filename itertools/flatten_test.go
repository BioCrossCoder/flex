// Package itertools provides iterator functions to create iterators and perform common operations on iterables.
package itertools

import (
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
