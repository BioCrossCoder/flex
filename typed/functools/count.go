package functools

import (
	"github.com/biocrosscoder/flex/common"
)

// Count function takes a slice of elements of type T and a value of type T, then returns the count of occurrences of the value in the slice.
func Count[T any](entry []T, value T) (count int) {
	for _, v := range entry {
		if common.Equal(v, value) {
			count++
		}
	}
	return
}

// CountBy function takes a slice of elements of type T and a condition function that takes an element of type T and returns a boolean, then returns the count of elements that satisfy the condition.
func CountBy[T any](entry []T, condition func(T) bool) (count int) {
	for _, v := range entry {
		if condition(v) {
			count++
		}
	}
	return
}
