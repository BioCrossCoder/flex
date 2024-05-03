package arraylist

import (
	"flex/common"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestSearchElement(t *testing.T) {
	convey.Convey("search element in list", t, func() {
		l := ArrayList{1, 2, 3, 2, 4}
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
			f := func(num any) bool {
				return num.(int)%2 == 0
			}
			assert.Equal(t, 2, l.Find(f))
			assert.Equal(t, 4, l.FindLast(f))
			assert.Equal(t, 1, l.FindIndex(f))
			assert.Equal(t, 4, l.FindLastIndex(f))
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
			f := func(num any) bool {
				return num.(int) < 0
			}
			assert.Nil(t, l.Find(f))
			assert.Nil(t, l.FindLast(f))
			assert.Equal(t, -1, l.FindIndex(f))
			assert.Equal(t, -1, l.FindLastIndex(f))
		})
		convey.Convey("empty list", func() {
			l := ArrayList{}
			head, err := l.Head()
			assert.Equal(t, err, common.ErrEmptyList)
			assert.Nil(t, head)
			tail, err := l.Tail()
			assert.Equal(t, err, common.ErrEmptyList)
			assert.Nil(t, tail)
		})
	})
}
