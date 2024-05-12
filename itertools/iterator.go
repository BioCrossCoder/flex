// Package itertools provides iterator functions to create iterators and perform common operations on iterables.
package itertools

// Iterator is an interface that defines the behavior of an iterator.
type Iterator interface {
	Next() bool
	Value() any
	Pour() any
}

// listIterator is an implementation of the Iterator interface for a slice of any type.
type listIterator struct {
	entry   []any
	length  int
	pointer int
	value   any
}

// NewListIterator creates a new Iterator for a slice of any type.
func NewListIterator(entry []any) Iterator {
	return &listIterator{
		entry:   entry,
		length:  len(entry),
		pointer: 0,
		value:   nil,
	}
}

// clear releases the resources used by the iterator.
func (iter *listIterator) clear() {
	iter.value = nil
	iter.entry = nil
}

// Next moves the pointer to the next element and returns true if there is a next element or false otherwise.
func (iter *listIterator) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.entry[iter.pointer]
	iter.pointer++
	return true
}

// Value returns the element pointed by the pointer.
func (iter *listIterator) Value() any {
	return iter.value
}

// Pour returns a slice of all remaining elements in the iterator.
func (iter *listIterator) Pour() any {
	length := iter.length - iter.pointer
	output := make([]any, length)
	i := 0
	for iter.Next() {
		output[i] = iter.Value()
		i++
	}
	return output
}
