package functools

import (
	"github.com/biocrosscoder/flex/common"
	"github.com/biocrosscoder/flex/itertools"
	"reflect"
)

func All(iter any, condition func(any) bool) (result bool, err error) {
	iterator, ok := iter.(itertools.Iterator)
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
			err = common.ErrNotBool
			return
		case reflect.Map:
			iterator = itertools.NewMapIterator(common.CopyMap(value, length))
		}
	}
	result = true
	for iterator.Next() {
		if !condition(iterator.Value()) {
			result = false
			break
		}
	}
	return
}
