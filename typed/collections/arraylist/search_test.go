package arraylist

import (
	"flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchElement(t *testing.T) {
	convey.Convey("search element in list", t, func() {
		l := ArrayList[int]{1, 2, 3, 2, 4}
		convey.Convey("search by index", func() {
			assert.Equal(t, 1, l.IndexOf(2))
			assert.Equal(t, 3, l.LastIndexOf(2))
		})
		convey.Convey("get element by index", func() {
			num, err := l.At(-3)
			assert.Nil(t, err)
			assert.Equal(t, l[2], num)
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
			assert.Equal(t, []int{1, 3, 4}, l.FindIndexes(f))
			assert.Equal(t, []int{4, 3, 1}, l.FindLastIndexes(f))
			assert.Equal(t, []int{2, 2, 4}, l.Finds(f))
			assert.Equal(t, []int{4, 2, 2}, l.FindLasts(f))
			assert.Equal(t, []int{1, 3}, l.FindIndexes(f, 2))

		})
		convey.Convey("get first/last element", func() {
			head, err := l.Head()
			assert.Nil(t, err)
			assert.Equal(t, l[0], head)
			tail, err := l.Tail()
			assert.Nil(t, err)
			assert.Equal(t, l[l.Len()-1], tail)
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
			l := ArrayList[any]{}
			head, err := l.Head()
			assert.Equal(t, err, common.ErrEmptyList)
			assert.Nil(t, head)
			tail, err := l.Tail()
			assert.Equal(t, err, common.ErrEmptyList)
			assert.Nil(t, tail)
		})
	})
}
