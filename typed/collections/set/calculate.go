package set

func (s Set[T]) Difference(another Set[T]) Set[T] {
	newSet := make(Set[T])
	for k := range s {
		if !another.Has(k) {
			newSet[k] = true
		}
	}
	return newSet
}

func (s Set[T]) Intersection(another Set[T]) Set[T] {
	newSet := make(Set[T])
	for k := range s {
		if another.Has(k) {
			newSet[k] = true
		}
	}
	return newSet
}

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
