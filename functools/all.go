package functools

import (
	"github.com/biocrosscoder/flex/common"
	"github.com/biocrosscoder/flex/itertools"
	"reflect"
)

// All returns true if all elements in the iterable satisfy the condition, and false otherwise.
func All(iter any, condition func(any) bool) (result bool, err error) {
	iterator, ok := iter.(itertools.Iterator)
	if !ok {
		err = common.IsSequence(iter)
		if err != nil {
			return
		}
		value := reflect.ValueOf(iter)
		length := value.Len()
		switch value.Kind() { //nolint
		case reflect.Array, reflect.Slice:
			iterator = itertools.NewListIterator(common.CopyList(value, length))
		case reflect.String:
			iterator = itertools.NewListIterator(common.ConvertStringToList(value.String()))
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
