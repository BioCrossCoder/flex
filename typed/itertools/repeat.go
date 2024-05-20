package itertools

// repeater is an struct that implements the ListIterator interface to iterate over a given element 'count' number of times.
type repeater[T any] struct {
	entry   T
	pointer int
	length  int
}

// Repeat creates a new repeater instance to iterate over the given element 'count' number of times.
func Repeat[T any](entry T, count int) ListIterator[T] {
	return &repeater[T]{
		entry,
		0,
		count,
	}
}

// Next moves the iterator to the next element in the list.
func (r *repeater[T]) Next() bool {
	if r.pointer == r.length {
		return false
	}
	r.pointer++
	return true
}

// Value returns the current element the iterator is pointing to.
func (r *repeater[T]) Value() T {
	return r.entry
}

// Pour returns the remaining elements in the list from the current position of the iterator.
func (r *repeater[T]) Pour() []T {
	output := make([]T, r.length-r.pointer)
	i := 0
	for r.Next() {
		output[i] = r.Value()
		i++
	}
	return output
}
