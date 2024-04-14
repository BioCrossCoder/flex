package itertools

import (
	"flex/common"
	"reflect"
)

type listFilter struct {
	listConvertor
}

func NewListFilter(entry []any, handler func(any) bool) common.Iterator {
	iter := new(listFilter)
	iter.entry = entry
	iter.length = len(entry)
	iter.handler = func(p any) any {
		return handler(p)
	}
	iter.pointer = 0
	iter.value = false
	return iter
}

func (iter *listFilter) Pour() any {
	output := make([]any, 0)
	for iter.Next() {
		if iter.Value().(bool) {
			value := iter.entry[iter.pointer-1]
			output = append(output, value)
		}
	}
	return output
}

type mapFilter struct {
	mapConvertor
}

func NewMapFilter(entry map[any]any, handler func(any) bool) common.Iterator {
	keys, values, length := common.ConvertMapToLists(entry)
	iter := new(mapFilter)
	iter.entryKeys = keys
	iter.entryValues = values
	iter.length = length
	iter.handler = func(p any) any {
		return handler(p)
	}
	iter.pointer = 0
	iter.value = false
	return iter
}

func (iter *mapFilter) Pour() any {
	output := make(map[any]any)
	for iter.Next() {
		if iter.Value().(bool) {
			key := iter.entryKeys[iter.pointer-1]
			value := iter.entryValues[iter.pointer-1]
			output[key] = value
		}
	}
	return output
}

func Filter(handler, entry any) (iterator common.Iterator, err error) {
	err = common.IsJudgeFunc(handler)
	if err != nil {
		return
	}
	err = common.IsIterable(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	length := value.Len()
	iterHandler := func(a any) bool {
		params := []reflect.Value{reflect.ValueOf(a)}
		return reflect.ValueOf(handler).Call(params)[0].Bool()
	}
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		iterator = NewListFilter(common.CopyList(value, length), iterHandler)
	case reflect.String:
		list := common.ConvertStringToList(entry.(string))
		iterator = NewListFilter(list, iterHandler)
	case reflect.Map:
		iterator = NewMapFilter(common.CopyMap(value, length), iterHandler)
	}
	return
}
