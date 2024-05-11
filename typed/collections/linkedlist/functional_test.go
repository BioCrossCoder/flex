package linkedlist

import (
	"github.com/biocrosscoder/flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	convey.Convey("mapping deque", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		f := func(x int) int {
			return x * 3
		}
		assert.True(t, NewLinkedList(3, 6, 9, 12, 15, 18, 21, 24, 27, 30).Equal(d.Map(f)))
	})
}

func TestReduce(t *testing.T) {
	convey.Convey("reduce deque", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		f := func(x, y int) int {
			return x - y
		}
		convey.Convey("normal reduce", func() {
			result, err := d.Reduce(f)
			assert.Nil(t, err)
			assert.Equal(t, -53, result)
		})
		convey.Convey("reduce with initial value", func() {
			result, err := d.Reduce(f, 100)
			assert.Nil(t, err)
			assert.Equal(t, 45, result)
		})
		convey.Convey("reduce from right", func() {
			result, err := d.ReduceRight(f)
			assert.Nil(t, err)
			assert.Equal(t, -35, result)
		})
		convey.Convey("reduce from right with initial value", func() {
			result, err := d.ReduceRight(f, 100)
			assert.Nil(t, err)
			assert.Equal(t, 45, result)
		})
		convey.Convey("reduce on empty deque", func() {
			d := NewLinkedList[int]()
			result, err := d.Reduce(f)
			assert.Equal(t, common.ErrEmptyList, err)
			assert.Zero(t, result)
			result, err = d.ReduceRight(f)
			assert.Equal(t, common.ErrEmptyList, err)
			assert.Zero(t, result)
		})
		convey.Convey("too many arguments", func() {
			result, err := d.Reduce(f, 1, 2)
			assert.Equal(t, common.ErrTooManyArguments, err)
			assert.Zero(t, result)
			result, err = d.ReduceRight(f, 1, 2)
			assert.Equal(t, common.ErrTooManyArguments, err)
			assert.Zero(t, result)
		})
	})
}

func TestFilter(t *testing.T) {
	convey.Convey("filter deque", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		f := func(x int) bool {
			return x > 5
		}
		assert.True(t, NewLinkedList(6, 7, 8, 9, 10).Equal(d.Filter(f)))
	})
}

func TestSomeAndAny(t *testing.T) {
	convey.Convey("check condition on deque", t, func() {
		d := NewLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		f := func(x int) bool {
			return x > 5
		}
		convey.Convey("at least one element in deque satisfies the condition", func() {
			assert.True(t, d.Some(f))
			assert.False(t, d.Some(func(x int) bool {
				return x > 10
			}))
		})
		convey.Convey("all elements in deque satisfy the condition", func() {
			assert.False(t, d.Every(f))
			assert.True(t, d.Every(func(x int) bool {
				return x > 0
			}))
		})
	})
}
