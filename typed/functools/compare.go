package functools

import "slices"

func Max[T any](cmp func(T, T) int, v ...T) T {
	if len(v) == 0 {
		return *new(T)
	}
	return slices.MaxFunc(v, cmp)
}

func Min[T any](cmp func(T, T) int, v ...T) T {
	if len(v) == 0 {
		return *new(T)
	}
	return slices.MinFunc(v, cmp)
}
