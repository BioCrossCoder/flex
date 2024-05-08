package itertools

type pairIterator[T any] struct {
	entry   []T
	length  int
	pointer int
	value   [2]T
}

func PairWise[T any](entry []T) ListIterator[[2]T] {
	return &pairIterator[T]{
		entry:   entry,
		length:  len(entry) - 1,
		pointer: 0,
		value:   *new([2]T),
	}
}

func (iter *pairIterator[T]) clear() {
	iter.value = *new([2]T)
	iter.entry = nil
}

func (iter *pairIterator[T]) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = [2]T{iter.entry[iter.pointer], iter.entry[iter.pointer+1]}
	iter.pointer++
	return true
}

func (iter *pairIterator[T]) Value() [2]T {
	return iter.value
}

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
