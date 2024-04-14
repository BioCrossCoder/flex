package functools

import (
	"flex/common"
	"flex/itertools"
	"reflect"
)

func Map(handler, entry any) (iterator itertools.Iterator, err error) {
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
		iterator = itertools.NewListIterator(common.CopyList(value, length), iterHandler)
	case reflect.String:
		list := common.ConvertStringToList(entry.(string))
		iterator = itertools.NewListIterator(list, iterHandler)
	case reflect.Map:
		iterator = itertools.NewMapIterator(common.CopyMap(value, length), iterHandler)
	}
	return
}

func MapResult(handler, entry any) (output any, err error) {
	iterator, err := Map(handler, entry)
	if err != nil {
		return
	}
	output = iterator.Pour()
	return
}
