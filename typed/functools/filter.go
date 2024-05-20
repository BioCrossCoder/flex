package functools

// Filter is a generic function that filters elements of type T based on the provided condition function.
// It takes a condition function and a slice of type T as input, and returns a new slice containing elements that satisfy the condition.
func Filter[T any](condition func(T) bool, entry []T) []T {
	output := make([]T, 0)
	for _, v := range entry {
		if condition(v) {
			output = append(output, v)
		}
	}
	return output
}
