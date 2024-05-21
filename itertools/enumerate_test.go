package itertools

import (
	"fmt"
	"github.com/biocrosscoder/flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnumerate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	convey.Convey("forward", t, func() {
		iter, err := Enumerate(arr, 1, 10, 2)
		assert.Nil(t, err)
		assert.Equal(t, [][2]any{{1, 2}, {3, 4}}, iter.Pour())
	})
	convey.Convey("string", t, func() {
		iter, err := Enumerate("hello", 0, 10, 1)
		assert.Nil(t, err)
		assert.Equal(t, [][2]any{{0, "h"}, {1, "e"}, {2, "l"}, {3, "l"}, {4, "o"}}, iter.Pour())
	})
	convey.Convey("backward", t, func() {
		iter, err := Enumerate(arr, 4, 0, -2)
		assert.Nil(t, err)
		assert.Equal(t, [][2]any{{4, 5}, {2, 3}, {0, 1}}, iter.Pour())
	})
	convey.Convey("zero step", t, func() {
		iter, err := Enumerate(arr, 1, 10, 0)
		assert.Equal(t, common.ErrZeroStep, err)
		assert.Nil(t, iter)
	})
	convey.Convey("invalid range", t, func() {
		iter, err := Enumerate(arr, 10, 0, 1)
		assert.Equal(t, common.ErrInvalidRange, err)
		assert.Nil(t, iter)
		iter, err = Enumerate(arr, 0, 10, -1)
		assert.Equal(t, common.ErrInvalidRange, err)
		assert.Nil(t, iter)
	})
	convey.Convey("out of range", t, func() {
		iter, err := Enumerate(arr, 5, 10, 2)
		assert.Equal(t, common.ErrOutOfRange, err)
		assert.Nil(t, iter)
	})
}

func ExampleEnumerate() {
	arr := []int{1, 2, 3, 4, 5}
	iter, _ := Enumerate(arr, 0, 4, 2)
	for iter.Next() {
		fmt.Println(iter.Value())
	}
	// Output:
	// [0 1]
	// [2 3]
	// [4 5]
}
