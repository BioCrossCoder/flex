package functools

import (
	"flex/typed/collections/sortedcontainers/sortedlist"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	convey.Convey("test sort", t, func() {
		a := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		b := Sorted(a, sortedlist.AscendOrder)
		Sort(a, sortedlist.AscendOrder)
		assert.Equal(t, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}, a)
		assert.Equal(t, a, b)
	})
}
