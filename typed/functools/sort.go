package functools

import "slices"

func Sort[T any](entry []T, cmps ...func(a, b T) int) {
	for _, cmp := range cmps {
		if !slices.IsSortedFunc(entry, cmp) {
			slices.SortStableFunc(entry, cmp)
		}
	}
}

func Sorted[T any](entry []T, cmps ...func(a, b T) int) []T {
	result := slices.Clone(entry)
	Sort(result, cmps...)
	return result
}
