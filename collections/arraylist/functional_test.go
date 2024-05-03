package arraylist

import (
	"flex/common"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	convey.Convey("mapping list", t, func() {
		l := ArrayList{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		f := func(x any) any {
			return x.(int) * 2
		}
		assert.Equal(t, ArrayList{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, l.Map(f))
	})
}

func TestReduce(t *testing.T) {
	convey.Convey("reduce list", t, func() {
		l := ArrayList{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		f := func(x, y any) any {
			if ix, ok := x.(int); ok {
				x = float64(ix)
			}
			return x.(float64) / float64(y.(int))
		}
		convey.Convey("normal reduce", func() {
			result, err := l.Reduce(f)
			assert.Nil(t, err)
			assert.Equal(t, 2.7557319223985894e-07, result)
		})
		convey.Convey("reduce with initial value", func() {
			result, err := l.Reduce(f, 10)
			assert.Nil(t, err)
			assert.Equal(t, 2.7557319223985897e-06, result)
		})
		convey.Convey("reduce from right", func() {
			result, err := l.ReduceRight(f)
			assert.Nil(t, err)
			assert.Equal(t, 2.7557319223985893e-05, result)
		})
		convey.Convey("reduce from right with initial value", func() {
			result, err := l.ReduceRight(f, 10)
			assert.Nil(t, err)
			assert.Equal(t, 2.755731922398589e-06, result)
		})
		convey.Convey("reduce on empty list", func() {
			l := ArrayList{}
			result, err := l.Reduce(f)
			assert.Equal(t, common.ErrEmptyList, err)
			assert.Nil(t, result)
			result, err = l.ReduceRight(f)
			assert.Equal(t, common.ErrEmptyList, err)
			assert.Nil(t, result)
		})
		convey.Convey("too many arguments", func() {
			result, err := l.Reduce(f, 1, 2)
			assert.Equal(t, common.ErrTooManyArguments, err)
			assert.Nil(t, result)
			result, err = l.ReduceRight(f, 1, 2)
			assert.Equal(t, common.ErrTooManyArguments, err)
			assert.Nil(t, result)
		})
	})
}

func TestFilter(t *testing.T) {
	convey.Convey("filter list", t, func() {
		l := ArrayList{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		f := func(x any) bool {
			return x.(int)%2 == 0
		}
		assert.Equal(t, ArrayList{2, 4, 6, 8, 10}, l.Filter(f))
	})
}

func TestSomeAndEvery(t *testing.T) {
	convey.Convey("check condition on list", t, func() {
		l := ArrayList{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		f := func(x any) bool {
			return x.(int) > 9
		}
		convey.Convey("at least one element in list satisfies the condition", func() {
			assert.True(t, l.Some(f))
			assert.False(t, l.Some(func(x any) bool {
				return x.(int) > 50
			}))
		})
		convey.Convey("all elements in list satisfy the condition", func() {
			assert.False(t, l.Every(f))
			assert.True(t, l.Every(func(x any) bool {
				return x.(int) > 0
			}))
		})
	})
}
