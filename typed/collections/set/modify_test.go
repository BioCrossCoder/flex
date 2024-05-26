package set

import (
	"fmt"
	"github.com/biocrosscoder/flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
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
		s := Set[int]{}
		element, err := s.Pop()
		assert.Equal(t, common.ErrEmptySet, err)
		assert.Zero(t, element)
	})
}

func ExampleSet_Add() {
	s := Of[int]()
	fmt.Println(s.Has(1))
	s.Add(1)
	fmt.Println(s.Has(1))
	// Output:
	// false
	// true
}

func ExampleSet_Discard() {
	s := Of(1, 2, 3)
	fmt.Println(s.Has(2))
	s.Discard(2)
	fmt.Println(s.Has(2))
	// Output:
	// true
	// false
}

func ExampleSet_Clear() {
	s := Of(1, 2)
	fmt.Println(s.Empty())
	s.Clear()
	fmt.Println(s.Empty())
	fmt.Println(s)
	// Output:
	// false
	// true
	// {}
}

func ExampleSet_Update() {
	s1 := Of(1)
	fmt.Println(s1)
	s2 := Of(2)
	fmt.Println(s2)
	s1.Update(s2)
	fmt.Println(s1.Has(2))
	// Output:
	// {1}
	// {2}
	// true
}

func ExampleSet_Pop() {
	s := Of(1)
	element, _ := s.Pop()
	fmt.Println(element)
	// Output: 1
}
