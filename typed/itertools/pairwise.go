package itertools

// pairIterator is a struct type that represents an iterator for generating pairs of elements from a given slice.
type pairIterator[T any] struct {
	entry   []T
	length  int
	pointer int
	value   [2]T
}

// PairWise is a function that creates and returns a new pairIterator for the given input slice.
func PairWise[T any](entry []T) ListIterator[[2]T] {
	return &pairIterator[T]{
		entry:   entry,
		length:  len(entry) - 1,
		pointer: 0,
		value:   *new([2]T),
	}
}

// clear is a method of pairIterator that resets the iterator's value and entry to their zero values.
func (iter *pairIterator[T]) clear() {
	iter.value = *new([2]T)
	iter.entry = nil
}

// Next is a method of pairIterator that advances the iterator to the next pair of elements and returns true if there are more pairs, or false if the end of the input slice is reached.
func (iter *pairIterator[T]) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = [2]T{iter.entry[iter.pointer], iter.entry[iter.pointer+1]}
	iter.pointer++
	return true
}

// Value is a method of pairIterator that returns the current pair of elements.
func (iter *pairIterator[T]) Value() [2]T {
	return iter.value
}

// Pour is a method of pairIterator that iterates through the remaining pairs and returns them as a slice of pairs.
func (iter *pairIterator[T]) Pour() [][2]T {
	length := iter.length - iter.pointer
	output := make([][2]T, length)
	i := 0
	for iter.Next() {
		output[i] = iter.Value()
		i++
	}
	return output
}
