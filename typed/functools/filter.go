package functools

func Filter[T any](condition func(T) bool, entry []T) []T {
	output := make([]T, 0)
	for _, v := range entry {
		if condition(v) {
			output = append(output, v)
		}
	}
	return output
}
