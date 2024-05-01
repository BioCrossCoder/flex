package set

import (
	"flex/common"
)

func (s *Set) Add(element any) *Set {
	(*s)[element] = true
	return s
}

func (s *Set) Discard(element any) bool {
	ok := s.Has(element)
	if ok {
		delete(*s, element)
	}
	return ok
}

func (s *Set) Clear() *Set {
	*s = make(Set)
	return s
}

func (s *Set) Update(another Set) *Set {
	count1 := s.Size()
	count2 := another.Size()
	if common.WillReHash(count1, count2) {
		capacity := common.GetMapInitialCapacity(count1 + count2)
		newSet := make(Set, capacity)
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

func (s *Set) Pop() (element any, err error) {
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
