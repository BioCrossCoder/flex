package linkedlist

import (
	"flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchElement(t *testing.T) {
	convey.Convey("search element in deque", t, func() {
		d := NewLinkedList(1, 2, 3, 2, 4)
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
			f := func(num int) bool {
				return num%2 == 0
			}
			v, found := d.Find(f)
			assert.True(t, found)
			assert.Equal(t, 2, v)
			v, found = d.FindLast(f)
			assert.True(t, found)
			assert.Equal(t, 4, v)
			assert.Equal(t, 1, d.FindIndex(f))
			assert.Equal(t, 4, d.FindLastIndex(f))
			assert.Equal(t, []int{1, 3, 4}, d.FindIndexes(f))
			assert.Equal(t, []int{4, 3, 1}, d.FindLastIndexes(f))
			assert.Equal(t, []int{2, 2, 4}, d.Finds(f))
			assert.Equal(t, []int{4, 2, 2}, d.FindLasts(f))
			assert.Equal(t, []int{1, 3}, d.FindIndexes(f, 2))

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
			f := func(num int) bool {
				return num > 5
			}
			v, found := d.Find(f)
			assert.False(t, found)
			assert.Zero(t, v)
			v, found = d.FindLast(f)
			assert.False(t, found)
			assert.Zero(t, v)
			assert.Equal(t, -1, d.FindIndex(f))
			assert.Equal(t, -1, d.FindLastIndex(f))
		})
		convey.Convey("empty deque", func() {
			d := NewLinkedList[any]()
			head, err := d.Head()
			assert.Equal(t, err, common.ErrEmptyList)
			assert.Nil(t, head)
			tail, err := d.Tail()
			assert.Equal(t, err, common.ErrEmptyList)
			assert.Nil(t, tail)
		})
	})
}
