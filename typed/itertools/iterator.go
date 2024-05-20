// Package itertools provides iterator functions to create iterators and perform common operations on iterables.
package itertools

// ListIterator is an interface for iterating over a list of elements
type ListIterator[T any] interface {
	// Next moves the iterator to the next element in the list and returns true if there is a next element, false otherwise
	Next() bool
	// Value returns the current value of the element the iterator is pointing to
	Value() T
	// Pour iterates over the remaining elements in the list and returns them as a slice
	Pour() []T
}

// listIterator is a struct that implements the ListIterator interface to iterate over a list of elements
type listIterator[T any] struct {
	entry   []T
	length  int
	pointer int
	value   T
}

// NewListIterator creates a new ListIterator for the given list of elements
func NewListIterator[T any](entry []T) ListIterator[T] {
	return &listIterator[T]{
		entry:   entry,
		length:  len(entry),
		pointer: 0,
		value:   *new(T),
	}
}

// clear resets the state of the iterator
func (iter *listIterator[T]) clear() {
	iter.value = *new(T)
	iter.entry = nil
}

// Next moves the iterator to the next element in the list and returns true if there is a next element, false otherwise
func (iter *listIterator[T]) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.entry[iter.pointer]
	iter.pointer++
	return true
}

// Value returns the current value of the element the iterator is pointing to
func (iter *listIterator[T]) Value() T {
	return iter.value
}

// Pour iterates over the remaining elements in the list and returns them as a slice
func (iter *listIterator[T]) Pour() []T {
	length := iter.length - iter.pointer
	output := make([]T, length)
	i := 0
	for iter.Next() {
		output[i] = iter.Value()
		i++
	}
	return output
}
