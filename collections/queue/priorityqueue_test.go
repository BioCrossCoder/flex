package queue

import (
	"fmt"
	"github.com/biocrosscoder/flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	convey.Convey("test priority queue", t, func() {
		s, err := NewPriorityQueue(5)
		assert.Nil(t, err)
		assert.True(t, s.Empty())
		head, ok := s.Peek()
		assert.False(t, ok)
		assert.Nil(t, head)
		element, ok := s.Dequeue()
		assert.False(t, ok)
		assert.Nil(t, element)
		for i := 0; i < 5; i++ {
			assert.False(t, s.Full())
			assert.True(t, s.Enqueue(i, i))
		}
		assert.True(t, s.Full())
		for i := 4; i >= 0; i-- {
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
	convey.Convey("create priority queue failed", t, func() {
		pq, err := NewPriorityQueue(0)
		assert.Equal(t, err, common.ErrInvalidCapacity)
		assert.Nil(t, pq)
	})
}

func ExamplePriorityQueue() {
	pq, _ := NewPriorityQueue(3)
	pq.Enqueue(100, 3)
	pq.Enqueue(80, 22)
	pq.Enqueue(75, 1)
	v, _ := pq.Dequeue()
	fmt.Println(v)
	v, _ = pq.Dequeue()
	fmt.Println(v)
	v, _ = pq.Dequeue()
	fmt.Println(v)
	// Output:
	// 80
	// 100
	// 75
}
