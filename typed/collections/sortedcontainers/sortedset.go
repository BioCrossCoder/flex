package sortedcontainers

import (
	"cmp"
	"github.com/biocrosscoder/flex/common"
	"github.com/biocrosscoder/flex/typed/collections/set"
	"github.com/biocrosscoder/flex/typed/collections/sortedcontainers/sortedlist"
)

// SortedSet represents a set of elements sorted according to the specified comparison function.
type SortedSet[T cmp.Ordered] struct {
	set.Set[T]
	sequence sortedlist.SortedList[T]
}

// NewSortedSet creates a new SortedSet instance with the given comparison function and initial entries.
func NewSortedSet[T cmp.Ordered](f func(a, b T) int, entries ...T) *SortedSet[T] {
	elements := set.Of(entries...)
	if f == nil {
		f = sortedlist.AscendOrder
	}
	sequence := sortedlist.NewSortedList(f)
	for element := range elements {
		_ = sequence.Insert(element)
	}
	return &SortedSet[T]{elements, *sequence}
}

// Add adds the given element to the SortedSet if it does not already exist.
func (s *SortedSet[T]) Add(element T) *SortedSet[T] {
	if !s.Has(element) {
		_ = s.Set.Add(element)
		_ = s.sequence.Insert(element)
	}
	return s
}

// Discard removes the specified element from the SortedSet.
func (s *SortedSet[T]) Discard(element T) bool {
	if s.Set.Discard(element) {
		_ = s.sequence.Remove(element)
		return true
	}
	return false
}

// Clear removes all elements from the SortedSet.
func (s *SortedSet[T]) Clear() *SortedSet[T] {
	_ = s.Set.Clear()
	_ = s.sequence.Clear()
	return s
}

// Update adds all elements from the specified set to the SortedSet.
func (s *SortedSet[T]) Update(another set.Set[T]) *SortedSet[T] {
	for element := range another {
		_ = s.Add(element)
	}
	return s
}

// Pop removes and returns the first element from the SortedSet.
func (s *SortedSet[T]) Pop() (element T, err error) {
	if s.Empty() {
		err = common.ErrEmptySet
		return
	}
	element, err = s.sequence.Pop()
	_ = s.Set.Discard(element)
	return
}

// Elements returns all elements in the SortedSet as a slice.
func (s SortedSet[T]) Elements() []T {
	return s.sequence.ToArray()
}

// Copy creates a shallow copy of the SortedSet.
func (s SortedSet[T]) Copy() SortedSet[T] {
	return SortedSet[T]{
		s.Set.Copy(),
		s.sequence.Copy(),
	}
}

// Equal checks if the SortedSet is equal to the specified SortedSet.
func (s SortedSet[T]) Equal(another SortedSet[T]) bool {
	return s.sequence.Equal(another.sequence)
}

// At returns the element at the specified index in the SortedSet's sorted sequence.
func (s SortedSet[T]) At(index int) (T, error) {
	return s.sequence.At(index)
}

// IndexOf returns the index of the specified element in the SortedSet's sorted sequence.
func (s SortedSet[T]) IndexOf(element T) int {
	return s.sequence.IndexOf(element)
}

// ToList returns a copy of the SortedSet's sorted sequence as a SortedList.
func (s SortedSet[T]) ToList() sortedlist.SortedList[T] {
	return s.sequence.Copy()
}
