package itertools

type Accumulator[T any] interface {
	Next() bool
	Value() T
	Pour() T
}

type accumulator[T any] struct {
	entry   []T
	length  int
	handler func(T, T) T
	pointer int
	value   T
}

func Accumulate[T any](entry []T, handler func(T, T) T) Accumulator[T] {
	return &accumulator[T]{
		entry:   entry,
		length:  len(entry),
		handler: handler,
		pointer: 1,
		value:   entry[0],
	}
}

func (iter *accumulator[T]) clear() {
	iter.value = *new(T)
	iter.entry = nil
	iter.handler = nil
}

func (iter *accumulator[T]) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.handler(iter.value, iter.entry[iter.pointer])
	iter.pointer++
	return true
}

func (iter *accumulator[T]) Value() T {
	return iter.value
}

func (iter *accumulator[T]) Pour() T {
	result := iter.Value()
	for iter.Next() {
		result = iter.Value()
	}
	return result
}
