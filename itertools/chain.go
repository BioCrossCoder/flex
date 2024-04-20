package itertools

import (
	"flex/common"
	"reflect"
)

func Chain(entries ...any) (iterator common.Iterator, err error) {
	elements := make([]any, 0)
	for _, entry := range entries {
		err = common.IsSequence(entry)
		if err != nil {
			return
		}
		value := reflect.ValueOf(entry)
		length := value.Len()
		switch value.Kind() {
		case reflect.Array, reflect.Slice:
			elements = append(elements, common.CopyList(value, length)...)
		case reflect.String:
			elements = append(elements, common.ConvertStringToList(entry.(string))...)
		}
	}
	return NewListIterator(elements), nil
}
