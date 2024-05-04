package functools

func All[T any](condition func(T) bool, entry []T) bool {
	for _, item := range entry {
		if !condition(item) {
			return false
		}
	}
	return true
}
