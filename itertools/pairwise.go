package itertools

import (
	"github.com/biocrosscoder/flex/common"
	"reflect"
)

type pairIterator struct {
	entry   []any
	length  int
	pointer int
	value   any
}

func NewPairIterator(entry []any) Iterator {
	return &pairIterator{
		entry:   entry,
		length:  len(entry) - 1,
		pointer: 0,
		value:   nil,
	}
}

func (iter *pairIterator) clear() {
	iter.value = nil
	iter.entry = nil
}

func (iter *pairIterator) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.entry[iter.pointer : iter.pointer+2]
	iter.pointer++
	return true
}

func (iter *pairIterator) Value() any {
	return iter.value
}

func (iter *pairIterator) Pour() any {
	length := iter.length - iter.pointer
	output := make([]any, length)
	i := 0
	for iter.Next() {
		output[i] = iter.Value()
		i++
	}
	return output
}

func PairWise(entry any) (iterator Iterator, err error) {
	err = common.IsSequence(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		iterator = NewPairIterator(common.CopyList(value, value.Len()))
	case reflect.String:
		iterator = NewPairIterator(common.ConvertStringToList(entry.(string)))
	}
	return
}
