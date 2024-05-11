package functools

import (
	"github.com/biocrosscoder/flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReduce(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	f := func(a, b int) int {
		return a + b
	}
	total := 1 + 2 + 3 + 4 + 5
	convey.Convey("normal reduce", t, func() {
		actual, err := Reduce(f, arr)
		assert.Nil(t, err)
		assert.Equal(t, total, actual)
	})
	convey.Convey("reduce with inital value", t, func() {
		expected := total - 2
		actual, err := Reduce(f, arr, -2)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	convey.Convey("reduce failed", t, func() {
		_, err := Reduce(f, nil)
		assert.Equal(t, err, common.ErrEmptyList)
		_, err = Reduce(f, arr, 1, 2)
		assert.Equal(t, err, common.ErrTooManyArguments)
	})
}
