package sortedlist

import (
	"flex/typed/collections/arraylist"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedList(t *testing.T) {
	list := arraylist.Of(1, 3, 2, 5, 6, 9)
	convey.Convey("create a sorted list from a slice", t, func() {
		l := NewSortedList(AscendOrder, list...)
		assert.False(t, l.Empty())
		assert.Equal(t, 6, l.Len())
		assert.Equal(t, []int{1, 2, 3, 5, 6, 9}, l.ToArray())
		for _, v := range list {
			assert.True(t, l.Includes(v))
			assert.Equal(t, list.Count(v), l.Count(v))
		}
	})
	convey.Convey("copy a sorted list from a sorted list", t, func() {
		l := NewSortedList(DescendOrder, list...)
		assert.Equal(t, []int{9, 6, 5, 3, 2, 1}, l.Copy().ToArray())
		assert.True(t, l.Slice(-1, -10, -1).Equal(l.ToReversed()))
	})
}
