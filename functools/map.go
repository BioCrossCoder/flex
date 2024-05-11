package functools

import (
	"github.com/biocrosscoder/flex/common"
	"github.com/biocrosscoder/flex/itertools"
	"reflect"
)

func Map(handler, entry any) (output any, err error) {
	iterator, err := itertools.Map(handler, entry)
	if err != nil {
		return
	}
	output = iterator.Pour()
	return
}

func Maps(handler any, entries ...any) (output []any, err error) {
	entryCount := len(entries)
	if entryCount < 1 {
		err = common.ErrIllegalParamCount
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
