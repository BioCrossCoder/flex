package itertools

// Accumulator is an interface for iterating over a sequence of values and accumulating them using a function.
type Accumulator[T any] interface {
	// Next advances the iterator and returns true if there are more elements.
	Next() bool
	// Value returns the current accumulated value.
	Value() T
	// Pour returns the final accumulated value by iterating over all elements.
	Pour() T
}

// Define a struct accumulator implementing the Accumulator interface.
type accumulator[T any] struct {
	entry   []T
	length  int
	handler func(T, T) T
	pointer int
	value   T
}

// Accumulate function returns a new Accumulator instance.
func Accumulate[T any](entry []T, handler func(T, T) T) Accumulator[T] {
	return &accumulator[T]{
		entry:   entry,
		length:  len(entry),
		handler: handler,
		pointer: 1,
		value:   entry[0],
	}
}

// Clear method resets the accumulator state by clearing its fields.
func (iter *accumulator[T]) clear() {
	iter.value = *new(T)
	iter.entry = nil
	iter.handler = nil
}

// Next method advances the pointer and updates the accumulated value.
func (iter *accumulator[T]) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.handler(iter.value, iter.entry[iter.pointer])
	iter.pointer++
	return true
}

// Value method returns the current accumulated value.
func (iter *accumulator[T]) Value() T {
	return iter.value
}

// Pour method returns the final accumulated value by iterating over all elements.
func (iter *accumulator[T]) Pour() T {
	result := iter.Value()
	for iter.Next() {
		result = iter.Value()
	}
	return result
}
