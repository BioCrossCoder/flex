package set

// Has checks if the set contains the given element.
func (s Set) Has(element any) bool {
	_, ok := s[element]
	return ok
}

// Empty checks if the set is empty.
func (s Set) Empty() bool {
	return s.Size() == 0
}

// IsDisjoint checks if the set is disjoint from another set.
func (s Set) IsDisjoint(another Set) bool {
	for k := range s {
		if another.Has(k) {
			return false
		}
	}
	return true
}

// IsSubset checks if the set is a subset of another set.
func (s Set) IsSubset(another Set) bool {
	for k := range s {
		if !another.Has(k) {
			return false
		}
	}
	return true
}

// IsSuperset checks if the set is a superset of another set.
func (s Set) IsSuperset(another Set) bool {
	for k := range another {
		if !s.Has(k) {
			return false
		}
	}
	return true
}

// Equal checks if the set is equal to another set.
func (s Set) Equal(another Set) bool {
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
