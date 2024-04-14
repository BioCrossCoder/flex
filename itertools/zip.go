package itertools

import (
	"flex/common"
	"reflect"
)

type zipIterator struct {
	entry1  []any
	entry2  []any
	length  int
	pointer int
	value   any
}

func NewZipIterator(entry1, entry2 []any, length int) common.Iterator {
	return &zipIterator{
		entry1:  entry1,
		entry2:  entry2,
		length:  length,
		pointer: 0,
		value:   nil,
	}
}

func (iter *zipIterator) clear() {
	iter.entry1 = nil
	iter.entry2 = nil
	iter.value = nil
}

func (iter *zipIterator) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = [2]any{iter.entry1[iter.pointer], iter.entry2[iter.pointer]}
	iter.pointer++
	return true
}

func (iter *zipIterator) Value() any {
	return iter.value
}

func (iter *zipIterator) Pour() any {
	length := iter.length - iter.pointer
	output := make([][2]any, length)
	i := 0
	for iter.Next() {
		output[i] = iter.Value().([2]any)
		i++
	}
	return output
}

func Zip(iter1, iter2 any) (iterator common.Iterator, err error) {
	err = common.IsSequence(iter1)
	if err != nil {
		return
	}
	value1 := reflect.ValueOf(iter1)
	length1 := value1.Len()
	if value1.Kind() == reflect.String {
		iter1 = common.ConvertStringToList(iter1.(string))
		length1 = len(iter1.([]any))
	} else {
		iter1 = common.CopyList(value1, length1)
	}
	err = common.IsSequence(iter2)
	if err != nil {
		return
	}
	value2 := reflect.ValueOf(iter2)
	length2 := value2.Len()
	if value2.Kind() == reflect.String {
		iter2 = common.ConvertStringToList(iter2.(string))
		length2 = len(iter2.([]any))
	} else {
		iter2 = common.CopyList(value2, length2)
	}
	length := length1
	if length2 < length1 {
		length = length2
	}
	iterator = NewZipIterator(iter1.([]any), iter2.([]any), length)
	return
}

func ZipResult(iter1, iter2 any) (output [][2]any, err error) {
	iterator, err := Zip(iter1, iter2)
	if err != nil {
		return
	}
	output = iterator.Pour().([][2]any)
	return
}
