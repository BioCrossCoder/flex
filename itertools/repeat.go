package itertools

type repeater struct {
	entry any
}

func Repeat(entry any) *repeater {
	return &repeater{entry}
}

func (r *repeater) Next() any {
	return r.entry
}

func (r *repeater) Repeat(n int) []any {
	output := make([]any, n)
	for i := 0; i < n; i++ {
		output[i] = r.Next()
	}
	return output
}
