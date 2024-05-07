package orderedcontainers

import (
	"encoding/json"
	"flex/common"
	"flex/typed/collections/linkedlist"
	"flex/typed/collections/set"
	"strings"
)

type OrderedChainSet[T comparable] struct {
	set.Set[T]
	sequence *linkedlist.LinkedList[T]
}

func NewOrderedChainSet[T comparable](entries ...T) *OrderedChainSet[T] {
	elements := set.Of[T]()
	sequence := linkedlist.NewLinkedList[T]()
	for _, entry := range entries {
		if elements.Has(entry) {
			continue
		}
		_ = elements.Add(entry)
		_ = sequence.Append(entry)
	}
	return &OrderedChainSet[T]{elements, sequence}
}

func (s *OrderedChainSet[T]) Add(element T) *OrderedChainSet[T] {
	if !s.Has(element) {
		_ = s.Set.Add(element)
		_ = s.sequence.Append(element)
	}
	return s
}

func (s *OrderedChainSet[T]) Discard(element T) bool {
	if s.Has(element) {
		_ = s.Set.Discard(element)
		_ = s.sequence.Remove(element)
		return true
	}
	return false
}

func (s *OrderedChainSet[T]) Clear() *OrderedChainSet[T] {
	_ = s.Set.Clear()
	_ = s.sequence.Clear()
	return s
}

func (s *OrderedChainSet[T]) Update(another OrderedChainSet[T]) *OrderedChainSet[T] {
	for _, element := range another.Elements() {
		_ = s.Add(element)
	}
	return s
}

func (s *OrderedChainSet[T]) Pop() (element T, err error) {
	if s.Empty() {
		err = common.ErrEmptySet
		return
	}
	element, err = s.sequence.Pop()
	_ = s.Set.Discard(element)
	return
}

func (s *OrderedChainSet[T]) Elements() []T {
	return s.sequence.ToArray()
}

func (s *OrderedChainSet[T]) Copy() OrderedChainSet[T] {
	newSeq := s.sequence.Copy()
	return OrderedChainSet[T]{
		s.Set.Copy(),
		&newSeq,
	}
}

func (s OrderedChainSet[T]) Equal(another OrderedChainSet[T]) bool {
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

func (s OrderedChainSet[T]) At(index int) (T, error) {
	return s.sequence.At(index)
}

func (s OrderedChainSet[T]) IndexOf(element T) int {
	return s.sequence.IndexOf(element)
}

func (s OrderedChainSet[T]) ToList() linkedlist.LinkedList[T] {
	return s.sequence.Copy()
}

func (s OrderedChainSet[T]) String() string {
	r := strings.NewReplacer("[", "{", "]", "}")
	return r.Replace(s.sequence.String())
}

func (s OrderedChainSet[T]) MarshalJSON() ([]byte, error) {
	return s.sequence.MarshalJSON()
}

func (s *OrderedChainSet[T]) UnmarshalJSON(data []byte) (err error) {
	var arr []T
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return
	}
	s.Set = set.Of(arr...)
	s.sequence = linkedlist.NewLinkedList(arr...)
	return
}
