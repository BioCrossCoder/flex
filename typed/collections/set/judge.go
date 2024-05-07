package set

func (s Set[T]) Has(element T) bool {
	_, ok := s[element]
	return ok
}

func (s Set[T]) Empty() bool {
	return s.Size() == 0
}

func (s Set[T]) IsDisjoint(another Set[T]) bool {
	for k := range s {
		if another.Has(k) {
			return false
		}
	}
	return true
}

func (s Set[T]) IsSubset(another Set[T]) bool {
	for k := range s {
		if !another.Has(k) {
			return false
		}
	}
	return true
}

func (s Set[T]) IsSuperset(another Set[T]) bool {
	for k := range another {
		if !s.Has(k) {
			return false
		}
	}
	return true
}

func (s Set[T]) Equal(another Set[T]) bool {
	if s.Size() != another.Size() {
		return false
	}
	for k := range s {
		if !another.Has(k) {
			return false
		}
	}
	return true
}
