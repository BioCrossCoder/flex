package sortedcontainers

import (
	"cmp"
	"github.com/biocrosscoder/flex/common"
	"github.com/biocrosscoder/flex/typed/collections/set"
	"github.com/biocrosscoder/flex/typed/collections/sortedcontainers/sortedlist"
)

type SortedSet[T cmp.Ordered] struct {
	set.Set[T]
	sequence sortedlist.SortedList[T]
}

func NewSortedSet[T cmp.Ordered](cmp func(a, b T) int, entries ...T) *SortedSet[T] {
	elements := set.Of(entries...)
	if cmp == nil {
		cmp = sortedlist.AscendOrder
	}
	sequence := sortedlist.NewSortedList(cmp)
	for element := range elements {
		_ = sequence.Insert(element)
	}
	return &SortedSet[T]{elements, *sequence}
}

func (s *SortedSet[T]) Add(element T) *SortedSet[T] {
	if !s.Has(element) {
		_ = s.Set.Add(element)
		_ = s.sequence.Insert(element)
	}
	return s
}

func (s *SortedSet[T]) Discard(element T) bool {
	if s.Set.Discard(element) {
		_ = s.sequence.Remove(element)
		return true
	}
	return false
}

func (s *SortedSet[T]) Clear() *SortedSet[T] {
	_ = s.Set.Clear()
	_ = s.sequence.Clear()
	return s
}

func (s *SortedSet[T]) Update(another set.Set[T]) *SortedSet[T] {
	for element := range another {
		_ = s.Add(element)
	}
	return s
}

func (s *SortedSet[T]) Pop() (element T, err error) {
	if s.Empty() {
		err = common.ErrEmptySet
		return
	}
	element, err = s.sequence.Pop()
	_ = s.Set.Discard(element)
	return
}

func (s SortedSet[T]) Elements() []T {
	return s.sequence.ToArray()
}

func (s SortedSet[T]) Copy() SortedSet[T] {
	return SortedSet[T]{
		s.Set.Copy(),
		s.sequence.Copy(),
	}
}

func (s SortedSet[T]) Equal(another SortedSet[T]) bool {
	return s.sequence.Equal(another.sequence)
}

func (s SortedSet[T]) At(index int) (T, error) {
	return s.sequence.At(index)
}

func (s SortedSet[T]) IndexOf(element T) int {
	return s.sequence.IndexOf(element)
}

func (s SortedSet[T]) ToList() sortedlist.SortedList[T] {
	return s.sequence.Copy()
}
