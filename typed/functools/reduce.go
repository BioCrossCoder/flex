package functools

import "github.com/biocrosscoder/flex/common"

// Reduce applies a binary function to the elements of an array to reduce them to a single value.
func Reduce[T any](handler func(T, T) T, entry []T, initial ...T) (result T, err error) {
	entryLength := len(entry)
	if entryLength == 0 {
		err = common.ErrEmptyList
		return
	}
	argCount := len(initial)
	if argCount > 1 {
		err = common.ErrTooManyArguments
		return
	}
	i := 0
	if argCount == 0 {
		result = entry[0]
		i++
	} else {
		result = initial[0]
	}
	for i < entryLength {
		result = handler(result, entry[i])
		i++
	}
	return
}
