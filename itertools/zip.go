package itertools

import (
	"flex/common"
	"reflect"
	"unicode/utf8"
)

type zipIterator struct {
	entries [][]any
	length  int
	pointer int
	value   any
	count   int
}

func NewZipIterator(entries [][]any, length int) Iterator {
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

func Zip(entries ...any) (iterator Iterator, err error) {
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
		var length int
		if value.Kind() == reflect.String {
			length = utf8.RuneCountInString(entry.(string))
		} else {
			length = value.Len()
		}
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

func ZipLongest(entries ...any) (iterator Iterator, err error) {
	entryCount := len(entries)
	if entryCount < 2 {
		err = common.ErrIllegalParamCount
		return
	}
	entryLength := 0
	iterEntries := make([][]any, entryCount)
	for _, entry := range entries {
		err = common.IsSequence(entry)
		if err != nil {
			return
		}
		value := reflect.ValueOf(entry)
		var length int
		if value.Kind() == reflect.String {
			length = utf8.RuneCountInString(entry.(string))
		} else {
			length = value.Len()
		}
		if length > entryLength {
			entryLength = length
		}
	}
	repeater := Repeat(nil)
	for i, entry := range entries {
		value := reflect.ValueOf(entry)
		var list []any
		var tailLength int
		if value.Kind() == reflect.String {
			list = common.ConvertStringToList(entry.(string))
			tailLength = entryLength - len(list)
		} else {
			list = common.CopyList(value, value.Len())
			tailLength = entryLength - value.Len()
		}
		iterEntries[i] = append(list, repeater.Repeat(tailLength)...)
	}
	iterator = NewZipIterator(iterEntries, entryLength)
	return
}
