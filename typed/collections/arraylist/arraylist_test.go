package arraylist

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLen(t *testing.T) {
	convey.Convey("check length", t, func() {
		l := ArrayList[int]{1, 2, 3}
		assert.Equal(t, len(l), l.Len())
	})
}

func TestCount(t *testing.T) {
	convey.Convey("count specific element", t, func() {
		l := ArrayList[int]{1, 2, 3, 2, 1, 4, 5, 4}
		assert.Equal(t, 2, l.Count(2))
		assert.Equal(t, 2, l.Count(4))
		assert.Equal(t, 0, l.Count(6))
	})
}

func TestIncludes(t *testing.T) {
	convey.Convey("check if element is included", t, func() {
		l := ArrayList[int]{1, 2, 3, 4, 5}
		assert.True(t, l.Includes(3))
		assert.False(t, l.Includes(6))
	})
}

func TestEmpty(t *testing.T) {
	convey.Convey("check if list is empty", t, func() {
		l1 := ArrayList[int]{1, 2, 3}
		l2 := ArrayList[any]{}
		assert.False(t, l1.Empty())
		assert.True(t, l2.Empty())
	})
}
