package set

// Difference returns a new set that contains elements from the receiver set that are not in the given set
func (s Set[T]) Difference(another Set[T]) Set[T] {
	newSet := make(Set[T])
	for k := range s {
		if !another.Has(k) {
			newSet[k] = true
		}
	}
	return newSet
}

// Intersection returns a new set that contains elements that are common to both the receiver set and the given set
func (s Set[T]) Intersection(another Set[T]) Set[T] {
	newSet := make(Set[T])
	for k := range s {
		if another.Has(k) {
			newSet[k] = true
		}
	}
	return newSet
}

// Union returns a new set that contains all the unique elements from both the receiver set and the given set
func (s Set[T]) Union(another Set[T]) Set[T] {
	newSet := make(Set[T])
	for k := range s {
		newSet[k] = true
	}
	for k := range another {
		newSet[k] = true
	}
	return newSet
}

// SymmetricDifference returns a new set that contains elements that are in either of the sets, but not in both
func (s Set[T]) SymmetricDifference(another Set[T]) Set[T] {
	newSet := make(Set[T])
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
