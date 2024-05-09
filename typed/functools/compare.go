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

func Equals[T any](cmp func(T, T) int, v ...T) bool {
	l := len(v)
	if l <= 1 {
		return true
	}
	for i := 1; i < l; i++ {
		if !Equal(v[i], v[i-1], cmp) {
			return false
		}
	}
	return true
}

func Equal[T any](a, b T, cmp func(T, T) int) bool {
	return cmp(a, b) == 0
}

func Less[T any](a, b T, cmp func(T, T) int) bool {
	return cmp(a, b) < 0
}

func Greater[T any](a, b T, cmp func(T, T) int) bool {
	return cmp(a, b) > 0
}

func IsIncreasing[T any](entry []T, cmp func(T, T) int, strict bool) bool {
	for i := 1; i < len(entry); i++ {
		if (strict && !Greater(entry[i], entry[i-1], cmp)) || Less(entry[i], entry[i-1], cmp) {
			return false
		}
	}
	return true
}

func IsDecreasing[T any](entry []T, cmp func(T, T) int, strict bool) bool {
	for i := 1; i < len(entry); i++ {
		if (strict && !Less(entry[i], entry[i-1], cmp)) || Greater(entry[i], entry[i-1], cmp) {
			return false
		}
	}
	return true
}

func IsMonotonic[T any](entry []T, cmp func(T, T) int, strict bool) bool {
	return IsIncreasing(entry, cmp, strict) || IsDecreasing(entry, cmp, strict)
}
