// Package functools provides functional programming tools.
package functools

import (
	"github.com/biocrosscoder/flex/common"
	"github.com/biocrosscoder/flex/itertools"
	"reflect"
)

// Map applies a conversion function to each element of an input sequence and returns the output list containing the converted elements.
func Map(handler, entry any) (output any, err error) {
	iterator, err := itertools.Map(handler, entry)
	if err != nil {
		return
	}
	output = iterator.Pour()
	return
}

// Maps applies a conversion function to each element of multiple input lists, each element in the returned output list is converted from elements with the same index in all input lists by the handler function.
func Maps(handler any, entries ...any) (output []any, err error) {
	entryCount := len(entries)
	if entryCount < 1 {
		err = common.ErrUnexpectedParamCount
		return
	}
	entryLength := 0
	for _, entry := range entries {
		err = common.IsList(entry)
		if err != nil {
			return
		}
		length := reflect.ValueOf(entry).Len()
		if entryLength == 0 {
			entryLength = length
			continue
		}
		if length != entryLength {
			err = common.ErrListLengthMismatch
			return
		}
	}
	err = common.IsInputFuncValid(handler, entryCount, 1)
	if err != nil {
		return
	}
	output = make([]any, entryLength)
	for i := 0; i < entryLength; i++ {
		args := make([]reflect.Value, entryCount)
		for j := 0; j < entryCount; j++ {
			args[j] = reflect.ValueOf(entries[j]).Index(i)
		}
		output[i] = reflect.ValueOf(handler).Call(args)[0].Interface()
	}
	return
}
