package set

func (s Set) Difference(another Set) Set {
	newSet := make(Set)
	for k := range s {
		if !another.Has(k) {
			newSet[k] = true
		}
	}
	return newSet
}

func (s Set) Intersection(another Set) Set {
	newSet := make(Set)
	for k := range s {
		if another.Has(k) {
			newSet[k] = true
		}
	}
	return newSet
}

func (s Set) Union(another Set) Set {
	newSet := make(Set)
	for k := range s {
		newSet[k] = true
	}
	for k := range another {
		newSet[k] = true
	}
	return newSet
}

func (s Set) SymmetricDifference(another Set) Set {
	newSet := make(Set)
	for k := range s {
		if !another.Has(k) {
			newSet[k] = true
		}
	}
	for k := range another {
		if !s.Has(k) {
			newSet[k] = true
		}
	}
	return newSet
}
