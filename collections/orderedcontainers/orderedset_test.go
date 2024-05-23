package orderedcontainers

import (
	"encoding/json"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderedSet(t *testing.T) {
	s := NewOrderedSet(1, 3, 2, 4, 5, 2, 3, 6, 9)
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
	convey.Convey("jsonify and stringify", t, func() {
		l := NewOrderedSet(1, 2, 3)
		data, err := json.Marshal(&l)
		assert.Nil(t, err)
		l2 := NewOrderedSet()
		assert.Nil(t, json.Unmarshal(data, &l2))
		assert.Equal(t, fmt.Sprint(l), fmt.Sprint(l2))
	})
}

func ExampleOrderedSet() {
	s := NewOrderedSet(1, 3, 2, 4, 5, 2, 3, 6, 9)
	fmt.Println(s)
	// Output: {1 3 2 4 5 6 9}
}

func ExampleOrderedSet_Add() {
	s := NewOrderedSet(1, 2, 3)
	fmt.Println(s)
	s.Add(4)
	fmt.Println(s)
	// Output:
	// {1 2 3}
	// {1 2 3 4}
}

func ExampleOrderedSet_Discard() {
	s := NewOrderedSet(1, 2, 3)
	fmt.Println(s)
	s.Discard(2)
	fmt.Println(s)
	// Output:
	// {1 2 3}
	// {1 3}
}

// ExampleOrderedSet_Pop demonstrates how to pop the last element from the OrderedSet.
func ExampleOrderedSet_Pop() {
	s := NewOrderedSet(1, 2, 3)
	fmt.Println(s)
	element, _ := s.Pop()
	fmt.Println(element)
	fmt.Println(s)
	// Output:
	// {1 2 3}
	// 3
	// {1 2}
}

func ExampleOrderedSet_Update() {
	s := NewOrderedSet(1, 2, 3)
	fmt.Println(s)
	as := NewOrderedSet(4, 3, 5)
	s.Update(*as)
	fmt.Println(s)
	// Output:
	// {1 2 3}
	// {1 2 3 4 5}
}

func ExampleOrderedSet_Elements() {
	s := NewOrderedSet(1, 2, 3)
	fmt.Println(s)
	fmt.Println(s.Elements())
	// Output:
	// {1 2 3}
	// [1 2 3]
}

func ExampleOrderedSet_Copy() {
	s := NewOrderedSet(1, 2, 3)
	fmt.Println(s)
	s2 := s.Copy()
	fmt.Println(s2)
	// Output:
	// {1 2 3}
	// {1 2 3}
}

func ExampleOrderedSet_Equal() {
	s := NewOrderedSet(1, 2, 3)
	s2 := NewOrderedSet(1, 2, 3)
	fmt.Println(s.Equal(*s2))
	s3 := NewOrderedSet(1, 3, 2)
	fmt.Println(s.Equal(*s3))
	// Output:
	// true
	// false
}

func ExampleOrderedSet_At() {
	s := NewOrderedSet(1, 2, 3)
	fmt.Println(s.At(0))
	fmt.Println(s.At(-2))
	fmt.Println(s.At(6))
	// Output:
	// 1 <nil>
	// 2 <nil>
	// <nil> the index is out of range
}

func ExampleOrderedSet_IndexOf() {
	s := NewOrderedSet(1, 2, 3)
	fmt.Println(s.IndexOf(2))
	fmt.Println(s.IndexOf(4))
	// Output:
	// 1
	// -1
}

func ExampleOrderedSet_ToList() {
	s := NewOrderedSet(1, 2, 3)
	fmt.Println(s)
	fmt.Println(s.ToList())
	// Output:
	// {1 2 3}
	// [1 2 3]
}

func ExampleOrderedSet_MarshalJSON() {
	s := NewOrderedSet(1, 2, 3)
	data, _ := json.Marshal(&s)
	fmt.Println(string(data))
	// Output: [1,2,3]
}

func ExampleOrderedSet_UnmarshalJSON() {
	s := NewOrderedSet()
	_ = json.Unmarshal([]byte("[1,2,3]"), &s)
	fmt.Println(s)
	// Output: {1 2 3}
}
