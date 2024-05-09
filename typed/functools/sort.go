package functools

import "slices"

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

func Sort[T any](entry []T, cmps ...func(a, b T) int) {
	if len(cmps) == 0 {
		return
	}
	f := cmpFunc(cmps...)
	if !slices.IsSortedFunc(entry, f) {
		slices.SortFunc(entry, f)
	}
}

func Sorted[T any](entry []T, cmps ...func(a, b T) int) []T {
	result := slices.Clone(entry)
	Sort(result, cmps...)
	return result
}

func IsSorted[T any](entry []T, cmps ...func(a, b T) int) bool {
	return slices.IsSortedFunc(entry, cmpFunc(cmps...))
}
