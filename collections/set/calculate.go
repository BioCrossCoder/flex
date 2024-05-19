package set

// Difference returns a new set containing elements that are present in the receiver set but not in the another set.
func (s Set) Difference(another Set) Set {
	newSet := make(Set)
	for k := range s {
		if !another.Has(k) {
			newSet[k] = true
		}
	}
	return newSet
}

// Intersection returns a new set containing elements that are present both in the receiver set and another set.
func (s Set) Intersection(another Set) Set {
	newSet := make(Set)
	for k := range s {
		if another.Has(k) {
			newSet[k] = true
		}
	}
	return newSet
}

// Union returns a new set containing all elements from both the receiver set and another set.
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

// SymmetricDifference returns a new set containing elements that are present in either the receiver set or another set, but not in both.
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
