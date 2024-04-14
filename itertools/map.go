package itertools

import (
	"flex/common"
	"reflect"
)

type listConvertor struct {
	entry   []any
	length  int
	handler func(any) any
	pointer int
	value   any
}

func NewListConvertor(entry []any, handler func(any) any) common.Iterator {
	return &listConvertor{
		entry:   entry,
		length:  len(entry),
		handler: handler,
		pointer: 0,
		value:   nil,
	}
}

func (iter *listConvertor) clear() {
	iter.value = nil
	iter.entry = nil
	iter.handler = nil
}

func (iter *listConvertor) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.handler(iter.entry[iter.pointer])
	iter.pointer++
	return true
}

func (iter *listConvertor) Value() any {
	return iter.value
}

func (iter *listConvertor) Pour() any {
	length := iter.length - iter.pointer
	output := make([]any, length)
	i := 0
	for iter.Next() {
		output[i] = iter.Value()
		i++
	}
	return output
}

type mapConvertor struct {
	entryKeys   []any
	entryValues []any
	length      int
	handler     func(any) any
	pointer     int
	value       any
}

func NewMapConvertor(entry map[any]any, handler func(any) any) common.Iterator {
	keys, values, length := common.ConvertMapToLists(entry)
	return &mapConvertor{
		entryKeys:   keys,
		entryValues: values,
		length:      length,
		handler:     handler,
		pointer:     0,
		value:       nil,
	}
}

func (iter *mapConvertor) clear() {
	iter.value = nil
	iter.entryKeys = nil
	iter.entryValues = nil
	iter.handler = nil
}

func (iter *mapConvertor) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.handler(iter.entryValues[iter.pointer])
	iter.pointer++
	return true
}

func (iter *mapConvertor) Value() any {
	return iter.value
}

func (iter *mapConvertor) Pour() any {
	length := iter.length - iter.pointer
	output := make(map[any]any, length)
	for iter.Next() {
		key := iter.entryKeys[iter.pointer-1]
		output[key] = iter.Value()
	}
	return output
}

func Map(handler, entry any) (iterator common.Iterator, err error) {
	err = common.IsInputFuncValid(handler, 1, 1)
	if err != nil {
		return
	}
	err = common.IsIterable(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	length := value.Len()
	iterHandler := func(a any) any {
		params := []reflect.Value{reflect.ValueOf(a)}
		return reflect.ValueOf(handler).Call(params)[0].Interface()
	}
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		iterator = NewListConvertor(common.CopyList(value, length), iterHandler)
	case reflect.String:
		list := common.ConvertStringToList(entry.(string))
		iterator = NewListConvertor(list, iterHandler)
	case reflect.Map:
		iterator = NewMapConvertor(common.CopyMap(value, length), iterHandler)
	}
	return
}
