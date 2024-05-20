package itertools

// DropWhile takes a condition function and a slice of type T, and returns a ListIterator[T] that iterates over the input slice after dropping elements while the condition is true.
func DropWhile[T any](condition func(T) bool, entry []T) ListIterator[T] {
	start := 0
	length := len(entry)
	for start < length && condition(entry[start]) {
		start++
	}
	return &listIterator[T]{entry, length, start, *new(T)}
}

// TakeWhile takes a condition function and a slice of type T, and returns a ListIterator[T] that iterates over the input slice only while the condition is true.
func TakeWhile[T any](condition func(T) bool, entry []T) ListIterator[T] {
	end := 0
	length := len(entry)
	for (end < length) && condition(entry[end]) {
		end++
	}
	return &listIterator[T]{entry, end, 0, *new(T)}
}
