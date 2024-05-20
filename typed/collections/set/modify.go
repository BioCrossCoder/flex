package set

import (
	"github.com/biocrosscoder/flex/common"
)

// Add inserts the specified element into the set and returns a pointer to the modified set.
func (s *Set[T]) Add(element T) *Set[T] {
	(*s)[element] = true
	return s
}

// Discard removes the specified element from the set and returns true if the element was present, false otherwise.
func (s *Set[T]) Discard(element T) bool {
	ok := s.Has(element)
	if ok {
		delete(*s, element)
	}
	return ok
}

// Clear removes all elements from the set and returns a pointer to the empty set.
func (s *Set[T]) Clear() *Set[T] {
	*s = make(Set[T])
	return s
}

// Update merges the elements from the given set 'another' into the current set and returns a pointer to the modified set.
func (s *Set[T]) Update(another Set[T]) *Set[T] {
	count1 := s.Size()
	count2 := another.Size()
	if common.WillReHash(count1, count2) {
		capacity := common.GetMapInitialCapacity(count1 + count2)
		newSet := make(Set[T], capacity)
		for k := range *s {
			newSet.Add(k)
		}
		*s = newSet
	}
	for k := range another {
		s.Add(k)
	}
	return s
}

// Pop removes and returns an arbitrary element from the set, along with an error if the set is empty.
func (s *Set[T]) Pop() (element T, err error) {
	if s.Empty() {
		err = common.ErrEmptySet
		return
	}
	for k := range *s {
		element = k
		_ = s.Discard(k)
		break
	}
	return
}
