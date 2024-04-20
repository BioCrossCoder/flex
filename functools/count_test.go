package functools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCount(t *testing.T) {
	convey.Convey("count in slice", t, func() {
		entry := []int{1, 1, 1, 2, 3, 5, 6, 1}
		result, err := Count(entry, 1)
		assert.Nil(t, err)
		assert.Equal(t, 4, result)
	})
	convey.Convey("count in array", t, func() {
		entry := [8]int{1, 0, 1, 2, 3, 5, 6, 1}
		result, err := Count(entry, 1)
		assert.Nil(t, err)
		assert.Equal(t, 3, result)
	})
	convey.Convey("count in string", t, func() {
		entry := "hello world"
		result, err := Count(entry, "l")
		assert.Nil(t, err)
		assert.Equal(t, 3, result)
	})
	convey.Convey("count in map", t, func() {
		entry := map[string]int{"a": 1, "b": 2, "c": 3, "d": 1}
		result, err := Count(entry, 1)
		assert.Nil(t, err)
		assert.Equal(t, 2, result)
	})
}
