package sortedlist

import (
	"flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchElement(t *testing.T) {
	convey.Convey("search element in list", t, func() {
		l := NewSortedList(AscendOrder, 1, 2, 3, 2, 4)
		convey.Convey("search by index", func() {
			assert.Equal(t, 1, l.IndexOf(2))
			assert.Equal(t, 2, l.LastIndexOf(2))
		})
		convey.Convey("get element by index", func() {
			num, err := l.At(-3)
			assert.Nil(t, err)
			assert.Equal(t, 2, num)
		})
		convey.Convey("search by condition", func() {
			f := func(num int) bool {
				return num%2 == 0
			}
			v, found := l.Find(f)
			assert.True(t, found)
			assert.Equal(t, 2, v)
			v, found = l.FindLast(f)
			assert.True(t, found)
			assert.Equal(t, 4, v)
			assert.Equal(t, 1, l.FindIndex(f))
			assert.Equal(t, 4, l.FindLastIndex(f))
		})
		convey.Convey("get max/min element", func() {
			head, err := l.Head()
			assert.Nil(t, err)
			minValue, err := l.Min()
			assert.Nil(t, err)
			assert.Equal(t, minValue, head)
			tail, err := l.Tail()
			assert.Nil(t, err)
			maxValue, err := l.Max()
			assert.Nil(t, err)
			assert.Equal(t, maxValue, tail)
		})
		convey.Convey("element not found", func() {
			assert.Equal(t, -1, l.IndexOf(0))
			assert.Equal(t, -1, l.LastIndexOf(0))
			f := func(num int) bool {
				return num < 0
			}
			v, found := l.Find(f)
			assert.False(t, found)
			assert.Zero(t, v)
			v, found = l.FindLast(f)
			assert.False(t, found)
			assert.Zero(t, v)
			assert.Equal(t, -1, l.FindIndex(f))
			assert.Equal(t, -1, l.FindLastIndex(f))
		})
		convey.Convey("empty list", func() {
			l := NewSortedList[int](nil)
			head, err := l.Head()
			assert.Equal(t, err, common.ErrEmptyList)
			assert.Zero(t, head)
			tail, err := l.Tail()
			assert.Equal(t, err, common.ErrEmptyList)
			assert.Zero(t, tail)
		})
	})
}
