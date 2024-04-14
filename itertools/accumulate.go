package itertools

import (
	"flex/common"
	"reflect"
)

type accumulator struct {
	entry   []any
	length  int
	handler func(any, any) any
	pointer int
	value   any
}

func NewAccumulator(entry []any, handler func(any, any) any) common.Iterator {
	return &accumulator{
		entry:   entry,
		length:  len(entry),
		handler: handler,
		pointer: 1,
		value:   entry[0],
	}
}

func (iter *accumulator) clear() {
	iter.value = nil
	iter.entry = nil
	iter.handler = nil
}

func (iter *accumulator) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.handler(iter.value, iter.entry[iter.pointer])
	iter.pointer++
	return true
}

func (iter *accumulator) Value() any {
	return iter.value
}

func (iter *accumulator) Pour() any {
	result := iter.Value()
	for iter.Next() {
		result = iter.Value()
	}
	return result
}

func Accumulate(handler, entry interface{}) (iterator common.Iterator, err error) {
	err = common.IsInputFuncValid(handler, 2, 1)
	if err != nil {
		return
	}
	err = common.IsSequence(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	length := value.Len()
	iterHandler := func(p1, p2 any) any {
		params := []reflect.Value{reflect.ValueOf(p1), reflect.ValueOf(p2)}
		return reflect.ValueOf(handler).Call(params)[0].Interface()
	}
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		iterator = NewAccumulator(common.CopyList(value, length), iterHandler)
	case reflect.String:
		list := common.ConvertStringToList(entry.(string))
		iterator = NewAccumulator(list, iterHandler)
	}
	return
}
