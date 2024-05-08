package itertools

type repeater[T any] struct {
	entry   T
	pointer int
	length  int
}

func Repeat[T any](entry T, count int) ListIterator[T] {
	return &repeater[T]{
		entry,
		0,
		count,
	}
}

func (r *repeater[T]) Next() bool {
	if r.pointer == r.length {
		return false
	}
	r.pointer++
	return true
}

func (r *repeater[T]) Value() T {
	return r.entry
}

func (r *repeater[T]) Pour() []T {
	output := make([]T, r.length-r.pointer)
	i := 0
	for r.Next() {
		output[i] = r.Value()
		i++
	}
	return output
}
