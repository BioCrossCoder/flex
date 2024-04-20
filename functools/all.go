package functools

import (
	"flex/common"
	"flex/itertools"
	"reflect"
)

func All(iter any) (result bool, err error) {
	iterator, ok := iter.(common.Iterator)
	if !ok {
		err = common.IsIterable(iter)
		if err != nil {
			return
		}
		value := reflect.ValueOf(iter)
		length := value.Len()
		switch value.Kind() {
		case reflect.Array, reflect.Slice:
			iterator = itertools.NewListIterator(common.CopyList(value, length))
		case reflect.String:
			iterator = itertools.NewListIterator(common.ConvertStringToList(iter.(string)))
		case reflect.Map:
			iterator = itertools.NewMapIterator(common.CopyMap(value, length))
		}
	}
	for iterator.Next() {
		value, ok := iterator.Value().(bool)
		if !ok {
			err = common.ErrNotBool
			return
		}
		if !value {
			result = false
			return
		}
	}
	result = true
	return
}
