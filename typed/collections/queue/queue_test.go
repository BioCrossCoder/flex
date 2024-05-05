package queue

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	convey.Convey("test queue", t, func() {
		s := NewQueue[int](5)
		assert.True(t, s.Empty())
		head, ok := s.Peek()
		assert.False(t, ok)
		assert.Zero(t, head)
		element, ok := s.Dequeue()
		assert.False(t, ok)
		assert.Zero(t, element)
		for i := 0; i < 5; i++ {
			assert.False(t, s.Full())
			assert.True(t, s.Enqueue(i))
		}
		assert.True(t, s.Full())
		for i := 0; i < 5; i++ {
			assert.False(t, s.Empty())
			head, ok = s.Peek()
			assert.True(t, ok)
			assert.Equal(t, i, head)
			element, ok = s.Dequeue()
			assert.True(t, ok)
			assert.Equal(t, i, element)
		}
		assert.True(t, s.Empty())
	})
}
