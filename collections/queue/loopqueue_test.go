package queue

import (
	"flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoopQueue(t *testing.T) {
	convey.Convey("test loop queue", t, func() {
		q, err := NewLoopQueue(5)
		assert.Nil(t, err)
		assert.True(t, q.Empty())
		head, ok := q.Peek()
		assert.False(t, ok)
		assert.Nil(t, head)
		element, ok := q.Dequeue()
		assert.False(t, ok)
		assert.Nil(t, element)
		for i := 0; i < 10; i++ {
			if !q.Full() {
				assert.True(t, q.Enqueue(i))
				continue
			}
			head, ok = q.Peek()
			assert.True(t, ok)
			assert.Equal(t, i-5, head)
			v, ok := q.Dequeue()
			assert.True(t, ok)
			assert.Equal(t, head, v)
			assert.True(t, q.Enqueue(i))
		}
	})
	convey.Convey("create loop queue failed", t, func() {
		q, err := NewLoopQueue(-1)
		assert.Equal(t, err, common.ErrInvalidCapacity)
		assert.Nil(t, q)
	})
}
