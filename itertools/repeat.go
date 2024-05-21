package itertools

// repeater is an iterator that repeats a single value
type repeater struct {
	entry any
}

// Repeat creates a repeater that repeats entry.
func Repeat(entry any) *repeater {
	return &repeater{entry}
}

// Next returns the entry value of the repeater.
func (r *repeater) Next() any {
	return r.entry
}

// Repeat returns a slice of n entries of the repeater.
func (r *repeater) Repeat(n int) []any {
	output := make([]any, n)
	for i := 0; i < n; i++ {
		output[i] = r.Next()
	}
	return output
}
