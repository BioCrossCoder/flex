package itertools

import (
	"fmt"
	"github.com/biocrosscoder/flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	convey.Convey("array slice", t, func() {
		s, err := Slice(arr, 5, 2, -1)
		assert.Nil(t, err)
		assert.Equal(t, []any{6, 5, 4, 3}, s)
	})
	convey.Convey("string slice", t, func() {
		entry := "hello world"
		s, err := Slice(entry, 2, 5, 1)
		assert.Nil(t, err)
		assert.Equal(t, "llo ", s)
	})
	convey.Convey("zero step", t, func() {
		s, err := Slice(arr, 1, 5, 0)
		assert.Equal(t, common.ErrZeroStep, err)
		assert.Nil(t, s)
	})
	convey.Convey("invalid range", t, func() {
		s, err := Slice(arr, 1, 5, -1)
		assert.Equal(t, common.ErrInvalidRange, err)
		assert.Nil(t, s)
	})
	convey.Convey("out of range", t, func() {
		s, err := Slice(arr, 20, 100, 1)
		assert.Equal(t, common.ErrOutOfRange, err)
		assert.Nil(t, s)
	})
}

func ExampleSlice() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s, _ := Slice(arr, 5, 2, -1)
	fmt.Println(s)
	// Output: [6 5 4 3]
}
