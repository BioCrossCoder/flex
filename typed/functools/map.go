package functools

import "flex/common"

func Map[E, R any](handler func(E) R, entry []E) []R {
	output := make([]R, len(entry))
	for i, item := range entry {
		output[i] = handler(item)
	}
	return output
}

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
