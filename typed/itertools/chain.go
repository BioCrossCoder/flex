package itertools

type chainIterator[T any] struct {
	entries    [][]T
	length     int
	pointer    int
	subPointer int
	value      T
}

func Chain[T any](entries ...[]T) ListIterator[T] {
	return &chainIterator[T]{
		entries:    entries,
		length:     len(entries),
		pointer:    0,
		subPointer: 0,
		value:      *new(T),
	}
}

func (iter *chainIterator[T]) clear() {
	iter.entries = nil
	iter.value = *new(T)
}

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

func (iter *chainIterator[T]) Value() T {
	return iter.value
}

func (iter *chainIterator[T]) Pour() []T {
	output := make([]T, 0)
	for iter.Next() {
		output = append(output, iter.Value())
	}
	return output
}
