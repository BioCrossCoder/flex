package functools

// Any checks if any element in the given slice satisfies the condition.
func Any[T any](condition func(T) bool, entry []T) bool {
	for _, item := range entry {
		if condition(item) {
			return true
		}
	}
	return false
}
