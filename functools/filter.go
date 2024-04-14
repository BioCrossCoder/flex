package functools

import (
	"flex/common"
	"flex/itertools"
	"reflect"
)

func Filter(handler, entry any) (iterator itertools.Iterator, err error) {
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
		iterator = itertools.NewListFilter(common.CopyList(value, length), iterHandler)
	case reflect.String:
		list := common.ConvertStringToList(entry.(string))
		iterator = itertools.NewListFilter(list, iterHandler)
	case reflect.Map:
		iterator = itertools.NewMapFilter(common.CopyMap(value, length), iterHandler)
	}
	return
}

func FilterResult(handler, entry any) (output any, err error) {
	iterator, err := Filter(handler, entry)
	if err != nil {
		return
	}
	output = iterator.Pour()
	return
}
