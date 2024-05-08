package itertools

func DropWhile[T any](condition func(T) bool, entry []T) ListIterator[T] {
	start := 0
	length := len(entry)
	for start < length && condition(entry[start]) {
		start++
	}
	return &listIterator[T]{entry, length, start, *new(T)}
}

func TakeWhile[T any](condition func(T) bool, entry []T) ListIterator[T] {
	end := 0
	length := len(entry)
	for (end < length) && condition(entry[end]) {
		end++
	}
	return &listIterator[T]{entry, end, 0, *new(T)}
}
