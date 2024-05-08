package functools

import (
	"flex/typed/collections/sortedcontainers/sortedlist"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompare(t *testing.T) {
	convey.Convey("get max/min element", t, func() {
		arr := []int{1, 2, 3, 4, 5}
		assert.Equal(t, 1, Max(sortedlist.DescendOrder, arr...))
		assert.Equal(t, 1, Min(sortedlist.AscendOrder, arr...))
		assert.Equal(t, 5, Max(sortedlist.AscendOrder, arr...))
		assert.Equal(t, 5, Min(sortedlist.DescendOrder, arr...))
	})
}
