package itertools

type cycler[T any] struct {
	entry   []T
	length  int
	pointer int
	value   T
}

func Cycle[T any](entry []T) ListIterator[T] {
	return &cycler[T]{
		entry:   entry,
		length:  len(entry),
		pointer: -1,
		value:   *new(T),
	}
}

func (iter *cycler[T]) Next() bool {
	iter.pointer = (iter.pointer + 1) % iter.length
	iter.value = iter.entry[iter.pointer]
	return true
}

func (iter *cycler[T]) Value() T {
	return iter.value
}

func (iter *cycler[T]) Pour() []T {
	output := make([]T, iter.length)
	for i := 0; i < iter.length; i++ {
		_ = iter.Next()
		output[i] = iter.Value()
	}
	return output
}
