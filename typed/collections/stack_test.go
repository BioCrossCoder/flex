package collections

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	convey.Convey("test stack", t, func() {
		s := NewStack[int](5)
		assert.True(t, s.Empty())
		top, ok := s.Peek()
		assert.False(t, ok)
		assert.Zero(t, top)
		element, ok := s.Pop()
		assert.False(t, ok)
		assert.Zero(t, element)
		for i := 0; i < 5; i++ {
			assert.False(t, s.Full())
			assert.True(t, s.Push(i))
		}
		assert.True(t, s.Full())
		for i := 4; i >= 0; i-- {
			assert.False(t, s.Empty())
			top, ok = s.Peek()
			assert.True(t, ok)
			assert.Equal(t, i, top)
			element, ok = s.Pop()
			assert.True(t, ok)
			assert.Equal(t, i, element)
		}
		assert.True(t, s.Empty())
	})
}
