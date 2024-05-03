package linkedlist

import (
	"flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchElement(t *testing.T) {
	convey.Convey("search element in deque", t, func() {
		d := NewDeque(1, 2, 3, 2, 4)
		convey.Convey("search by index", func() {
			assert.Equal(t, 1, d.IndexOf(2))
			assert.Equal(t, 3, d.LastIndexOf(2))
		})
		convey.Convey("get element by index", func() {
			num, err := d.At(-3)
			assert.Nil(t, err)
			assert.Equal(t, 3, num)
		})
		convey.Convey("search by condition", func() {
			f := func(num any) bool {
				return num.(int)%2 == 0
			}
			assert.Equal(t, 2, d.Find(f))
			assert.Equal(t, 4, d.FindLast(f))
			assert.Equal(t, 1, d.FindIndex(f))
			assert.Equal(t, 4, d.FindLastIndex(f))
		})
		convey.Convey("get first/last element", func() {
			head, err := d.Head()
			assert.Nil(t, err)
			assert.Equal(t, 1, head)
			tail, err := d.Tail()
			assert.Nil(t, err)
			assert.Equal(t, 4, tail)
		})
		convey.Convey("element not found", func() {
			assert.Equal(t, -1, d.IndexOf(5))
			assert.Equal(t, -1, d.LastIndexOf(5))
			f := func(num any) bool {
				return num.(int) > 5
			}
			assert.Nil(t, d.Find(f))
			assert.Nil(t, d.FindLast(f))
			assert.Equal(t, -1, d.FindIndex(f))
			assert.Equal(t, -1, d.FindLastIndex(f))
		})
		convey.Convey("empty deque", func() {
			d := NewDeque()
			head, err := d.Head()
			assert.Equal(t, err, common.ErrEmptyList)
			assert.Nil(t, head)
			tail, err := d.Tail()
			assert.Equal(t, err, common.ErrEmptyList)
			assert.Nil(t, tail)
		})
	})
}
