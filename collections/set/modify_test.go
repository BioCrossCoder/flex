package set

import (
	"flex/common"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	convey.Convey("add an element to the set", t, func() {
		s := newTestSet()
		_ = s.Discard(6)
		assert.False(t, s.Has(6))
		oldSize := s.Size()
		_ = s.Add(6)
		assert.True(t, s.Has(6))
		newSize := s.Size()
		assert.Equal(t, oldSize+1, newSize)
		_ = s.Add(6)
		assert.Equal(t, newSize, s.Size())
	})
}

func TestDisCard(t *testing.T) {
	convey.Convey("discard an element from the set", t, func() {
		s := newTestSet()
		_ = s.Add(6)
		assert.True(t, s.Has(6))
		oldSize := s.Size()
		_ = s.Discard(6)
		assert.False(t, s.Has(6))
		newSize := s.Size()
		assert.Equal(t, oldSize-1, newSize)
		_ = s.Discard(6)
		assert.Equal(t, newSize, s.Size())
	})
}

func TestClear(t *testing.T) {
	convey.Convey("remove all elements from the set", t, func() {
		s := newTestSet()
		for i := 0; i < 10; i++ {
			s.Add(i)
		}
		assert.True(t, s.Size() >= 10)
		_ = s.Clear()
		assert.True(t, s.Empty())
	})
}

func TestUpdate(t *testing.T) {
	convey.Convey("update the set with another set", t, func() {
		s := newTestSet()
		_ = s.Clear()
		s2 := Of(1, 2, 3)
		_ = s.Update(s2)
		assert.Equal(t, 3, s.Size())
		assert.Equal(t, s, s2)
	})
}

func TestPop(t *testing.T) {
	convey.Convey("pop an element from the set", t, func() {
		s := Of(1, 2, 3, 4, 5, 6)
		element, err := s.Pop()
		assert.Nil(t, err)
		assert.False(t, s.Has(element))
	})
	convey.Convey("pop an element from an empty set", t, func() {
		s := Set{}
		element, err := s.Pop()
		assert.Equal(t, common.ErrEmptySet, err)
		assert.Nil(t, element)
	})
}
