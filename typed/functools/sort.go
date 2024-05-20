package functools

import "slices"

// cmpFunc creates a composite comparison function by combining multiple comparison functions.
func cmpFunc[T any](cmps ...func(a, b T) int) func(a, b T) int {
	return func(a, b T) (v int) {
		for _, cmp := range cmps {
			v = cmp(a, b)
			if v != 0 {
				break
			}
		}
		return
	}
}

// Sort sorts the input slice using the provided comparison functions.
func Sort[T any](entry []T, cmps ...func(a, b T) int) {
	if len(cmps) == 0 {
		return
	}
	f := cmpFunc(cmps...)
	if !slices.IsSortedFunc(entry, f) {
		slices.SortFunc(entry, f)
	}
}

// Sorted returns a new sorted slice without modifying the original input slice.
func Sorted[T any](entry []T, cmps ...func(a, b T) int) []T {
	result := slices.Clone(entry)
	Sort(result, cmps...)
	return result
}

// IsSorted checks if the input slice is sorted according to the provided comparison functions.
func IsSorted[T any](entry []T, cmps ...func(a, b T) int) bool {
	return slices.IsSortedFunc(entry, cmpFunc(cmps...))
}
