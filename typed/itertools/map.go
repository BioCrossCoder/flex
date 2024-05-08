package itertools

type convertor[E, R any] struct {
	entry   []E
	length  int
	handler func(E) R
	pointer int
	value   R
}

func Map[E, R any](handler func(E) R, entry []E) ListIterator[R] {
	return &convertor[E, R]{
		entry:   entry,
		length:  len(entry),
		handler: handler,
		pointer: 0,
		value:   *new(R),
	}
}

func (iter *convertor[E, R]) clear() {
	iter.value = *new(R)
	iter.entry = nil
	iter.handler = nil
}

func (iter *convertor[E, R]) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.handler(iter.entry[iter.pointer])
	iter.pointer++
	return true
}

func (iter *convertor[E, R]) Value() R {
	return iter.value
}

func (iter *convertor[E, R]) Pour() []R {
	length := iter.length - iter.pointer
	output := make([]R, length)
	i := 0
	for iter.Next() {
		output[i] = iter.Value()
		i++
	}
	return output
}
