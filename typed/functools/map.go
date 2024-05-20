// Package functools provides functional programming tools.
package functools

import "github.com/biocrosscoder/flex/common"

// Map applies the given handler function to each element in the input slice and returns a new slice of results.
func Map[E, R any](handler func(E) R, entry []E) []R {
	output := make([]R, len(entry))
	for i, item := range entry {
		output[i] = handler(item)
	}
	return output
}

// Maps applies the given handler function to corresponding elements in the two input slices and returns a new slice of results,
// as well as an error if the input slices have different lengths.
func Maps[T, U, R any](handler func(T, U) R, entry1 []T, entry2 []U) (output []R, err error) {
	length := len(entry1)
	if length != len(entry2) {
		err = common.ErrListLengthMismatch
		return
	}
	output = make([]R, length)
	for i := 0; i < length; i++ {
		output[i] = handler(entry1[i], entry2[i])
	}
	return
}
