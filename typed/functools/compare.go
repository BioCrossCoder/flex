package functools

import "slices"

// Max finds the maximum element in the given slice using the custom comparison function cmp.
func Max[T any](cmp func(T, T) int, v ...T) T {
	if len(v) == 0 {
		return *new(T)
	}
	return slices.MaxFunc(v, cmp)
}

// Min finds the minimum element in the given slice using the custom comparison function cmp.
func Min[T any](cmp func(T, T) int, v ...T) T {
	if len(v) == 0 {
		return *new(T)
	}
	return slices.MinFunc(v, cmp)
}

// Equals checks if all elements in the slice are equal based on the custom comparison function cmp.
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

// Equal checks if two elements are equal based on the custom comparison function cmp.
func Equal[T any](a, b T, cmp func(T, T) int) bool {
	return cmp(a, b) == 0
}

// Less checks if a is less than b based on the custom comparison function cmp.
func Less[T any](a, b T, cmp func(T, T) int) bool {
	return cmp(a, b) < 0
}

// Greater checks if a is greater than b based on the custom comparison function cmp.
func Greater[T any](a, b T, cmp func(T, T) int) bool {
	return cmp(a, b) > 0
}

// IsIncreasing checks if the elements in the slice are in increasing order based on the custom comparison function cmp.
func IsIncreasing[T any](entry []T, cmp func(T, T) int, strict bool) bool {
	for i := 1; i < len(entry); i++ {
		if (strict && !Greater(entry[i], entry[i-1], cmp)) || Less(entry[i], entry[i-1], cmp) {
			return false
		}
	}
	return true
}

// IsDecreasing checks if the elements in the slice are in decreasing order based on the custom comparison function cmp.
func IsDecreasing[T any](entry []T, cmp func(T, T) int, strict bool) bool {
	for i := 1; i < len(entry); i++ {
		if (strict && !Less(entry[i], entry[i-1], cmp)) || Greater(entry[i], entry[i-1], cmp) {
			return false
		}
	}
	return true
}

// IsMonotonic checks if the elements in the slice are either increasing or decreasing based on the custom comparison function cmp.
func IsMonotonic[T any](entry []T, cmp func(T, T) int, strict bool) bool {
	return IsIncreasing(entry, cmp, strict) || IsDecreasing(entry, cmp, strict)
}
