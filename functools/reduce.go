package functools

import (
	"flex/common"
	"flex/itertools"
	"reflect"
)

func Reduce(handler, entry interface{}) (iterator itertools.Iterator, err error) {
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
		iterator = itertools.NewAccumulator(common.CopyList(value, length), iterHandler)
	case reflect.String:
		list := common.ConvertStringToList(entry.(string))
		iterator = itertools.NewAccumulator(list, iterHandler)
	}
	return
}

func ReduceResult(handler, entry any) (output any, err error) {
	iterator, err := Reduce(handler, entry)
	if err != nil {
		return
	}
	output = iterator.Pour()
	return
}
