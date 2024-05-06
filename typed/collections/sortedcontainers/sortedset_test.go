package sortedcontainers

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderedSet(t *testing.T) {
	s := NewSortedSet(nil, 1, 3, 2, 4, 5, 2, 3, 6, 9)
	size := s.Size()
	convey.Convey("add an element to the set", t, func() {
		s2 := s.Copy()
		assert.Equal(t, s2.Size(), size)
		assert.False(t, s2.Has(-1))
		_ = s2.Add(-1)
		assert.True(t, s2.Has(-1))
		assert.Equal(t, s2.Size(), size+1)
		assert.Equal(t, s2.IndexOf(-1), 0)
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
		expected, err := s2.At(-1)
		assert.Nil(t, err)
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
		_ = s2.Update(s.Set)
		assert.True(t, s.ToList().Equal(s2.ToList()))
	})
}