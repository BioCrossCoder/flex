package orderedcontainers

import (
	"flex/collections/linkedlist"
	"flex/collections/set"
	"flex/common"
)

type OrderedChainSet struct {
	set.Set
	sequence *linkedlist.LinkedList
}

func NewOrderedChainSet(entries ...any) *OrderedChainSet {
	elements := set.Of()
	sequence := linkedlist.NewLinkedList()
	for _, entry := range entries {
		if elements.Has(entry) {
			continue
		}
		_ = elements.Add(entry)
		_ = sequence.Append(entry)
	}
	return &OrderedChainSet{elements, sequence}
}

func (s *OrderedChainSet) Add(element any) *OrderedChainSet {
	if !s.Has(element) {
		_ = s.Set.Add(element)
		_ = s.sequence.Append(element)
	}
	return s
}

func (s *OrderedChainSet) Discard(element any) bool {
	if s.Has(element) {
		_ = s.Set.Discard(element)
		_ = s.sequence.Remove(element)
		return true
	}
	return false
}

func (s *OrderedChainSet) Clear() *OrderedChainSet {
	_ = s.Set.Clear()
	_ = s.sequence.Clear()
	return s
}

func (s *OrderedChainSet) Update(another OrderedChainSet) *OrderedChainSet {
	for _, element := range another.Elements() {
		_ = s.Add(element)
	}
	return s
}

func (s *OrderedChainSet) Pop() (element any, err error) {
	if s.Empty() {
		err = common.ErrEmptySet
		return
	}
	element, err = s.sequence.Pop()
	_ = s.Set.Discard(element)
	return
}

func (s *OrderedChainSet) Elements() []any {
	return s.sequence.ToArray()
}

func (s *OrderedChainSet) Copy() OrderedChainSet {
	newSeq := s.sequence.Copy()
	return OrderedChainSet{
		s.Set.Copy(),
		&newSeq,
	}
}

func (s OrderedChainSet) Equal(another OrderedChainSet) bool {
	if s.Size() != another.Size() {
		return false
	}
	elements1 := s.Elements()
	elements2 := another.Elements()
	for i := 0; i < s.Size(); i++ {
		if elements1[i] != elements2[i] {
			return false
		}
	}
	return true
}

func (s OrderedChainSet) At(index int) (any, error) {
	return s.sequence.At(index)
}

func (s OrderedChainSet) IndexOf(element any) int {
	return s.sequence.IndexOf(element)
}

func (s OrderedChainSet) ToList() linkedlist.LinkedList {
	return s.sequence.Copy()
}