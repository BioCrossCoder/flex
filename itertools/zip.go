package itertools

import (
	"flex/common"
	"reflect"
)

type zipIterator struct {
	entries [][]any
	length  int
	pointer int
	value   any
	count   int
}

func NewZipIterator(entries [][]any, length int) common.Iterator {
	return &zipIterator{
		entries: entries,
		length:  length,
		pointer: 0,
		value:   nil,
		count:   len(entries),
	}
}

func (iter *zipIterator) clear() {
	iter.entries = nil
	iter.value = nil
}

func (iter *zipIterator) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	values := make([]any, iter.count)
	for i := 0; i < iter.count; i++ {
		values[i] = iter.entries[i][iter.pointer]
	}
	iter.value = values
	iter.pointer++
	return true
}

func (iter *zipIterator) Value() any {
	return iter.value
}

func (iter *zipIterator) Pour() any {
	length := iter.length - iter.pointer
	output := make([]any, length)
	i := 0
	for iter.Next() {
		output[i] = iter.Value()
		i++
	}
	return output
}

func Zip(entries ...any) (iterator common.Iterator, err error) {
	entryCount := len(entries)
	if entryCount < 2 {
		err = common.ErrIllegalParamCount
		return
	}
	entryLength := -1
	iterEntries := make([][]any, entryCount)
	for i, entry := range entries {
		err = common.IsSequence(entry)
		if err != nil {
			return
		}
		value := reflect.ValueOf(entry)
		length := value.Len()
		if entryLength == -1 {
			entryLength = length
		} else if length < entryLength {
			entryLength = length
		}
		if value.Kind() == reflect.String {
			iterEntries[i] = common.ConvertStringToList(entry.(string))
		} else {
			iterEntries[i] = common.CopyList(value, length)
		}
	}
	iterator = NewZipIterator(iterEntries, entryLength)
	return
}
