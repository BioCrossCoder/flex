package sortedlist

// Reduce applies a function against an accumulator and each element in the list (from left to right) to reduce it to a single value.
func (l SortedList[T]) Reduce(handler func(T, T) T, initial ...T) (T, error) {
	return l.elements.Reduce(handler, initial...)
}

// ReduceRight applies a function against an accumulator and each element in the list (from right to left) to reduce it to a single value.
func (l SortedList[T]) ReduceRight(handler func(T, T) T, initial ...T) (T, error) {
	return l.elements.ReduceRight(handler, initial...)
}

// Filter creates a new list with all elements that pass the condition implemented by the provided function.
func (l SortedList[T]) Filter(condition func(T) bool) SortedList[T] {
	return SortedList[T]{
		l.elements.Filter(condition),
		l.cmp,
	}
}

// Some checks if at least one element in the list passes the condition implemented by the provided function.
func (l SortedList[T]) Some(condition func(T) bool) bool {
	return l.elements.Some(condition)
}

// Every checks if all elements in the list pass the condition implemented by the provided function.
func (l SortedList[T]) Every(condition func(T) bool) bool {
	return l.elements.Every(condition)
}
