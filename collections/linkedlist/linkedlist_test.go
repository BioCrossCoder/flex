package linkedlist

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLen(t *testing.T) {
	convey.Convey("check length", t, func() {
		d := NewDeque(1, 2, 3, 4, 5)
		assert.Equal(t, 5, d.Len())
	})
}

func TestCount(t *testing.T) {
	convey.Convey("count specific element", t, func() {
		d := NewDeque(1, 2, 3, 2, 1, 4, 5, 4)
		assert.Equal(t, 2, d.Count(2))
		assert.Equal(t, 2, d.Count(4))
		assert.Equal(t, 0, d.Count(6))
	})
}

func TestIncludes(t *testing.T) {
	convey.Convey("check if deque includes element", t, func() {
		d := NewDeque(1, 2, 3, 2, 1, 4, 5, 4)
		assert.True(t, d.Includes(2))
		assert.False(t, d.Includes(6))
	})
}

func TestEmpty(t *testing.T) {
	convey.Convey("check if deque is empty", t, func() {
		d1 := NewDeque()
		d2 := NewDeque(1, 2, 3, 4, 5)
		assert.True(t, d1.Empty())
		assert.False(t, d2.Empty())
	})
}

func TestToList(t *testing.T) {
	convey.Convey("convert deque to list", t, func() {
		d := NewDeque(1, 2, 3, 4, 5)
		assert.Equal(t, []any{1, 2, 3, 4, 5}, d.ToArray())
	})
}
