package set

// Has checks if the set contains the given element and returns true if found, otherwise false.
func (s Set[T]) Has(element T) bool {
	_, ok := s[element]
	return ok
}

// Empty returns true if the set is empty, otherwise false.
func (s Set[T]) Empty() bool {
	return s.Size() == 0
}

// IsDisjoint checks if the set is disjoint with another set and returns true if they have no common elements, otherwise false.
func (s Set[T]) IsDisjoint(another Set[T]) bool {
	for k := range s {
		if another.Has(k) {
			return false
		}
	}
	return true
}

// IsSubset checks if the set is a subset of another set and returns true if every element in the set is also in the other set, otherwise false.
func (s Set[T]) IsSubset(another Set[T]) bool {
	for k := range s {
		if !another.Has(k) {
			return false
		}
	}
	return true
}

// IsSuperset checks if the set is a superset of another set and returns true if it contains every element of the other set, otherwise false.
func (s Set[T]) IsSuperset(another Set[T]) bool {
	for k := range another {
		if !s.Has(k) {
			return false
		}
	}
	return true
}

// Equal checks if the set is equal to another set by comparing the size and elements of the sets, and returns true if they are equal, otherwise false.
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
