package orderedcontainers

import (
	"encoding/json"
	"flex/common"
	"flex/typed/collections/arraylist"
	"flex/typed/collections/set"
	"fmt"
	"strings"
)

type OrderedSet[T comparable] struct {
	set.Set[T]
	sequence arraylist.ArrayList[T]
}

func NewOrderedSet[T comparable](entries ...T) *OrderedSet[T] {
	elements := set.Of[T]()
	sequence := arraylist.Of[T]()
	for _, entry := range entries {
		if elements.Has(entry) {
			continue
		}
		_ = elements.Add(entry)
		_ = sequence.Push(entry)
	}
	return &OrderedSet[T]{elements, sequence}
}

func (s *OrderedSet[T]) Add(element T) *OrderedSet[T] {
	if !s.Has(element) {
		_ = s.Set.Add(element)
		_ = s.sequence.Push(element)
	}
	return s
}

func (s *OrderedSet[T]) Discard(element T) bool {
	if s.Set.Discard(element) {
		_ = s.sequence.Remove(element)
		return true
	}
	return false
}

func (s *OrderedSet[T]) Clear() *OrderedSet[T] {
	_ = s.Set.Clear()
	_ = s.sequence.Clear()
	return s
}

func (s *OrderedSet[T]) Update(another OrderedSet[T]) *OrderedSet[T] {
	for _, element := range another.Elements() {
		_ = s.Add(element)
	}
	return s
}

func (s *OrderedSet[T]) Pop() (element T, err error) {
	if s.Empty() {
		err = common.ErrEmptySet
		return
	}
	element, err = s.sequence.Pop()
	_ = s.Set.Discard(element)
	return
}

func (s OrderedSet[T]) Elements() []T {
	return s.sequence.Copy()
}

func (s OrderedSet[T]) Copy() OrderedSet[T] {
	return OrderedSet[T]{
		s.Set.Copy(),
		s.sequence.Copy(),
	}
}

func (s OrderedSet[T]) Equal(another OrderedSet[T]) bool {
	if s.Size() != another.Size() {
		return false
	}
	elements1 := s.Elements()
	elements2 := another.Elements()
	for i := 0; i < s.Size(); i++ {
		if !common.Equal(elements1[i], elements2[i]) {
			return false
		}
	}
	return true
}

func (s OrderedSet[T]) At(index int) (T, error) {
	return s.sequence.At(index)
}

func (s OrderedSet[T]) IndexOf(element T) int {
	return s.sequence.IndexOf(element)
}

func (s OrderedSet[T]) ToList() arraylist.ArrayList[T] {
	return s.sequence.Copy()
}

func (s OrderedSet[T]) String() string {
	r := strings.NewReplacer("[", "{", "]", "}")
	return r.Replace(fmt.Sprint(s.sequence))
}

func (s OrderedSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.sequence)
}

func (s *OrderedSet[T]) UnmarshalJSON(data []byte) (err error) {
	var arr []T
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return
	}
	s.Set = set.Of(arr...)
	s.sequence = arraylist.Of(arr...)
	return
}
