package collections

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	convey.Convey("test stack", t, func() {
		s := NewStack(5)
		assert.True(t, s.Empty())
		top, ok := s.Peek()
		assert.False(t, ok)
		assert.Nil(t, top)
		element, ok := s.Pop()
		assert.False(t, ok)
		assert.Nil(t, element)
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

func ExampleStack() {
	s := NewStack(3)
	for i := 0; i < 3; i++ {
		s.Push(i)
	}
	for i := 0; i < 3; i++ {
		fmt.Println(s.Pop())
	}
	// Output:
	// 2 true
	// 1 true
	// 0 true
}
