package functools

import (
	"flex/common"
	"flex/itertools"
	"reflect"
)

func Any(iter any) (result bool, err error) {
	iterator, ok := iter.(common.Iterator)
	if !ok {
		err = common.IsIterable(iter)
		if err != nil {
			return
		}
		value := reflect.ValueOf(iter)
		length := value.Len()
		iterHandler := func(a any) any {
			return a
		}
		switch value.Kind() {
		case reflect.Array, reflect.Slice:
			iterator = itertools.NewListConvertor(common.CopyList(value, length), iterHandler)
		case reflect.String:
			list := common.ConvertStringToList(iter.(string))
			iterator = itertools.NewListConvertor(list, iterHandler)
		case reflect.Map:
			iterator = itertools.NewMapConvertor(common.CopyMap(value, length), iterHandler)
		}
	}
	for iterator.Next() {
		value, ok := iterator.Value().(bool)
		if !ok {
			continue
		}
		if value {
			result = true
			return
		}
	}
	result = false
	return
}
