package sortedlist

func (l SortedList[T]) Reduce(handler func(T, T) T, initial ...T) (T, error) {
	return l.elements.Reduce(handler, initial...)
}

func (l SortedList[T]) ReduceRight(handler func(T, T) T, initial ...T) (T, error) {
	return l.elements.ReduceRight(handler, initial...)
}

func (l SortedList[T]) Filter(condition func(T) bool) SortedList[T] {
	return SortedList[T]{
		l.elements.Filter(condition),
		l.cmp,
	}
}

func (l SortedList[T]) Some(condition func(T) bool) bool {
	return l.elements.Some(condition)
}

func (l SortedList[T]) Every(condition func(T) bool) bool {
	return l.elements.Every(condition)
}
