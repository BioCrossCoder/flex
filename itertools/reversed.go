package itertools

import (
	"flex/common"
	"reflect"
)

func Reversed(entry any) (output any, err error) {
	err = common.IsSequence(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	var length int
	var iterator common.Iterator
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		length = value.Len()
		iterator = NewSliceIterator(common.CopyList(value, length), length-1, 0, -1)
		output = iterator.Pour()
	case reflect.String:
		length = len([]rune(entry.(string)))
		iterator = NewSliceIterator(common.ConvertStringToList(entry.(string)), length-1, 0, -1)
		output = ""
		for _, c := range iterator.Pour().([]any) {
			output = output.(string) + c.(string)
		}
	}
	return
}
