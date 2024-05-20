package itertools

import "github.com/biocrosscoder/flex/common"

// sliceIterator represents an iterator for iterating over a slice.
type sliceIterator[T any] struct {
	entry   []T
	end     int
	pointer int
	value   T
	step    int
}

// NewSliceIterator creates and returns a new slice iterator.
func NewSliceIterator[T any](entry []T, start, end, step int) ListIterator[T] {
	return &sliceIterator[T]{
		entry:   entry,
		end:     end,
		pointer: start,
		value:   *new(T),
		step:    step,
	}
}

// clear resets the iterator to its initial state.
func (iter *sliceIterator[T]) clear() {
	iter.value = *new(T)
	iter.entry = nil
}

// Next moves the iterator to the next position and returns true if the iterator has not reached the end.
func (iter *sliceIterator[T]) Next() bool {
	if (iter.step > 0 && iter.pointer > iter.end) || (iter.step < 0 && iter.pointer < iter.end) {
		iter.clear()
		return false
	}
	iter.value = iter.entry[iter.pointer]
	iter.pointer += iter.step
	return true
}

// Value returns the current value at the iterator's position.
func (iter *sliceIterator[T]) Value() T {
	return iter.value
}

// Pour iterates through the slice and returns a new slice.
func (iter *sliceIterator[T]) Pour() []T {
	output := make([]T, 0)
	for iter.Next() {
		output = append(output, iter.Value())
	}
	return output
}

// Slice returns a new slice by iterating over the input slice based on the start, end, and step parameters.
func Slice[T any](entry []T, start, end, step int) (slice []T, err error) {
	length := len(entry)
	err = common.CheckRange(start, end, step, length)
	if err != nil {
		return
	}
	slice = NewSliceIterator[T](entry, common.ParseIndex(start, length), common.ParseIndex(end, length), step).Pour()
	return
}

// Reversed returns a new iterator for iterating over a slice in reverse order.
func Reversed[T any](entry []T) ListIterator[T] {
	return NewSliceIterator(entry, len(entry)-1, 0, -1)
}
