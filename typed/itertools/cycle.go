package itertools

// cycler is a generic type representing an iterator that cycles through a slice of elements of any type.
type cycler[T any] struct {
	entry   []T
	length  int
	pointer int
	value   T
}

// Cycle creates and returns a ListIterator using the given slice as the initial entries.
func Cycle[T any](entry []T) ListIterator[T] {
	return &cycler[T]{
		entry:   entry,
		length:  len(entry),
		pointer: -1,
		value:   *new(T),
	}
}

// Next updates the iterator to the next element in the slice and returns true.
func (iter *cycler[T]) Next() bool {
	iter.pointer = (iter.pointer + 1) % iter.length
	iter.value = iter.entry[iter.pointer]
	return true
}

// Value returns the current value pointed to by the iterator.
func (iter *cycler[T]) Value() T {
	return iter.value
}

// Pour cycles through the iterator and returns a new slice containing all the elements in the order they were visited.
func (iter *cycler[T]) Pour() []T {
	output := make([]T, iter.length)
	for i := 0; i < iter.length; i++ {
		_ = iter.Next()
		output[i] = iter.Value()
	}
	return output
}
