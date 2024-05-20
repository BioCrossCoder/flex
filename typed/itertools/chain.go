package itertools

// chainIterator represents an iterator for iterating over a chain of slices.
type chainIterator[T any] struct {
	entries    [][]T
	length     int
	pointer    int
	subPointer int
	value      T
}

// Chain returns a ListIterator for iterating over the given slices of type T.
func Chain[T any](entries ...[]T) ListIterator[T] {
	return &chainIterator[T]{
		entries:    entries,
		length:     len(entries),
		pointer:    0,
		subPointer: 0,
		value:      *new(T),
	}
}

// clear resets the chainIterator by clearing its entries and value.
func (iter *chainIterator[T]) clear() {
	iter.entries = nil
	iter.value = *new(T)
}

// Next moves the iterator to the next element in the chain and returns true if successful.
func (iter *chainIterator[T]) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.entries[iter.pointer][iter.subPointer]
	iter.subPointer++
	if iter.subPointer == len(iter.entries[iter.pointer]) {
		iter.subPointer = 0
		iter.pointer++
	}
	return true
}

// Value returns the current value of the iterator.
func (iter *chainIterator[T]) Value() T {
	return iter.value
}

// Pour iterates through the entire chain and returns all the values in a new slice.
func (iter *chainIterator[T]) Pour() []T {
	output := make([]T, 0)
	for iter.Next() {
		output = append(output, iter.Value())
	}
	return output
}
