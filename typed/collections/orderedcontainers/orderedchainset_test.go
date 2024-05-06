package orderedcontainers

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderedChainSet(t *testing.T) {
	s := NewOrderedChainSet(1, 3, 2, 4, 5, 2, 3, 6, 9)
	size := s.Size()
	convey.Convey("add an element to the set", t, func() {
		s2 := s.Copy()
		assert.Equal(t, s2.Size(), size)
		assert.False(t, s2.Has(-1))
		_ = s2.Add(-1)
		assert.True(t, s2.Has(-1))
		assert.Equal(t, s2.Size(), size+1)
		assert.Equal(t, s2.IndexOf(-1), size)
		element, err := s2.At(size)
		assert.Nil(t, err)
		assert.Equal(t, element, -1)
	})
	convey.Convey("remove an element from the set", t, func() {
		s2 := s.Copy()
		assert.Equal(t, s2.Size(), size)
		assert.True(t, s2.Has(9))
		assert.True(t, s2.Discard(9))
		assert.False(t, s2.Has(9))
		assert.Equal(t, s2.Size(), size-1)
	})
	convey.Convey("pop an element from the set", t, func() {
		s2 := s.Copy()
		assert.Equal(t, s2.Size(), size)
		expected := s2.Elements()[s2.Size()-1]
		element, err := s2.Pop()
		assert.Nil(t, err)
		assert.Equal(t, element, expected)
		assert.Equal(t, s2.Size(), size-1)
	})
	convey.Convey("clear and update a set", t, func() {
		s2 := s.Copy()
		assert.True(t, s.Equal(s2))
		assert.False(t, s2.Empty())
		_ = s2.Clear()
		assert.True(t, s2.Empty())
		_ = s2.Update(*s)
		assert.Equal(t, s2.ToList(), s.ToList())
	})
}