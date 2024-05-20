package functools

// All checks whether all elements in the given array satisfy the specified condition.
func All[T any](condition func(T) bool, entry []T) bool {
	for _, item := range entry {
		if !condition(item) {
			return false
		}
	}
	return true
}
