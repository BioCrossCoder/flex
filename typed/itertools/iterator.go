package itertools

type ListIterator[T any] interface {
	Next() bool
	Value() T
	Pour() []T
}

type listIterator[T any] struct {
	entry   []T
	length  int
	pointer int
	value   T
}

func NewListIterator[T any](entry []T) ListIterator[T] {
	return &listIterator[T]{
		entry:   entry,
		length:  len(entry),
		pointer: 0,
		value:   *new(T),
	}
}

func (iter *listIterator[T]) clear() {
	iter.value = *new(T)
	iter.entry = nil
}

func (iter *listIterator[T]) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.entry[iter.pointer]
	iter.pointer++
	return true
}

func (iter *listIterator[T]) Value() T {
	return iter.value
}

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
