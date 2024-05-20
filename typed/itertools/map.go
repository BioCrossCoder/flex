package itertools

// convertor struct is used to convert elements of type E to type R
type convertor[E, R any] struct {
	entry   []E
	length  int
	handler func(E) R
	pointer int
	value   R
}

// Map is a function that takes a handler function and a slice of elements of type E,
// and returns a ListIterator of type R
func Map[E, R any](handler func(E) R, entry []E) ListIterator[R] {
	return &convertor[E, R]{
		entry:   entry,
		length:  len(entry),
		handler: handler,
		pointer: 0,
		value:   *new(R),
	}
}

// clear is a method of convertor struct that resets its internal values and handler
func (iter *convertor[E, R]) clear() {
	iter.value = *new(R)
	iter.entry = nil
	iter.handler = nil
}

// Next is a method of convertor struct that moves to the next element and performs the conversion
func (iter *convertor[E, R]) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.handler(iter.entry[iter.pointer])
	iter.pointer++
	return true
}

// Value is a method of convertor struct that returns the current converted value
func (iter *convertor[E, R]) Value() R {
	return iter.value
}

// Pour is a method of convertor struct that iterates through the remaining elements,
// converts them, and returns a slice of converted values
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
