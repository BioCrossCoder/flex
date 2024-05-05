package set

import (
	"flex/common"
)

func (s *Set[T]) Add(element T) *Set[T] {
	(*s)[element] = true
	return s
}

func (s *Set[T]) Discard(element T) bool {
	ok := s.Has(element)
	if ok {
		delete(*s, element)
	}
	return ok
}

func (s *Set[T]) Clear() *Set[T] {
	*s = make(Set[T])
	return s
}

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
