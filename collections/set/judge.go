package set

func (s Set) Has(element any) bool {
	_, ok := s[element]
	return ok
}

func (s Set) Empty() bool {
	return s.Size() == 0
}

func (s Set) IsDisjoint(another Set) bool {
	for k := range s {
		if another.Has(k) {
			return false
		}
	}
	return true
}

func (s Set) IsSubset(another Set) bool {
	for k := range s {
		if !another.Has(k) {
			return false
		}
	}
	return true
}

func (s Set) IsSuperset(another Set) bool {
	for k := range another {
		if !s.Has(k) {
			return false
		}
	}
	return true
}

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
