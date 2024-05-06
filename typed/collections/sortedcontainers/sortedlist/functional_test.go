package sortedlist

import (
	"flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReduce(t *testing.T) {
	convey.Convey("reduce list", t, func() {
		l := NewSortedList(nil, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		f := func(x, y int) int {
			return x - y
		}
		convey.Convey("normal reduce", func() {
			result, err := l.Reduce(f)
			assert.Nil(t, err)
			assert.Equal(t, -53, result)
		})
		convey.Convey("reduce with initial value", func() {
			result, err := l.Reduce(f, 10)
			assert.Nil(t, err)
			assert.Equal(t, -45, result)
		})
		convey.Convey("reduce from right", func() {
			result, err := l.ReduceRight(f)
			assert.Nil(t, err)
			assert.Equal(t, -35, result)
		})
		convey.Convey("reduce from right with initial value", func() {
			result, err := l.ReduceRight(f, 10)
			assert.Nil(t, err)
			assert.Equal(t, -45, result)
		})
		convey.Convey("reduce on empty list", func() {
			l := NewSortedList[int](nil)
			result, err := l.Reduce(f)
			assert.Equal(t, common.ErrEmptyList, err)
			assert.Zero(t, result)
			result, err = l.ReduceRight(f)
			assert.Equal(t, common.ErrEmptyList, err)
			assert.Zero(t, result)
		})
		convey.Convey("too many arguments", func() {
			result, err := l.Reduce(f, 1, 2)
			assert.Equal(t, common.ErrTooManyArguments, err)
			assert.Zero(t, result)
			result, err = l.ReduceRight(f, 1, 2)
			assert.Equal(t, common.ErrTooManyArguments, err)
			assert.Zero(t, result)
		})
	})
}

func TestFilter(t *testing.T) {
	convey.Convey("filter list", t, func() {
		l := NewSortedList(DescendOrder, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		f := func(x int) bool {
			return x%2 == 0
		}
		assert.True(t, NewSortedList(DescendOrder, 2, 4, 6, 8, 10).Equal(l.Filter(f)))
	})
}

func TestSomeAndEvery(t *testing.T) {
	convey.Convey("check condition on list", t, func() {
		l := NewSortedList(nil, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		f := func(x int) bool {
			return x > 9
		}
		convey.Convey("at least one element in list satisfies the condition", func() {
			assert.True(t, l.Some(f))
			assert.False(t, l.Some(func(x int) bool {
				return x > 50
			}))
		})
		convey.Convey("all elements in list satisfy the condition", func() {
			assert.False(t, l.Every(f))
			assert.True(t, l.Every(func(x int) bool {
				return x > 0
			}))
		})
	})
}
