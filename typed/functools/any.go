package functools

func Any[T any](condition func(T) bool, entry []T) bool {
	for _, item := range entry {
		if condition(item) {
			return true
		}
	}
	return false
}
